package score

import (
	"cmp"
	"slices"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/algo"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// LineTolerance widens each ground-truth line range on both sides before
// testing it for overlap with a prediction's range. It is 0: the ranges
// themselves must overlap.
const LineTolerance = 0

// Pair is one matched ground-truth asset and the prediction assigned to it.
type Pair struct {
	Truth groundtruth.Asset
	Pred  prediction.Asset
}

// Match is the one-to-one assignment of one entry's predictions to its ground truth.
type Match struct {
	Pairs      []Pair
	FP         []prediction.Asset // unmatched, in scope, not duplicates
	FN         []groundtruth.Asset
	Duplicates []prediction.Asset // would match only ground truth already matched
	Ignored    int                // predictions outside scope.files or scope.components
}

// MatchEntry assigns preds to gt by the pre-registered matching rules.
func MatchEntry(gt groundtruth.File, preds []prediction.Asset) Match {
	var m Match
	inScope := func(prediction.Asset) bool { return true }
	if gt.Scope.Kind == groundtruth.Files {
		files := map[string]bool{}
		for _, f := range gt.Scope.Files {
			files[prediction.NormalizePath(f)] = true
		}
		inScope = func(p prediction.Asset) bool { return files[p.File] }
	}
	if comps := gt.Scope.Components; len(comps) > 0 {
		listed := map[string]bool{}
		for _, c := range comps {
			listed[normalizePURL(c)] = true
		}
		inFiles := inScope
		inScope = func(p prediction.Asset) bool {
			if d, ok := p.Origin.(groundtruth.Dependency); ok && !listed[normalizePURL(d.Component)] {
				return false
			}
			return inFiles(p)
		}
	}
	var scoped []prediction.Asset
	for _, p := range preds {
		if inScope(p) {
			scoped = append(scoped, p)
		} else {
			m.Ignored++
		}
	}
	slices.SortStableFunc(scoped, comparePredictions)

	type candidate struct {
		truth, pred, distance int
		attributed            bool
	}
	var cands []candidate
	hasCandidate := make([]bool, len(scoped))
	for ti, t := range gt.Assets {
		for pi, p := range scoped {
			if d, ok := lineDistance(t, p); ok {
				attributed, _ := Attributed(Pair{t, p})
				cands = append(cands, candidate{ti, pi, d, attributed})
				hasCandidate[pi] = true
			}
		}
	}
	unattributed := func(c candidate) int {
		if c.attributed {
			return 0
		}
		return 1
	}
	slices.SortFunc(cands, func(a, b candidate) int {
		return cmp.Or(
			cmp.Compare(a.distance, b.distance),
			cmp.Compare(unattributed(a), unattributed(b)),
			cmp.Compare(gt.Assets[a.truth].ID, gt.Assets[b.truth].ID),
			cmp.Compare(a.pred, b.pred),
		)
	})
	truthTaken := make([]bool, len(gt.Assets))
	predTaken := make([]bool, len(scoped))
	for _, c := range cands {
		if truthTaken[c.truth] || predTaken[c.pred] {
			continue
		}
		truthTaken[c.truth], predTaken[c.pred] = true, true
		m.Pairs = append(m.Pairs, Pair{gt.Assets[c.truth], scoped[c.pred]})
	}
	for pi, p := range scoped {
		switch {
		case predTaken[pi]:
		case hasCandidate[pi]:
			m.Duplicates = append(m.Duplicates, p)
		default:
			m.FP = append(m.FP, p)
		}
	}
	for ti, t := range gt.Assets {
		if !truthTaken[ti] {
			m.FN = append(m.FN, t)
		}
	}
	return m
}

// lineDistance reports whether p is a candidate match for t and, if so, the
// number of lines between p and the nearest of t's ranges (its lines and
// anchors; 0 when an unwidened range overlaps p).
func lineDistance(t groundtruth.Asset, p prediction.Asset) (int, bool) {
	if t.Type != p.Type || prediction.NormalizePath(t.File) != p.File || !sameOrigin(t.Origin, p.Origin) {
		return 0, false
	}
	best, found := 0, false
	for _, r := range append([]groundtruth.Lines{t.Lines}, t.Anchors...) {
		if p.Lines.Start > r.End+LineTolerance || p.Lines.End < r.Start-LineTolerance {
			continue
		}
		d := max(0, p.Lines.Start-r.End, r.Start-p.Lines.End)
		if !found || d < best {
			best, found = d, true
		}
	}
	return best, found
}

func sameOrigin(a, b groundtruth.Origin) bool {
	switch a := a.(type) {
	case groundtruth.Project:
		_, ok := b.(groundtruth.Project)
		return ok
	case groundtruth.Dependency:
		b, ok := b.(groundtruth.Dependency)
		return ok && normalizePURL(a.Component) == normalizePURL(b.Component)
	}
	return false
}

// normalizePURL drops a package URL's qualifiers and subpath.
func normalizePURL(purl string) string {
	if i := strings.IndexAny(purl, "?#"); i >= 0 {
		return purl[:i]
	}
	return purl
}

func comparePredictions(a, b prediction.Asset) int {
	return cmp.Or(
		cmp.Compare(a.File, b.File),
		cmp.Compare(a.Lines.Start, b.Lines.Start),
		cmp.Compare(a.Lines.End, b.Lines.End),
		cmp.Compare(a.Type, b.Type),
		cmp.Compare(a.Algorithm, b.Algorithm),
	)
}

// Attributed reports whether a matched pair also agrees on the algorithm, on
// every parameter the ground truth sets, both sides canonicalized by
// algo.Canonicalize, and on the material type when the ground truth sets it. conflicts lists, for either side, each parameter its
// algorithm name implied that an explicit field overrode.
func Attributed(p Pair) (attributed bool, conflicts []string) {
	t := algo.Canonicalize(p.Truth.Algorithm, p.Truth.Parameters)
	g := algo.Canonicalize(p.Pred.Algorithm, p.Pred.Params)
	for _, c := range t.Conflicts {
		conflicts = append(conflicts, p.Truth.ID+" ground truth: "+c)
	}
	for _, c := range g.Conflicts {
		conflicts = append(conflicts, p.Truth.ID+" prediction: "+c)
	}
	if t.Algorithm != "" && t.Algorithm != g.Algorithm {
		return false, conflicts
	}
	if p.Truth.MaterialType != "" && p.Truth.MaterialType != p.Pred.MaterialType {
		return false, conflicts
	}
	tp, gp := t.Params, g.Params
	for _, f := range [][2]string{{tp.Mode, gp.Mode}, {tp.Padding, gp.Padding}, {tp.Curve, gp.Curve}, {tp.Hash, gp.Hash}} {
		if f[0] != "" && f[0] != f[1] {
			return false, conflicts
		}
	}
	return tp.KeySize == 0 || tp.KeySize == gp.KeySize, conflicts
}
