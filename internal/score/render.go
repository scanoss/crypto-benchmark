package score

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

func (r Ratio) String() string {
	if !r.OK {
		return "n/a"
	}
	return fmt.Sprintf("%.3f", r.Value)
}

func (r Ratio) MarshalJSON() ([]byte, error) {
	if !r.OK {
		return []byte("null"), nil
	}
	return json.Marshal(r.Value)
}

var answerLabel = map[prediction.Answer]string{
	prediction.Reachable: "reachable", prediction.Unreachable: "unreachable", prediction.NoAnswer: "no answer",
}

func configLine(t prediction.Tool) string {
	var kv []string
	for _, k := range slices.Sorted(maps.Keys(t.Config)) {
		kv = append(kv, k+"="+t.Config[k])
	}
	if len(kv) == 0 {
		return "(none)"
	}
	return strings.Join(kv, ", ")
}

func list(ids []string) string {
	if len(ids) == 0 {
		return "none"
	}
	return strings.Join(ids, ", ")
}

// Markdown renders r as a markdown report.
func Markdown(r Report) string {
	var b strings.Builder
	p := func(format string, args ...any) { fmt.Fprintf(&b, format, args...) }
	p("# Score: %s\n\n", r.Tool)
	if len(r.Configs) > 1 {
		p("Predictions come from %d tool configurations:\n\n", len(r.Configs))
	}
	for _, c := range r.Configs {
		p("- Tool: %s %s\n- Config: %s\n", c.Name, c.Version, configLine(c))
	}
	p("\nEntries scored (%d): %s\n\n", len(r.Scored), list(r.Scored))
	p("Not run, no prediction file (%d): %s\n\n", len(r.NotRun), list(r.NotRun))
	for _, n := range r.Skipped {
		p("Skipped %s: %s\n\n", n.Entry, n.Reason)
	}
	if r.Ignored > 0 {
		p("Predictions outside scope.files or scope.components, ignored: %d\n\n", r.Ignored)
	}
	if r.OutsideFrame > 0 {
		p("Predictions in pool entries outside the entry's language source files or in excluded paths, ignored: %d\n\n", r.OutsideFrame)
	}

	p("## Detection\n\nCandidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = %d). Ties prefer an attributed pair.\n", LineTolerance)
	if r.Overall.Pooled {
		p("Rows marked † include entries with scope.kind pool; their recall is relative to pool.\n")
	}
	detection := []string{"Precision", "Recall", "F1", "TP", "FP", "FN", "Duplicates"}
	detRow := func(c Counts) []string {
		return []string{c.Precision().String(), c.Recall().String(), c.F1().String(),
			fmt.Sprint(c.TP), fmt.Sprint(c.FP), fmt.Sprint(c.FN), fmt.Sprint(c.Duplicates)}
	}
	sections := []struct {
		title  string
		groups []Group
	}{
		{"Overall", []Group{{"all", r.Overall}}}, {"By suite", r.BySuite}, {"By language", r.ByLanguage},
		{"By asset type", r.ByType}, {"By language × asset type", r.ByLanguageType}, {"By entry", r.ByEntry},
	}
	for _, s := range sections {
		p("\n### %s\n\n", s.title)
		table(&b, s.groups, detection, detRow)
	}

	p("\n## Attribution\n\nA matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).\n")
	for _, n := range r.NameConflicts {
		p("\nName and field disagree in %s, field used: %s\n", n.Entry, n.Reason)
	}
	attribution := []string{"Attributed TP / TP", "Attributed precision", "Attributed recall"}
	attrRow := func(c Counts) []string {
		return []string{fmt.Sprintf("%d / %d (%s)", c.AttributedTP, c.TP, c.AttributedShare()),
			c.AttributedPrecision().String(), c.AttributedRecall().String()}
	}
	for _, s := range append(sections[:4:4], sections[5]) {
		p("\n### %s\n\n", s.title)
		table(&b, s.groups, attribution, attrRow)
	}

	p("\n## Reachability\n\n")
	if r.Reachability == nil {
		p("n/a: no scored entry's predictions report reachability.\n")
	} else {
		rc := r.Reachability
		p("Over matched pairs whose ground truth reachable is not unknown, in %d of %d scored entries.\n\n### Overall\n\n", rc.Entries, len(r.Scored))
		confusion(&b, rc.Overall)
		for _, d := range DepthBuckets {
			p("\n### Ground-truth depth %s\n\n", d)
			confusion(&b, rc.ByDepth[d])
		}
	}

	p("\n## Suite B pair discrimination\n\n")
	if r.Pairs == nil {
		p("n/a: no scored entry's predictions report reachability.\n")
	} else {
		p("Correct pairs: %d / %d (%s)\n", r.Pairs.Correct(), len(r.Pairs.Results), ratio(r.Pairs.Correct(), len(r.Pairs.Results)))
		if len(r.Pairs.Results) > 0 {
			p("\n| Reachable variant | Unreachable variant | Result | Why |\n|---|---|---|---|\n")
			for _, pr := range r.Pairs.Results {
				result := "correct"
				if !pr.Correct {
					result = "wrong"
				}
				p("| %s | %s | %s | %s |\n", pr.Reachable, pr.Unreachable, result, pr.Why)
			}
		}
		for _, n := range r.Pairs.Skipped {
			p("\nSkipped pair %s: %s\n", n.Entry, n.Reason)
		}
	}

	markdownSuiteB(&b, r.SuiteB)

	p("\n## Object-level view\n\nEach ground-truth object counts once: the assets that share an `object` id, or one asset without one. An object is detected when any of its assets is matched above, and attributed when any of its matched assets is attributed. A prediction matched to an object already detected is a duplicate; FPs are the same as above.\n")
	objCols := append(slices.Clone(detection), attribution...)
	objRow := func(c Counts) []string { return append(detRow(c), attrRow(c)...) }
	o := r.Objects
	for _, s := range []struct {
		title  string
		groups []Group
	}{{"Overall", []Group{{"all", o.Overall}}}, {"By language", o.ByLanguage}, {"By asset type", o.ByType}, {"By entry", o.ByEntry}} {
		p("\n### %s\n\n", s.title)
		table(&b, s.groups, objCols, objRow)
	}
	return b.String()
}

func table(b *strings.Builder, groups []Group, cols []string, row func(Counts) []string) {
	if len(groups) == 0 {
		b.WriteString("n/a: nothing scored.\n")
		return
	}
	fmt.Fprintf(b, "| Group | %s |\n|---|%s\n", strings.Join(cols, " | "), strings.Repeat("---|", len(cols)))
	for _, g := range groups {
		key := g.Key
		if g.Pooled {
			key += " †"
		}
		fmt.Fprintf(b, "| %s | %s |\n", key, strings.Join(row(g.Counts), " | "))
	}
}

func confusion(b *strings.Builder, c Confusion) {
	b.WriteString("| Ground truth \\ Predicted | reachable | unreachable | no answer |\n|---|---|---|---|\n")
	for _, truth := range []bool{true, false} {
		label := "reachable"
		if !truth {
			label = "unreachable"
		}
		fmt.Fprintf(b, "| %s | %d | %d | %d |\n", label, c[truth][prediction.Reachable], c[truth][prediction.Unreachable], c[truth][prediction.NoAnswer])
	}
	fmt.Fprintf(b, "\nReachable precision %s, reachable recall %s.\n", c.ReachablePrecision(), c.ReachableRecall())
}

type jsonCounts struct {
	Group                string `json:"group,omitempty"`
	TP                   int    `json:"tp"`
	FP                   int    `json:"fp"`
	FN                   int    `json:"fn"`
	Duplicates           int    `json:"duplicates"`
	Precision            Ratio  `json:"precision"`
	Recall               Ratio  `json:"recall"`
	RecallRelativeToPool bool   `json:"recall_relative_to_pool"`
	F1                   Ratio  `json:"f1"`
	AttributedTP         int    `json:"attributed_tp"`
	AttributedShare      Ratio  `json:"attributed_share"`
	AttributedPrecision  Ratio  `json:"attributed_precision"`
	AttributedRecall     Ratio  `json:"attributed_recall"`
}

func toJSONCounts(key string, c Counts) jsonCounts {
	return jsonCounts{key, c.TP, c.FP, c.FN, c.Duplicates, c.Precision(), c.Recall(), c.Pooled, c.F1(),
		c.AttributedTP, c.AttributedShare(), c.AttributedPrecision(), c.AttributedRecall()}
}

func toJSONGroups(gs []Group) []jsonCounts {
	out := []jsonCounts{}
	for _, g := range gs {
		out = append(out, toJSONCounts(g.Key, g.Counts))
	}
	return out
}

type jsonConfusion struct {
	Matrix             map[string]map[string]int `json:"matrix"` // ground truth -> predicted -> pairs
	ReachablePrecision Ratio                     `json:"reachable_precision"`
	ReachableRecall    Ratio                     `json:"reachable_recall"`
}

func toJSONConfusion(c Confusion) jsonConfusion {
	m := map[string]map[string]int{}
	for truth, label := range map[bool]string{true: "reachable", false: "unreachable"} {
		m[label] = map[string]int{}
		for _, a := range Answers {
			m[label][answerLabel[a]] = c[truth][a]
		}
	}
	return jsonConfusion{m, c.ReachablePrecision(), c.ReachableRecall()}
}

type objectLevel struct {
	Overall    jsonCounts   `json:"overall"`
	ByLanguage []jsonCounts `json:"by_language"`
	ByType     []jsonCounts `json:"by_asset_type"`
	ByEntry    []jsonCounts `json:"by_entry"`
}

// JSON renders r as an indented JSON report. Undefined metrics are null.
func JSON(r Report) ([]byte, error) {
	type tool struct {
		Name    string            `json:"name"`
		Version string            `json:"version"`
		Config  map[string]string `json:"config"`
	}
	type reach struct {
		Entries int                      `json:"entries"`
		Overall jsonConfusion            `json:"overall"`
		ByDepth map[string]jsonConfusion `json:"by_depth"`
	}
	type pair struct {
		Reachable   string `json:"reachable"`
		Unreachable string `json:"unreachable"`
		Correct     bool   `json:"correct"`
		Why         string `json:"why,omitempty"`
	}
	type pairs struct {
		Correct int    `json:"correct"`
		Total   int    `json:"total"`
		Results []pair `json:"results"`
		Skipped []Note `json:"skipped"`
	}
	out := struct {
		Tool           string       `json:"tool"`
		Configs        []tool       `json:"configs"`
		Scored         []string     `json:"scored"`
		NotRun         []string     `json:"not_run"`
		Skipped        []Note       `json:"skipped"`
		Ignored        int          `json:"ignored_out_of_scope"`
		OutsideFrame   int          `json:"ignored_outside_pool_frame"`
		NameConflicts  []Note       `json:"name_conflicts"`
		LineTolerance  int          `json:"line_tolerance"`
		Overall        jsonCounts   `json:"overall"`
		BySuite        []jsonCounts `json:"by_suite"`
		ByLanguage     []jsonCounts `json:"by_language"`
		ByType         []jsonCounts `json:"by_asset_type"`
		ByLanguageType []jsonCounts `json:"by_language_asset_type"`
		ByEntry        []jsonCounts `json:"by_entry"`
		Reachability   *reach       `json:"reachability"`
		Pairs          *pairs       `json:"pair_discrimination"`
		SuiteB         *jsonSuiteB  `json:"suite_b"`
		ObjectLevel    objectLevel  `json:"object_level"`
	}{
		Tool: r.Tool, Scored: nonNil(r.Scored), NotRun: nonNil(r.NotRun), Skipped: nonNil(r.Skipped),
		Ignored: r.Ignored, OutsideFrame: r.OutsideFrame, NameConflicts: nonNil(r.NameConflicts), LineTolerance: LineTolerance, Overall: toJSONCounts("", r.Overall),
		BySuite: toJSONGroups(r.BySuite), ByLanguage: toJSONGroups(r.ByLanguage),
		ByType: toJSONGroups(r.ByType), ByLanguageType: toJSONGroups(r.ByLanguageType), ByEntry: toJSONGroups(r.ByEntry),
		ObjectLevel: objectLevel{toJSONCounts("", r.Objects.Overall), toJSONGroups(r.Objects.ByLanguage),
			toJSONGroups(r.Objects.ByType), toJSONGroups(r.Objects.ByEntry)},
		SuiteB: toJSONSuiteB(r.SuiteB),
	}
	for _, c := range r.Configs {
		out.Configs = append(out.Configs, tool{c.Name, c.Version, c.Config})
	}
	if rc := r.Reachability; rc != nil {
		out.Reachability = &reach{rc.Entries, toJSONConfusion(rc.Overall), map[string]jsonConfusion{}}
		for _, d := range DepthBuckets {
			out.Reachability.ByDepth[d] = toJSONConfusion(rc.ByDepth[d])
		}
	}
	if pr := r.Pairs; pr != nil {
		out.Pairs = &pairs{Correct: pr.Correct(), Total: len(pr.Results), Results: []pair{}, Skipped: nonNil(pr.Skipped)}
		for _, x := range pr.Results {
			out.Pairs.Results = append(out.Pairs.Results, pair{x.Reachable, x.Unreachable, x.Correct, x.Why})
		}
	}
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("score: %w", err)
	}
	return append(data, '\n'), nil
}

func nonNil[T any](s []T) []T {
	if s == nil {
		return []T{}
	}
	return s
}
