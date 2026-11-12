package libraries

import (
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
)

const valid = `schema_version: "1"
libraries:
  - ecosystem: python
    package: cryptography
    metric: pypistats_last_month
    value: 200
    source: https://pypistats.org/api/packages/cryptography/recent
    retrieved: 2026-11-04
    dependents:
      metric: depsdev_dependents
      value: 9
      source: https://api.deps.dev/v3alpha/systems/pypi/packages/cryptography
      retrieved: 2026-11-04
  - ecosystem: python
    package: bcrypt
    metric: pypistats_last_month
    value: 100
    source: https://pypistats.org/api/packages/bcrypt/recent
    retrieved: 2026-11-04
  - ecosystem: go
    package: golang.org/x/crypto/bcrypt
    metric: pkggodev_imported_by
    value: 5
    source: https://pkg.go.dev/golang.org/x/crypto/bcrypt
    retrieved: 2026-11-04
stdlib:
  - ecosystem: python
    package: python-stdlib
    modules: [hashlib, hmac]
    source: https://docs.python.org/3/library/crypto.html
    retrieved: 2026-11-04
excluded:
  - ecosystem: python
    package: zeroize
    reason: memory wiping helper
    metric: pypistats_last_month
    value: 150
    source: https://pypistats.org/api/packages/zeroize/recent
    retrieved: 2026-11-04
`

func TestValidList(t *testing.T) {
	r := &check.Report{}
	l := Parse("libraries.yaml", []byte(valid), r)
	if len(r.Problems) > 0 {
		t.Fatalf("unexpected problems: %v", r.Problems)
	}
	if len(l.Libraries) != 3 || len(l.Stdlib) != 1 || len(l.Excluded) != 1 {
		t.Fatalf("got %d libraries, %d stdlib, %d excluded", len(l.Libraries), len(l.Stdlib), len(l.Excluded))
	}
	if d := l.Libraries[0].Dependents; d == nil || d.Value != 9 {
		t.Fatalf("dependents not parsed: %+v", d)
	}
	if l.Libraries[1].Dependents != nil {
		t.Fatalf("bcrypt has no dependents block, got %+v", l.Libraries[1].Dependents)
	}
}

func TestEachRuleDetectsItsDefect(t *testing.T) {
	cases := []struct {
		name, old, new, path, msg string
	}{
		{"schema version", `schema_version: "1"`, `schema_version: "2"`, "schema_version", `got "2", want "1"`},
		{"unknown field", "    package: bcrypt\n", "    package: bcrypt\n    stars: 5\n", "", "field stars not found"},
		{"unknown ecosystem", "ecosystem: go", "ecosystem: cobol", "libraries[2].ecosystem", `"cobol" is not one of`},
		{"duplicate package", "package: bcrypt", "package: cryptography", "libraries[1].package", `python "cryptography" is already listed at libraries[0]`},
		{"duplicate across sections", "package: zeroize", "package: bcrypt", "excluded[0].package", `already listed at libraries[1]`},
		{"missing package", "    package: bcrypt\n", "", "libraries[1].package", "required"},
		{"missing source", "    source: https://pypistats.org/api/packages/bcrypt/recent\n", "", "libraries[1].source", "not an absolute http(s) URL"},
		{"relative source", "source: https://pypistats.org/api/packages/bcrypt/recent", "source: pypistats", "libraries[1].source", "not an absolute http(s) URL"},
		{"missing date", "    retrieved: 2026-11-04\n  - ecosystem: go", "  - ecosystem: go", "libraries[1].retrieved", "not a YYYY-MM-DD date"},
		{"bad date", "retrieved: 2026-11-04\n  - ecosystem: go", "retrieved: 4 Nov\n  - ecosystem: go", "libraries[1].retrieved", "not a YYYY-MM-DD date"},
		{"missing value", "    value: 100\n", "", "libraries[1].value", "required"},
		{"wrong primary metric", "metric: pkggodev_imported_by", "metric: npm_weekly_downloads", "libraries[2].metric", `want "pkggodev_imported_by"`},
		{"wrong dependents metric", "metric: depsdev_dependents", "metric: pypistats_last_month", "libraries[0].dependents.metric", `want "depsdev_dependents"`},
		{"not descending", "value: 100\n", "value: 300\n", "libraries[1].value", "above the previous python library's 200"},
		{"stdlib without modules", "    modules: [hashlib, hmac]\n", "", "stdlib[0].modules", "required"},
		{"stdlib without source", "    source: https://docs.python.org/3/library/crypto.html\n", "", "stdlib[0].source", "not an absolute http(s) URL"},
		{"excluded without reason", "    reason: memory wiping helper\n", "", "excluded[0].reason", "required"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !strings.Contains(valid, c.old) {
				t.Fatalf("fixture does not contain %q", c.old)
			}
			r := &check.Report{}
			Parse("libraries.yaml", []byte(strings.Replace(valid, c.old, c.new, 1)), r)
			for _, p := range r.Problems {
				if p.Path == c.path && strings.Contains(p.Msg, c.msg) {
					return
				}
			}
			t.Fatalf("no problem at %q containing %q; got %v", c.path, c.msg, r.Problems)
		})
	}
}

func TestMoreThanFifteenPerEcosystem(t *testing.T) {
	var b strings.Builder
	b.WriteString("schema_version: \"1\"\nlibraries:\n")
	for i := 0; i < MaxPerEcosystem+1; i++ {
		b.WriteString("  - ecosystem: rust\n    package: crate" + string(rune('a'+i)) +
			"\n    metric: crates_recent_downloads\n    value: 1\n    source: https://crates.io/\n    retrieved: 2026-11-04\n")
	}
	r := &check.Report{}
	Parse("libraries.yaml", []byte(b.String()), r)
	if len(r.Problems) != 1 || r.Problems[0].Path != "libraries[15].package" || !strings.Contains(r.Problems[0].Msg, "more than 15 rust libraries") {
		t.Fatalf("got %v", r.Problems)
	}
}
