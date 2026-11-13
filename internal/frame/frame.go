// Package frame computes the tool-independent crypto signal of a suite C
// repository: the files of the entry's language that import a crypto module,
// the seeded sample of them that the `files` scope annotates, and the grep
// candidate lines that join the `pool` scope. METHODOLOGY.md ("Suite C")
// pre-registers every pattern here; changing one is a change to those rules.
package frame

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

// SampleSize is how many signal files the `files` scope annotates per repository.
const SampleSize = 15

// Extensions are the source files of each language. A JavaScript entry
// covers JavaScript and TypeScript.
var Extensions = map[corpus.Language][]string{
	corpus.Java:       {".java"},
	corpus.Python:     {".py"},
	corpus.Go:         {".go"},
	corpus.Rust:       {".rs"},
	corpus.JavaScript: {".js", ".jsx", ".mjs", ".cjs", ".ts", ".tsx", ".mts", ".cts"},
}

// vendoredDirs are path components whose subtree is vendored or installed
// third-party code, never the project's own source.
var vendoredDirs = map[string]bool{
	"vendor": true, "third_party": true, "third-party": true, "thirdparty": true,
	"node_modules": true, "bower_components": true, ".git": true,
}

// Excluded reports whether a slash-separated path relative to the repository
// root is outside every suite C scope: vendored code, minified bundles and
// TypeScript declaration files.
func Excluded(path string) bool {
	parts := strings.Split(path, "/")
	for _, p := range parts[:len(parts)-1] {
		if vendoredDirs[p] {
			return true
		}
	}
	base := parts[len(parts)-1]
	return strings.HasSuffix(base, ".min.js") || strings.HasSuffix(base, ".d.ts") ||
		strings.HasSuffix(base, ".d.mts") || strings.HasSuffix(base, ".d.cts")
}

// InScope reports whether a slash-separated path relative to the repository
// root is a source file of lang outside excluded paths: the files every suite
// C scope covers.
func InScope(lang corpus.Language, path string) bool {
	return isSource(lang, path) && !Excluded(path)
}

func isSource(lang corpus.Language, path string) bool {
	ext := filepath.Ext(path)
	for _, e := range Extensions[lang] {
		if ext == e {
			return true
		}
	}
	return false
}

// cryptoSegments are module-path segments that name a crypto module, after
// lowercasing and turning '-' into '_'. cryptoSubstrings match anywhere in
// a segment. A module path is split into segments at '.', '/', ':' and '@'.
var cryptoSegments = map[string]bool{}

func init() {
	for _, s := range strings.Fields(`
		hashlib hmac secrets ssl tls jose joserfc nacl pynacl tweetnacl sodium libsodium sodiumoxide
		paramiko ssh ssh2 jsch sshj rsa ecdsa ecdh ed25519 x25519 eddsa curve25519
		curve25519_dalek ed25519_dalek x25519_dalek p256 p384 k256
		jsonwebtoken jsrsasign otp pyotp totp hotp pgp gpg openpgp gnupg argon2 argon2_cffi passlib security
		bouncycastle tink conscrypt nimbusds jasypt truststore ring webpki
		sha1 sha2 sha3 sha256 sha512 md5 md_5 blake2 blake2b blake3 hkdf pbkdf2
		aes aes_gcm chacha20 chacha20poly1305 keccak ripemd160 hashes noble node_forge
		digest signature cipher ciphers webcrypto pki`) {
		cryptoSegments[s] = true
	}
}

var cryptoSubstrings = []string{"crypt", "jwt", "jws", "jwe", "jwk", "x509", "pkcs", "openssl", "rustls"}

// CryptoModule reports whether an imported module path names a crypto module.
func CryptoModule(module string) bool {
	segs := strings.FieldsFunc(strings.ToLower(module), func(r rune) bool {
		return r == '.' || r == '/' || r == ':' || r == '@'
	})
	for _, s := range segs {
		s = strings.ReplaceAll(s, "-", "_")
		if cryptoSegments[s] {
			return true
		}
		for _, sub := range cryptoSubstrings {
			if strings.Contains(s, sub) {
				return true
			}
		}
	}
	return false
}

var (
	pyFrom      = regexp.MustCompile(`^\s*from\s+([\w.]+)\s+import\b`)
	pyImport    = regexp.MustCompile(`^\s*import\s+(.+)$`)
	javaImport  = regexp.MustCompile(`^\s*import\s+(?:static\s+)?([\w.]+)`)
	goSingle    = regexp.MustCompile(`^\s*import\s+(?:[\w.]+\s+)?"([^"]+)"`)
	goBlock     = regexp.MustCompile(`^\s*import\s*\(\s*$`)
	goBlockLine = regexp.MustCompile(`^\s*(?:[\w.]+\s+)?"([^"]+)"`)
	rustUse     = regexp.MustCompile(`^\s*(?:pub(?:\([^)]*\))?\s+)?use\s+([\w:]+)`)
	rustCrate   = regexp.MustCompile(`^\s*extern\s+crate\s+(\w+)`)
	rustPath    = regexp.MustCompile(`\b([A-Za-z_][A-Za-z0-9_]*)::`)
	jsFrom      = regexp.MustCompile(`\bfrom\s+['"]([^'"]+)['"]`)
	jsBare      = regexp.MustCompile(`^\s*import\s+['"]([^'"]+)['"]`)
	jsRequire   = regexp.MustCompile(`\b(?:require|import)\s*\(\s*['"]([^'"]+)['"]\s*\)`)
)

// tokens are uses of a crypto API that need no import of a crypto module.
var tokens = map[corpus.Language]*regexp.Regexp{
	corpus.Python:     regexp.MustCompile(`\bos\.urandom\b|\brandom\.SystemRandom\b`),
	corpus.Java:       regexp.MustCompile(`\bjavax?\.(?:security|crypto|net\.ssl)\.`),
	corpus.JavaScript: regexp.MustCompile(`\bcrypto\.(?:subtle|getRandomValues)\b`),
}

// Signal reports whether a source file of lang carries the crypto signal:
// it imports a crypto module (CryptoModule), or uses one of the few crypto
// APIs that need no import. own lists the Go module paths the repository
// itself declares; an import below one of them is judged on the part after
// it, so a repository whose own name contains "crypt" does not flag every
// file that imports its own packages.
func Signal(lang corpus.Language, src []byte, own ...string) bool {
	if t := tokens[lang]; t != nil && t.Match(src) {
		return true
	}
	for _, m := range Modules(lang, src) {
		for _, o := range own {
			if m == o || strings.HasPrefix(m, o+"/") {
				m = strings.TrimPrefix(strings.TrimPrefix(m, o), "/")
				break
			}
		}
		if CryptoModule(m) {
			return true
		}
	}
	return false
}

// Modules lists the module paths a source file of lang imports. For Rust it
// also lists the first segment of every path expression, since a crate is
// often used by its full path without a `use`.
func Modules(lang corpus.Language, src []byte) []string {
	var out []string
	inGoBlock := false
	sc := bufio.NewScanner(bytes.NewReader(src))
	sc.Buffer(make([]byte, 1024*1024), 16*1024*1024)
	for sc.Scan() {
		line := sc.Text()
		switch lang {
		case corpus.Python:
			if m := pyFrom.FindStringSubmatch(line); m != nil {
				out = append(out, m[1])
			} else if m := pyImport.FindStringSubmatch(line); m != nil {
				for _, part := range strings.Split(m[1], ",") {
					if f := strings.Fields(part); len(f) > 0 {
						out = append(out, f[0])
					}
				}
			}
		case corpus.Java:
			if m := javaImport.FindStringSubmatch(line); m != nil {
				out = append(out, m[1])
			}
		case corpus.Go:
			switch {
			case inGoBlock && strings.HasPrefix(strings.TrimSpace(line), ")"):
				inGoBlock = false
			case inGoBlock:
				if m := goBlockLine.FindStringSubmatch(line); m != nil {
					out = append(out, m[1])
				}
			case goBlock.MatchString(line):
				inGoBlock = true
			default:
				if m := goSingle.FindStringSubmatch(line); m != nil {
					out = append(out, m[1])
				}
			}
		case corpus.Rust:
			if m := rustUse.FindStringSubmatch(line); m != nil {
				out = append(out, m[1])
			}
			if m := rustCrate.FindStringSubmatch(line); m != nil {
				out = append(out, m[1])
			}
			for _, m := range rustPath.FindAllStringSubmatch(line, -1) {
				out = append(out, m[1])
			}
		case corpus.JavaScript:
			for _, re := range []*regexp.Regexp{jsFrom, jsBare, jsRequire} {
				for _, m := range re.FindAllStringSubmatch(line, -1) {
					out = append(out, m[1])
				}
			}
		}
	}
	return out
}

// Frame is the crypto signal of one repository checkout.
type Frame struct {
	SourceFiles int      // files of the language outside excluded paths
	Signal      []string // those that carry the crypto signal, sorted
}

// Scan walks dir and computes its frame for lang. Symbolic links are not followed.
func Scan(dir string, lang corpus.Language) (Frame, error) {
	if _, ok := Extensions[lang]; !ok {
		return Frame{}, fmt.Errorf("frame: no source extensions for language %q", lang)
	}
	var own []string
	if lang == corpus.Go {
		var err error
		if own, err = goModules(dir); err != nil {
			return Frame{}, err
		}
	}
	var f Frame
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			if rel != "." && vendoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() || !InScope(lang, rel) {
			return nil
		}
		f.SourceFiles++
		src, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if Signal(lang, src, own...) {
			f.Signal = append(f.Signal, rel)
		}
		return nil
	})
	if err != nil {
		return Frame{}, fmt.Errorf("frame: %w", err)
	}
	sort.Strings(f.Signal)
	return f, nil
}

var goModuleLine = regexp.MustCompile(`(?m)^\s*module\s+"?([^\s"]+)"?`)

// goModules lists the module paths of every go.mod under dir outside
// excluded paths, longest first so a nested module wins over its parent.
func goModules(dir string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if p != dir && vendoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if d.Name() != "go.mod" || !d.Type().IsRegular() {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if m := goModuleLine.FindSubmatch(data); m != nil {
			out = append(out, string(m[1]))
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("frame: %w", err)
	}
	sort.Slice(out, func(i, j int) bool { return len(out[i]) > len(out[j]) })
	return out, nil
}

// Sample draws the `files` scope of a repository from its signal files: each
// file is ranked by the hex SHA-256 of "<commit>:<path>", and the first n are
// taken, or every file when there are no more than n. The pinned commit is
// the seed, so nobody chooses it, and the draw is the same on every machine.
func Sample(signal []string, commit string, n int) []string {
	type ranked struct{ key, path string }
	rs := make([]ranked, len(signal))
	for i, p := range signal {
		sum := sha256.Sum256([]byte(commit + ":" + p))
		rs[i] = ranked{hex.EncodeToString(sum[:]), p}
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i].key < rs[j].key })
	if len(rs) > n {
		rs = rs[:n]
	}
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.path
	}
	sort.Strings(out)
	return out
}

// grepTerms is the pool's grep candidate pattern: a line of a source file
// that names a crypto operation, algorithm or API. It is deliberately broad;
// every line it matches is adjudicated.
var grepTerms = regexp.MustCompile(`(?i)crypt|cipher|hmac|digest|hashlib|\bsha-?(?:1|224|256|384|512|3)|\bmd5|\baes|\brsa\b|ecdsa|ecdh|ed25519|x25519|pbkdf2|argon2|hkdf|chacha|x509|\btls|\bssl|jwt|\bjw[sek]\b|urandom|securerandom|getrandomvalues|token_bytes|token_hex|token_urlsafe|keypair|private_?key|\bsign\s*\(|\bverify\s*\(`)

// Line is one grep candidate.
type Line struct {
	File string
	Line int
}

// Grep lists the grep candidate lines of the source files of lang under dir:
// every line, outside import lines, that grepTerms matches.
func Grep(dir string, lang corpus.Language) ([]Line, error) {
	var out []Line
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			if rel != "." && vendoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() || !InScope(lang, rel) {
			return nil
		}
		src, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		sc := bufio.NewScanner(bytes.NewReader(src))
		sc.Buffer(make([]byte, 1024*1024), 16*1024*1024)
		for n := 1; sc.Scan(); n++ {
			line := sc.Text()
			if importLine(lang, line) {
				continue
			}
			if grepTerms.MatchString(line) {
				out = append(out, Line{rel, n})
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("frame: %w", err)
	}
	return out, nil
}

var importLines = map[corpus.Language]*regexp.Regexp{
	corpus.Python:     regexp.MustCompile(`^\s*(?:from\s+\S+\s+)?import\s`),
	corpus.Java:       regexp.MustCompile(`^\s*import\s`),
	corpus.Go:         regexp.MustCompile(`^\s*(?:import\s+)?(?:[\w.]+\s+)?"[^"]+"\s*$`),
	corpus.Rust:       regexp.MustCompile(`^\s*(?:pub(?:\([^)]*\))?\s+)?use\s`),
	corpus.JavaScript: regexp.MustCompile(`^\s*import\s|^\s*(?:const|let|var)\s+[^=]+=\s*require\s*\(`),
}

func importLine(lang corpus.Language, line string) bool {
	re := importLines[lang]
	return re != nil && re.MatchString(line)
}
