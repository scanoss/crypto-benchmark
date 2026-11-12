package score

import (
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// typedCounts is one object-level outcome and the asset type it is broken down by.
type typedCounts struct {
	Type groundtruth.AssetType
	Counts
}

// objectCounts derives one entry's object-level outcomes from its
// operation-level match. An object is the set of assets that share an
// object id, or a single asset without one. It is detected when any of its
// assets is matched; it then counts one TP, and every other prediction
// matched to it, as well as every operation-level duplicate (all of whose
// candidates are matched assets), counts as a duplicate. It is attributed
// when any of its matched assets is in attributed. Operation-level FPs stay
// FPs. Objects are broken down by their asset type, which the annotation
// guide makes the same for all their assets; FPs and duplicates by the
// prediction's.
func objectCounts(gt groundtruth.File, m Match, attributed map[string]bool) []typedCounts {
	type object struct {
		typ        groundtruth.AssetType
		matched    int
		attributed bool
	}
	key := func(a groundtruth.Asset) string {
		if a.Object != "" {
			return "object:" + a.Object
		}
		return "asset:" + a.ID
	}
	objects := map[string]*object{}
	var order []string
	for _, a := range gt.Assets {
		k := key(a)
		if objects[k] == nil {
			objects[k] = &object{typ: a.Type}
			order = append(order, k)
		}
	}
	for _, p := range m.Pairs {
		o := objects[key(p.Truth)]
		o.matched++
		o.attributed = o.attributed || attributed[p.Truth.ID]
	}
	var out []typedCounts
	for _, k := range order {
		o := objects[k]
		c := Counts{FN: 1}
		if o.matched > 0 {
			c = Counts{TP: 1, Duplicates: o.matched - 1}
			if o.attributed {
				c.AttributedTP = 1
			}
		}
		out = append(out, typedCounts{o.typ, c})
	}
	pred := func(ps []prediction.Asset, c Counts) {
		for _, p := range ps {
			out = append(out, typedCounts{p.Type, c})
		}
	}
	pred(m.FP, Counts{FP: 1})
	pred(m.Duplicates, Counts{Duplicates: 1})
	return out
}
