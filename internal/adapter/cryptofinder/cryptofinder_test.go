package cryptofinder

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// The fixtures are a real crypto-finder scan of its
// testdata/projects/transitive_reachability/go fixture (app -> liba -> libb,
// crypto in libb and in the app) with --scan-dependencies and
// --export-callgraph; the call graph is trimmed to the fields the adapter reads.
func fixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func describe(a prediction.Asset) string {
	origin := "project"
	if d, ok := a.Origin.(groundtruth.Dependency); ok {
		origin = d.Component
	}
	depth := "-"
	if a.Depth != nil {
		depth = fmt.Sprint(*a.Depth)
	}
	s := fmt.Sprintf("%s %d-%d %s %s %s %s %+v reachable=%s depth=%s",
		a.File, a.Lines.Start, a.Lines.End, a.Type, a.Primitive, a.Algorithm, origin, a.Params, a.Reachable, depth)
	if a.MaterialType != "" {
		s += " material_type=" + string(a.MaterialType)
	}
	return s
}

func adapt(t *testing.T, interim, callgraph string) Result {
	t.Helper()
	var graph []byte
	if callgraph != "" {
		graph = []byte(callgraph)
	}
	res, err := Adapt("go-transitive", []byte(interim), graph)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func describeAll(res Result) string {
	var lines []string
	for _, a := range res.File.Assets {
		lines = append(lines, describe(a))
	}
	return strings.Join(append(lines, res.Skipped...), "\n")
}

func TestRealScanWithCallgraph(t *testing.T) {
	res := adapt(t, fixture(t, "go-transitive.interim.json"), fixture(t, "go-transitive.callgraph.json"))
	want := strings.Join([]string{
		"main.go 18-18 algorithm hash SHA-256 project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=true depth=0",
		"main.go 22-22 algorithm hash MD5 project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=false depth=0",
		"libb.go 11-11 algorithm hash SHA-256 pkg:golang/example.com/translibb@v0.0.0 {Mode: Padding: KeySize:0 Curve: Hash:} reachable=true depth=2",
		"libb.go 17-17 algorithm hash SHA-512 pkg:golang/example.com/translibb@v0.0.0 {Mode: Padding: KeySize:0 Curve: Hash:} reachable=false depth=2",
	}, "\n")
	if got := describeAll(res); got != want {
		t.Fatalf("assets:\n%s\nwant:\n%s", got, want)
	}
	f := res.File
	if got := fmt.Sprint(f.Tool.Name, " ", f.Tool.Version, " ", f.Tool.Config, " ", f.ReachabilityReported); got !=
		"crypto-finder dev map[ruleset_checksum:120961b10cb5a6d359ceabf4f43552fd01c04d733e8fa26ecef1aba1815321e4 ruleset_source:local] true" {
		t.Fatalf("tool and flag: %s", got)
	}

	data, err := prediction.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	r := &check.Report{}
	if _, ok := prediction.Parse("out.json", data, r); !ok || len(r.Problems) > 0 {
		t.Fatalf("adapter output is not a valid prediction file: %v\n%s", r.Problems, data)
	}
}

// A remote ruleset is recorded by source, version and checksum. Its name is
// not written to the prediction file.
func TestRemoteRulesetConfigLeavesOutName(t *testing.T) {
	interim := strings.Replace(fixture(t, "go-transitive.interim.json"),
		`"source": "local",`, `"source": "remote", "name": "example-ruleset", "version": "v9.9.9",`, 1)
	res := adapt(t, interim, fixture(t, "go-transitive.callgraph.json"))
	if got, want := fmt.Sprint(res.File.Tool.Config),
		"map[ruleset_checksum:120961b10cb5a6d359ceabf4f43552fd01c04d733e8fa26ecef1aba1815321e4 ruleset_source:remote ruleset_version:v9.9.9]"; got != want {
		t.Fatalf("config: %s, want %s", got, want)
	}
}

// Format 1.8 adds a pqc block to an asset that names an algorithm. It must not
// change what is scored: the same scan as 1.7 gives the same assets.
func TestInterim18WithPQCScoresLikeInterim17(t *testing.T) {
	interim17 := fixture(t, "go-transitive.interim.json")
	graph := fixture(t, "go-transitive.callgraph.json")
	pqc := `"oid": "2.16.840.1.101.3.4.2.1",
          "pqc": {"status": "vulnerable", "family": "example-family", "deadline": 2030,
            "deadlines": [{"source": "example source", "deprecated_after": 2030, "disallowed_after": 2035, "migrate_by": 2032, "citation": "example citation"}],
            "dataset_version": "0.0.0"},`
	interim18 := strings.Replace(interim17, `"version": "1.7"`, `"version": "1.8"`, 1)
	interim18 = strings.Replace(interim18, `"oid": "2.16.840.1.101.3.4.2.1",`, pqc, 1)
	if !strings.Contains(interim18, `"pqc"`) {
		t.Fatal("fixture has no asset to attach pqc to")
	}
	want := describeAll(adapt(t, interim17, graph))
	if got := describeAll(adapt(t, interim18, graph)); got != want {
		t.Fatalf("1.8 report scored differently:\n%s\nwant:\n%s", got, want)
	}
}

func TestWithoutCallgraphNoReachabilityIsReported(t *testing.T) {
	res := adapt(t, fixture(t, "go-transitive.interim.json"), "")
	if res.File.ReachabilityReported {
		t.Fatal("reachability_reported is true without a call graph")
	}
	for _, a := range res.File.Assets {
		if a.Reachable != prediction.NoAnswer || a.Depth != nil {
			t.Fatalf("asset answers reachability without a call graph: %s", describe(a))
		}
	}
	if len(res.File.Assets) != 4 {
		t.Fatalf("got %d assets, want 4", len(res.File.Assets))
	}
}

func TestReachabilityStatesAndDepthSources(t *testing.T) {
	graph := fixture(t, "go-transitive.callgraph.json")
	cases := []struct{ name, old, new, want string }{
		{"not_applicable is no answer", `"reachability": "reachable",
      "reachable": true
    }`, `"reachability": "not_applicable"
    }`, "main.go 18-18 reachable=unknown depth=-"},
		{"relationship without path", `"relationship": "transitive",
        "path": [`, `"relationship": "direct",
        "unused": [`, "libb.go 11-11 reachable=true depth=1"},
		{"no dependency block", `"reachable": false,
      "dependency": {`, `"reachable": false,
      "other": {`, "libb.go 17-17 reachable=false depth=-"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !strings.Contains(graph, c.old) {
				t.Fatalf("fixture lacks %q", c.old)
			}
			res := adapt(t, fixture(t, "go-transitive.interim.json"), strings.Replace(graph, c.old, c.new, 1))
			var got []string
			for _, a := range res.File.Assets {
				d := "-"
				if a.Depth != nil {
					d = fmt.Sprint(*a.Depth)
				}
				got = append(got, fmt.Sprintf("%s %d-%d reachable=%s depth=%s", a.File, a.Lines.Start, a.Lines.End, a.Reachable, d))
			}
			if !strings.Contains(strings.Join(got, "\n"), c.want) {
				t.Fatalf("assets:\n%s\nwant a line %q", strings.Join(got, "\n"), c.want)
			}
		})
	}
}

func TestMetadataMapping(t *testing.T) {
	interim := fixture(t, "go-transitive.interim.json")
	sha256Project := `"algorithmFamily": "SHA-2",
            "algorithmName": "SHA-256",
            "algorithmPrimitive": "hash",
            "assetType": "algorithm"
          },
          "oid": "2.16.840.1.101.3.4.2.1",
          "finding_id": "7855dad3"`
	cases := []struct{ name, new, want string }{
		{"mode, padding and numeric key size", `"algorithmFamily": "AES", "algorithmPrimitive": "AE", "algorithmMode": "GCM", "algorithmPadding": "NoPadding",
            "algorithmParameterSetIdentifier": "256", "assetType": "Algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm ae AES project {Mode:GCM Padding:NoPadding KeySize:256 Curve: Hash:} reachable=unknown depth=-"},
		{"curve parameter set and unknown primitive", `"algorithmFamily": "EC", "algorithmName": "ECDSA", "algorithmPrimitive": "elliptic",
            "algorithmParameterSetIdentifier": "prime256v1", "assetType": "algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm  ECDSA project {Mode: Padding: KeySize:0 Curve:P-256 Hash:} reachable=unknown depth=-"},
		{"other parameter set is omitted", `"algorithmFamily": "ML-KEM", "algorithmName": "ML-KEM-768", "algorithmPrimitive": "kem",
            "algorithmParameterSetIdentifier": "ML-KEM-768", "assetType": "algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm kem ML-KEM-768 project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=-"},
		{"name over family", `"algorithmFamily": "AES", "algorithmName": "AES-256-GCM", "algorithmPrimitive": "ae", "assetType": "algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm ae AES-256-GCM project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=-"},
		{"unresolved name is kept for the scorer to treat as absent", `"algorithmFamily": "SHA-2", "algorithmName": "SHA-$bits", "algorithmPrimitive": "hash", "assetType": "algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm hash SHA-$bits project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=-"},
		{"related crypto material size and curve", `"assetType": "related-crypto-material", "materialType": "private-key", "materialSize": "384", "materialCurve": "secp384r1"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 related-crypto-material   project {Mode: Padding: KeySize:384 Curve:secp384r1 Hash:} reachable=unknown depth=- material_type=private-key"},
		{"material type is lowercased", `"assetType": "related-crypto-material", "materialType": "Initialization-Vector"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 related-crypto-material   project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=- material_type=initialization-vector"},
		{"material type outside CycloneDX is left out", `"assetType": "related-crypto-material", "materialType": "iv"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 related-crypto-material   project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=-\n"},
		{"material type only on related crypto material", `"algorithmName": "AES", "assetType": "algorithm", "materialType": "nonce"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm  AES project {Mode: Padding: KeySize:0 Curve: Hash:} reachable=unknown depth=-\n"},
		{"hash function", `"algorithmName": "ECDSA-SHA-256", "algorithmPrimitive": "signature", "algorithmHashFunction": "SHA-256", "assetType": "algorithm"
          },
          "finding_id": "7855dad3"`,
			"main.go 18-18 algorithm signature ECDSA-SHA-256 project {Mode: Padding: KeySize:0 Curve: Hash:SHA-256} reachable=unknown depth=-"},
		{"unknown asset type is skipped", `"assetType": "cipher-suite"
          },
          "finding_id": "7855dad3"`,
			`main.go:18: assetType "cipher-suite" is not a CycloneDX asset type`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !strings.Contains(interim, sha256Project) {
				t.Fatal("fixture changed")
			}
			res := adapt(t, strings.Replace(interim, sha256Project, c.new, 1), "")
			if got := describeAll(res); !strings.Contains(got, c.want) {
				t.Fatalf("assets:\n%s\nwant a line %q", got, c.want)
			}
		})
	}

	noPURL := strings.Replace(interim, `,
            "purl": "pkg:golang/example.com/translibb@v0.0.0"`, "", 1)
	if got := describeAll(adapt(t, noPURL, "")); !strings.Contains(got, "libb.go:11: dependency finding without dependency_info.purl") {
		t.Fatalf("dependency without purl not skipped:\n%s", got)
	}
}

func TestContractDriftIsRefused(t *testing.T) {
	interim, graph := fixture(t, "go-transitive.interim.json"), fixture(t, "go-transitive.callgraph.json")
	cases := []struct{ name, interim, graph, want string }{
		{"undeclared interim field", strings.Replace(interim, `"status": "pending"`, `"status": "pending", "confidence": "high"`, 1), graph,
			`cryptofinder: interim report: json: unknown field "confidence"`},
		{"interim version", strings.Replace(interim, `"version": "1.7"`, `"version": "2.0"`, 1), graph,
			`cryptofinder: interim report version "2.0" is not one of 1.5|1.6|1.7|1.8`},
		{"call graph without dependency blocks", interim, strings.Replace(graph, `"schema_version": "6.18"`, `"schema_version": "6.14"`, 1),
			`cryptofinder: call graph schema_version "6.14" is not one of 6.17|6.18`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := Adapt("e", []byte(c.interim), []byte(c.graph))
			if err == nil || err.Error() != c.want {
				t.Fatalf("err = %v, want %s", err, c.want)
			}
		})
	}
}

func TestDistanceIsTheDependencyPathLength(t *testing.T) {
	graph := fixture(t, "go-transitive.callgraph.json")
	cases := []struct{ name, graph, want string }{
		{"path of two", graph, "main.go:18 depth=0 distance=0; main.go:22 depth=0 distance=0; libb.go:11 depth=2 distance=2; libb.go:17 depth=2 distance=2"},
		{"relationship without path gives no distance", strings.Replace(graph, `"relationship": "transitive",
        "path": [`, `"relationship": "transitive",
        "unused": [`, 1), "libb.go:11 depth=2 distance=-"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := adapt(t, fixture(t, "go-transitive.interim.json"), c.graph)
			var got []string
			for _, a := range res.File.Assets {
				d, n := "-", "-"
				if a.Depth != nil {
					d = fmt.Sprint(*a.Depth)
				}
				if a.Distance != nil {
					n = fmt.Sprint(*a.Distance)
				}
				got = append(got, fmt.Sprintf("%s:%d depth=%s distance=%s", a.File, a.Lines.Start, d, n))
			}
			if !strings.Contains(strings.Join(got, "; "), c.want) {
				t.Fatalf("got %s, want %s", strings.Join(got, "; "), c.want)
			}
		})
	}
}
