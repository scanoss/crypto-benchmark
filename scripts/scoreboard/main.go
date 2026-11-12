// Command scoreboard scores every published pilot under results/ and prints
// the consolidated tables of results/README.md.
//
//	go run ./scripts/scoreboard > tables.md
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/score"
)

var (
	tools     = []string{"crypto-finder", "sonar-cryptography", "cdxgen"}
	languages = []string{"python", "javascript", "rust", "go", "java"}
)

// run is one tool's score on one language pilot.
type run struct {
	tool, lang string
	rep        score.Report
}

// pair holds the operation-level and object-level counts of one slice.
type pair struct{ op, obj score.Counts }

func main() {
	runs, err := load(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, "scoreboard:", err)
		os.Exit(1)
	}
	write(os.Stdout, runs)
}

func load(root string) ([]run, error) {
	var runs []run
	for _, lang := range languages {
		for _, tool := range tools {
			dir := filepath.Join(root, "results", tool, lang+"-pilot", "predictions")
			if _, err := os.Stat(filepath.Join(dir, tool)); err != nil {
				continue
			}
			in, r := score.Load(root, "", dir, tool)
			if r.HasErrors() {
				return nil, fmt.Errorf("%s %s: invalid inputs", tool, lang)
			}
			runs = append(runs, run{tool, lang, score.Score(tool, in.Manifest, in.Truth, in.Preds)})
		}
	}
	return runs, nil
}

func sum(a, b score.Counts) score.Counts {
	return score.Counts{TP: a.TP + b.TP, FP: a.FP + b.FP, FN: a.FN + b.FN, Duplicates: a.Duplicates + b.Duplicates, AttributedTP: a.AttributedTP + b.AttributedTP}
}

func byType(r score.Report) map[string]pair {
	out := map[string]pair{}
	for _, g := range r.ByType {
		p := out[g.Key]
		p.op = g.Counts
		out[g.Key] = p
	}
	for _, g := range r.Objects.ByType {
		p := out[g.Key]
		p.obj = g.Counts
		out[g.Key] = p
	}
	return out
}

func prf(c score.Counts) string {
	return fmt.Sprintf("%s | %s | %s", c.Precision(), c.Recall(), c.F1())
}

func share(c score.Counts) string {
	return fmt.Sprintf("%d / %d (%s)", c.AttributedTP, c.TP, c.AttributedShare())
}

func write(w io.Writer, runs []run) {
	p := func(format string, a ...any) { fmt.Fprintf(w, format, a...) }

	p("### Language × tool\n\n")
	p("| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |\n")
	p("|---|---|---:|---:|---:|---:|---:|---:|---|---|---|\n")
	for _, r := range runs {
		op, obj := r.rep.Overall, r.rep.Objects.Overall
		p("| %s | %s | %s | %s | %d / %d / %d | %s | %s |\n", r.lang, r.tool, prf(op), prf(obj), op.TP, op.FP, op.FN, share(op), share(obj))
	}

	p("\n### Asset type × tool, all languages\n\n")
	p("Each tool's rows sum the pilots it ran on.\n\n")
	p("| Asset type | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op |\n")
	p("|---|---|---:|---:|---:|---:|---:|---:|---|---|\n")
	totals := map[string]map[string]pair{}
	var types []string
	for _, r := range runs {
		if totals[r.tool] == nil {
			totals[r.tool] = map[string]pair{}
		}
		for typ, c := range byType(r.rep) {
			t := totals[r.tool][typ]
			totals[r.tool][typ] = pair{sum(t.op, c.op), sum(t.obj, c.obj)}
			if !slices.Contains(types, typ) {
				types = append(types, typ)
			}
		}
	}
	slices.Sort(types)
	for _, typ := range types {
		for _, tool := range tools {
			c, ok := totals[tool][typ]
			if !ok {
				continue
			}
			p("| %s | %s | %s | %s | %d / %d / %d | %s |\n", typ, tool, prf(c.op), prf(c.obj), c.op.TP, c.op.FP, c.op.FN, share(c.op))
		}
	}

	p("\n### crypto-finder, language × asset type\n\n")
	p("| Language | Asset type | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op |\n")
	p("|---|---|---:|---:|---:|---:|---:|---:|---|---|\n")
	for _, r := range runs {
		if r.tool != "crypto-finder" {
			continue
		}
		bt := byType(r.rep)
		keys := make([]string, 0, len(bt))
		for k := range bt {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, typ := range keys {
			c := bt[typ]
			p("| %s | %s | %s | %s | %d / %d / %d | %s |\n", r.lang, typ, prf(c.op), prf(c.obj), c.op.TP, c.op.FP, c.op.FN, share(c.op))
		}
	}

	writeTargets(w, runs)

	p("\n### Tool runs\n\n")
	for _, r := range runs {
		var cfg []string
		for _, c := range r.rep.Configs {
			cfg = append(cfg, strings.TrimSpace(c.Name+" "+c.Version))
		}
		p("- %s %s: %s, %d entries\n", r.lang, r.tool, strings.Join(slices.Compact(cfg), "; "), len(r.rep.Scored))
	}
}

// The pre-registered targets (METHODOLOGY.md, Targets): precision >= 0.90
// and recall >= 0.85, compared in integers so no float rounding decides a cell.
func meetsPrecision(c score.Counts) bool { return c.TP+c.FP == 0 || 10*c.TP >= 9*(c.TP+c.FP) }
func meetsRecall(c score.Counts) bool    { return c.TP+c.FN == 0 || 100*c.TP >= 85*(c.TP+c.FN) }

// target renders one cell: precision / recall and pass, fail, or n/a when
// the cell has neither ground truth nor predictions. An undefined ratio
// counts as met: no predictions cannot miss precision, no ground truth
// cannot miss recall.
func target(c score.Counts) (string, bool, bool) {
	if c.TP+c.FP+c.FN == 0 {
		return "n/a", false, false
	}
	verdict := "fail"
	pass := meetsPrecision(c) && meetsRecall(c)
	if pass {
		verdict = "pass"
	}
	return fmt.Sprintf("%s / %s %s", c.Precision(), c.Recall(), verdict), pass, true
}

// writeTargets prints, for each tool and view, the language x asset type
// table of the targets, with a row and a column that sum the other cells.
func writeTargets(w io.Writer, runs []run) {
	p := func(format string, a ...any) { fmt.Fprintf(w, format, a...) }
	types := make([]string, len(groundtruth.AssetTypes))
	for i, t := range groundtruth.AssetTypes {
		types[i] = string(t)
	}
	views := []struct {
		name string
		pick func(pair) score.Counts
	}{
		{"object level (headline)", func(c pair) score.Counts { return c.obj }},
		{"operation level", func(c pair) score.Counts { return c.op }},
	}
	p("\n### Targets\n\n")
	p("Precision >= 0.90 and recall >= 0.85 in every cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). Each cell is `precision / recall` and its verdict; `n/a` has neither ground truth nor predictions.\n")
	for _, tool := range tools {
		var mine []run
		for _, r := range runs {
			if r.tool == tool {
				mine = append(mine, r)
			}
		}
		if len(mine) == 0 {
			continue
		}
		for _, v := range views {
			p("\n#### %s, %s\n\n", tool, v.name)
			p("| Language | %s | all types |\n", strings.Join(types, " | "))
			p("|---|%s---|\n", strings.Repeat("---|", len(types)))
			colTotal := map[string]score.Counts{}
			var allTotal score.Counts
			passed, scored := 0, 0
			cell := func(c score.Counts) string {
				s, pass, ok := target(c)
				if ok {
					scored++
					if pass {
						passed++
					}
				}
				return s
			}
			for _, r := range mine {
				bt := byType(r.rep)
				var row []string
				var rowTotal score.Counts
				for _, typ := range types {
					c := v.pick(bt[typ])
					rowTotal = sum(rowTotal, c)
					colTotal[typ] = sum(colTotal[typ], c)
					row = append(row, cell(c))
				}
				allTotal = sum(allTotal, rowTotal)
				p("| %s | %s | %s |\n", r.lang, strings.Join(row, " | "), cell(rowTotal))
			}
			var row []string
			for _, typ := range types {
				row = append(row, cell(colTotal[typ]))
			}
			p("| all languages | %s | %s |\n", strings.Join(row, " | "), cell(allTotal))
			p("\n%d of %d cells pass.\n", passed, scored)
		}
	}
}
