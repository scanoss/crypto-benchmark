package frame

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

func TestSignal(t *testing.T) {
	cases := []struct {
		lang corpus.Language
		src  string
		want bool
	}{
		{corpus.Python, "import hashlib\n", true},
		{corpus.Python, "from cryptography.hazmat.primitives import hashes\n", true},
		{corpus.Python, "import os, ssl as s\n", true},
		{corpus.Python, "key = os.urandom(32)\n", true},
		{corpus.Python, "import os\nimport json\n", false},
		{corpus.Java, "import javax.crypto.Cipher;\n", true},
		{corpus.Java, "import org.springframework.security.core.Authentication;\n", true},
		{corpus.Java, "import java.util.HashMap;\n", false},
		{corpus.Java, "var md = java.security.MessageDigest.getInstance(\"SHA-256\");\n", true},
		{corpus.Go, "package x\n\nimport (\n\t\"fmt\"\n\tsha \"crypto/sha256\"\n)\n", true},
		{corpus.Go, "package x\n\nimport \"golang.org/x/crypto/bcrypt\"\n", true},
		{corpus.Go, "package x\n\nimport (\n\t\"fmt\"\n\t\"hash/fnv\"\n)\n\nvar s = []string{\n\t\"crypto/tls\"\n}\n", false},
		{corpus.Rust, "use ring::{digest, hmac};\n", true},
		{corpus.Rust, "let d = sha2::Sha256::digest(b\"x\");\n", true},
		{corpus.Rust, "use std::collections::HashMap;\n", false},
		{corpus.JavaScript, "import { createHash } from 'node:crypto';\n", true},
		{corpus.JavaScript, "const jwt = require(\"jsonwebtoken\");\n", true},
		{corpus.JavaScript, "const k = await crypto.subtle.generateKey(a, true, u);\n", true},
		{corpus.JavaScript, "import React from 'react';\n", false},
	}
	for _, c := range cases {
		if got := Signal(c.lang, []byte(c.src)); got != c.want {
			t.Errorf("Signal(%s, %q) = %v, want %v", c.lang, c.src, got, c.want)
		}
	}
}

func TestSignalOwnGoModule(t *testing.T) {
	src := []byte("package x\n\nimport \"example.com/cryptoapp/internal/store\"\n")
	if Signal(corpus.Go, src, "example.com/cryptoapp") {
		t.Error("an import of the repository's own module was judged by the module's name")
	}
	src = []byte("package x\n\nimport \"example.com/cryptoapp/internal/crypto\"\n")
	if !Signal(corpus.Go, src, "example.com/cryptoapp") {
		t.Error("an own package named crypto lost its signal")
	}
}

func TestExcluded(t *testing.T) {
	for path, want := range map[string]bool{
		"src/a.go":               false,
		"vendor/x/a.go":          true,
		"web/node_modules/a.js":  true,
		"lib/third_party/a.py":   true,
		"static/app.min.js":      true,
		"types/index.d.ts":       true,
		"src/vendor.go":          false,
		"internal/thirdparty.rs": false,
	} {
		if got := Excluded(path); got != want {
			t.Errorf("Excluded(%q) = %v, want %v", path, got, want)
		}
	}
}

func TestScanAndSample(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"a.py":                "import hashlib\n",
		"b.py":                "import json\n",
		"pkg/c.py":            "from Crypto.Cipher import AES\n",
		"vendor/d.py":         "import hashlib\n",
		"node_modules/e.py":   "import ssl\n",
		"README.md":           "import hashlib\n",
		"pkg/sub/f.py":        "import secrets\n",
		"pkg/sub/g.py":        "x = 1\n",
		"third_party/x/h.py":  "import hmac\n",
		"tests/test_crypt.py": "import bcrypt\n",
	}
	for name, body := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	f, err := Scan(dir, corpus.Python)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"a.py", "pkg/c.py", "pkg/sub/f.py", "tests/test_crypt.py"}
	if f.SourceFiles != 6 || !reflect.DeepEqual(f.Signal, want) {
		t.Fatalf("Scan = %d files, %v; want 6, %v", f.SourceFiles, f.Signal, want)
	}

	commit := "0123456789abcdef0123456789abcdef01234567"
	s := Sample(f.Signal, commit, 2)
	if len(s) != 2 || !reflect.DeepEqual(s, Sample(f.Signal, commit, 2)) {
		t.Fatalf("Sample is not a deterministic draw of 2: %v", s)
	}
	if all := Sample(f.Signal, commit, 15); !reflect.DeepEqual(all, want) {
		t.Fatalf("Sample with n above the frame = %v, want every file", all)
	}
	other := Sample(f.Signal, "89abcdef0123456789abcdef0123456789abcdef", 4)
	if !reflect.DeepEqual(other, want) {
		t.Fatalf("Sample of the whole frame depends on the seed: %v", other)
	}

	lines, err := Grep(dir, corpus.Python)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Fatalf("Grep matched import lines: %v", lines)
	}
}

// TestSampleKnownDraw pins the ranking function, so the published samples
// cannot drift with a refactor.
func TestSampleKnownDraw(t *testing.T) {
	signal := []string{"a.go", "b.go", "c.go", "d.go", "e.go"}
	got := Sample(signal, "0123456789abcdef0123456789abcdef01234567", 2)
	// sha256sum of "<commit>:<path>" ranks d.go (18c1...) then a.go (3844...).
	want := []string{"a.go", "d.go"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Sample = %v, want %v", got, want)
	}
}

func TestGrep(t *testing.T) {
	dir := t.TempDir()
	src := "import hashlib\nh = hashlib.sha256(b'x')\nprint('hello')\nsig = key.sign(data)\n"
	if err := os.WriteFile(filepath.Join(dir, "a.py"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := Grep(dir, corpus.Python)
	if err != nil {
		t.Fatal(err)
	}
	want := []Line{{"a.py", 2}, {"a.py", 4}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Grep = %v, want %v", got, want)
	}
}
