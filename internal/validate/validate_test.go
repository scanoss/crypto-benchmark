package validate

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
)

const manifest = `schema_version: "1"
entries:
  - id: go-remote
    language: go
    suite: C
    split: dev
    url: https://example.com/remote
    sha: 0123456789abcdef0123456789abcdef01234567
  - id: java-pair-r
    language: java
    suite: B
    split: dev
    path: fixtures/r
    pair: java-pair-u
    variant: reachable
    scenario: transitive
  - id: java-pair-u
    language: java
    suite: B
    split: dev
    path: fixtures/u
    pair: java-pair-r
    variant: unreachable
    scenario: transitive
`

const gtRemote = `schema_version: "1"
entry: go-remote
scope:
  kind: files
  files: [crypto/a.go]
assets:
  - id: gt-0001
    file: crypto/a.go
    lines: [3, 4]
    asset_type: algorithm
    primitive: ae
    algorithm: AES
    parameters:
      mode: GCM
      key_size: 256
    origin: dependency
    component: pkg:golang/example.com/lib
    reachable: true
    depth: 1
`

const gtPair = `schema_version: "1"
entry: java-pair-r
scope:
  kind: exhaustive
assets:
  - id: gt-0001
    file: src/Main.java
    lines: [7, 7]
    asset_type: protocol
    origin: project
    reachable: false
    depth: 0
    distance: 0
`

const overlayManifest = `schema_version: "1"
entries:
  - id: py-secret
    language: python
    suite: C
    split: holdout
    url: https://example.com/secret
    sha: 89abcdef0123456789abcdef0123456789abcdef
`

const gtOverlay = `schema_version: "1"
entry: py-secret
scope:
  kind: pool
assets: []
`

const libraryList = `schema_version: "1"
stdlib:
  - ecosystem: go
    package: go-stdlib
    modules: [crypto/sha256]
    source: https://pkg.go.dev/crypto
    retrieved: 2026-11-04
`

func baseFiles() map[string]string {
	return map[string]string{
		"corpus/libraries.yaml":          libraryList,
		"corpus/manifest.yaml":           manifest,
		"groundtruth/C/go-remote.yaml":   gtRemote,
		"groundtruth/B/java-pair-r.yaml": gtPair,
		"fixtures/r/.gitkeep":            "",
		"fixtures/u/.gitkeep":            "",
	}
}

func overlayFiles() map[string]string {
	return map[string]string{
		"corpus/manifest.yaml":         overlayManifest,
		"groundtruth/C/py-secret.yaml": gtOverlay,
	}
}

func writeTree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for name, body := range files {
		p := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func rel(root string, p check.Problem) string {
	r, err := filepath.Rel(root, p.File)
	if err != nil {
		return p.File
	}
	return filepath.ToSlash(r)
}

func TestValidBaseAndOverlay(t *testing.T) {
	base, overlay := writeTree(t, baseFiles()), writeTree(t, overlayFiles())
	r := Run(base, overlay)
	var got []string
	for _, p := range r.Problems {
		got = append(got, p.Severity.String()+" "+rel(base, p)+" "+p.Path+" "+p.Msg)
	}
	want := []string{`warning corpus/manifest.yaml entries[2] "java-pair-u" is not annotated yet (no groundtruth/B/java-pair-u.yaml)`}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("problems:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

type edit struct {
	file, old, new string
}

func TestEachRuleDetectsItsDefect(t *testing.T) {
	cases := []struct {
		name      string
		base      []edit
		overlay   []edit // non-nil (even empty) means pass an overlay
		file      string // expected problem, file relative to its root
		path, msg string
	}{
		{name: "unknown manifest field", base: []edit{{"corpus/manifest.yaml", "    split: dev\n    url", "    split: dev\n    colour: red\n    url"}},
			file: "corpus/manifest.yaml", msg: "field colour not found"},
		{name: "manifest schema version", base: []edit{{"corpus/manifest.yaml", `schema_version: "1"`, `schema_version: "2"`}},
			file: "corpus/manifest.yaml", path: "schema_version", msg: `got "2", want "1"`},
		{name: "unknown language", base: []edit{{"corpus/manifest.yaml", "language: go", "language: cobol"}},
			file: "corpus/manifest.yaml", path: "entries[0].language", msg: `"cobol" is not one of java|python|go|rust|javascript`},
		{name: "unknown suite", base: []edit{{"corpus/manifest.yaml", "suite: C", "suite: D"}},
			file: "corpus/manifest.yaml", path: "entries[0].suite", msg: `"D" is not one of A|B|C`},
		{name: "unknown split", base: []edit{{"corpus/manifest.yaml", "split: dev\n    url", "split: test\n    url"}},
			file: "corpus/manifest.yaml", path: "entries[0].split", msg: `"test" is not one of dev|holdout`},
		{name: "bad id", base: []edit{{"corpus/manifest.yaml", "id: go-remote", "id: Go_Remote"}},
			file: "corpus/manifest.yaml", path: "entries[0].id", msg: "must match [a-z0-9-]+"},
		{name: "duplicate id", base: []edit{{"corpus/manifest.yaml", "id: java-pair-u", "id: go-remote"}},
			file: "corpus/manifest.yaml", path: "entries[2].id", msg: `duplicate id "go-remote"`},
		{name: "url and path", base: []edit{{"corpus/manifest.yaml", "    sha: 0123456789abcdef0123456789abcdef01234567\n", "    sha: 0123456789abcdef0123456789abcdef01234567\n    path: fixtures/r\n"}},
			file: "corpus/manifest.yaml", path: "entries[0].path", msg: "not both"},
		{name: "neither url nor path", base: []edit{{"corpus/manifest.yaml", "    path: fixtures/u\n", ""}},
			file: "corpus/manifest.yaml", path: "entries[2].url", msg: "set either url+sha or path"},
		{name: "short sha", base: []edit{{"corpus/manifest.yaml", "sha: 0123456789abcdef0123456789abcdef01234567", "sha: 0123456"}},
			file: "corpus/manifest.yaml", path: "entries[0].sha", msg: "40-character"},
		{name: "url without sha", base: []edit{{"corpus/manifest.yaml", "    sha: 0123456789abcdef0123456789abcdef01234567\n", ""}},
			file: "corpus/manifest.yaml", path: "entries[0].sha", msg: "40-character"},
		{name: "path missing", base: []edit{{"corpus/manifest.yaml", "path: fixtures/u", "path: fixtures/missing"}},
			file: "corpus/manifest.yaml", path: "entries[2].path", msg: `"fixtures/missing" is not an existing directory`},
		{name: "path escapes repo", base: []edit{{"corpus/manifest.yaml", "path: fixtures/u", "path: ../outside"}},
			file: "corpus/manifest.yaml", path: "entries[2].path", msg: "relative path inside the repository"},
		{name: "pair not symmetric", base: []edit{{"corpus/manifest.yaml", "pair: java-pair-r", "pair: go-remote"}},
			file: "corpus/manifest.yaml", path: "entries[1].pair", msg: `partner "java-pair-u" does not pair back with "java-pair-r"`},
		{name: "pair partner unknown", base: []edit{{"corpus/manifest.yaml", "pair: java-pair-u", "pair: nobody"}},
			file: "corpus/manifest.yaml", path: "entries[1].pair", msg: `partner "nobody" is not a manifest entry`},
		{name: "pair both reachable", base: []edit{{"corpus/manifest.yaml", "variant: unreachable", "variant: reachable"}},
			file: "corpus/manifest.yaml", path: "entries[1].variant", msg: "both reachable"},
		{name: "pair without variant", base: []edit{{"corpus/manifest.yaml", "    variant: unreachable\n", ""}},
			file: "corpus/manifest.yaml", path: "entries[2].variant", msg: "required with pair"},
		{name: "variant without pair", base: []edit{{"corpus/manifest.yaml", "    sha: 0123456789abcdef0123456789abcdef01234567\n", "    sha: 0123456789abcdef0123456789abcdef01234567\n    variant: reachable\n"}},
			file: "corpus/manifest.yaml", path: "entries[0].variant", msg: "only allowed together with pair"},
		{name: "pair outside suite B", base: []edit{{"corpus/manifest.yaml", "suite: B\n    split: dev\n    path: fixtures/u", "suite: A\n    split: dev\n    path: fixtures/u"}},
			file: "corpus/manifest.yaml", path: "entries[2].pair", msg: "only allowed in suite B"},
		{name: "pair languages differ", base: []edit{{"corpus/manifest.yaml", "language: java\n    suite: B\n    split: dev\n    path: fixtures/u", "language: rust\n    suite: B\n    split: dev\n    path: fixtures/u"}},
			file: "corpus/manifest.yaml", path: "entries[1].language", msg: `its partner "java-pair-u" is rust`},
		{name: "holdout in base", base: []edit{{"corpus/manifest.yaml", "split: dev\n    url", "split: holdout\n    url"}},
			file: "corpus/manifest.yaml", path: "entries[0].split", msg: "only allowed in the private overlay"},
		{name: "id duplicated across base and overlay", overlay: []edit{{"corpus/manifest.yaml", "id: py-secret", "id: go-remote"}},
			file: "corpus/manifest.yaml", path: "entries[0].id", msg: `duplicate id "go-remote", first declared at`},
		{name: "holdout ground truth in base", base: []edit{{"groundtruth/C/py-secret.yaml", "", gtOverlay}}, overlay: []edit{{"groundtruth/C/py-secret.yaml", "", ""}},
			file: "groundtruth/C/py-secret.yaml", path: "entry", msg: "its ground truth must live in the overlay"},

		{name: "library list missing", base: []edit{{"corpus/libraries.yaml", "", ""}},
			file: "corpus/libraries.yaml", msg: "no such file"},
		{name: "library list checked", base: []edit{{"corpus/libraries.yaml", "    modules: [crypto/sha256]\n", ""}},
			file: "corpus/libraries.yaml", path: "stdlib[0].modules", msg: "required"},

		{name: "unknown ground truth field", base: []edit{{"groundtruth/C/go-remote.yaml", "    depth: 1\n", "    depth: 1\n    confidence: high\n"}},
			file: "groundtruth/C/go-remote.yaml", msg: "field confidence not found"},
		{name: "ground truth schema version", base: []edit{{"groundtruth/B/java-pair-r.yaml", `schema_version: "1"`, `schema_version: "0"`}},
			file: "groundtruth/B/java-pair-r.yaml", path: "schema_version", msg: `got "0", want "1"`},
		{name: "entry does not match file name", base: []edit{{"groundtruth/B/java-pair-r.yaml", "entry: java-pair-r", "entry: java-pair-u"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "entry", msg: `"java-pair-u" does not match the file name "java-pair-r"`},
		{name: "entry not in manifest", base: []edit{{"groundtruth/B/ghost.yaml", "", strings.Replace(gtPair, "entry: java-pair-r", "entry: ghost", 1)}},
			file: "groundtruth/B/ghost.yaml", path: "entry", msg: `"ghost" is not a manifest entry`},
		{name: "wrong suite directory", base: []edit{{"groundtruth/A/java-pair-u.yaml", "", strings.Replace(gtPair, "java-pair-r", "java-pair-u", 1)}},
			file: "groundtruth/A/java-pair-u.yaml", path: "entry", msg: "must be groundtruth/B/java-pair-u.yaml"},
		{name: "annotated twice across base and overlay", overlay: []edit{{"groundtruth/C/go-remote.yaml", "", gtRemote}},
			file: "groundtruth/C/go-remote.yaml", path: "entry", msg: `"go-remote" is already annotated in`},
		{name: "unknown scope kind", base: []edit{{"groundtruth/B/java-pair-r.yaml", "kind: exhaustive", "kind: partial"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "scope.kind", msg: `"partial" is not one of exhaustive|files|pool`},
		{name: "files scope without files", base: []edit{{"groundtruth/B/java-pair-r.yaml", "kind: exhaustive", "kind: files"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "scope.files", msg: "required when scope.kind is files"},
		{name: "files listed on exhaustive scope", base: []edit{{"groundtruth/B/java-pair-r.yaml", "kind: exhaustive", "kind: exhaustive\n  files: [src/Main.java]"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "scope.files", msg: "only allowed when scope.kind is files"},
		{name: "asset outside files scope", base: []edit{{"groundtruth/C/go-remote.yaml", "    file: crypto/a.go", "    file: crypto/b.go"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].file", msg: `"crypto/b.go" is outside scope.files`},
		{name: "absolute asset file", base: []edit{{"groundtruth/B/java-pair-r.yaml", "file: src/Main.java", "file: /etc/Main.java"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].file", msg: "must be relative to the entry root"},
		{name: "duplicate asset id", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", "    distance: 0\n" + strings.SplitN(gtPair, "assets:\n", 2)[1]}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[1].id", msg: `duplicate asset id "gt-0001"`},
		{name: "anchor not a pair", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [3, 4]\n    anchors: [[1]]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].anchors[0]", msg: "must be [start, end]"},
		{name: "anchor start after end", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [3, 4]\n    anchors: [[2, 1]]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].anchors[0]", msg: "need 1 <= start <= end"},
		{name: "anchor overlapping lines", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [3, 4]\n    anchors: [[1, 3]]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].anchors[0]", msg: "overlaps lines [3, 4]"},
		{name: "anchors overlapping each other", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [3, 4]\n    anchors: [[1, 1], [1, 2]]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].anchors[1]", msg: "overlaps another anchor [1, 1]"},
		{name: "anchor on a protocol asset", base: []edit{{"groundtruth/B/java-pair-r.yaml", "lines: [7, 7]", "lines: [7, 7]\n    anchors: [[5, 5]]"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].anchors[0]", msg: "only allowed when asset_type is algorithm"},
		{name: "start after end", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [5, 4]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].lines", msg: "need 1 <= start <= end"},
		{name: "zero start line", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [0, 4]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].lines", msg: "need 1 <= start <= end"},
		{name: "lines not a pair", base: []edit{{"groundtruth/C/go-remote.yaml", "lines: [3, 4]", "lines: [3]"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].lines", msg: "must be [start, end]"},
		{name: "unknown asset type", base: []edit{{"groundtruth/B/java-pair-r.yaml", "asset_type: protocol", "asset_type: key"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].asset_type", msg: `"key" is not one of algorithm|certificate|protocol|related-crypto-material`},
		{name: "unknown primitive", base: []edit{{"groundtruth/C/go-remote.yaml", "primitive: ae", "primitive: aead"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].primitive", msg: `"aead" is not one of drbg|mac|`},
		{name: "primitive on non-algorithm", base: []edit{{"groundtruth/B/java-pair-r.yaml", "asset_type: protocol", "asset_type: protocol\n    primitive: hash"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].primitive", msg: "only allowed when asset_type is algorithm"},
		{name: "unknown material type", base: []edit{{"groundtruth/B/java-pair-r.yaml", "asset_type: protocol", "asset_type: related-crypto-material\n    material_type: iv"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].material_type", msg: `"iv" is not one of private-key|public-key|`},
		{name: "material type on non-material", base: []edit{{"groundtruth/C/go-remote.yaml", "primitive: ae", "primitive: ae\n    material_type: nonce"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].material_type", msg: "only allowed when asset_type is related-crypto-material"},
		{name: "non-positive key size", base: []edit{{"groundtruth/C/go-remote.yaml", "key_size: 256", "key_size: 0"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].parameters.key_size", msg: "must be > 0, got 0"},
		{name: "unknown parameter", base: []edit{{"groundtruth/C/go-remote.yaml", "key_size: 256", "key_size: 256\n      rounds: 10"}},
			file: "groundtruth/C/go-remote.yaml", msg: "field rounds not found"},
		{name: "unknown origin", base: []edit{{"groundtruth/B/java-pair-r.yaml", "origin: project", "origin: vendor"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].origin", msg: `"vendor" is not one of project|dependency`},
		{name: "dependency without component", base: []edit{{"groundtruth/C/go-remote.yaml", "    component: pkg:golang/example.com/lib\n", ""}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].component", msg: "required when origin is dependency"},
		{name: "component on project asset", base: []edit{{"groundtruth/B/java-pair-r.yaml", "origin: project", "origin: project\n    component: pkg:maven/x/y"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].component", msg: "only allowed when origin is dependency"},
		{name: "unknown reachable value", base: []edit{{"groundtruth/B/java-pair-r.yaml", "reachable: false", "reachable: maybe"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].reachable", msg: `"maybe" is not one of true|false|unknown`},
		{name: "known reachability without depth", base: []edit{{"groundtruth/C/go-remote.yaml", "    depth: 1\n", ""}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].depth", msg: "required when reachable is true"},
		{name: "depth with unknown reachability", base: []edit{{"groundtruth/C/go-remote.yaml", "reachable: true", "reachable: unknown"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].depth", msg: "not allowed when reachable is unknown"},
		{name: "object with one asset", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", "    distance: 0\n    object: main/ctx\n"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].object", msg: `object "main/ctx" has only one asset; leave object out`},
		{name: "object mixing asset types", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", "    distance: 0\n    object: main/ctx\n" + objectTwin("gt-0002", "asset_type: certificate")}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[1].object", msg: `object "main/ctx" mixes asset types protocol (gt-0001) and certificate`},
		{name: "object on material with a material type", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    asset_type: protocol\n    origin: project\n    reachable: false\n    depth: 0\n    distance: 0\n", "    asset_type: related-crypto-material\n    material_type: nonce\n    origin: project\n    reachable: false\n    depth: 0\n    distance: 0\n    object: main/ctx\n" + objectTwin("gt-0002", "asset_type: related-crypto-material")}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].object", msg: `"main/ctx" is set on related crypto material with material_type nonce, which is its own object`},
		{name: "pair without scenario", base: []edit{{"corpus/manifest.yaml", "    variant: reachable\n    scenario: transitive\n", "    variant: reachable\n"}},
			file: "corpus/manifest.yaml", path: "entries[1].scenario", msg: "required on a suite B pair entry"},
		{name: "pair scenarios differ", base: []edit{{"corpus/manifest.yaml", "    variant: unreachable\n    scenario: transitive\n", "    variant: unreachable\n    scenario: dispatch\n"}},
			file: "corpus/manifest.yaml", path: "entries[1].scenario", msg: `partner "java-pair-u" has "dispatch"`},
		{name: "scenario without pair", base: []edit{{"corpus/manifest.yaml", "    sha: 0123456789abcdef0123456789abcdef01234567\n", "    sha: 0123456789abcdef0123456789abcdef01234567\n    scenario: transitive\n"}},
			file: "corpus/manifest.yaml", path: "entries[0].scenario", msg: "only allowed together with pair"},
		{name: "pair asset with unknown reachability", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    reachable: false\n    depth: 0\n    distance: 0\n", "    reachable: unknown\n"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].reachable", msg: "must be true or false in a suite B pair"},
		{name: "pair asset without distance", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", ""}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].distance", msg: "required in a suite B pair"},
		{name: "reachable pair asset without path", base: []edit{{"groundtruth/B/java-pair-r.yaml", "reachable: false", "reachable: true"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].path", msg: "required when reachable is true"},
		{name: "unreachable pair asset with path", base: []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", "    distance: 0\n    path: [Main.main]\n"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].path", msg: "not allowed when reachable is false"},
		{name: "distance disagrees with depth", base: []edit{{"groundtruth/C/go-remote.yaml", "    depth: 1\n", "    depth: 1\n    distance: 3\n"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].distance", msg: "disagrees with depth 1"},
		{name: "project asset with nonzero distance", base: []edit{{"groundtruth/B/java-pair-r.yaml", "distance: 0", "distance: 1"}},
			file: "groundtruth/B/java-pair-r.yaml", path: "assets[0].distance", msg: "must be 0 exactly when origin is project"},
		{name: "component outside scope.components", base: []edit{{"groundtruth/C/go-remote.yaml", "  files: [crypto/a.go]\n", "  files: [crypto/a.go]\n  components: [pkg:golang/example.com/other]\n"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].component", msg: "is not in scope.components"},
		{name: "scope component not a purl", base: []edit{{"groundtruth/C/go-remote.yaml", "  files: [crypto/a.go]\n", "  files: [crypto/a.go]\n  components: [example.com/lib]\n"}},
			file: "groundtruth/C/go-remote.yaml", path: "scope.components[0]", msg: "is not a package URL"},
		{name: "depth out of range", base: []edit{{"groundtruth/C/go-remote.yaml", "depth: 1", "depth: 3"}},
			file: "groundtruth/C/go-remote.yaml", path: "assets[0].depth", msg: "must be 0, 1 or 2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base := writeTree(t, applyEdits(t, baseFiles(), tc.base))
			overlay := ""
			if tc.overlay != nil {
				overlay = writeTree(t, applyEdits(t, overlayFiles(), tc.overlay))
			}
			r := Run(base, overlay)
			for _, p := range r.Problems {
				f := rel(base, p)
				if overlay != "" && strings.HasPrefix(p.File, overlay) {
					f = rel(overlay, p)
				}
				if p.Severity == check.Error && f == tc.file && p.Path == tc.path && strings.Contains(p.Msg, tc.msg) {
					return
				}
			}
			var got []string
			for _, p := range r.Problems {
				got = append(got, p.String())
			}
			t.Fatalf("want error %s: %s: ...%s...\ngot:\n%s", tc.file, tc.path, tc.msg, strings.Join(got, "\n"))
		})
	}
}

// objectTwin is a second asset for gtPair, on object main/ctx, with id and asset type replaced.
func objectTwin(id, assetType string) string {
	twin := strings.SplitN(gtPair, "assets:\n", 2)[1]
	twin = strings.Replace(twin, "id: gt-0001", "id: "+id, 1)
	return strings.Replace(twin, "asset_type: protocol", assetType, 1) + "    object: main/ctx\n"
}

// applyEdits replaces old with new in file. With an empty old, a missing file is
// created with new as its body, and an existing file is deleted when new is empty.
func applyEdits(t *testing.T, files map[string]string, edits []edit) map[string]string {
	t.Helper()
	for _, e := range edits {
		body, exists := files[e.file]
		switch {
		case e.old == "" && !exists:
			files[e.file] = e.new
		case e.old == "" && e.new == "":
			delete(files, e.file)
		case !strings.Contains(body, e.old):
			t.Fatalf("edit target %q not found in %s", e.old, e.file)
		default:
			files[e.file] = strings.Replace(body, e.old, e.new, 1)
		}
	}
	return files
}

func TestTwoAssetsOfOneTypeFormAnObject(t *testing.T) {
	edits := []edit{{"groundtruth/B/java-pair-r.yaml", "    distance: 0\n", "    distance: 0\n    object: main/ctx\n" + objectTwin("gt-0002", "asset_type: protocol")}}
	r := Run(writeTree(t, applyEdits(t, baseFiles(), edits)), "")
	if r.HasErrors() {
		for _, p := range r.Problems {
			t.Error(p.String())
		}
	}
}
