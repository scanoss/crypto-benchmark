// Package libraries parses corpus/libraries.yaml, the frozen list of crypto
// libraries that suite A covers. METHODOLOGY.md ("Suite A library list")
// states how the list was chosen; this package only checks that the file
// follows those rules.
package libraries

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

const SchemaVersion = "1"

// MaxPerEcosystem is how many third-party libraries each ecosystem keeps.
const MaxPerEcosystem = 15

// Metric names a popularity measure.
type Metric string

const (
	PyPIStatsLastMonth      Metric = "pypistats_last_month"
	NPMWeeklyDownloads      Metric = "npm_weekly_downloads"
	CratesRecentDownloads   Metric = "crates_recent_downloads"
	PkgGoDevImportedBy      Metric = "pkggodev_imported_by"
	DepsDevDirectDependents Metric = "depsdev_direct_dependents"
	DepsDevDependents       Metric = "depsdev_dependents"
)

// PrimaryMetric is the metric that ranks each ecosystem's libraries.
var PrimaryMetric = map[corpus.Language]Metric{
	corpus.Python:     PyPIStatsLastMonth,
	corpus.JavaScript: NPMWeeklyDownloads,
	corpus.Rust:       CratesRecentDownloads,
	corpus.Go:         PkgGoDevImportedBy,
	corpus.Java:       DepsDevDirectDependents,
}

// Measure is one retrieved metric value and where it came from.
type Measure struct {
	Metric    Metric
	Value     int64
	Source    string
	Retrieved time.Time
}

// Library is a ranked third-party library. Rank follows file order.
type Library struct {
	Ecosystem  corpus.Language
	Package    string
	Primary    Measure
	Dependents *Measure // deps.dev dependents; nil where deps.dev has none
}

// Stdlib is an ecosystem's standard-library crypto, always included.
type Stdlib struct {
	Ecosystem corpus.Language
	Package   string
	Modules   []string
	Source    string
	Retrieved time.Time
}

// Excluded is a candidate that ranked high enough but is not a crypto library.
type Excluded struct {
	Ecosystem corpus.Language
	Package   string
	Reason    string
	Primary   Measure
}

type List struct {
	Libraries []Library
	Stdlib    []Stdlib
	Excluded  []Excluded
}

type rawList struct {
	SchemaVersion string        `yaml:"schema_version"`
	Libraries     []rawLibrary  `yaml:"libraries"`
	Stdlib        []rawStdlib   `yaml:"stdlib"`
	Excluded      []rawExcluded `yaml:"excluded"`
}

type rawMeasure struct {
	Metric    string `yaml:"metric"`
	Value     *int64 `yaml:"value"`
	Source    string `yaml:"source"`
	Retrieved string `yaml:"retrieved"`
}

type rawLibrary struct {
	Ecosystem  string `yaml:"ecosystem"`
	Package    string `yaml:"package"`
	rawMeasure `yaml:",inline"`
	Dependents *rawMeasure `yaml:"dependents"`
}

type rawStdlib struct {
	Ecosystem string   `yaml:"ecosystem"`
	Package   string   `yaml:"package"`
	Modules   []string `yaml:"modules"`
	Source    string   `yaml:"source"`
	Retrieved string   `yaml:"retrieved"`
}

type rawExcluded struct {
	Ecosystem  string `yaml:"ecosystem"`
	Package    string `yaml:"package"`
	Reason     string `yaml:"reason"`
	rawMeasure `yaml:",inline"`
}

func Path(root string) string { return filepath.Join(root, "corpus", "libraries.yaml") }

// Load parses root's library list, recording every problem in r.
func Load(root string, r *check.Report) List {
	file := Path(root)
	data, err := os.ReadFile(file)
	if err != nil {
		r.Errorf(file, "", "%v", err)
		return List{}
	}
	return Parse(file, data, r)
}

// Parse decodes and checks one library list.
func Parse(file string, data []byte, r *check.Report) List {
	var raw rawList
	if err := check.DecodeStrict(data, &raw); err != nil {
		r.Errorf(file, "", "%v", err)
		return List{}
	}
	p := parser{file: file, r: r, seen: map[string]string{}}
	if raw.SchemaVersion != SchemaVersion {
		r.Errorf(file, "schema_version", "got %q, want %q", raw.SchemaVersion, SchemaVersion)
	}
	var l List
	for i, rl := range raw.Libraries {
		l.Libraries = append(l.Libraries, p.library(rl, fmt.Sprintf("libraries[%d]", i)))
	}
	for i, rs := range raw.Stdlib {
		l.Stdlib = append(l.Stdlib, p.stdlib(rs, fmt.Sprintf("stdlib[%d]", i)))
	}
	for i, re := range raw.Excluded {
		l.Excluded = append(l.Excluded, p.excluded(re, fmt.Sprintf("excluded[%d]", i)))
	}
	p.checkRanking(l)
	return l
}

type parser struct {
	file string
	r    *check.Report
	seen map[string]string // ecosystem/package -> where first declared
}

func (p parser) fail(where, field, format string, args ...any) {
	p.r.Errorf(p.file, where+"."+field, format, args...)
}

// row checks the fields every row has and returns its ecosystem.
func (p parser) row(ecosystem, pkg, where string) corpus.Language {
	eco, err := check.OneOf(ecosystem, corpus.Python, corpus.JavaScript, corpus.Rust, corpus.Go, corpus.Java)
	if err != nil {
		p.fail(where, "ecosystem", "%v", err)
	}
	if pkg == "" {
		p.fail(where, "package", "required")
		return eco
	}
	key := ecosystem + "/" + pkg
	if prev, dup := p.seen[key]; dup {
		p.fail(where, "package", "%s %q is already listed at %s", ecosystem, pkg, prev)
	} else {
		p.seen[key] = where
	}
	return eco
}

func (p parser) source(s, where, field string) string {
	if u, err := url.Parse(s); err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.Host == "" {
		p.fail(where, field, "%q is not an absolute http(s) URL", s)
	}
	return s
}

func (p parser) date(s, where, field string) time.Time {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		p.fail(where, field, "%q is not a YYYY-MM-DD date", s)
	}
	return t
}

func (p parser) measure(rm rawMeasure, want Metric, where string) Measure {
	m := Measure{
		Metric:    Metric(rm.Metric),
		Source:    p.source(rm.Source, where, "source"),
		Retrieved: p.date(rm.Retrieved, where, "retrieved"),
	}
	if want != "" && m.Metric != want {
		p.fail(where, "metric", "got %q, want %q", rm.Metric, want)
	}
	switch {
	case rm.Value == nil:
		p.fail(where, "value", "required")
	case *rm.Value < 0:
		p.fail(where, "value", "must be >= 0, got %d", *rm.Value)
	default:
		m.Value = *rm.Value
	}
	return m
}

func (p parser) library(rl rawLibrary, where string) Library {
	eco := p.row(rl.Ecosystem, rl.Package, where)
	lib := Library{Ecosystem: eco, Package: rl.Package, Primary: p.measure(rl.rawMeasure, PrimaryMetric[eco], where)}
	if rl.Dependents != nil {
		d := p.measure(*rl.Dependents, DepsDevDependents, where+".dependents")
		lib.Dependents = &d
	}
	return lib
}

func (p parser) stdlib(rs rawStdlib, where string) Stdlib {
	s := Stdlib{
		Ecosystem: p.row(rs.Ecosystem, rs.Package, where),
		Package:   rs.Package,
		Modules:   rs.Modules,
		Source:    p.source(rs.Source, where, "source"),
		Retrieved: p.date(rs.Retrieved, where, "retrieved"),
	}
	if len(rs.Modules) == 0 {
		p.fail(where, "modules", "required")
	}
	return s
}

func (p parser) excluded(re rawExcluded, where string) Excluded {
	eco := p.row(re.Ecosystem, re.Package, where)
	if re.Reason == "" {
		p.fail(where, "reason", "required")
	}
	return Excluded{Ecosystem: eco, Package: re.Package, Reason: re.Reason, Primary: p.measure(re.rawMeasure, PrimaryMetric[eco], where)}
}

// checkRanking enforces at most MaxPerEcosystem libraries per ecosystem,
// listed by descending primary metric, so file order is the rank.
func (p parser) checkRanking(l List) {
	count := map[corpus.Language]int{}
	last := map[corpus.Language]int64{}
	for i, lib := range l.Libraries {
		where := fmt.Sprintf("libraries[%d]", i)
		count[lib.Ecosystem]++
		if count[lib.Ecosystem] == MaxPerEcosystem+1 {
			p.fail(where, "package", "more than %d %s libraries", MaxPerEcosystem, lib.Ecosystem)
		}
		if prev, ok := last[lib.Ecosystem]; ok && lib.Primary.Value > prev {
			p.fail(where, "value", "%d is above the previous %s library's %d; list each ecosystem by descending value", lib.Primary.Value, lib.Ecosystem, prev)
		}
		last[lib.Ecosystem] = lib.Primary.Value
	}
}
