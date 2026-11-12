package score

import (
	"maps"
	"slices"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// NotDetected is the suite B outcome of a ground-truth asset that no
// prediction matched: the tool gave no answer because it found nothing.
const NotDetected prediction.Answer = "not-detected"

// SuiteBAnswers are the outcome columns of the suite B view.
var SuiteBAnswers = []prediction.Answer{prediction.Reachable, prediction.Unreachable, prediction.NoAnswer, NotDetected}

// SuiteBCounts are the suite B outcomes over one group of pair entries. Unlike
// the Reachability confusion, every ground-truth asset counts, matched or not.
type SuiteBCounts struct {
	// Outcome counts assets by ground-truth reachable (row) and outcome.
	Outcome map[bool]map[prediction.Answer]int
	// FalseReachable counts the reachable: false assets of unreachable
	// variants answered reachable; UnreachableDetected those detected at all.
	FalseReachable, UnreachableDetected int
	// Depth and Distance compare matched assets whose prediction states one.
	DepthChecked, DepthCorrect, DepthMissing int
	DistanceChecked, DistanceCorrect         int
	PairsCorrect, PairsScored                int
}

func newSuiteBCounts() *SuiteBCounts {
	return &SuiteBCounts{Outcome: map[bool]map[prediction.Answer]int{true: {}, false: {}}}
}

func (c *SuiteBCounts) add(o *SuiteBCounts) {
	for t, row := range o.Outcome {
		for a, n := range row {
			c.Outcome[t][a] += n
		}
	}
	c.FalseReachable += o.FalseReachable
	c.UnreachableDetected += o.UnreachableDetected
	c.DepthChecked += o.DepthChecked
	c.DepthCorrect += o.DepthCorrect
	c.DepthMissing += o.DepthMissing
	c.DistanceChecked += o.DistanceChecked
	c.DistanceCorrect += o.DistanceCorrect
	c.PairsCorrect += o.PairsCorrect
	c.PairsScored += o.PairsScored
}

func (c *SuiteBCounts) row(truth bool) int {
	n := 0
	for _, v := range c.Outcome[truth] {
		n += v
	}
	return n
}

// ReachablePrecision is reachable assets answered reachable over all assets answered reachable.
func (c *SuiteBCounts) ReachablePrecision() Ratio {
	tp := c.Outcome[true][prediction.Reachable]
	return ratio(tp, tp+c.Outcome[false][prediction.Reachable])
}

// ReachableRecall is reachable assets answered reachable over all reachable
// assets; an undetected asset and no answer both count as misses.
func (c *SuiteBCounts) ReachableRecall() Ratio {
	return ratio(c.Outcome[true][prediction.Reachable], c.row(true))
}

// FalseReachableRate is FalseReachable over UnreachableDetected.
func (c *SuiteBCounts) FalseReachableRate() Ratio {
	return ratio(c.FalseReachable, c.UnreachableDetected)
}

func (c *SuiteBCounts) DepthAccuracy() Ratio    { return ratio(c.DepthCorrect, c.DepthChecked) }
func (c *SuiteBCounts) DistanceAccuracy() Ratio { return ratio(c.DistanceCorrect, c.DistanceChecked) }

// SuiteBGroup is one row of a suite B breakdown.
type SuiteBGroup struct {
	Key string
	*SuiteBCounts
}

// SuiteB is the suite B view: reachability, depth and pair discrimination
// over the pair entries, by language and by scenario.
type SuiteB struct {
	// Reported is false when no scored pair entry's predictions report
	// reachability; the tool then has no reachability, which is not zero.
	Reported           bool
	Entries            int
	Overall            *SuiteBCounts
	ByLanguage         []SuiteBGroup
	ByScenario         []SuiteBGroup
	ByLanguageScenario []SuiteBGroup
}

// scoreSuiteB builds the suite B view, or nil when no pair entry was scored.
func scoreSuiteB(m corpus.Manifest, truth map[string]groundtruth.File, preds map[string]prediction.File, matches map[string]Match, pairs *Pairs) *SuiteB {
	out := &SuiteB{Overall: newSuiteBCounts()}
	dims := map[string]map[string]*SuiteBCounts{"language": {}, "scenario": {}, "language-scenario": {}}
	addTo := func(e corpus.Entry, c *SuiteBCounts) {
		out.Overall.add(c)
		for dim, key := range map[string]string{
			"language": string(e.Language), "scenario": e.Pair.Scenario,
			"language-scenario": string(e.Language) + " / " + e.Pair.Scenario,
		} {
			if dims[dim][key] == nil {
				dims[dim][key] = newSuiteBCounts()
			}
			dims[dim][key].add(c)
		}
	}
	byID := map[string]corpus.Entry{}
	for _, e := range m.Entries {
		byID[e.ID] = e
		mt, scored := matches[e.ID]
		if e.Pair == nil || !scored {
			continue
		}
		out.Entries++
		pf := preds[e.ID]
		out.Reported = out.Reported || pf.ReachabilityReported
		addTo(e, entrySuiteB(e, truth[e.ID], mt))
	}
	if out.Entries == 0 {
		return nil
	}
	if pairs != nil {
		for _, pr := range pairs.Results {
			c := newSuiteBCounts()
			c.PairsScored = 1
			if pr.Correct {
				c.PairsCorrect = 1
			}
			addTo(byID[pr.Reachable], c)
		}
	}
	groups := func(dim string) []SuiteBGroup {
		var g []SuiteBGroup
		for _, k := range slices.Sorted(maps.Keys(dims[dim])) {
			g = append(g, SuiteBGroup{k, dims[dim][k]})
		}
		return g
	}
	out.ByLanguage, out.ByScenario, out.ByLanguageScenario = groups("language"), groups("scenario"), groups("language-scenario")
	return out
}

func entrySuiteB(e corpus.Entry, gt groundtruth.File, mt Match) *SuiteBCounts {
	c := newSuiteBCounts()
	matched := map[string]prediction.Asset{}
	for _, p := range mt.Pairs {
		matched[p.Truth.ID] = p.Pred
	}
	for _, t := range gt.Assets {
		k, ok := t.Reachability.(groundtruth.Known)
		if !ok {
			continue
		}
		p, found := matched[t.ID]
		answer := NotDetected
		if found {
			answer = p.Reachable
		}
		c.Outcome[k.Reachable][answer]++
		if e.Pair.Variant == corpus.Unreachable && !k.Reachable && found {
			c.UnreachableDetected++
			if p.Reachable == prediction.Reachable {
				c.FalseReachable++
			}
		}
		if !found || p.Reachable == prediction.NoAnswer {
			continue
		}
		if p.Depth == nil {
			c.DepthMissing++
			continue
		}
		c.DepthChecked++
		if *p.Depth == k.Depth {
			c.DepthCorrect++
		}
		if p.Distance != nil && k.Distance >= 0 {
			c.DistanceChecked++
			if *p.Distance == k.Distance {
				c.DistanceCorrect++
			}
		}
	}
	return c
}
