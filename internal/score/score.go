// Package score matches one tool's predictions to the ground truth and
// computes detection, attribution, reachability and pair-discrimination metrics.
package score

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/frame"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// Counts are detection outcomes over a set of entries.
type Counts struct {
	TP, FP, FN, Duplicates int
	AttributedTP           int
	// Pooled is true when any counted entry has scope.kind pool, so recall is
	// relative to the pooled candidates.
	Pooled bool
}

// Ratio is a metric that is undefined (n/a) when its denominator is zero.
type Ratio struct {
	Value float64
	OK    bool
}

func ratio(num, den int) Ratio {
	if den == 0 {
		return Ratio{}
	}
	return Ratio{float64(num) / float64(den), true}
}

func (c Counts) Precision() Ratio { return ratio(c.TP, c.TP+c.FP) }
func (c Counts) Recall() Ratio    { return ratio(c.TP, c.TP+c.FN) }

// F1 is 2TP/(2TP+FP+FN), the harmonic mean of precision and recall wherever
// both are defined, and 0 rather than n/a when TP is 0 but FP or FN is not.
func (c Counts) F1() Ratio                  { return ratio(2*c.TP, 2*c.TP+c.FP+c.FN) }
func (c Counts) AttributedShare() Ratio     { return ratio(c.AttributedTP, c.TP) }
func (c Counts) AttributedPrecision() Ratio { return ratio(c.AttributedTP, c.TP+c.FP) }
func (c Counts) AttributedRecall() Ratio    { return ratio(c.AttributedTP, c.TP+c.FN) }
func (c *Counts) add(o Counts) {
	*c = Counts{c.TP + o.TP, c.FP + o.FP, c.FN + o.FN, c.Duplicates + o.Duplicates, c.AttributedTP + o.AttributedTP, c.Pooled || o.Pooled}
}
func (c Counts) total() int { return c.TP + c.FP + c.FN + c.Duplicates }

// Answer columns of a reachability confusion matrix.
var Answers = []prediction.Answer{prediction.Reachable, prediction.Unreachable, prediction.NoAnswer}

// Confusion counts matched pairs by ground-truth reachable (row) and the
// tool's answer (column). A missing answer is prediction.NoAnswer.
type Confusion map[bool]map[prediction.Answer]int

func (c Confusion) add(truth bool, got prediction.Answer) {
	if c[truth] == nil {
		c[truth] = map[prediction.Answer]int{}
	}
	c[truth][got]++
}

// ReachablePrecision is the share of pairs answered reachable that are reachable.
func (c Confusion) ReachablePrecision() Ratio {
	return ratio(c[true][prediction.Reachable], c[true][prediction.Reachable]+c[false][prediction.Reachable])
}

// ReachableRecall is the share of reachable pairs answered reachable; no answer counts as a miss.
func (c Confusion) ReachableRecall() Ratio {
	row := c[true]
	return ratio(row[prediction.Reachable], row[prediction.Reachable]+row[prediction.Unreachable]+row[prediction.NoAnswer])
}

// DepthBuckets label ground-truth depth 0, 1 and 2 (two or more hops).
var DepthBuckets = []string{"0", "1", "2+"}

type Reachability struct {
	Entries int // scored entries whose predictions report reachability
	Overall Confusion
	ByDepth map[string]Confusion
}

type Group struct {
	Key string
	Counts
}

type Note struct {
	Entry  string `json:"entry"`
	Reason string `json:"reason"`
}

type PairResult struct {
	Reachable, Unreachable string // entry ids of the two variants
	Correct                bool
	Why                    string // set when Correct is false
}

type Pairs struct {
	Results []PairResult
	Skipped []Note
}

func (p Pairs) Correct() int {
	n := 0
	for _, r := range p.Results {
		if r.Correct {
			n++
		}
	}
	return n
}

type Report struct {
	Tool    string
	Configs []prediction.Tool // distinct version and config found in the prediction files
	Scored  []string
	NotRun  []string
	Skipped []Note
	Ignored int // predictions outside scope.files
	// OutsideFrame are predictions of pool entries outside the entry's
	// language source files or in excluded paths (internal/frame), dropped
	// before matching.
	OutsideFrame int
	// NameConflicts are matched pairs whose algorithm name implied a
	// parameter that an explicit field contradicted; the field was used.
	NameConflicts []Note

	Overall        Counts
	BySuite        []Group
	ByLanguage     []Group
	ByType         []Group
	ByLanguageType []Group
	ByEntry        []Group

	Reachability *Reachability // nil when no scored entry reports reachability
	Pairs        *Pairs        // nil when no scored entry reports reachability
	SuiteB       *SuiteB       // nil when no suite B pair entry was scored

	Objects ObjectView
}

// ObjectView counts detection and attribution over ground-truth objects
// instead of assets (see objectCounts).
type ObjectView struct {
	Overall    Counts
	ByLanguage []Group
	ByType     []Group
	ByEntry    []Group
}

// Score computes the report for tool from validated inputs keyed by entry id.
func Score(tool string, m corpus.Manifest, truth map[string]groundtruth.File, preds map[string]prediction.File) Report {
	rep := Report{Tool: tool}
	reach := Reachability{Overall: Confusion{}, ByDepth: map[string]Confusion{}}
	for _, b := range DepthBuckets {
		reach.ByDepth[b] = Confusion{}
	}
	ops := newTally("suite", "language", "type", "language-type", "entry")
	objects := newTally("language", "type", "entry")
	matches := map[string]Match{}
	for _, e := range m.Entries {
		gt, hasTruth := truth[e.ID]
		pf, hasPred := preds[e.ID]
		switch {
		case !hasTruth:
			rep.Skipped = append(rep.Skipped, Note{e.ID, "no ground truth"})
			continue
		case !hasPred:
			rep.NotRun = append(rep.NotRun, e.ID)
			continue
		}
		rep.Scored = append(rep.Scored, e.ID)
		rep.Configs = addConfig(rep.Configs, pf.Tool)
		assets := pf.Assets
		if gt.Scope.Kind == groundtruth.Pool {
			assets = inFrame(e.Language, assets)
			rep.OutsideFrame += len(pf.Assets) - len(assets)
		}
		mt := MatchEntry(gt, assets)
		matches[e.ID] = mt
		rep.Ignored += mt.Ignored
		pooled := gt.Scope.Kind == groundtruth.Pool
		attributed := map[string]bool{}
		for _, p := range mt.Pairs {
			c := Counts{TP: 1, Pooled: pooled}
			ok, conflicts := Attributed(p)
			if ok {
				c.AttributedTP = 1
				attributed[p.Truth.ID] = true
			}
			for _, why := range conflicts {
				rep.NameConflicts = append(rep.NameConflicts, Note{e.ID, why})
			}
			ops.add(e, p.Truth.Type, c)
			if k, ok := p.Truth.Reachability.(groundtruth.Known); ok && pf.ReachabilityReported {
				reach.Overall.add(k.Reachable, p.Pred.Reachable)
				reach.ByDepth[DepthBuckets[k.Depth]].add(k.Reachable, p.Pred.Reachable)
			}
		}
		for _, p := range mt.FP {
			ops.add(e, p.Type, Counts{FP: 1, Pooled: pooled})
		}
		for _, p := range mt.Duplicates {
			ops.add(e, p.Type, Counts{Duplicates: 1, Pooled: pooled})
		}
		for _, t := range mt.FN {
			ops.add(e, t.Type, Counts{FN: 1, Pooled: pooled})
		}
		for _, o := range objectCounts(gt, mt, attributed) {
			o.Pooled = pooled
			objects.add(e, o.Type, o.Counts)
		}
		if pf.ReachabilityReported {
			reach.Entries++
		}
	}
	rep.Overall = ops.overall
	rep.BySuite, rep.ByLanguage = ops.groups("suite"), ops.groups("language")
	rep.ByType, rep.ByLanguageType = ops.groups("type"), ops.groups("language-type")
	rep.ByEntry = ops.groups("entry")
	rep.Objects = ObjectView{objects.overall, objects.groups("language"), objects.groups("type"), objects.groups("entry")}
	if reach.Entries > 0 {
		rep.Reachability = &reach
		pairs := scorePairs(m, truth, preds, matches)
		rep.Pairs = &pairs
	}
	rep.SuiteB = scoreSuiteB(m, truth, preds, matches, rep.Pairs)
	return rep
}

// inFrame keeps the predictions in source files of lang outside excluded
// paths, the files a pool entry covers (METHODOLOGY.md, "Suite C decisions").
func inFrame(lang corpus.Language, preds []prediction.Asset) []prediction.Asset {
	var out []prediction.Asset
	for _, p := range preds {
		if frame.InScope(lang, p.File) {
			out = append(out, p)
		}
	}
	return out
}

func addConfig(configs []prediction.Tool, t prediction.Tool) []prediction.Tool {
	for _, c := range configs {
		if c.Name == t.Name && c.Version == t.Version && maps.Equal(c.Config, t.Config) {
			return configs
		}
	}
	return append(configs, t)
}

// tally adds counts to an overall total and to one group per breakdown
// dimension it was created with.
type tally struct {
	overall Counts
	dims    map[string]map[string]*Counts
}

func newTally(dims ...string) *tally {
	t := &tally{dims: map[string]map[string]*Counts{}}
	for _, d := range dims {
		t.dims[d] = map[string]*Counts{}
	}
	return t
}

func (t *tally) add(e corpus.Entry, typ groundtruth.AssetType, c Counts) {
	for dim, key := range map[string]string{
		"suite": string(e.Suite), "language": string(e.Language), "type": string(typ),
		"language-type": string(e.Language) + " / " + string(typ), "entry": e.ID,
	} {
		g, ok := t.dims[dim]
		if !ok {
			continue
		}
		if g[key] == nil {
			g[key] = &Counts{}
		}
		g[key].add(c)
	}
	t.overall.add(c)
}

func (t *tally) groups(dim string) []Group {
	g := t.dims[dim]
	out := make([]Group, 0, len(g))
	for _, k := range slices.Sorted(maps.Keys(g)) {
		if g[k].total() > 0 {
			out = append(out, Group{k, *g[k]})
		}
	}
	return out
}

// scorePairs judges every suite B pair. Only the assets that make the pair a
// pair count: the reachable variant's ground truth annotated reachable=true,
// which must all be matched and answered reachable, and the unreachable
// variant's ground truth annotated reachable=false, none of which may be
// matched and answered reachable.
func scorePairs(m corpus.Manifest, truth map[string]groundtruth.File, preds map[string]prediction.File, matches map[string]Match) Pairs {
	var out Pairs
	for _, e := range m.Entries {
		if e.Pair == nil || e.Pair.Variant != corpus.Reachable {
			continue
		}
		r, u := e.ID, e.Pair.Partner
		if why := pairUnscorable(r, u, truth, preds, matches); why != "" {
			out.Skipped = append(out.Skipped, Note{r + " / " + u, why})
			continue
		}
		answered := map[string]prediction.Answer{}
		for _, p := range matches[r].Pairs {
			answered[p.Truth.ID] = p.Pred.Reachable
		}
		var wrong []string
		for _, t := range assetsWithReachable(truth[r], true) {
			if answered[t.ID] != prediction.Reachable {
				wrong = append(wrong, fmt.Sprintf("%s: %s not found reachable", r, t.ID))
			}
		}
		for _, p := range matches[u].Pairs {
			if k, ok := p.Truth.Reachability.(groundtruth.Known); ok && !k.Reachable && p.Pred.Reachable == prediction.Reachable {
				wrong = append(wrong, fmt.Sprintf("%s: %s reported reachable", u, p.Truth.ID))
			}
		}
		out.Results = append(out.Results, PairResult{r, u, len(wrong) == 0, strings.Join(wrong, "; ")})
	}
	return out
}

func pairUnscorable(r, u string, truth map[string]groundtruth.File, preds map[string]prediction.File, matches map[string]Match) string {
	for _, id := range []string{r, u} {
		switch {
		case truth[id].Entry == "":
			return id + " has no ground truth"
		case !hasKey(matches, id):
			return id + " was not run"
		case !preds[id].ReachabilityReported:
			return id + " predictions do not report reachability"
		}
	}
	switch {
	case len(assetsWithReachable(truth[r], true)) == 0:
		return r + " has no ground-truth asset with reachable: true"
	case len(assetsWithReachable(truth[u], false)) == 0:
		return u + " has no ground-truth asset with reachable: false"
	}
	return ""
}

func hasKey[V any](m map[string]V, k string) bool { _, ok := m[k]; return ok }

func assetsWithReachable(f groundtruth.File, want bool) []groundtruth.Asset {
	var out []groundtruth.Asset
	for _, a := range f.Assets {
		if k, ok := a.Reachability.(groundtruth.Known); ok && k.Reachable == want {
			out = append(out, a)
		}
	}
	return out
}
