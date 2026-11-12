package prediction

import (
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
)

const valid = `{"schema_version":"1","tool":{"name":"crypto-finder","version":"1.0","config":{"ruleset_version":"v1.0.0"}},
 "entry":"java-x","reachability_reported":true,
 "assets":[{"file":"./org/x/Cipher.java","lines":[12,14],"asset_type":"algorithm","primitive":"ae","algorithm":"AES",
            "parameters":{"mode":"GCM","key_size":256},"origin":"dependency","component":"pkg:maven/org.x/lib@1.0",
            "reachable":"true","depth":2},
           {"file":"org/x/Mac.java","lines":[20,20],"asset_type":"algorithm","primitive":"mac","algorithm":"HMAC",
            "parameters":{"hash":"SHA-256"},"origin":"project","reachable":"true","depth":0},
           {"file":"org/x/Iv.java","lines":[30,30],"asset_type":"related-crypto-material","material_type":"initialization-vector",
            "origin":"project","reachable":"true","depth":0}],
 "metadata":{"components_without_location":"1"}}`

func problems(r *check.Report) string {
	var out []string
	for _, p := range r.Problems {
		out = append(out, p.Path+": "+p.Msg)
	}
	return strings.Join(out, "\n")
}

func TestParseValidAndRoundTrip(t *testing.T) {
	r := &check.Report{}
	f, ok := Parse("p.json", []byte(valid), r)
	if !ok || len(r.Problems) > 0 {
		t.Fatalf("problems: %s", problems(r))
	}
	a := f.Assets[0]
	if a.File != "org/x/Cipher.java" || a.Origin != (groundtruth.Dependency{Component: "pkg:maven/org.x/lib@1.0"}) || a.Params.KeySize != 256 || *a.Depth != 2 || a.Reachable != Reachable {
		t.Fatalf("parsed asset: %+v", a)
	}
	data, err := Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	again, ok := Parse("p.json", data, r)
	if !ok || len(r.Problems) > 0 {
		t.Fatalf("re-parse problems: %s\n%s", problems(r), data)
	}
	if got, want := again.Assets[0].Lines, (groundtruth.Lines{Start: 12, End: 14}); got != want || again.Tool.Config["ruleset_version"] != "v1.0.0" || again.Metadata["components_without_location"] != "1" ||
		again.Assets[1].Params.Hash != "SHA-256" || again.Assets[2].MaterialType != "initialization-vector" {
		t.Fatalf("round trip lost data:\n%s", data)
	}
}

func TestEachRuleDetectsItsDefect(t *testing.T) {
	cases := []struct{ name, old, new, want string }{
		{"unknown field", `"entry":"java-x"`, `"entry":"java-x","colour":"red"`, `: json: unknown field "colour"`},
		{"trailing document", `"1"}}`, `"1"}}{}`, ": trailing data after the JSON document"},
		{"schema version", `"schema_version":"1"`, `"schema_version":"2"`, `schema_version: got "2", want "1"`},
		{"missing reachability_reported", `"reachability_reported":true,`, ``, "reachability_reported: required"},
		{"reachable without reachability", `"reachability_reported":true`, `"reachability_reported":false`, "assets[0].reachable: must be unknown when reachability_reported is false\nassets[1].reachable: must be unknown when reachability_reported is false\nassets[2].reachable: must be unknown when reachability_reported is false"},
		{"depth with unknown", `"reachable":"true"`, `"reachable":"unknown"`, "assets[0].depth: not allowed when reachable is unknown"},
		{"depth out of range", `"depth":2`, `"depth":3`, "assets[0].depth: must be 0, 1 or 2 (2 means two or more hops), got 3"},
		{"distance disagrees with depth", `"depth":2`, `"depth":2,"distance":1`, "assets[0].distance: 1 disagrees with depth 2 (depth is distance capped at 2)"},
		{"distance without depth", `"depth":2`, `"distance":3`, "assets[0].distance: only allowed together with depth"},
		{"bad reachable", `"reachable":"true"`, `"reachable":"yes"`, `assets[0].reachable: "yes" is not one of true|false|unknown`},
		{"bad asset type", `"asset_type":"algorithm"`, `"asset_type":"cipher"`, `assets[0].asset_type: "cipher" is not one of algorithm|certificate|protocol|related-crypto-material` + "\nassets[0].primitive: only allowed when asset_type is algorithm"},
		{"bad material type", `"material_type":"initialization-vector"`, `"material_type":"iv"`, `assets[2].material_type: "iv" is not one of private-key|public-key|secret-key|key|ciphertext|signature|digest|initialization-vector|nonce|seed|salt|shared-secret|tag|additional-data|password|credential|token|other|unknown`},
		{"material type on an algorithm", `"primitive":"mac",`, `"primitive":"mac","material_type":"tag",`, "assets[1].material_type: only allowed when asset_type is related-crypto-material"},
		{"component on project", `"origin":"dependency"`, `"origin":"project"`, "assets[0].component: only allowed when origin is dependency"},
		{"absolute file", `"./org/x/Cipher.java"`, `"/org/x/Cipher.java"`, `assets[0].file: "/org/x/Cipher.java" must be relative to the entry or component root`},
		{"inverted lines", `[12,14]`, `[14,12]`, "assets[0].lines: need 1 <= start <= end, got [14 12]"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !strings.Contains(valid, c.old) {
				t.Fatalf("fixture lacks %q", c.old)
			}
			r := &check.Report{}
			Parse("p.json", []byte(strings.Replace(valid, c.old, c.new, 1)), r)
			if got := problems(r); got != c.want {
				t.Fatalf("problems:\n%s\nwant:\n%s", got, c.want)
			}
		})
	}
}
