package algo

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
)

// show renders a Spec as one literal: algorithm, then each set parameter.
func show(s Spec) string {
	out := []string{"alg=" + s.Algorithm}
	p := s.Params
	if p.KeySize != 0 {
		out = append(out, fmt.Sprintf("key_size=%d", p.KeySize))
	}
	for _, f := range [][2]string{{"mode", p.Mode}, {"padding", p.Padding}, {"curve", p.Curve}, {"hash", p.Hash}} {
		if f[1] != "" {
			out = append(out, f[0]+"="+f[1])
		}
	}
	return strings.Join(out, " ")
}

func TestCanonicalizeName(t *testing.T) {
	cases := []struct{ name, want string }{
		{"SHA256", "alg=SHA-256"},
		{"SHA-256", "alg=SHA-256"},
		{"sha_256", "alg=SHA-256"},
		{"SHA-512", "alg=SHA-512"},
		{"md5", "alg=MD5"},
		{"3DES", "alg=3DES"},
		{"DESede", "alg=3DES"},
		{"TripleDES", "alg=3DES"},
		{"Curve25519", "alg=Curve25519"},
		{"X25519", "alg=X25519"},
		{"chacha20-poly1305", "alg=ChaCha20-Poly1305"},
		{"ML-KEM-768", "alg=ML-KEM-768"},

		{"AES-256-GCM", "alg=AES key_size=256 mode=GCM"},
		{"AES-GCM", "alg=AES mode=GCM"},
		{"AES-128", "alg=AES key_size=128"},
		{"aes_256_cbc", "alg=AES key_size=256 mode=CBC"},
		{"AES256", "alg=AES key_size=256"},
		{"AES-GCM-SIV", "alg=AES mode=GCM-SIV"},
		{"AES-KWP", "alg=AES mode=KWP"},
		{"DESede-CBC", "alg=3DES mode=CBC"},
		{"3DES-168-CBC", "alg=3DES key_size=168 mode=CBC"},
		{"Camellia-128-CTR", "alg=Camellia key_size=128 mode=CTR"},
		{"SM4-ECB", "alg=SM4 mode=ECB"},
		{"AES-CFB8", "alg=AES mode=CFB"},
		{"AES-128-CFB1", "alg=AES key_size=128 mode=CFB"},
		{"aes-256-ofb128", "alg=AES key_size=256 mode=OFB"},
		{"DES-CFB64", "alg=DES mode=CFB"},
		{"AES-CFB7", "alg=AESCFB7"},
		{"AES-CTR8", "alg=AESCTR8"},
		{"Blowfish-448-CBC", "alg=Blowfish key_size=448 mode=CBC"},

		{"AES128-CBC-PKCS5", "alg=AES key_size=128 mode=CBC padding=PKCS7"},
		{"AES-128-CBC-PKCS5", "alg=AES key_size=128 mode=CBC padding=PKCS7"},
		{"AES256-GCM", "alg=AES key_size=256 mode=GCM"},
		{"AES192-CTR-NoPadding", "alg=AES key_size=192 mode=CTR padding=none"},
		{"AES-CBC-PKCS7", "alg=AES mode=CBC padding=PKCS7"},
		{"AES256-CFB8-PKCS5", "alg=AES key_size=256 mode=CFB padding=PKCS7"},
		{"AES256-GCM-SIV", "alg=AES key_size=256 mode=GCM-SIV"},
		{"DESede168-CBC-PKCS5", "alg=3DES key_size=168 mode=CBC padding=PKCS7"},
		{"Camellia256-CBC-PKCS5", "alg=Camellia key_size=256 mode=CBC padding=PKCS7"},
		{"SM4128-CBC-PKCS5", "alg=SM4 key_size=128 mode=CBC padding=PKCS7"},
		{"AES-512", "alg=AES512"},
		{"AES512-GCM", "alg=AES512GCM"},
		{"SM4-256", "alg=SM4256"},
		{"DES-128-CBC", "alg=DES128CBC"},
		{"AES128-CBC-ISO10126", "alg=AES128CBCISO10126"},
		{"AES128-PKCS5", "alg=AES128PKCS5"},
		{"AES128-CBC-PKCS5-GCM", "alg=AES128CBCPKCS5GCM"},

		{"RSA-OAEP", "alg=RSA padding=OAEP"},
		{"RSAES-OAEP", "alg=RSA padding=OAEP"},
		{"RSA-PSS", "alg=RSA padding=PSS"},
		{"RSASSA-PSS", "alg=RSA padding=PSS"},
		{"rsassa_pss", "alg=RSA padding=PSS"},
		{"RSA-PKCS1", "alg=RSA padding=PKCS1v15"},
		{"RSASSA-PKCS1-v1_5", "alg=RSA padding=PKCS1v15"},
		{"RSAES-PKCS1-v1_5", "alg=RSA padding=PKCS1v15"},
		{"RSA-2048", "alg=RSA key_size=2048"},
		{"RSA4096", "alg=RSA key_size=4096"},
		{"RSA-OAEP-3072", "alg=RSA key_size=3072 padding=OAEP"},
		{"RSASSA-PSS-4096", "alg=RSA key_size=4096 padding=PSS"},
		{"RSA-PKCS1-2048", "alg=RSA key_size=2048 padding=PKCS1v15"},
		{"RSA-0", "alg=RSA0"},
		{"RSA-KEM", "alg=RSAKEM"},
		{"RSA/ECB/PSSPadding", "alg=RSA mode=ECB padding=PSS"},

		{"AES/GCM/NoPadding", "alg=AES mode=GCM padding=none"},
		{"AES/CBC/PKCS5Padding", "alg=AES mode=CBC padding=PKCS7"},
		{"AES/CBC/PKCS7Padding", "alg=AES mode=CBC padding=PKCS7"},
		{"aes/ctr/nopadding", "alg=AES mode=CTR padding=none"},
		{"AES/CFB8/NoPadding", "alg=AES mode=CFB padding=none"},
		{"AES_256/GCM/NoPadding", "alg=AES key_size=256 mode=GCM padding=none"},
		{"DESede/CBC/PKCS5Padding", "alg=3DES mode=CBC padding=PKCS7"},
		{"RSA/ECB/PKCS1Padding", "alg=RSA mode=ECB padding=PKCS1v15"},
		{"RSA/ECB/OAEPWithSHA-256AndMGF1Padding", "alg=RSA mode=ECB padding=OAEP"},
		{"RSA/ECB/OAEPWithSHA-1AndMGF1Padding", "alg=RSA mode=ECB padding=OAEP"},
		{"RSA/ECB/OAEPPadding", "alg=RSA mode=ECB padding=OAEP"},
		{"RSA/NONE/OAEPWithSHA-512/256AndMGF1Padding", "alg=RSA mode=NONE padding=OAEP"},
		{"RSA/NONE/NoPadding", "alg=RSA mode=NONE padding=none"},
		{"Blowfish/CBC/ISO10126Padding", "alg=Blowfish mode=CBC padding=ISO10126PADDING"},
		{"SHA-512/256", "alg=SHA-512/256"},
		{"AES/GCM", "alg=AES/GCM"},

		{"ECDSA-secp256k1", "alg=ECDSA curve=secp256k1"},
		{"ECDSA-P-256", "alg=ECDSA curve=P-256"},
		{"ECDH-prime256v1", "alg=ECDH curve=P-256"},
		{"ECIES-secp384r1", "alg=ECIES curve=P-384"},
		{"EC-brainpoolP256r1", "alg=EC curve=brainpoolP256r1"},
		{"ECDSA", "alg=ECDSA"},

		{"HMAC-SHA-256", "alg=HMAC hash=SHA-256"},
		{"HmacSHA256", "alg=HMAC hash=SHA-256"},
		{"HMAC-SHA3-256", "alg=HMAC hash=SHA3-256"},
		{"hmac_sha1", "alg=HMAC hash=SHA-1"},
		{"HMAC-SHA-512/256", "alg=HMAC hash=SHA-512/256"},
		{"HMAC", "alg=HMAC"},
		{"HMACSHA2", "alg=HMACSHA2"},
		{"HMAC-SHAKE128", "alg=HMACSHAKE128"},
		{"AES-CMAC", "alg=AESCMAC"},
		{"AES-0", "alg=AES0"},
		{"ECDHE", "alg=ECDHE"},

		{"ECDSA-SHA256", "alg=ECDSA hash=SHA-256"},
		{"ECDSA-SHA-384", "alg=ECDSA hash=SHA-384"},
		{"DSA-SHA1", "alg=DSA hash=SHA-1"},
		{"RSA-SHA256", "alg=RSA hash=SHA-256"},
		{"RSA-PSS-SHA256", "alg=RSA padding=PSS hash=SHA-256"},
		{"RSASSA-PSS-SHA-512", "alg=RSA padding=PSS hash=SHA-512"},
		{"RSA-PKCS1-SHA256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RSA-PKCS1v15-SHA-256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RSASSA-PKCS1-SHA-256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RSASSA-PKCS1-v1_5-SHA-256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RSA-PKCS1v15", "alg=RSA padding=PKCS1v15"},
		{"ECDSA-SHA2", "alg=ECDSASHA2"},
		{"ECDSA-P-256-SHA256", "alg=ECDSAP256SHA256"},
		{"Ed25519-SHA512", "alg=ED25519SHA512"},
		{"RSA-OAEP-SHA2", "alg=RSAOAEPSHA2"},
		{"PBKDF2-SHA2", "alg=PBKDF2SHA2"},
		{"PBKDF2-HMAC-MD", "alg=PBKDF2HMACMD"},
		{"PBKDF2WithHmac", "alg=PBKDF2"},
		{"PBKDF2-HMAC", "alg=PBKDF2"},
		{"HKDF-SHA256", "alg=HKDFSHA256"},

		{"RSA-OAEP-SHA256", "alg=RSA padding=OAEP hash=SHA-256"},
		{"RSA-OAEP-SHA-1", "alg=RSA padding=OAEP hash=SHA-1"},
		{"RSAES-OAEP-SHA-512", "alg=RSA padding=OAEP hash=SHA-512"},
		{"PBKDF2-SHA256", "alg=PBKDF2 hash=SHA-256"},
		{"PBKDF2-SHA-512", "alg=PBKDF2 hash=SHA-512"},
		{"PBKDF2-HMAC-SHA256", "alg=PBKDF2 hash=SHA-256"},
		{"pbkdf2_hmac_sha1", "alg=PBKDF2 hash=SHA-1"},
		{"PBKDF2WithHmacSHA256", "alg=PBKDF2 hash=SHA-256"},
		{"PBKDF2WithHmacSHA3-512", "alg=PBKDF2 hash=SHA3-512"},

		{"SHA256withRSA", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"SHA384withECDSA", "alg=ECDSA hash=SHA-384"},
		{"SHA1withDSA", "alg=DSA hash=SHA-1"},
		{"SHA256withRSA/PSS", "alg=RSA padding=PSS hash=SHA-256"},
		{"SHA512withRSAandMGF1", "alg=RSA padding=PSS hash=SHA-512"},
		{"SHA3-256withECDSA", "alg=ECDSA hash=SHA3-256"},
		{"SHA256withEd25519", "alg=SHA256WITHED25519"},
		{"SHA2withRSA", "alg=SHA2WITHRSA"},
		{"OAEPWithSHA-256AndMGF1Padding", "alg=OAEPWITHSHA256ANDMGF1PADDING"},
		{"Some Vendor Cipher", "alg=SOMEVENDORCIPHER"},

		{"$ALG", "alg="},
		{"SHA-$bits", "alg="},
		{"AES-$size-GCM", "alg="},
		{"", "alg="},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := Canonicalize(c.name, groundtruth.Parameters{})
			if got := show(s); got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
			if len(s.Conflicts) > 0 {
				t.Fatalf("unexpected conflicts %v", s.Conflicts)
			}
		})
	}
}

// TestCanonicalizeUpstreamSonarNames pins every algorithm name in the
// Keycloak CBOM that CBOMkit publishes as its sonar-cryptography example,
// https://github.com/cbomkit/cbomkit/blob/eea204303f53bd13192cc9770a917c938238f6ed/example/keycloak-cbom.json.
// A name kept whole is kept whole on purpose: HMACSHA2 names no digest
// size, EdDSA no curve, and MGF1, RAW and ConcatenationKDF have no other
// spelling in the tables.
func TestCanonicalizeUpstreamSonarNames(t *testing.T) {
	cases := []struct{ name, want string }{
		{"AES", "alg=AES"},
		{"AES128", "alg=AES key_size=128"},
		{"AES128-CBC-PKCS5", "alg=AES key_size=128 mode=CBC padding=PKCS7"},
		{"AES128-GCM", "alg=AES key_size=128 mode=GCM"},
		{"ConcatenationKDF", "alg=CONCATENATIONKDF"},
		{"DSA", "alg=DSA"},
		{"EC", "alg=EC"},
		{"EC-secp256r1", "alg=EC curve=P-256"},
		{"EC-secp384r1", "alg=EC curve=P-384"},
		{"EC-secp521r1", "alg=EC curve=P-521"},
		{"ECDH", "alg=ECDH"},
		{"Ed25519", "alg=Ed25519"},
		{"EdDSA", "alg=EdDSA"},
		{"HMAC-SHA256", "alg=HMAC hash=SHA-256"},
		{"HMACSHA2", "alg=HMACSHA2"},
		{"MGF1", "alg=MGF1"},
		{"RAW", "alg=RAW"},
		{"RSA-2048", "alg=RSA key_size=2048"},
		{"RSASSA-PSS", "alg=RSA padding=PSS"},
		{"SHA1", "alg=SHA-1"},
		{"SHA256", "alg=SHA-256"},
		{"SHA512", "alg=SHA-512"},
		{"TLS", "alg=TLS"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := show(Canonicalize(c.name, groundtruth.Parameters{})); got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

// The JWS and JWA names of RFC 7518 and RFC 8812, and the prefixed
// spelling crypto-finder writes for python-jose (RSA-RS256).
func TestCanonicalizeJWSNames(t *testing.T) {
	cases := []struct{ name, want string }{
		{"HS256", "alg=HMAC hash=SHA-256"},
		{"HS384", "alg=HMAC hash=SHA-384"},
		{"HS512", "alg=HMAC hash=SHA-512"},
		{"RS256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RS384", "alg=RSA padding=PKCS1v15 hash=SHA-384"},
		{"RS512", "alg=RSA padding=PKCS1v15 hash=SHA-512"},
		{"PS256", "alg=RSA padding=PSS hash=SHA-256"},
		{"PS384", "alg=RSA padding=PSS hash=SHA-384"},
		{"PS512", "alg=RSA padding=PSS hash=SHA-512"},
		{"ES256", "alg=ECDSA curve=P-256 hash=SHA-256"},
		{"ES384", "alg=ECDSA curve=P-384 hash=SHA-384"},
		{"ES512", "alg=ECDSA curve=P-521 hash=SHA-512"},
		{"ES256K", "alg=ECDSA curve=secp256k1 hash=SHA-256"},
		{"EdDSA", "alg=EdDSA"},
		{"RSA-OAEP-256", "alg=RSA padding=OAEP hash=SHA-256"},
		{"RSA-OAEP-384", "alg=RSA padding=OAEP hash=SHA-384"},
		{"RSA-OAEP-512", "alg=RSA padding=OAEP hash=SHA-512"},
		{"rs256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},

		{"RSA-RS256", "alg=RSA padding=PKCS1v15 hash=SHA-256"},
		{"RSA-PS512", "alg=RSA padding=PSS hash=SHA-512"},
		{"ECDSA-ES256", "alg=ECDSA curve=P-256 hash=SHA-256"},
		{"HMAC-HS256", "alg=HMAC hash=SHA-256"},
		{"HMAC-RS256", "alg=HMACRS256"},
		{"ECDSA-RS256", "alg=ECDSARS256"},
		{"RSA-ES256", "alg=RSAES256"},
		{"XRS256", "alg=XRS256"},
		{"RS1", "alg=RS1"},
		{"HS128", "alg=HS128"},
		{"ES256-SHA256", "alg=ES256SHA256"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := show(Canonicalize(c.name, groundtruth.Parameters{})); got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestTablesAreConsistent(t *testing.T) {
	for _, d := range append(slices.Clone(digestSchemes), jcaSignatures...) {
		if e, ok := nameIndex[key(d.algorithm)]; !ok || e.canonical != d.algorithm {
			t.Errorf("digest scheme %s names algorithm %q, not a canonical name", d.spelling, d.algorithm)
		}
		if e, ok := paddingIndex[key(d.padding)]; d.padding != "" && (!ok || e.canonical != d.padding) {
			t.Errorf("digest scheme %s names padding %q, not a canonical padding", d.spelling, d.padding)
		}
		if _, ok := digestName(key(d.spelling)); ok {
			t.Errorf("digest scheme %s is itself a digest", d.spelling)
		}
	}
	for _, j := range jwsNames {
		if e, ok := nameIndex[key(j.algorithm)]; !ok || e.canonical != j.algorithm {
			t.Errorf("JWS name %s names algorithm %q, not a canonical name", j.name, j.algorithm)
		}
		if _, ok := nameIndex[key(j.name)]; ok {
			t.Errorf("JWS name %s is also in the algorithm table", j.name)
		}
		p := j.params
		if canonical(paddingIndex, p.Padding) != p.Padding && p.Padding != "" || canonical(curveIndex, p.Curve) != p.Curve && p.Curve != "" {
			t.Errorf("JWS name %s has a padding or curve that is not canonical: %+v", j.name, p)
		}
		if h, ok := digestName(key(p.Hash)); !ok || h != p.Hash {
			t.Errorf("JWS name %s has hash %q, not a canonical digest", j.name, p.Hash)
		}
	}
	for _, e := range names {
		if e.kind == blockCipher && len(keySizes[e.canonical]) == 0 {
			t.Errorf("block cipher %s has no key sizes", e.canonical)
		}
	}
	for _, r := range rsaSchemes {
		if e, ok := paddingIndex[key(r.padding)]; r.padding != "" && (!ok || e.canonical != r.padding) {
			t.Errorf("RSA scheme %s names padding %q, not a canonical padding", r.spelling, r.padding)
		}
	}
}

func TestCanonicalizeExplicitParameters(t *testing.T) {
	cases := []struct {
		label     string
		name      string
		explicit  groundtruth.Parameters
		want      string
		conflicts string
	}{
		{"fields are canonicalized", "AES", groundtruth.Parameters{Mode: "gcm", Padding: "PKCS5Padding", KeySize: 256}, "alg=AES key_size=256 mode=GCM padding=PKCS7", ""},
		{"CycloneDX padding spellings", "RSA", groundtruth.Parameters{Padding: "pkcs1v15"}, "alg=RSA padding=PKCS1v15", ""},
		{"CycloneDX pkcs5 is PKCS7", "AES", groundtruth.Parameters{Padding: "pkcs5"}, "alg=AES padding=PKCS7", ""},
		{"size-suffixed mode field", "AES", groundtruth.Parameters{Mode: "cfb8"}, "alg=AES mode=CFB", ""},
		{"unknown padding is only normalized", "AES", groundtruth.Parameters{Padding: "ISO10126Padding"}, "alg=AES padding=ISO10126PADDING", ""},
		{"curve alias in a field", "ECDSA", groundtruth.Parameters{Curve: "prime256v1"}, "alg=ECDSA curve=P-256", ""},
		{"unknown curve is only normalized", "ECDSA", groundtruth.Parameters{Curve: "sect233k1"}, "alg=ECDSA curve=SECT233K1", ""},
		{"name fills empty fields", "AES-256-GCM", groundtruth.Parameters{Padding: "NoPadding"}, "alg=AES key_size=256 mode=GCM padding=none", ""},
		{"transformation fills empty fields", "AES/CBC/PKCS5Padding", groundtruth.Parameters{KeySize: 128}, "alg=AES key_size=128 mode=CBC padding=PKCS7", ""},
		{"agreeing transformation padding is no conflict", "AES/CBC/PKCS5Padding", groundtruth.Parameters{Padding: "PKCS7"}, "alg=AES mode=CBC padding=PKCS7", ""},
		{"explicit padding wins over the transformation", "RSA/ECB/PKCS1Padding", groundtruth.Parameters{Padding: "OAEP"}, "alg=RSA mode=ECB padding=OAEP",
			`name "RSA/ECB/PKCS1Padding" implies padding PKCS1v15, explicit padding OAEP wins`},
		{"agreeing field is no conflict", "AES-256-GCM", groundtruth.Parameters{Mode: "gcm", KeySize: 256}, "alg=AES key_size=256 mode=GCM", ""},
		{"explicit key_size wins", "AES-128-GCM", groundtruth.Parameters{KeySize: 256}, "alg=AES key_size=256 mode=GCM",
			`name "AES-128-GCM" implies key_size 128, explicit key_size 256 wins`},
		{"explicit mode wins", "AES-GCM", groundtruth.Parameters{Mode: "CBC"}, "alg=AES mode=CBC",
			`name "AES-GCM" implies mode GCM, explicit mode CBC wins`},
		{"explicit padding wins over an RSA scheme", "RSA-PSS", groundtruth.Parameters{Padding: "pkcs1v15"}, "alg=RSA padding=PKCS1v15",
			`name "RSA-PSS" implies padding PSS, explicit padding PKCS1v15 wins`},
		{"explicit key_size wins over an RSA name", "RSA-OAEP-2048", groundtruth.Parameters{KeySize: 4096}, "alg=RSA key_size=4096 padding=OAEP",
			`name "RSA-OAEP-2048" implies key_size 2048, explicit key_size 4096 wins`},
		{"explicit padding wins over a glued cipher name", "AES128-CBC-PKCS5", groundtruth.Parameters{Padding: "NoPadding"}, "alg=AES key_size=128 mode=CBC padding=none",
			`name "AES128-CBC-PKCS5" implies padding PKCS7, explicit padding none wins`},
		{"CycloneDX pss padding", "RSA", groundtruth.Parameters{Padding: "pss"}, "alg=RSA padding=PSS", ""},
		{"explicit curve wins", "ECDSA-secp256k1", groundtruth.Parameters{Curve: "secp256r1"}, "alg=ECDSA curve=P-256",
			`name "ECDSA-secp256k1" implies curve secp256k1, explicit curve P-256 wins`},
		{"unresolved name keeps explicit fields", "$ALG", groundtruth.Parameters{KeySize: 128}, "alg= key_size=128", ""},
		{"hash field is canonicalized", "ECDSA", groundtruth.Parameters{Hash: "sha256"}, "alg=ECDSA hash=SHA-256", ""},
		{"unknown hash is only normalized", "HMAC", groundtruth.Parameters{Hash: "Whirlpool"}, "alg=HMAC hash=WHIRLPOOL", ""},
		{"name fills an empty hash", "HMAC-SHA-256", groundtruth.Parameters{}, "alg=HMAC hash=SHA-256", ""},
		{"agreeing hash is no conflict", "SHA256withRSA", groundtruth.Parameters{Hash: "SHA-256"}, "alg=RSA padding=PKCS1v15 hash=SHA-256", ""},
		{"explicit hash wins", "ECDSA-SHA256", groundtruth.Parameters{Hash: "SHA-384"}, "alg=ECDSA hash=SHA-384",
			`name "ECDSA-SHA256" implies hash SHA-256, explicit hash SHA-384 wins`},
		{"explicit padding wins over a JCA name", "SHA256withRSA", groundtruth.Parameters{Padding: "PSS"}, "alg=RSA padding=PSS hash=SHA-256",
			`name "SHA256withRSA" implies padding PKCS1v15, explicit padding PSS wins`},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			s := Canonicalize(c.name, c.explicit)
			if got := show(s); got != c.want {
				t.Fatalf("got %q, want %q", got, c.want)
			}
			if got := strings.Join(s.Conflicts, "; "); got != c.conflicts {
				t.Fatalf("conflicts %q, want %q", got, c.conflicts)
			}
		})
	}
}

func TestCurve(t *testing.T) {
	for in, want := range map[string]string{"secp256r1": "P-256", "prime256v1": "P-256", "P-521": "P-521", "SECP256K1": "secp256k1"} {
		if got, ok := Curve(in); !ok || got != want {
			t.Errorf("Curve(%q) = %q, %v; want %q", in, got, ok, want)
		}
	}
	for _, in := range []string{"256", "ML-KEM-768", ""} {
		if got, ok := Curve(in); ok {
			t.Errorf("Curve(%q) = %q, want no curve", in, got)
		}
	}
}
