// Package selection parses corpus/suite-c.yaml, the record of every suite C
// candidate examined and its outcome. METHODOLOGY.md ("Suite C selection")
// states the rule; this package checks that the record follows its shape and
// agrees with corpus/manifest.yaml.
package selection

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

const SchemaVersion = "1"

// PerLanguage is how many repositories each language selects.
const PerLanguage = 2

// PoolSuffix names the second entry of a repository, which holds its pool
// ground truth; the plain id holds the files scope.
const PoolSuffix = "-pool"

// Filters is the number of selection filters; a rejection names one of 1..Filters.
const Filters = 6

type Outcome string

const (
	Selected Outcome = "selected"
	Rejected Outcome = "rejected"
)

type Candidate struct {
	Language corpus.Language
	Rank     int
	Repo     string
	URL      string
	SHA      string
	Outcome  Outcome
	Filter   int
	ID       string
	Where    string
}

type rawFile struct {
	SchemaVersion string         `yaml:"schema_version"`
	Candidates    []rawCandidate `yaml:"candidates"`
}

type rawCandidate struct {
	Language    string `yaml:"language"`
	Rank        int    `yaml:"rank"`
	Repo        string `yaml:"repo"`
	URL         string `yaml:"url"`
	Source      string `yaml:"source"`
	Retrieved   string `yaml:"retrieved"`
	Stars       *int64 `yaml:"stars"`
	License     string `yaml:"license"`
	SHA         string `yaml:"sha"`
	LastCommit  string `yaml:"last_commit"`
	Outcome     string `yaml:"outcome"`
	Filter      int    `yaml:"filter"`
	Reason      string `yaml:"reason"`
	SourceFiles *int   `yaml:"source_files"`
	SignalFiles *int   `yaml:"signal_files"`
	ID          string `yaml:"id"`
}

var (
	shaPattern  = regexp.MustCompile(`^[0-9a-f]{40}$`)
	repoPattern = regexp.MustCompile(`^[A-Za-z0-9_.-]+/[A-Za-z0-9_.-]+$`)
)

func Path(root string) string { return filepath.Join(root, "corpus", "suite-c.yaml") }

// Load parses root's selection record and checks it against m. A tree
// without the file (a test fixture, an overlay) is not checked; the
// repository's own record is required by TestRepositoryRecord.
func Load(root string, m corpus.Manifest, r *check.Report) []Candidate {
	file := Path(root)
	data, err := os.ReadFile(file)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		r.Errorf(file, "", "%v", err)
		return nil
	}
	cs := Parse(file, data, r)
	CheckManifest(file, cs, m, r)
	return cs
}

// Parse decodes one selection record and checks each row and the ranking.
func Parse(file string, data []byte, r *check.Report) []Candidate {
	var raw rawFile
	if err := check.DecodeStrict(data, &raw); err != nil {
		r.Errorf(file, "", "%v", err)
		return nil
	}
	if raw.SchemaVersion != SchemaVersion {
		r.Errorf(file, "schema_version", "got %q, want %q", raw.SchemaVersion, SchemaVersion)
	}
	var out []Candidate
	lastRank := map[corpus.Language]int{}
	selected := map[corpus.Language]int{}
	seen := map[string]string{}
	for i, rc := range raw.Candidates {
		where := fmt.Sprintf("candidates[%d]", i)
		fail := func(field, format string, args ...any) { r.Errorf(file, where+"."+field, format, args...) }
		c := Candidate{Rank: rc.Rank, Repo: rc.Repo, URL: rc.URL, SHA: rc.SHA, Filter: rc.Filter, ID: rc.ID, Where: where}
		var err error
		if c.Language, err = check.OneOf(rc.Language, corpus.Java, corpus.Python, corpus.Go, corpus.Rust, corpus.JavaScript); err != nil {
			fail("language", "%v", err)
		}
		if !repoPattern.MatchString(rc.Repo) {
			fail("repo", "%q must be OWNER/NAME", rc.Repo)
		} else if prev, dup := seen[rc.Repo]; dup {
			fail("repo", "%q is already listed at %s", rc.Repo, prev)
		} else {
			seen[rc.Repo] = where
		}
		if rc.URL != "https://github.com/"+rc.Repo {
			fail("url", "%q must be https://github.com/%s", rc.URL, rc.Repo)
		}
		if rc.Source == "" {
			fail("source", "required")
		}
		for field, v := range map[string]string{"retrieved": rc.Retrieved, "last_commit": rc.LastCommit} {
			if v == "" && field == "last_commit" {
				continue
			}
			if _, err := time.Parse(time.DateOnly, v); err != nil {
				fail(field, "%q is not a YYYY-MM-DD date", v)
			}
		}
		if rc.Stars == nil {
			fail("stars", "required")
		}
		if rc.License == "" {
			fail("license", "required")
		}
		if rc.Reason == "" {
			fail("reason", "required")
		}
		if rc.SHA != "" && !shaPattern.MatchString(rc.SHA) {
			fail("sha", "%q must be a full 40-character lowercase hex commit", rc.SHA)
		}
		if rc.Rank != lastRank[c.Language]+1 {
			fail("rank", "got %d, want %d: list every candidate examined, in rank order", rc.Rank, lastRank[c.Language]+1)
		}
		lastRank[c.Language] = rc.Rank
		if selected[c.Language] == PerLanguage {
			fail("rank", "%s already has %d selected repositories; the search stops there", c.Language, PerLanguage)
		}
		if c.Outcome, err = check.OneOf(rc.Outcome, Selected, Rejected); err != nil {
			fail("outcome", "%v", err)
		}
		switch c.Outcome {
		case Selected:
			selected[c.Language]++
			if rc.Filter != 0 {
				fail("filter", "only a rejected candidate names a filter")
			}
			if rc.ID == "" || rc.SHA == "" || rc.SourceFiles == nil || rc.SignalFiles == nil {
				fail("id", "a selected candidate needs id, sha, source_files and signal_files")
			}
		case Rejected:
			if rc.Filter < 1 || rc.Filter > Filters {
				fail("filter", "got %d, want the failing filter, 1 to %d", rc.Filter, Filters)
			}
			if rc.ID != "" {
				fail("id", "only a selected candidate has an id")
			}
		}
		out = append(out, c)
	}
	for _, lang := range []corpus.Language{corpus.Java, corpus.Python, corpus.Go, corpus.Rust, corpus.JavaScript} {
		if n := selected[lang]; n != PerLanguage {
			r.Errorf(file, "candidates", "%s has %d selected repositories, want %d", lang, n, PerLanguage)
		}
	}
	return out
}

// CheckManifest checks that the selected candidates and the base manifest's
// suite C entries are the same repositories at the same commits.
func CheckManifest(file string, cs []Candidate, m corpus.Manifest, r *check.Report) {
	byID := map[string]Candidate{}
	for _, c := range cs {
		if c.Outcome != Selected {
			continue
		}
		byID[c.ID] = c
		e, ok := m.Lookup(c.ID)
		if !ok {
			r.Errorf(file, c.Where+".id", "%q is not a manifest entry", c.ID)
			continue
		}
		src, remote := e.Source.(corpus.Remote)
		if e.Suite != corpus.SuiteC || !remote || src.URL != c.URL || src.SHA != c.SHA || e.Language != c.Language {
			r.Errorf(file, c.Where+".id", "%q does not match its manifest entry (suite C, %s, %s @ %s)", c.ID, c.Language, c.URL, c.SHA)
		}
	}
	for _, e := range m.Entries {
		if e.Suite != corpus.SuiteC || e.Overlay {
			continue
		}
		c, ok := byID[strings.TrimSuffix(e.ID, PoolSuffix)]
		if !ok {
			r.Errorf(e.File, e.Where, "suite C entry %q is not a selected candidate in %s, or one with %q appended", e.ID, file, PoolSuffix)
			continue
		}
		src, remote := e.Source.(corpus.Remote)
		if e.ID != c.ID && (!remote || src.URL != c.URL || src.SHA != c.SHA || e.Language != c.Language) {
			r.Errorf(e.File, e.Where, "pool entry %q must have the language, url and sha of %q", e.ID, c.ID)
		}
	}
}
