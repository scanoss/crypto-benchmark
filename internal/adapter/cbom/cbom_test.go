package cbom

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// Fixtures, trimmed to the components and fields the adapter reads:
//
//   - cdxgen-js.cbom.json: cdxgen 13.6.0 `cbom` on crypto-finder's
//     testdata/projects/polyglot_reachability.
//   - cdxgen-go.cbom.json and cdxgen-go-all.cbom.json: cdxgen 13.6.0
//     `cdxgen -t go` then `evinse -l go` on crypto-finder's
//     testdata/projects/transitive_reachability/go/app, the second with
//     --golem-patterns all. Absolute paths are rebased onto /work and
//     /usr/local/go.
//   - sonar-keycloak.cbom.json: sonar-cryptography's CBOM of Keycloak as
//     published at
//     https://github.com/cbomkit/cbomkit/blob/eea204303f53bd13192cc9770a917c938238f6ed/example/keycloak-cbom.json
//     (nine of its 56 components, at most two occurrences each).
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
	s := fmt.Sprintf("%s %d %s %s %s %s %+v", a.File, a.Lines.Start, a.Type, a.Primitive, a.Algorithm, origin, a.Params)
	if a.MaterialType != "" {
		s += " material_type=" + string(a.MaterialType)
	}
	return s
}

func adapt(t *testing.T, bom string, opts Options) Result {
	t.Helper()
	opts.Tool, opts.ToolVersion, opts.Entry = "cdxgen", "13.6.0", "e"
	res, err := Adapt([]byte(bom), opts)
	if err != nil {
		t.Fatal(err)
	}
	data, err := prediction.Marshal(res.File)
	if err != nil {
		t.Fatal(err)
	}
	r := &check.Report{}
	if _, ok := prediction.Parse("out.json", data, r); !ok || len(r.Problems) > 0 {
		t.Fatalf("adapter output is not a valid prediction file: %v\n%s", r.Problems, data)
	}
	return res
}

func describeAll(res Result) string {
	var lines []string
	for _, a := range res.File.Assets {
		lines = append(lines, describe(a))
	}
	m := res.File.Metadata
	lines = append(lines, fmt.Sprintf("without location %s, dropped %s", m["components_without_location"], m["locations_dropped"]))
	return strings.Join(append(lines, res.Notes...), "\n")
}

func expect(t *testing.T, res Result, want ...string) {
	t.Helper()
	if got := describeAll(res); got != strings.Join(want, "\n") {
		t.Fatalf("got:\n%s\nwant:\n%s", got, strings.Join(want, "\n"))
	}
}

func TestCdxgenJavaScriptOccurrences(t *testing.T) {
	res := adapt(t, fixture(t, "cdxgen-js.cbom.json"), Options{})
	expect(t, res,
		"web/src/lib/avatarHash.ts 4 algorithm hash md5 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"web/src/lib/avatarHash.ts 12 algorithm hash sha-1 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"web/src/lib/etag.ts 4 algorithm hash sha-256 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"without location 0, dropped 0")
	if got, want := describeAll(adapt(t, fixture(t, "cdxgen-js.cbom.json"), Options{Cdxgen: true})), describeAll(res); got != want {
		t.Fatalf("with --cdxgen, occurrences and location properties both counted:\n%s", got)
	}
	f := res.File
	if got := fmt.Sprint(f.Tool.Name, " ", f.Tool.Version, " ", f.Tool.Config, " ", f.Entry, " ", f.ReachabilityReported); got != "cdxgen 13.6.0 map[] e false" {
		t.Fatalf("tool and flag: %s", got)
	}
	for _, a := range f.Assets {
		if a.Reachable != prediction.NoAnswer || a.Depth != nil {
			t.Fatalf("asset answers reachability: %s", describe(a))
		}
	}
}

func TestCdxgenGoLocationsNeedTheFlagAndARoot(t *testing.T) {
	bom := fixture(t, "cdxgen-go.cbom.json")
	expect(t, adapt(t, bom, Options{}),
		"without location 2, dropped 0",
		"md5 (crypto/algorithm/md5@1.2.840.113549.2.5): no occurrence location; it cannot be matched",
		"sha-256 (crypto/algorithm/sha-256@2.16.840.1.101.3.4.2.1): no occurrence location; it cannot be matched")
	expect(t, adapt(t, bom, Options{Cdxgen: true}),
		"without location 2, dropped 2",
		"md5 (crypto/algorithm/md5@1.2.840.113549.2.5): /work/go/app/main.go:22 is absolute and no source root was given",
		"sha-256 (crypto/algorithm/sha-256@2.16.840.1.101.3.4.2.1): /work/go/app/main.go:18 is absolute and no source root was given")
	res := adapt(t, bom, Options{Cdxgen: true, SourceRoot: "/work/go/app"})
	expect(t, res,
		"main.go 22 algorithm hash md5 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"main.go 18 algorithm hash sha-256 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"without location 0, dropped 0")
	if res.File.Tool.Config["cdxgen_locations"] != "true" {
		t.Fatalf("config does not record the flag: %v", res.File.Tool.Config)
	}
}

func TestCdxgenGoDependencySourceRoots(t *testing.T) {
	bom := fixture(t, "cdxgen-go-all.cbom.json")
	expect(t, adapt(t, bom, Options{Cdxgen: true, SourceRoot: "/work/go/app"}),
		"main.go 22 algorithm hash md5 project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"libb.go 11 algorithm hash sha-256 pkg:golang/example.com/translibb {Mode: Padding: KeySize:0 Curve: Hash:}",
		"libb.go 17 algorithm hash sha-512 pkg:golang/example.com/translibb {Mode: Padding: KeySize:0 Curve: Hash:}",
		"without location 2, dropped 4",
		"key (crypto/material/key@key): /usr/local/go/src/reflect/map_swiss.go:163 is outside the source root and no dependency's source root",
		"key (crypto/material/key@key): /usr/local/go/src/reflect/map_swiss.go:295 is outside the source root and no dependency's source root",
		"key (crypto/material/key@key): /usr/local/go/src/reflect/map_swiss.go:431 is outside the source root and no dependency's source root",
		"vdsoSymbolKeys (crypto/material/vdsoSymbolKeys@key): /usr/local/go/src/runtime/vdso_linux_amd64.go:17 is outside the source root and no dependency's source root")

	noLocalDir := strings.Replace(bom, `"name": "cdx:go:local_dir",
          "value": "../libb"`, `"name": "cdx:go:other",
          "value": "../libb"`, 1)
	if noLocalDir == bom {
		t.Fatal("fixture changed")
	}
	const occurrence = `{"bomFormat":"CycloneDX","specVersion":"1.7","components":[
		{"type":"library","bom-ref":"b","name":"b","purl":"pkg:golang/example.com/translibb","properties":[
			{"name":"cdx:go:local_dir","value":"../libb"},{"name":"internal:SrcGoMod","value":"/work/go/libb/go.mod"}]},
		{"type":"cryptographic-asset","bom-ref":"c","name":"sha-512","cryptoProperties":{"assetType":"algorithm"},
			"evidence":{"occurrences":[{"location":"/work/go/libb/libb.go","line":17}]}}]}`
	if got := describeAll(adapt(t, occurrence, Options{SourceRoot: "/work/go/app"})); !strings.Contains(got, "/work/go/libb/libb.go:17 is outside the source root") {
		t.Fatalf("a local Go module claimed a source root without --cdxgen:\n%s", got)
	}
	if got := describeAll(adapt(t, occurrence, Options{SourceRoot: "/work/go/app", Cdxgen: true})); !strings.HasPrefix(got, "libb.go 17 algorithm  sha-512 pkg:golang/example.com/translibb") {
		t.Fatalf("a local Go module did not claim its source root with --cdxgen:\n%s", got)
	}
	got := describeAll(adapt(t, noLocalDir, Options{Cdxgen: true, SourceRoot: "/work/go/app"}))
	if !strings.Contains(got, "sha-512 (crypto/algorithm/sha-512@2.16.840.1.101.3.4.2.3): /work/go/libb/libb.go:17 is outside the source root") {
		t.Fatalf("a module that is not a local replacement still claimed a source root:\n%s", got)
	}
}

func TestSonarCryptographyCBOM(t *testing.T) {
	expect(t, adapt(t, fixture(t, "sonar-keycloak.cbom.json"), Options{}),
		"crypto/default/src/main/java/org/keycloak/crypto/def/BCEcdhEsAlgorithmProvider.java 208 algorithm key-agree ECDH project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"crypto/elytron/src/main/java/org/keycloak/crypto/elytron/ElytronEcdhEsAlgorithmProvider.java 183 algorithm key-agree ECDH project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"testsuite/integration-arquillian/servers/auth-server/services/testsuite-providers/src/main/java/org/keycloak/testsuite/rest/resource/TestingOIDCEndpointsApplicationResource.java 152 algorithm pke EC-secp521r1 project {Mode: Padding: KeySize:0 Curve:secp521r1 Hash:}",
		"core/src/main/java/org/keycloak/jose/jwk/JWKParser.java 136 algorithm pke RSA-2048 project {Mode: Padding: KeySize:2048 Curve: Hash:}",
		"saml-core-api/src/main/java/org/keycloak/dom/xmlsec/w3/xmldsig/RSAKeyValueType.java 103 algorithm pke RSA-2048 project {Mode: Padding: KeySize:2048 Curve: Hash:}",
		"testsuite/integration-arquillian/servers/auth-server/services/testsuite-providers/src/main/java/org/keycloak/testsuite/rest/resource/TestingOIDCEndpointsApplicationResource.java 159 algorithm signature Ed25519 project {Mode: Padding: KeySize:0 Curve:Edwards25519 Hash:}",
		"quarkus/runtime/src/main/java/org/keycloak/quarkus/runtime/storage/infinispan/jgroups/impl/JpaJGroupsTlsConfigurator.java 80 protocol  TLS project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"quarkus/runtime/src/main/java/org/keycloak/quarkus/runtime/storage/infinispan/CacheManagerFactory.java 370 protocol  TLS project {Mode: Padding: KeySize:0 Curve: Hash:}",
		"crypto/elytron/src/main/java/org/keycloak/crypto/elytron/WildFlyElytronProvider.java 143 algorithm block-cipher AES128-CBC-PKCS5 project {Mode:cbc Padding:pkcs5 KeySize:128 Curve: Hash:}",
		"crypto/elytron/src/main/java/org/keycloak/crypto/elytron/WildFlyElytronProvider.java 148 algorithm ae AES128-GCM project {Mode:gcm Padding: KeySize:128 Curve: Hash:}",
		"server-spi-private/src/main/java/org/keycloak/broker/provider/AbstractIdentityProvider.java 118 algorithm hash SHA256 project {Mode: Padding: KeySize:256 Curve: Hash:}",
		"server-spi-private/src/main/java/org/keycloak/models/utils/KeycloakModelUtils.java 213 algorithm hash SHA256 project {Mode: Padding: KeySize:256 Curve: Hash:}",
		"saml-core-api/src/main/java/org/keycloak/dom/xmlsec/w3/xmldsig/DSAKeyValueType.java 234 related-crypto-material  key@b627000e-ed4e-449c-acb9-4e9547d6ee93 project {Mode: Padding: KeySize:0 Curve: Hash:} material_type=secret-key",
		"without location 0, dropped 0")
}

func TestOriginRules(t *testing.T) {
	const head = `{"bomFormat":"CycloneDX","specVersion":"1.7","metadata":{"component":{"type":"application","bom-ref":"app","name":"app"}},"components":[`
	const libs = `{"type":"library","bom-ref":"liba","name":"liba","purl":"pkg:npm/liba@1.0.0"},
		{"type":"library","bom-ref":"scoped","name":"x","purl":"pkg:npm/%40acme/x@2.0.0"},`
	asset := func(ref, loc string) string {
		return `{"type":"cryptographic-asset","bom-ref":"` + ref + `","name":"sha-256","cryptoProperties":{"assetType":"algorithm"},
			"evidence":{"occurrences":[{"location":"` + loc + `","line":3}]}}`
	}
	cases := []struct{ name, bom, want string }{
		{"project by default", head + libs + asset("c", "src/a.js") + `]}`, "src/a.js 3 project"},
		{"npm package path", head + libs + asset("c", "node_modules/liba/lib/h.js") + `]}`, "lib/h.js 3 pkg:npm/liba@1.0.0"},
		{"nested npm package path", head + libs + asset("c", "node_modules/q/node_modules/@acme/x/i.js") + `]}`, "i.js 3 pkg:npm/%40acme/x@2.0.0"},
		{"package path of an unlisted package", head + libs + asset("c", "node_modules/other/i.js") + `]}`, "node_modules/other/i.js 3 project"},
		{"nested under a library", head + libs + `{"type":"library","bom-ref":"n","name":"n","purl":"pkg:maven/g/n@1","components":[` + asset("c", "N.java") + `]}]}`, "N.java 3 pkg:maven/g/n@1"},
		{"provided by one library", head + libs + asset("c", "h.js") + `],"dependencies":[{"ref":"liba","provides":["c"]}]}`, "h.js 3 pkg:npm/liba@1.0.0"},
		{"depended on by one library", head + libs + asset("c", "h.js") + `],"dependencies":[{"ref":"liba","dependsOn":["c"]}]}`, "h.js 3 pkg:npm/liba@1.0.0"},
		{"provided by a library and the project", head + libs + asset("c", "h.js") + `],"dependencies":[{"ref":"liba","provides":["c"]},{"ref":"app","provides":["c"]}]}`, "h.js 3 project"},
		{"provided by two libraries", head + libs + asset("c", "h.js") + `],"dependencies":[{"ref":"liba","provides":["c"]},{"ref":"scoped","provides":["c"]}]}`, "h.js 3 project"},
		{"crypto assets linking each other", head + libs + asset("c", "h.js") + `,` + asset("d", "k.js") + `],"dependencies":[{"ref":"d","dependsOn":["c"]}]}`, "h.js 3 project\nk.js 3 project"},
		{"package path of a package listed twice", head + libs + `{"type":"library","bom-ref":"liba2","name":"liba","purl":"pkg:npm/liba@2.0.0"},` + asset("c", "node_modules/liba/h.js") + `]}`, "node_modules/liba/h.js 3 project"},
		{"path wins over a link", head + libs + asset("c", "node_modules/@acme/x/i.js") + `],"dependencies":[{"ref":"liba","provides":["c"]}]}`, "i.js 3 pkg:npm/%40acme/x@2.0.0"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got []string
			for _, a := range adapt(t, c.bom, Options{}).File.Assets {
				origin := "project"
				if d, ok := a.Origin.(groundtruth.Dependency); ok {
					origin = d.Component
				}
				got = append(got, fmt.Sprintf("%s %d %s", a.File, a.Lines.Start, origin))
			}
			if g := strings.Join(got, "\n"); g != c.want {
				t.Fatalf("got %q, want %q", g, c.want)
			}
		})
	}
}

func TestParametersAndUnmatchableLocations(t *testing.T) {
	const head = `{"bomFormat":"CycloneDX","specVersion":"1.6","components":[`
	cases := []struct{ name, component, want string }{
		{"1.7 ellipticCurve drops its namespace", `{"type":"cryptographic-asset","bom-ref":"r","name":"ECDSA","cryptoProperties":{"assetType":"algorithm",
			"algorithmProperties":{"primitive":"signature","ellipticCurve":"nist/P-256","curve":"secp256r1"}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 algorithm signature ECDSA project {Mode: Padding: KeySize:0 Curve:P-256 Hash:}"},
		{"curve parameter set", `{"type":"cryptographic-asset","bom-ref":"r","name":"EC","cryptoProperties":{"assetType":"Algorithm",
			"algorithmProperties":{"primitive":"elliptic","parameterSetIdentifier":"secp384r1"}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 algorithm  EC project {Mode: Padding: KeySize:0 Curve:P-384 Hash:}"},
		{"other parameter set is left out", `{"type":"cryptographic-asset","bom-ref":"r","name":"ML-KEM-768","cryptoProperties":{"assetType":"algorithm",
			"algorithmProperties":{"primitive":"kem","parameterSetIdentifier":"ML-KEM-768"}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 algorithm kem ML-KEM-768 project {Mode: Padding: KeySize:0 Curve: Hash:}"},
		{"related crypto material size", `{"type":"cryptographic-asset","bom-ref":"r","name":"key","cryptoProperties":{"assetType":"related-crypto-material",
			"relatedCryptoMaterialProperties":{"type":"secret-key","size":256}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 related-crypto-material  key project {Mode: Padding: KeySize:256 Curve: Hash:} material_type=secret-key"},
		{"material type is lowercased", `{"type":"cryptographic-asset","bom-ref":"r","name":"iv","cryptoProperties":{"assetType":"related-crypto-material",
			"relatedCryptoMaterialProperties":{"type":"Initialization-Vector"}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 related-crypto-material  iv project {Mode: Padding: KeySize:0 Curve: Hash:} material_type=initialization-vector"},
		{"material type outside CycloneDX is left out", `{"type":"cryptographic-asset","bom-ref":"r","name":"iv","cryptoProperties":{"assetType":"related-crypto-material",
			"relatedCryptoMaterialProperties":{"type":"iv"}},"evidence":{"occurrences":[{"location":"a.go","line":1}]}}`,
			"a.go 1 related-crypto-material  iv project {Mode: Padding: KeySize:0 Curve: Hash:}"},
		{"unknown asset type", `{"type":"cryptographic-asset","bom-ref":"r","name":"x","cryptoProperties":{"assetType":"cipher-suite"}}`,
			"without location 0, dropped 0\nx (r): assetType \"cipher-suite\" is not a CycloneDX asset type"},
		{"occurrence without line", `{"type":"cryptographic-asset","bom-ref":"r","name":"x","cryptoProperties":{"assetType":"algorithm"},"evidence":{"occurrences":[{"location":"a.go"}]}}`,
			"without location 1, dropped 1\nx (r): a.go:0 has no line"},
		{"relative path leaving the root", `{"type":"cryptographic-asset","bom-ref":"r","name":"x","cryptoProperties":{"assetType":"algorithm"},"evidence":{"occurrences":[{"location":"../b.go","line":2},{"location":"c.go","line":4}]}}`,
			"c.go 4 algorithm  x project {Mode: Padding: KeySize:0 Curve: Hash:}\nwithout location 0, dropped 1\nx (r): ../b.go:2 is outside the source root and no dependency's source root"},
		{"location properties are ignored without the flag", `{"type":"cryptographic-asset","bom-ref":"r","name":"x","cryptoProperties":{"assetType":"algorithm"},"properties":[{"name":"cdx:crypto:sourceLocation","value":"a.js:1:2"}]}`,
			"without location 1, dropped 0\nx (r): no occurrence location; it cannot be matched"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := describeAll(adapt(t, head+c.component+`]}`, Options{}))
			got = strings.TrimSuffix(got, "\nwithout location 0, dropped 0")
			if got != c.want {
				t.Fatalf("got:\n%s\nwant:\n%s", got, c.want)
			}
		})
	}
}

func TestParseCdxgenLocation(t *testing.T) {
	cases := map[string]string{
		"web/src/lib/etag.ts:4:9":   "web/src/lib/etag.ts 4",
		"/work/go/app/main.go:22:9": "/work/go/app/main.go 22",
		"src/main.rs#4":             "src/main.rs 4",
		"a.py:7":                    "a.py 7",
		"a.py":                      "invalid",
		"a.py:x:9":                  "a.py:x 9",
		"#4":                        "invalid",
	}
	for in, want := range cases {
		got := "invalid"
		if l, ok := parseCdxgenLocation(in); ok {
			got = fmt.Sprintf("%s %d", l.path, l.line)
		}
		if got != want {
			t.Errorf("parseCdxgenLocation(%q) = %s, want %s", in, got, want)
		}
	}
}

func TestRefusesOtherDocuments(t *testing.T) {
	cases := []struct{ bom, want string }{
		{`{"bomFormat":"CycloneDX","specVersion":"1.5","components":[]}`, `cbom: specVersion "1.5" is not one of 1.6|1.7`},
		{`{"spdxVersion":"SPDX-2.3"}`, `cbom: bomFormat is "", want CycloneDX`},
		{`[]`, `cbom: json: cannot unmarshal array into Go value of type cbom.bom`},
	}
	for _, c := range cases {
		if _, err := Adapt([]byte(c.bom), Options{Tool: "t", Entry: "e"}); err == nil || err.Error() != c.want {
			t.Errorf("err = %v, want %s", err, c.want)
		}
	}
}
