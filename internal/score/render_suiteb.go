package score

import (
	"fmt"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

var suiteBLabels = map[prediction.Answer]string{
	prediction.Reachable: "reachable", prediction.Unreachable: "unreachable",
	prediction.NoAnswer: "no answer", NotDetected: "not detected",
}

// markdownSuiteB writes the suite B section of the markdown report.
func markdownSuiteB(b *strings.Builder, s *SuiteB) {
	p := func(format string, args ...any) { fmt.Fprintf(b, format, args...) }
	p("\n## Suite B reachability\n\n")
	if s == nil {
		p("n/a: no suite B pair entry was scored.\n")
		return
	}
	if !s.Reported {
		p("No reachability: none of the %d scored pair entries' predictions report reachability. The tool is not scored on reachability, depth or pairs; the outcome table below only shows which assets it detected.\n", s.Entries)
	} else {
		p("Over every ground-truth asset of the %d scored pair entries, matched or not. An asset no prediction matched is `not detected` and counts as a miss for reachable recall. Reachable precision is reachable assets answered reachable over all assets answered reachable. False reachable counts the `reachable: false` assets of unreachable variants answered reachable, over those detected. Depth (0, 1, 2+) and distance (exact) compare matched assets the tool answered with a depth.\n", s.Entries)
	}
	for _, sec := range []struct {
		title  string
		groups []SuiteBGroup
	}{{"Overall", []SuiteBGroup{{"all", s.Overall}}}, {"By language", s.ByLanguage}, {"By scenario", s.ByScenario}, {"By language and scenario", s.ByLanguageScenario}} {
		p("\n### %s\n\n", sec.title)
		p("| Group | Truth | Assets |")
		for _, a := range SuiteBAnswers {
			p(" %s |", suiteBLabels[a])
		}
		if s.Reported {
			p(" Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs |")
		}
		p("\n|---|---|---|")
		for range SuiteBAnswers {
			p("---|")
		}
		if s.Reported {
			p("---|---|---|---|---|---|")
		}
		p("\n")
		for _, g := range sec.groups {
			for _, truth := range []bool{true, false} {
				label := "reachable"
				if !truth {
					label = "unreachable"
				}
				p("| %s | %s | %d |", g.Key, label, g.row(truth))
				for _, a := range SuiteBAnswers {
					p(" %d |", g.Outcome[truth][a])
				}
				if s.Reported {
					if truth {
						missing := ""
						if g.DepthMissing > 0 {
							missing = fmt.Sprintf(", %d without depth", g.DepthMissing)
						}
						p(" %s | %s | %d / %d (%s) | %d / %d (%s%s) | %d / %d (%s) | %d / %d |",
							g.ReachablePrecision(), g.ReachableRecall(),
							g.FalseReachable, g.UnreachableDetected, g.FalseReachableRate(),
							g.DepthCorrect, g.DepthChecked, g.DepthAccuracy(), missing,
							g.DistanceCorrect, g.DistanceChecked, g.DistanceAccuracy(),
							g.PairsCorrect, g.PairsScored)
					} else {
						p(" | | | | | |")
					}
				}
				p("\n")
			}
		}
	}
}

type jsonSuiteBCounts struct {
	Key                 string                    `json:"key,omitempty"`
	Outcome             map[string]map[string]int `json:"outcome"`
	ReachablePrecision  Ratio                     `json:"reachable_precision"`
	ReachableRecall     Ratio                     `json:"reachable_recall"`
	FalseReachable      int                       `json:"false_reachable"`
	UnreachableDetected int                       `json:"unreachable_detected"`
	FalseReachableRate  Ratio                     `json:"false_reachable_rate"`
	DepthCorrect        int                       `json:"depth_correct"`
	DepthChecked        int                       `json:"depth_checked"`
	DepthMissing        int                       `json:"depth_missing"`
	DepthAccuracy       Ratio                     `json:"depth_accuracy"`
	DistanceCorrect     int                       `json:"distance_correct"`
	DistanceChecked     int                       `json:"distance_checked"`
	DistanceAccuracy    Ratio                     `json:"distance_accuracy"`
	PairsCorrect        int                       `json:"pairs_correct"`
	PairsScored         int                       `json:"pairs_scored"`
}

type jsonSuiteB struct {
	Reported           bool               `json:"reachability_reported"`
	Entries            int                `json:"entries"`
	Overall            jsonSuiteBCounts   `json:"overall"`
	ByLanguage         []jsonSuiteBCounts `json:"by_language"`
	ByScenario         []jsonSuiteBCounts `json:"by_scenario"`
	ByLanguageScenario []jsonSuiteBCounts `json:"by_language_scenario"`
}

func toJSONSuiteBCounts(key string, c *SuiteBCounts) jsonSuiteBCounts {
	out := map[string]map[string]int{}
	for truth, label := range map[bool]string{true: "reachable", false: "unreachable"} {
		out[label] = map[string]int{}
		for _, a := range SuiteBAnswers {
			out[label][string(a)] = c.Outcome[truth][a]
		}
	}
	return jsonSuiteBCounts{key, out, c.ReachablePrecision(), c.ReachableRecall(),
		c.FalseReachable, c.UnreachableDetected, c.FalseReachableRate(),
		c.DepthCorrect, c.DepthChecked, c.DepthMissing, c.DepthAccuracy(),
		c.DistanceCorrect, c.DistanceChecked, c.DistanceAccuracy(), c.PairsCorrect, c.PairsScored}
}

func toJSONSuiteB(s *SuiteB) *jsonSuiteB {
	if s == nil {
		return nil
	}
	groups := func(gs []SuiteBGroup) []jsonSuiteBCounts {
		out := []jsonSuiteBCounts{}
		for _, g := range gs {
			out = append(out, toJSONSuiteBCounts(g.Key, g.SuiteBCounts))
		}
		return out
	}
	return &jsonSuiteB{s.Reported, s.Entries, toJSONSuiteBCounts("", s.Overall),
		groups(s.ByLanguage), groups(s.ByScenario), groups(s.ByLanguageScenario)}
}
