// Package corpus parses corpus/manifest.yaml from the public repository and,
// optionally, from a private overlay with the same layout.
package corpus

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/scanoss/crypto-benchmark/internal/check"
)

const SchemaVersion = "1"

type Language string

const (
	Java       Language = "java"
	Python     Language = "python"
	Go         Language = "go"
	Rust       Language = "rust"
	JavaScript Language = "javascript"
)

type Suite string

const (
	SuiteA Suite = "A"
	SuiteB Suite = "B"
	SuiteC Suite = "C"
)

type Split string

const (
	Dev     Split = "dev"
	Holdout Split = "holdout"
)

type Variant string

const (
	Reachable   Variant = "reachable"
	Unreachable Variant = "unreachable"
)

// Source is either Remote (fetched by URL at a pinned commit) or Local (an in-repo fixture).
type Source interface{ isSource() }

type Remote struct {
	URL string
	SHA string
}

type Local struct {
	Path string // relative to Entry.Root
}

func (Remote) isSource() {}
func (Local) isSource()  {}

// Pair links a suite B fixture to its counterpart variant.
type Pair struct {
	Partner string
	Variant Variant
	// Scenario names what the pair tests (transitive, dispatch, ...); both
	// variants carry the same one. Suite B scores are broken down by it.
	Scenario string
}

type Entry struct {
	ID       string
	Language Language
	Suite    Suite
	Split    Split
	Source   Source
	Pair     *Pair
	Notes    string

	Root    string // directory holding corpus/ for this entry
	File    string // manifest file the entry came from
	Where   string // YAML path inside File, e.g. entries[2]
	Overlay bool
}

type Manifest struct {
	Entries  []Entry
	declared map[string]bool
}

// Declared reports whether any manifest declares id, even if that entry failed to parse.
func (m Manifest) Declared(id string) bool { return m.declared[id] }

func (m Manifest) Lookup(id string) (Entry, bool) {
	for _, e := range m.Entries {
		if e.ID == id {
			return e, true
		}
	}
	return Entry{}, false
}

type rawManifest struct {
	SchemaVersion string     `yaml:"schema_version"`
	Entries       []rawEntry `yaml:"entries"`
}

type rawEntry struct {
	ID       string `yaml:"id"`
	Language string `yaml:"language"`
	Suite    string `yaml:"suite"`
	Split    string `yaml:"split"`
	URL      string `yaml:"url"`
	SHA      string `yaml:"sha"`
	Path     string `yaml:"path"`
	Pair     string `yaml:"pair"`
	Variant  string `yaml:"variant"`
	Scenario string `yaml:"scenario"`
	Notes    string `yaml:"notes"`
}

var (
	idPattern  = regexp.MustCompile(`^[a-z0-9-]+$`)
	shaPattern = regexp.MustCompile(`^[0-9a-f]{40}$`)
)

func ManifestPath(root string) string { return filepath.Join(root, "corpus", "manifest.yaml") }

// Load parses the base manifest and, when overlay is not empty, the overlay's,
// recording every problem in r. Only entries that parsed cleanly are returned.
func Load(base, overlay string, r *check.Report) Manifest {
	m := Manifest{declared: map[string]bool{}}
	seen := map[string]string{}
	roots := []string{base}
	if overlay != "" {
		roots = append(roots, overlay)
	}
	for i, root := range roots {
		m.Entries = append(m.Entries, loadFile(root, i == 1, seen, m.declared, r)...)
	}
	checkPairs(m, r)
	return m
}

func loadFile(root string, overlay bool, seen map[string]string, declared map[string]bool, r *check.Report) []Entry {
	file := ManifestPath(root)
	data, err := os.ReadFile(file)
	if err != nil {
		r.Errorf(file, "", "%v", err)
		return nil
	}
	var raw rawManifest
	if err := check.DecodeStrict(data, &raw); err != nil {
		r.Errorf(file, "", "%v", err)
		return nil
	}
	if raw.SchemaVersion != SchemaVersion {
		r.Errorf(file, "schema_version", "got %q, want %q", raw.SchemaVersion, SchemaVersion)
	}
	var out []Entry
	for i, re := range raw.Entries {
		where := fmt.Sprintf("entries[%d]", i)
		loc := file + ": " + where
		if re.ID != "" {
			if prev, dup := seen[re.ID]; dup {
				r.Errorf(file, where+".id", "duplicate id %q, first declared at %s", re.ID, prev)
			} else {
				seen[re.ID] = loc
			}
			declared[re.ID] = true
		}
		if e, ok := parseEntry(re, root, file, where, overlay, r); ok {
			out = append(out, e)
		}
	}
	return out
}

func parseEntry(re rawEntry, root, file, where string, overlay bool, r *check.Report) (Entry, bool) {
	before := len(r.Problems)
	fail := func(field, format string, args ...any) {
		r.Errorf(file, where+"."+field, format, args...)
	}
	e := Entry{ID: re.ID, Notes: re.Notes, Root: root, File: file, Where: where, Overlay: overlay}
	if !idPattern.MatchString(re.ID) {
		fail("id", "%q must match [a-z0-9-]+", re.ID)
	}
	var err error
	if e.Language, err = check.OneOf(re.Language, Java, Python, Go, Rust, JavaScript); err != nil {
		fail("language", "%v", err)
	}
	if e.Suite, err = check.OneOf(re.Suite, SuiteA, SuiteB, SuiteC); err != nil {
		fail("suite", "%v", err)
	}
	if e.Split, err = check.OneOf(re.Split, Dev, Holdout); err != nil {
		fail("split", "%v", err)
	}
	if e.Split == Holdout && !overlay {
		fail("split", "holdout entries are only allowed in the private overlay")
	}

	remote, local := re.URL != "" || re.SHA != "", re.Path != ""
	switch {
	case remote && local:
		fail("path", "set either url+sha or path, not both")
	case local:
		e.Source = Local{Path: re.Path}
		if !filepath.IsLocal(re.Path) {
			fail("path", "%q must be a relative path inside the repository", re.Path)
		} else if fi, err := os.Stat(filepath.Join(root, re.Path)); err != nil || !fi.IsDir() {
			fail("path", "%q is not an existing directory", re.Path)
		}
	case remote:
		e.Source = Remote{URL: re.URL, SHA: re.SHA}
		if u, err := url.Parse(re.URL); err != nil || u.Scheme == "" {
			fail("url", "%q is not an absolute URL", re.URL)
		}
		if !shaPattern.MatchString(re.SHA) {
			fail("sha", "%q must be a full 40-character lowercase hex commit", re.SHA)
		}
	default:
		fail("url", "set either url+sha or path")
	}

	switch {
	case re.Pair == "" && re.Variant != "":
		fail("variant", "variant is only allowed together with pair")
	case re.Pair == "" && re.Scenario != "":
		fail("scenario", "scenario is only allowed together with pair")
	case re.Pair != "":
		e.Pair = &Pair{Partner: re.Pair, Scenario: re.Scenario}
		if re.Scenario != "" && !idPattern.MatchString(re.Scenario) {
			fail("scenario", "%q must match [a-z0-9-]+", re.Scenario)
		}
		if e.Pair.Variant, err = check.OneOf(re.Variant, Reachable, Unreachable); err != nil {
			fail("variant", "required with pair: %v", err)
		}
		if e.Suite != SuiteB {
			fail("pair", "pairs are only allowed in suite B")
		}
		if !local {
			fail("pair", "pairs are only allowed on in-repo fixtures (path)")
		}
		if re.Pair == re.ID {
			fail("pair", "an entry cannot pair with itself")
		}
	}
	return e, len(r.Problems) == before
}

func checkPairs(m Manifest, r *check.Report) {
	byID := map[string]Entry{}
	for _, e := range m.Entries {
		byID[e.ID] = e
	}
	ids := make([]string, 0, len(byID))
	for id := range byID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		e := byID[id]
		if e.Pair == nil {
			continue
		}
		path := e.Where + ".pair"
		partner, ok := byID[e.Pair.Partner]
		switch {
		case !ok && !m.declared[e.Pair.Partner]:
			r.Errorf(e.File, path, "partner %q is not a manifest entry", e.Pair.Partner)
		case !ok:
			// The partner is declared but invalid; its own errors explain why.
		case partner.Pair == nil || partner.Pair.Partner != e.ID:
			r.Errorf(e.File, path, "partner %q does not pair back with %q", partner.ID, e.ID)
		case e.ID < partner.ID && partner.Pair.Variant == e.Pair.Variant:
			r.Errorf(e.File, e.Where+".variant", "%q and %q are both %s; a pair needs one reachable and one unreachable", e.ID, partner.ID, e.Pair.Variant)
		case e.ID < partner.ID && partner.Language != e.Language:
			r.Errorf(e.File, e.Where+".language", "%q is %s but its partner %q is %s", e.ID, e.Language, partner.ID, partner.Language)
		case e.ID < partner.ID && partner.Pair.Scenario != e.Pair.Scenario:
			r.Errorf(e.File, e.Where+".scenario", "%q has scenario %q but its partner %q has %q", e.ID, e.Pair.Scenario, partner.ID, partner.Pair.Scenario)
		}
	}
}
