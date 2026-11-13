package score

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// Run is one tool run's predictions, keyed by entry id, under a label that
// tells runs of the same tool apart.
type Run struct {
	Label string
	Preds map[string]prediction.File
}

// UniqueRow counts, for one run, the pool assets it matched and those that
// no other compared run matched, broken down by language and asset type.
type UniqueRow struct {
	Label        string         `json:"run"`
	Found        int            `json:"found"`
	Only         int            `json:"only_this_run"`
	OnlyByLang   map[string]int `json:"only_by_language"`
	OnlyByType   map[string]int `json:"only_by_asset_type"`
	FoundByLang  map[string]int `json:"found_by_language"`
	FoundByType  map[string]int `json:"found_by_asset_type"`
	EntriesNoRun []string       `json:"entries_not_run"`
}

// UniqueReport answers METHODOLOGY.md "pool scope": per tool, the number of
// adjudicated pool assets that only that tool found.
type UniqueReport struct {
	Entries []string    `json:"entries"`
	Assets  int         `json:"assets"`
	Runs    []UniqueRow `json:"runs"`
	// None counts the pool assets no compared run matched: the grep or a run
	// left out of the comparison brought them into the pool.
	None       int            `json:"found_by_no_compared_run"`
	NoneByLang map[string]int `json:"none_by_language"`
	NoneByType map[string]int `json:"none_by_asset_type"`
}

// Unique matches every run against the ground truth of every pool entry, by
// the same rules as Score (predictions outside the pool frame dropped, then
// MatchEntry), and counts which assets each run matched alone. An entry a run
// has no prediction file for counts as nothing found by that run.
func Unique(m corpus.Manifest, truth map[string]groundtruth.File, runs []Run) UniqueReport {
	rep := UniqueReport{NoneByLang: map[string]int{}, NoneByType: map[string]int{}}
	for _, r := range runs {
		rep.Runs = append(rep.Runs, UniqueRow{Label: r.Label, OnlyByLang: map[string]int{}, OnlyByType: map[string]int{},
			FoundByLang: map[string]int{}, FoundByType: map[string]int{}})
	}
	for _, e := range m.Entries {
		gt, ok := truth[e.ID]
		if !ok || gt.Scope.Kind != groundtruth.Pool {
			continue
		}
		rep.Entries = append(rep.Entries, e.ID)
		rep.Assets += len(gt.Assets)
		foundBy := map[string][]int{}
		for i, r := range runs {
			pf, ok := r.Preds[e.ID]
			if !ok {
				rep.Runs[i].EntriesNoRun = append(rep.Runs[i].EntriesNoRun, e.ID)
				continue
			}
			for _, p := range MatchEntry(gt, inFrame(e.Language, pf.Assets)).Pairs {
				foundBy[p.Truth.ID] = append(foundBy[p.Truth.ID], i)
			}
		}
		lang := string(e.Language)
		for _, a := range gt.Assets {
			typ := string(a.Type)
			by := foundBy[a.ID]
			for _, i := range by {
				row := &rep.Runs[i]
				row.Found++
				row.FoundByLang[lang]++
				row.FoundByType[typ]++
				if len(by) == 1 {
					row.Only++
					row.OnlyByLang[lang]++
					row.OnlyByType[typ]++
				}
			}
			if len(by) == 0 {
				rep.None++
				rep.NoneByLang[lang]++
				rep.NoneByType[typ]++
			}
		}
	}
	return rep
}

// UniqueJSON renders r as indented JSON.
func UniqueJSON(r UniqueReport) ([]byte, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	return append(data, '\n'), err
}

// UniqueMarkdown renders r as a markdown report.
func UniqueMarkdown(r UniqueReport) string {
	var b strings.Builder
	p := func(format string, args ...any) { fmt.Fprintf(&b, format, args...) }
	p("# Assets only one run found\n\n")
	p("Pool entries (%d): %s\n\n", len(r.Entries), list(r.Entries))
	p("Each run is matched against the pool ground truth by the score's rules. An asset counts for a run when the run matches it, and as found only by that run when no other run in this comparison matches it. Recall is relative to the pool: %d adjudicated assets, not every asset in the repositories.\n\n", r.Assets)
	p("| Run | Pool assets found | Found only by this run | Entries not run |\n|---|---|---|---|\n")
	for _, row := range r.Runs {
		p("| %s | %d | %d | %s |\n", row.Label, row.Found, row.Only, list(row.EntriesNoRun))
	}
	p("| found by none of these runs | | %d | |\n\n", r.None)
	breakdown := func(title string, only func(UniqueRow) map[string]int, none map[string]int) {
		keys := map[string]bool{}
		for k := range none {
			keys[k] = true
		}
		for _, row := range r.Runs {
			for k := range only(row) {
				keys[k] = true
			}
		}
		p("## Found only by one run, by %s\n\n| %s |", title, title)
		for _, row := range r.Runs {
			p(" %s |", row.Label)
		}
		p(" none of these runs |\n|---|")
		for range r.Runs {
			p("---|")
		}
		p("---|\n")
		for _, k := range slices.Sorted(maps.Keys(keys)) {
			p("| %s |", k)
			for _, row := range r.Runs {
				p(" %d |", only(row)[k])
			}
			p(" %d |\n", none[k])
		}
		p("\n")
	}
	breakdown("language", func(u UniqueRow) map[string]int { return u.OnlyByLang }, r.NoneByLang)
	breakdown("asset type", func(u UniqueRow) map[string]int { return u.OnlyByType }, r.NoneByType)
	return b.String()
}
