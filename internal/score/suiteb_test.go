package score

import (
	"fmt"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

const libB = "pkg:golang/example.com/libb@v1.0.0"

func depAsset(id string, line int, reachable bool, distance int) groundtruth.Asset {
	a := truthAsset(id, "libb.go", line, line)
	a.Origin = groundtruth.Dependency{Component: libB}
	a.Reachability = groundtruth.Known{Reachable: reachable, Depth: groundtruth.Depth(min(distance, 2)), Distance: distance}
	return a
}

func depPred(line int, answer prediction.Answer, distance int) prediction.Asset {
	p := predAsset("libb.go", line, line)
	p.Origin = groundtruth.Dependency{Component: libB}
	p.Reachable = answer
	if answer != prediction.NoAnswer {
		d, n := groundtruth.Depth(min(distance, 2)), distance
		p.Depth, p.Distance = &d, &n
	}
	return p
}

func pairEntries(scenario string) (corpus.Entry, corpus.Entry) {
	r := corpus.Entry{ID: scenario + "-r", Language: corpus.Go, Suite: corpus.SuiteB, Pair: &corpus.Pair{Partner: scenario + "-u", Variant: corpus.Reachable, Scenario: scenario}}
	u := corpus.Entry{ID: scenario + "-u", Language: corpus.Go, Suite: corpus.SuiteB, Pair: &corpus.Pair{Partner: scenario + "-r", Variant: corpus.Unreachable, Scenario: scenario}}
	return r, u
}

func predFile(entry string, reported bool, assets ...prediction.Asset) prediction.File {
	return prediction.File{Tool: prediction.Tool{Name: "t"}, Entry: entry, ReachabilityReported: reported, Assets: assets}
}

// suiteBLine renders one suite B row as the counts the tests assert on.
func suiteBLine(c *SuiteBCounts) string {
	return fmt.Sprintf("R:%d/%d/%d/%d U:%d/%d/%d/%d P=%s R=%s false=%d/%d depth=%d/%d dist=%d/%d pairs=%d/%d",
		c.Outcome[true][prediction.Reachable], c.Outcome[true][prediction.Unreachable], c.Outcome[true][prediction.NoAnswer], c.Outcome[true][NotDetected],
		c.Outcome[false][prediction.Reachable], c.Outcome[false][prediction.Unreachable], c.Outcome[false][prediction.NoAnswer], c.Outcome[false][NotDetected],
		c.ReachablePrecision(), c.ReachableRecall(), c.FalseReachable, c.UnreachableDetected,
		c.DepthCorrect, c.DepthChecked, c.DistanceCorrect, c.DistanceChecked, c.PairsCorrect, c.PairsScored)
}

func TestSuiteBCountsEveryAssetAndBreaksDownByScenario(t *testing.T) {
	tr, tu := pairEntries("transitive")
	dr, du := pairEntries("depth3")
	truth := map[string]groundtruth.File{
		// transitive: one used asset and one dead asset in the reachable
		// variant; both dead in the unreachable variant.
		tr.ID: exhaustive(depAsset("gt-1", 5, true, 2), depAsset("gt-2", 9, false, 2)),
		tu.ID: exhaustive(depAsset("gt-1", 5, false, 2), depAsset("gt-2", 9, false, 2)),
		dr.ID: exhaustive(depAsset("gt-1", 5, true, 3)),
		du.ID: exhaustive(depAsset("gt-1", 5, false, 3)),
	}
	preds := map[string]prediction.File{
		// Right in the reachable variant; the used asset is wrongly reachable
		// in the unreachable variant.
		tr.ID: predFile(tr.ID, true, depPred(5, prediction.Reachable, 2), depPred(9, prediction.Unreachable, 2)),
		tu.ID: predFile(tu.ID, true, depPred(5, prediction.Reachable, 2), depPred(9, prediction.Unreachable, 2)),
		// depth3: the reachable asset is not detected; the distance is off by one.
		dr.ID: predFile(dr.ID, true),
		du.ID: predFile(du.ID, true, depPred(5, prediction.Unreachable, 2)),
	}
	rep := Score("t", manifest(tr, tu, dr, du), truth, preds)
	s := rep.SuiteB
	if s == nil || !s.Reported || s.Entries != 4 {
		t.Fatalf("suite B view = %+v", s)
	}
	cases := map[string]string{
		"all":        "R:1/0/0/1 U:1/3/0/0 P=0.500 R=0.500 false=1/3 depth=5/5 dist=4/5 pairs=0/2",
		"transitive": "R:1/0/0/0 U:1/2/0/0 P=0.500 R=1.000 false=1/2 depth=4/4 dist=4/4 pairs=0/1",
		"depth3":     "R:0/0/0/1 U:0/1/0/0 P=n/a R=0.000 false=0/1 depth=1/1 dist=0/1 pairs=0/1",
	}
	got := map[string]string{"all": suiteBLine(s.Overall)}
	for _, g := range s.ByScenario {
		got[g.Key] = suiteBLine(g.SuiteBCounts)
	}
	for k, want := range cases {
		if got[k] != want {
			t.Errorf("%s: got %s, want %s", k, got[k], want)
		}
	}
	md := Markdown(rep)
	for _, want := range []string{"## Suite B reachability", "| go / depth3 | reachable | 1 | 0 | 0 | 0 | 1 |"} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown lacks %q", want)
		}
	}
	if _, err := JSON(rep); err != nil {
		t.Fatal(err)
	}
}

func TestSuiteBWithoutReachabilityIsNotZeroFilled(t *testing.T) {
	r, u := pairEntries("transitive")
	truth := map[string]groundtruth.File{r.ID: exhaustive(depAsset("gt-1", 5, true, 2)), u.ID: exhaustive(depAsset("gt-1", 5, false, 2))}
	preds := map[string]prediction.File{r.ID: predFile(r.ID, false, depPred(5, prediction.NoAnswer, 2)), u.ID: predFile(u.ID, false)}
	rep := Score("t", manifest(r, u), truth, preds)
	if rep.SuiteB == nil || rep.SuiteB.Reported {
		t.Fatalf("want a suite B view that reports no reachability, got %+v", rep.SuiteB)
	}
	md := Markdown(rep)
	if !strings.Contains(md, "No reachability: none of the 2 scored pair entries") || strings.Contains(md, "Reachable P") {
		t.Fatalf("markdown must say no reachability and print no reachability metrics:\n%s", md)
	}
}

func TestScopeComponentsIgnoresOtherDependencies(t *testing.T) {
	gt := exhaustive(depAsset("gt-1", 5, true, 2))
	gt.Scope.Components = []string{libB}
	sha2 := predAsset("src/sha256.rs", 40, 40)
	sha2.Origin = groundtruth.Dependency{Component: "pkg:cargo/sha2@0.10.8"}
	project := predAsset("main.go", 3, 3)
	m := MatchEntry(gt, []prediction.Asset{depPred(5, prediction.Reachable, 2), sha2, project})
	if got, want := summary(m), "TP=1 FP=1 FN=0 dup=0 ignored=1 [gt-1@5]"; got != want {
		t.Fatalf("got %s, want %s (the sha2 finding is out of scope, project code stays in)", got, want)
	}
}
