package selection

import (
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

// TestRepositoryRecord checks the repository's own record: it must exist,
// parse cleanly and agree with the manifest.
func TestRepositoryRecord(t *testing.T) {
	r := &check.Report{}
	m := corpus.Load("../..", "", r)
	cs := Load("../..", m, r)
	if len(cs) == 0 {
		t.Fatal("corpus/suite-c.yaml is missing or empty")
	}
	for _, p := range r.Problems {
		if p.Severity == check.Error {
			t.Error(p)
		}
	}
}

func row(lang string, rank int, repo, outcome, extra string) string {
	return "  - language: " + lang + "\n    rank: " + itoa(rank) + "\n    repo: " + repo +
		"\n    url: https://github.com/" + repo + "\n    source: https://api.github.com/search\n    retrieved: 2026-11-13\n    stars: 10\n    license: MIT\n    outcome: " + outcome + "\n    reason: \"r\"\n" + extra
}

func itoa(n int) string { return string(rune('0' + n)) }

func sel(id string) string {
	return "    sha: 0123456789abcdef0123456789abcdef01234567\n    source_files: 60\n    signal_files: 6\n    id: " + id + "\n"
}

func valid() string {
	var b strings.Builder
	b.WriteString("schema_version: \"1\"\ncandidates:\n")
	for _, l := range []string{"java", "python", "go", "rust", "javascript"} {
		b.WriteString(row(l, 1, "o/"+l+"-x", "rejected", "    filter: 3\n"))
		b.WriteString(row(l, 2, "o/"+l+"-a", "selected", sel(l+"-a")))
		b.WriteString(row(l, 3, "o/"+l+"-b", "selected", sel(l+"-b")))
	}
	return b.String()
}

func TestParse(t *testing.T) {
	cases := []struct {
		name, from, to, want string
	}{
		{"valid", "", "", ""},
		{"rank gap", "rank: 3\n    repo: o/go-b", "rank: 4\n    repo: o/go-b", "want 3"},
		{"no filter", "    filter: 3\n", "", "want the failing filter"},
		{"bad url", "url: https://github.com/o/rust-a", "url: https://example.com/o/rust-a", "must be https://github.com/o/rust-a"},
		{"one selected", "outcome: selected\n    reason: \"r\"\n    sha: 0123456789abcdef0123456789abcdef01234567\n    source_files: 60\n    signal_files: 6\n    id: java-b", "outcome: rejected\n    reason: \"r\"\n    filter: 6", "java has 1 selected"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := strings.Replace(valid(), c.from, c.to, 1)
			r := &check.Report{}
			Parse("suite-c.yaml", []byte(data), r)
			var got []string
			for _, p := range r.Problems {
				got = append(got, p.String())
			}
			all := strings.Join(got, "\n")
			if c.want == "" && all != "" {
				t.Fatalf("unexpected problems:\n%s", all)
			}
			if c.want != "" && !strings.Contains(all, c.want) {
				t.Fatalf("want a problem containing %q, got:\n%s", c.want, all)
			}
		})
	}
}

func TestCheckManifestPoolEntry(t *testing.T) {
	cs := []Candidate{{Language: corpus.Go, URL: "https://github.com/o/a", SHA: "0123456789abcdef0123456789abcdef01234567", Outcome: Selected, ID: "go-c-a"}}
	manifest := func(poolSHA string) corpus.Manifest {
		var m corpus.Manifest
		for id, sha := range map[string]string{"go-c-a": cs[0].SHA, "go-c-a-pool": poolSHA} {
			m.Entries = append(m.Entries, corpus.Entry{ID: id, Language: corpus.Go, Suite: corpus.SuiteC, Source: corpus.Remote{URL: cs[0].URL, SHA: sha}})
		}
		return m
	}
	r := &check.Report{}
	CheckManifest("suite-c.yaml", cs, manifest(cs[0].SHA), r)
	if len(r.Problems) != 0 {
		t.Fatalf("a matching pool entry was refused: %v", r.Problems)
	}
	r = &check.Report{}
	CheckManifest("suite-c.yaml", cs, manifest("89abcdef0123456789abcdef0123456789abcdef"), r)
	if len(r.Problems) != 1 || !strings.Contains(r.Problems[0].String(), "pool entry") {
		t.Fatalf("a pool entry at another commit was accepted: %v", r.Problems)
	}
}
