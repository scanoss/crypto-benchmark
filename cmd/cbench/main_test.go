package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

func benchmark() map[string]string {
	return map[string]string{
		"corpus/libraries.yaml": `schema_version: "1"
stdlib:
  - ecosystem: go
    package: go-stdlib
    modules: [crypto/sha256]
    source: https://pkg.go.dev/crypto
    retrieved: 2026-11-04
`,
		"corpus/manifest.yaml": `schema_version: "1"
entries:
  - id: go-lib
    language: go
    suite: C
    split: dev
    url: https://example.com/lib
    sha: 0123456789abcdef0123456789abcdef01234567
  - id: go-app
    language: go
    suite: C
    split: dev
    url: https://example.com/app
    sha: 89abcdef0123456789abcdef0123456789abcdef
`,
		"groundtruth/C/go-lib.yaml": `schema_version: "1"
entry: go-lib
scope:
  kind: exhaustive
assets:
  - id: gt-0001
    file: hash.go
    lines: [4, 4]
    asset_type: algorithm
    algorithm: SHA-2
    origin: project
    reachable: true
    depth: 0
`,
		"groundtruth/C/go-app.yaml": `schema_version: "1"
entry: go-app
scope:
  kind: exhaustive
assets: []
`,
		"predictions/crypto-finder/go-lib.json": `{"schema_version":"1","tool":{"name":"crypto-finder","version":"0.9.0","config":{"ruleset_source":"remote","ruleset_version":"v1.40.0"}},
"entry":"go-lib","reachability_reported":false,
"assets":[{"file":"hash.go","lines":[4,4],"asset_type":"algorithm","algorithm":"SHA-2","origin":"project","reachable":"unknown"}]}`,
	}
}

func TestScoreMarkdownHeader(t *testing.T) {
	root := writeTree(t, benchmark())
	var stdout, stderr bytes.Buffer
	code := run([]string{"score", "--root", root, "--tool", "crypto-finder", "--predictions", filepath.Join(root, "predictions")}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit %d, stderr:\n%s", code, stderr.String())
	}
	head, _, _ := strings.Cut(stdout.String(), "## Detection")
	want := `# Score: crypto-finder

- Tool: crypto-finder 0.9.0
- Config: ruleset_source=remote, ruleset_version=v1.40.0

Entries scored (1): go-lib

Not run, no prediction file (1): go-app

`
	if head != want {
		t.Fatalf("header:\n%s\nwant:\n%s", head, want)
	}
	if !strings.Contains(stdout.String(), "| all | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |") {
		t.Fatalf("overall row missing:\n%s", stdout.String())
	}
}

func TestScoreJSONWritesFile(t *testing.T) {
	root := writeTree(t, benchmark())
	out := filepath.Join(root, "score.json")
	var stdout, stderr bytes.Buffer
	code := run([]string{"score", "--root", root, "--tool", "crypto-finder", "--predictions", filepath.Join(root, "predictions"), "--format", "json", "--out", out}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit %d, stderr:\n%s", code, stderr.String())
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"not_run": [
    "go-app"
  ]`, `"reachability": null`, `"pair_discrimination": null`, `"precision": 1,`} {
		if !strings.Contains(string(data), want) {
			t.Errorf("JSON lacks %s", want)
		}
	}
	if t.Failed() {
		t.Log(string(data))
	}
}

func TestScoreRefusesMislabeledPredictions(t *testing.T) {
	files := benchmark()
	files["predictions/crypto-finder/go-lib.json"] = strings.Replace(files["predictions/crypto-finder/go-lib.json"], `"entry":"go-lib"`, `"entry":"go-app"`, 1)
	root := writeTree(t, files)
	var stdout, stderr bytes.Buffer
	code := run([]string{"score", "--root", root, "--tool", "crypto-finder", "--predictions", filepath.Join(root, "predictions")}, &stdout, &stderr)
	if code != 1 || !strings.Contains(stderr.String(), `entry: "go-app" does not match the file name "go-lib"`) {
		t.Fatalf("exit %d, stderr:\n%s", code, stderr.String())
	}
}

func TestAdaptCBOMThenScore(t *testing.T) {
	files := benchmark()
	files["bom.json"] = `{"bomFormat":"CycloneDX","specVersion":"1.6","components":[
		{"type":"cryptographic-asset","bom-ref":"a","name":"SHA-256","cryptoProperties":{"assetType":"algorithm","algorithmProperties":{"primitive":"hash"}},
		 "evidence":{"occurrences":[{"location":"hash.go","line":4}]}},
		{"type":"cryptographic-asset","bom-ref":"b","name":"MD5","cryptoProperties":{"assetType":"algorithm"}}]}`
	root := writeTree(t, files)
	var stdout, stderr bytes.Buffer
	out := filepath.Join(root, "predictions/sonar/go-lib.json")
	args := []string{"adapt", "cbom", "--root", root, "--tool", "sonar", "--tool-version", "2.0", "--entry", "go-lib", "--input", filepath.Join(root, "bom.json"), "--out", out}
	if code := run(args, &stdout, &stderr); code != 0 {
		t.Fatalf("exit %d, stderr:\n%s", code, stderr.String())
	}
	if got := stderr.String(); got != "skipped MD5 (b): no occurrence location; it cannot be matched\n" {
		t.Fatalf("stderr: %q", got)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"name": "sonar"`, `"version": "2.0"`, `"reachability_reported": false`, `"components_without_location": "1"`} {
		if !strings.Contains(string(data), want) {
			t.Fatalf("prediction lacks %s:\n%s", want, data)
		}
	}
	stdout.Reset()
	stderr.Reset()
	if code := run([]string{"score", "--root", root, "--predictions", filepath.Join(root, "predictions"), "--tool", "sonar"}, &stdout, &stderr); code != 0 {
		t.Fatalf("score exit %d, stderr:\n%s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "| all | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |") {
		t.Fatalf("score report:\n%s", stdout.String())
	}

	stderr.Reset()
	noVersion := []string{"adapt", "cbom", "--root", root, "--tool", "sonar", "--entry", "go-lib", "--input", filepath.Join(root, "bom.json"), "--out", out}
	if code := run(noVersion, &stdout, &stderr); code != 1 || stderr.String() != "adapt cbom: --tool and --tool-version are required\n" {
		t.Fatalf("missing --tool-version: exit %d, stderr %q", code, stderr.String())
	}
}
