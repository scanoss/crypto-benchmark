// Package algo canonicalizes algorithm names and parameters before
// attribution, so that annotators and adapters may write either the compact
// form (AES-256-GCM) or the split form (AES with key_size 256 and mode GCM).
// The tables in this file are the normative lists METHODOLOGY.md refers to.
package algo

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
)

type kind int

const (
	plain       kind = iota
	blockCipher      // a trailing key size and mode split off
	ecScheme         // a trailing curve splits off
	digest           // a hash; it splits off a signature or MAC name
)

type entry struct {
	canonical string
	kind      kind
	aliases   []string
}

// names lists every algorithm with a canonical spelling. A name matches an
// entry when it equals the canonical name or an alias after key.
var names = []entry{
	{"MD4", digest, nil},
	{"MD5", digest, nil},
	{"SHA-1", digest, nil},
	{"SHA-224", digest, nil},
	{"SHA-256", digest, nil},
	{"SHA-384", digest, nil},
	{"SHA-512", digest, nil},
	{"SHA-512/224", digest, nil},
	{"SHA-512/256", digest, nil},
	{"SHA3-224", digest, nil},
	{"SHA3-256", digest, nil},
	{"SHA3-384", digest, nil},
	{"SHA3-512", digest, nil},
	{"SHAKE128", plain, nil},
	{"SHAKE256", plain, nil},
	{"RIPEMD-160", digest, nil},
	{"BLAKE2b", digest, nil},
	{"BLAKE2s", digest, nil},
	{"BLAKE3", digest, nil},
	{"SM3", digest, nil},
	{"HMAC", plain, nil},
	{"AES", blockCipher, nil},
	{"ARIA", blockCipher, nil},
	{"Camellia", blockCipher, nil},
	{"SM4", blockCipher, nil},
	{"DES", blockCipher, nil},
	{"3DES", blockCipher, []string{"DESede", "TripleDES", "TDEA"}},
	{"Blowfish", blockCipher, nil},
	{"Twofish", blockCipher, nil},
	{"SEED", blockCipher, nil},
	{"RC4", plain, []string{"ARCFOUR"}},
	{"ChaCha20", plain, nil},
	{"ChaCha20-Poly1305", plain, nil},
	{"XChaCha20-Poly1305", plain, nil},
	{"Poly1305", plain, nil},
	{"RSA", plain, nil},
	{"DSA", plain, nil},
	{"DH", plain, nil},
	{"EC", ecScheme, nil},
	{"ECDSA", ecScheme, nil},
	{"ECDH", ecScheme, nil},
	{"ECIES", ecScheme, nil},
	{"Ed25519", plain, nil},
	{"EdDSA", plain, nil}, // the JWS name; it names no curve, so it is neither Ed25519 nor Ed448
	{"Ed448", plain, nil},
	{"X25519", plain, nil},
	{"X448", plain, nil},
	{"Curve25519", plain, nil},
	{"PBKDF2", plain, []string{"PBKDF2-HMAC", "PBKDF2WithHmac"}}, // HMAC is PBKDF2's PRF
	{"HKDF", plain, nil},
	{"scrypt", plain, nil},
	{"bcrypt", plain, nil},
	{"Argon2", plain, nil},
	{"Argon2d", plain, nil},
	{"Argon2i", plain, nil},
	{"Argon2id", plain, nil},
	{"ML-KEM-512", plain, nil},
	{"ML-KEM-768", plain, nil},
	{"ML-KEM-1024", plain, nil},
	{"ML-DSA-44", plain, nil},
	{"ML-DSA-65", plain, nil},
	{"ML-DSA-87", plain, nil},
}

// keySizes lists, for each block cipher, the key sizes in bits that split
// off its name. A number not listed keeps the name whole, so AES-512 is not
// AES with key_size 512. 3DES lists its effective sizes and the sizes with
// parity bits; Blowfish every multiple of 8 from 32 to 448.
var keySizes = map[string][]int{
	"AES":      {128, 192, 256},
	"ARIA":     {128, 192, 256},
	"Camellia": {128, 192, 256},
	"SM4":      {128},
	"DES":      {56, 64},
	"3DES":     {112, 128, 168, 192},
	"Blowfish": multiples(8, 32, 448),
	"Twofish":  {128, 192, 256},
	"SEED":     {128},
}

// rsaScheme is a spelling of an RSA scheme and the padding it names, ""
// for plain RSA. Every padding is a canonical name from the padding table.
type rsaScheme struct{ spelling, padding string }

// rsaSchemes are the RSA scheme names. Each is RSA with its padding, and
// may end in a key size: RSA-OAEP-3072 is RSA, OAEP and 3072.
var rsaSchemes = []rsaScheme{
	{"RSA", ""},
	{"RSA-OAEP", "OAEP"},
	{"RSAES-OAEP", "OAEP"},
	{"RSA-PSS", "PSS"},
	{"RSASSA-PSS", "PSS"},
	{"RSA-PKCS1", "PKCS1v15"},
	{"RSA-PKCS1v15", "PKCS1v15"},
	{"RSASSA-PKCS1-v1_5", "PKCS1v15"},
	{"RSAES-PKCS1-v1_5", "PKCS1v15"},
}

// digestScheme is a scheme that a digest name splits off: the algorithm it
// is and the padding its spelling names, "" for none.
type digestScheme struct{ spelling, algorithm, padding string }

// digestSchemes may be followed by a digest: HMAC-SHA256 is HMAC with hash
// SHA-256, RSASSA-PSS-SHA-256 is RSA with padding PSS and hash SHA-256,
// RSA-OAEP-SHA256 is RSA with padding OAEP and hash SHA-256, and
// PBKDF2WithHmacSHA256 is PBKDF2 with hash SHA-256.
var digestSchemes = []digestScheme{
	{"HMAC", "HMAC", ""},
	{"ECDSA", "ECDSA", ""},
	{"DSA", "DSA", ""},
	{"RSA", "RSA", ""},
	{"RSA-PSS", "RSA", "PSS"},
	{"RSASSA-PSS", "RSA", "PSS"},
	{"RSA-PKCS1", "RSA", "PKCS1v15"},
	{"RSA-PKCS1v15", "RSA", "PKCS1v15"},
	{"RSASSA-PKCS1", "RSA", "PKCS1v15"},
	{"RSASSA-PKCS1-v1_5", "RSA", "PKCS1v15"},
	{"RSA-OAEP", "RSA", "OAEP"},
	{"RSAES-OAEP", "RSA", "OAEP"},
	{"PBKDF2", "PBKDF2", ""},
	{"PBKDF2-HMAC", "PBKDF2", ""},
	{"PBKDF2WithHmac", "PBKDF2", ""},
}

// jwsName is a JWS or JWA algorithm name and what it computes.
type jwsName struct {
	name, algorithm string
	params          groundtruth.Parameters
}

// jwsNames are the JWS and JWA algorithm names (RFC 7518, RFC 8812, and
// the RSA-OAEP sizes of the IANA JOSE registry). RS* is PKCS#1 v1.5, PS*
// is PSS, and ES* fixes its curve. The number in RSA-OAEP-256 is the OAEP
// digest, not a key size. EdDSA names no curve and is in the algorithm
// table instead.
var jwsNames = []jwsName{
	{"HS256", "HMAC", groundtruth.Parameters{Hash: "SHA-256"}},
	{"HS384", "HMAC", groundtruth.Parameters{Hash: "SHA-384"}},
	{"HS512", "HMAC", groundtruth.Parameters{Hash: "SHA-512"}},
	{"RS256", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-256"}},
	{"RS384", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-384"}},
	{"RS512", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-512"}},
	{"PS256", "RSA", groundtruth.Parameters{Padding: "PSS", Hash: "SHA-256"}},
	{"PS384", "RSA", groundtruth.Parameters{Padding: "PSS", Hash: "SHA-384"}},
	{"PS512", "RSA", groundtruth.Parameters{Padding: "PSS", Hash: "SHA-512"}},
	{"ES256", "ECDSA", groundtruth.Parameters{Curve: "P-256", Hash: "SHA-256"}},
	{"ES384", "ECDSA", groundtruth.Parameters{Curve: "P-384", Hash: "SHA-384"}},
	{"ES512", "ECDSA", groundtruth.Parameters{Curve: "P-521", Hash: "SHA-512"}},
	{"ES256K", "ECDSA", groundtruth.Parameters{Curve: "secp256k1", Hash: "SHA-256"}},
	{"RSA-OAEP-256", "RSA", groundtruth.Parameters{Padding: "OAEP", Hash: "SHA-256"}},
	{"RSA-OAEP-384", "RSA", groundtruth.Parameters{Padding: "OAEP", Hash: "SHA-384"}},
	{"RSA-OAEP-512", "RSA", groundtruth.Parameters{Padding: "OAEP", Hash: "SHA-512"}},
}

// jcaSignatures are the signature algorithms of a Java (JCA) name
// <digest>with<signature>. JCA's plain RSA signature is PKCS#1 v1.5.
var jcaSignatures = []digestScheme{
	{"RSA", "RSA", "PKCS1v15"},
	{"RSA/PSS", "RSA", "PSS"},
	{"RSAandMGF1", "RSA", "PSS"},
	{"ECDSA", "ECDSA", ""},
	{"DSA", "DSA", ""},
}

// modes are the block cipher modes that split off a block cipher's name or
// a Java transformation's middle part. A size-suffixed CFB or OFB (CFB8,
// OFB128) is the base mode; the feedback size is dropped.
var modes = append([]entry{
	{"CFB", plain, sizeSuffixed("CFB")},
	{"OFB", plain, sizeSuffixed("OFB")},
}, plainEntries("ECB", "CBC", "CTR", "GCM", "CCM", "XTS", "SIV", "GCM-SIV", "OCB", "KW", "KWP")...)

// paddings are the canonical paddings, with the Java (JCA) and CycloneDX
// spellings that mean each. An OAEP padding loses its digest and MGF.
var paddings = []entry{
	{"none", plain, []string{"NoPadding"}},
	{"PKCS7", plain, []string{"PKCS5Padding", "PKCS7Padding", "PKCS5"}},
	{"PKCS1v15", plain, []string{"PKCS1Padding"}},
	{"OAEP", plain, oaepPaddings()},
	{"PSS", plain, []string{"PSSPadding"}},
}

func multiples(step, from, to int) []int {
	var out []int
	for n := from; n <= to; n += step {
		out = append(out, n)
	}
	return out
}

// sizeSuffixed returns mode followed by each feedback size in bits that
// JCA and OpenSSL spell out: CFB1, CFB8, ..., CFB128.
func sizeSuffixed(mode string) []string {
	var out []string
	for _, bits := range []int{1, 8, 16, 32, 64, 128} {
		out = append(out, mode+strconv.Itoa(bits))
	}
	return out
}

// oaepPaddings returns the JCA OAEP padding names: OAEPPadding and
// OAEPWith<digest>AndMGF1Padding for every JCA OAEP digest.
func oaepPaddings() []string {
	out := []string{"OAEPPadding"}
	for _, d := range []string{"MD5", "SHA-1", "SHA-224", "SHA-256", "SHA-384", "SHA-512", "SHA-512/224", "SHA-512/256", "SHA3-224", "SHA3-256", "SHA3-384", "SHA3-512"} {
		out = append(out, "OAEPWith"+d+"AndMGF1Padding")
	}
	return out
}

// curves are the named curves that split off an EC scheme's name and that
// curve parameters are canonicalized to.
var curves = []entry{
	{"P-192", plain, []string{"secp192r1", "prime192v1"}},
	{"P-224", plain, []string{"secp224r1"}},
	{"P-256", plain, []string{"secp256r1", "prime256v1"}},
	{"P-384", plain, []string{"secp384r1"}},
	{"P-521", plain, []string{"secp521r1"}},
	{"secp256k1", plain, nil},
	{"brainpoolP256r1", plain, nil},
	{"brainpoolP320r1", plain, nil},
	{"brainpoolP384r1", plain, nil},
	{"brainpoolP512r1", plain, nil},
}

// spelling is one way to write an entry, keyed by key.
type spelling struct {
	key string
	entry
}

var (
	nameIndex    = index(names)
	modeIndex    = index(modes)
	paddingIndex = index(paddings)
	curveIndex   = index(curves)
	// prefixes are the spellings of block ciphers and EC schemes, longest
	// first, so DESede is tried before DES and ECDSA before EC.
	prefixes = func() []spelling {
		var out []spelling
		for k, e := range nameIndex {
			if e.kind == blockCipher || e.kind == ecScheme {
				out = append(out, spelling{k, e})
			}
		}
		slices.SortFunc(out, func(a, b spelling) int {
			return cmp.Or(cmp.Compare(len(b.key), len(a.key)), cmp.Compare(a.key, b.key))
		})
		return out
	}()
	// digestPrefixes are digestSchemes keyed, longest first, so RSASSA-PSS
	// is tried before RSA.
	digestPrefixes = func() []digestScheme {
		out := slices.Clone(digestSchemes)
		slices.SortFunc(out, func(a, b digestScheme) int {
			return cmp.Or(cmp.Compare(len(key(b.spelling)), len(key(a.spelling))), cmp.Compare(a.spelling, b.spelling))
		})
		return out
	}()
	// rsaKeys maps each RSA scheme's key to its padding.
	rsaKeys = func() map[string]string {
		m := map[string]string{}
		for _, r := range rsaSchemes {
			m[key(r.spelling)] = r.padding
		}
		return m
	}()
)

func plainEntries(canonical ...string) []entry {
	out := make([]entry, len(canonical))
	for i, c := range canonical {
		out[i] = entry{canonical: c}
	}
	return out
}

func index(entries []entry) map[string]entry {
	m := map[string]entry{}
	for _, e := range entries {
		for _, s := range append([]string{e.canonical}, e.aliases...) {
			k := key(s)
			if prev, dup := m[k]; dup {
				panic(fmt.Sprintf("algo: %q and %q share the key %q", prev.canonical, e.canonical, k))
			}
			m[k] = e
		}
	}
	return m
}

// key uppercases s and removes '-', '_' and spaces: two spellings with the
// same key are the same name.
func key(s string) string {
	return strings.ToUpper(strings.NewReplacer("-", "", "_", "", " ", "").Replace(strings.TrimSpace(s)))
}

// Spec is an algorithm and its parameters, both canonical.
type Spec struct {
	Algorithm string // "" when absent or unresolved
	Params    groundtruth.Parameters
	// Conflicts describes each parameter the name implied that disagreed
	// with an explicit field; the explicit field won.
	Conflicts []string
}

// Canonicalize resolves name and the explicit parameters into a Spec. A
// parameter the name carries fills its field only when the explicit field
// is empty. A name containing '$' is an unresolved placeholder and counts as
// absent.
func Canonicalize(name string, explicit groundtruth.Parameters) Spec {
	var s Spec
	var implied groundtruth.Parameters
	if !strings.Contains(name, "$") {
		s.Algorithm, implied = parseName(name)
	}
	s.Params = groundtruth.Parameters{
		Mode:    canonical(modeIndex, explicit.Mode),
		Padding: canonical(paddingIndex, explicit.Padding),
		KeySize: explicit.KeySize,
		Curve:   canonical(curveIndex, explicit.Curve),
		Hash:    canonical(nameIndex, explicit.Hash),
	}
	fill := func(field string, dst *string, v string) {
		switch {
		case v == "":
		case *dst == "":
			*dst = v
		case *dst != v:
			s.Conflicts = append(s.Conflicts, fmt.Sprintf("name %q implies %s %s, explicit %s %s wins", name, field, v, field, *dst))
		}
	}
	fill("mode", &s.Params.Mode, implied.Mode)
	fill("padding", &s.Params.Padding, implied.Padding)
	fill("curve", &s.Params.Curve, implied.Curve)
	fill("hash", &s.Params.Hash, implied.Hash)
	if ks := implied.KeySize; ks != 0 {
		switch s.Params.KeySize {
		case 0:
			s.Params.KeySize = ks
		case ks:
		default:
			s.Conflicts = append(s.Conflicts, fmt.Sprintf("name %q implies key_size %d, explicit key_size %d wins", name, ks, s.Params.KeySize))
		}
	}
	return s
}

// Curve returns the canonical name of a curve in the curves table.
func Curve(s string) (string, bool) {
	e, ok := curveIndex[key(s)]
	return e.canonical, ok
}

// canonical returns the table's spelling of s, or key(s) when the table does
// not list it.
func canonical(table map[string]entry, s string) string {
	k := key(s)
	if e, ok := table[k]; ok {
		return e.canonical
	}
	return k
}

// parseName returns the canonical algorithm in name and the parameters its
// spelling carries. A Java transformation ALG/MODE/PADDING splits into its
// three parts; the padding may itself hold a slash (OAEPWithSHA-512/256AndMGF1Padding). A block cipher may end in a key size, a mode and a
// padding, an EC scheme in a curve, an RSA scheme in a key size, a
// signature or MAC scheme in a digest, and a JCA signature name starts with
// one. Any other name is only normalized by key.
func parseName(name string) (string, groundtruth.Parameters) {
	k := key(name)
	if e, ok := nameIndex[k]; ok {
		return e.canonical, groundtruth.Parameters{}
	}
	if j, ok := jws(k); ok {
		return j.algorithm, j.params
	}
	if parts := strings.SplitN(name, "/", 3); len(parts) == 3 {
		alg, p := parseName(parts[0])
		if p.Mode == "" {
			p.Mode = canonical(modeIndex, parts[1])
		}
		p.Padding = canonical(paddingIndex, parts[2])
		return alg, p
	}
	for _, p := range prefixes {
		rest, ok := strings.CutPrefix(k, p.key)
		if !ok || rest == "" {
			continue
		}
		switch p.kind {
		case blockCipher:
			if params, ok := cipherParams(p.canonical, rest); ok {
				return p.canonical, params
			}
		case ecScheme:
			if c, ok := curveIndex[rest]; ok {
				return p.canonical, groundtruth.Parameters{Curve: c.canonical}
			}
		}
	}
	for scheme, padding := range rsaKeys {
		rest, ok := strings.CutPrefix(k, scheme)
		if !ok {
			continue
		}
		if size, rest := leadingNumber(rest); rest == "" && (size > 0 || padding != "") {
			return "RSA", groundtruth.Parameters{KeySize: size, Padding: padding}
		}
	}
	for _, d := range digestPrefixes {
		if rest, ok := strings.CutPrefix(k, key(d.spelling)); ok {
			if h, ok := digestName(rest); ok {
				return d.algorithm, groundtruth.Parameters{Padding: d.padding, Hash: h}
			}
		}
	}
	if h, sig, ok := strings.Cut(k, "WITH"); ok {
		for _, j := range jcaSignatures {
			if d, ok := digestName(h); ok && sig == key(j.spelling) {
				return j.algorithm, groundtruth.Parameters{Padding: j.padding, Hash: d}
			}
		}
	}
	return k, groundtruth.Parameters{}
}

// jws returns the JWS or JWA name whose key is k, or that k ends in after
// the name of the algorithm it computes: RSA-RS256 is RS256.
func jws(k string) (jwsName, bool) {
	for _, j := range jwsNames {
		prefix, ok := strings.CutSuffix(k, key(j.name))
		if ok && (prefix == "" || nameIndex[prefix].canonical == j.algorithm) {
			return j, true
		}
	}
	return jwsName{}, false
}

// digestName returns the canonical digest whose key is k.
func digestName(k string) (string, bool) {
	e, ok := nameIndex[k]
	return e.canonical, ok && e.kind == digest
}

// cipherParams parses rest, what follows cipher's name in a key, as an
// optional key size from keySizes, an optional mode from the modes table
// and, after a mode, an optional padding from the padding table. At least
// one of the three is present: AES128CBCPKCS5 is 128, CBC and PKCS7.
func cipherParams(cipher, rest string) (groundtruth.Parameters, bool) {
	var p groundtruth.Parameters
	p.KeySize, rest = leadingNumber(rest)
	if p.KeySize != 0 && !slices.Contains(keySizes[cipher], p.KeySize) {
		return p, false
	}
	if rest == "" {
		return p, p.KeySize != 0
	}
	for m, e := range modeIndex {
		pad, ok := strings.CutPrefix(rest, m)
		if !ok {
			continue
		}
		p.Mode = e.canonical
		if pad == "" {
			return p, true
		}
		if e, ok := paddingIndex[pad]; ok {
			p.Padding = e.canonical
			return p, true
		}
	}
	return p, false
}

// leadingNumber splits the decimal digits off the front of s. When there
// are none, or they are not a positive number, it returns 0 and s
// unchanged, so AES0 has no key size.
func leadingNumber(s string) (int, string) {
	digits := len(s) - len(strings.TrimLeft(s, "0123456789"))
	n, err := strconv.Atoi(s[:digits])
	if err != nil || n <= 0 {
		return 0, s
	}
	return n, s[digits:]
}
