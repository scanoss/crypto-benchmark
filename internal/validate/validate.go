// Package validate checks the manifest and every ground-truth file of the
// base repository and an optional overlay against each other.
package validate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/libraries"
)

// Run validates base and, when overlay is not empty, the overlay. The suite A
// library list lives only in base.
func Run(base, overlay string) *check.Report {
	r := &check.Report{}
	libraries.Load(base, r)
	m := corpus.Load(base, overlay, r)
	annotated := map[string]string{}
	checkTree(base, false, m, annotated, r)
	if overlay != "" {
		checkTree(overlay, true, m, annotated, r)
	}
	for _, e := range m.Entries {
		if _, ok := annotated[e.ID]; !ok {
			r.Warnf(e.File, e.Where, "%q is not annotated yet (no groundtruth/%s/%s.yaml)", e.ID, e.Suite, e.ID)
		}
	}
	return r
}

func checkTree(root string, overlay bool, m corpus.Manifest, annotated map[string]string, r *check.Report) {
	dir := filepath.Join(root, "groundtruth")
	err := filepath.WalkDir(dir, func(file string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(file) != ".yaml" {
			return err
		}
		data, err := os.ReadFile(file)
		if err != nil {
			r.Errorf(file, "", "%v", err)
			return nil
		}
		f, ok := groundtruth.Parse(file, data, r)
		if !ok {
			return nil
		}
		checkPlacement(file, dir, overlay, f, m, annotated, r)
		return nil
	})
	if err != nil && !os.IsNotExist(err) {
		r.Errorf(dir, "", "%v", err)
	}
}

func checkPlacement(file, dir string, overlay bool, f groundtruth.File, m corpus.Manifest, annotated map[string]string, r *check.Report) {
	rel, _ := filepath.Rel(dir, file)
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if stem := strings.TrimSuffix(filepath.Base(file), ".yaml"); f.Entry != stem {
		r.Errorf(file, "entry", "%q does not match the file name %q", f.Entry, stem)
		return
	}
	e, ok := m.Lookup(f.Entry)
	if !ok {
		if !m.Declared(f.Entry) {
			r.Errorf(file, "entry", "%q is not a manifest entry", f.Entry)
		}
		return
	}
	if len(parts) != 2 || parts[0] != string(e.Suite) {
		r.Errorf(file, "entry", "%q is in suite %s; the file must be groundtruth/%s/%s.yaml", e.ID, e.Suite, e.Suite, e.ID)
	}
	if e.Split == corpus.Holdout && !overlay {
		r.Errorf(file, "entry", "%q is a holdout entry; its ground truth must live in the overlay", e.ID)
	}
	if prev, dup := annotated[e.ID]; dup {
		r.Errorf(file, "entry", "%q is already annotated in %s", e.ID, prev)
		return
	}
	annotated[e.ID] = file
	if e.Pair != nil {
		checkPairTruth(file, e, f, r)
	}
}

// checkPairTruth enforces what suite B scoring relies on: a scenario on the
// entry, and on every asset a known reachability, the exact dependency
// distance, and a call path exactly when the asset is reachable.
func checkPairTruth(file string, e corpus.Entry, f groundtruth.File, r *check.Report) {
	if e.Pair.Scenario == "" {
		r.Errorf(e.File, e.Where+".scenario", "required on a suite B pair entry")
	}
	for i, a := range f.Assets {
		where := fmt.Sprintf("assets[%d]", i)
		k, ok := a.Reachability.(groundtruth.Known)
		switch {
		case !ok:
			r.Errorf(file, where+".reachable", "must be true or false in a suite B pair")
			continue
		case k.Distance < 0:
			r.Errorf(file, where+".distance", "required in a suite B pair")
		}
		switch {
		case k.Reachable && len(a.Path) == 0:
			r.Errorf(file, where+".path", "required when reachable is true in a suite B pair")
		case !k.Reachable && len(a.Path) > 0:
			r.Errorf(file, where+".path", "not allowed when reachable is false: there is no call path")
		}
	}
}
