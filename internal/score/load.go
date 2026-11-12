package score

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/check"
	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
	"github.com/scanoss/crypto-benchmark/internal/validate"
)

// Inputs are the validated manifest, ground truth and one tool's predictions, keyed by entry id.
type Inputs struct {
	Manifest corpus.Manifest
	Truth    map[string]groundtruth.File
	Preds    map[string]prediction.File
}

// Load validates the benchmark under root (and overlay), then reads the
// ground truth and the predictions in predDir/<tool>/<entry-id>.json.
// Scoring proceeds only when the returned report has no errors.
func Load(root, overlay, predDir, tool string) (Inputs, *check.Report) {
	r := validate.Run(root, overlay)
	if r.HasErrors() {
		return Inputs{}, r
	}
	r = &check.Report{}
	in := Inputs{Manifest: corpus.Load(root, overlay, r), Truth: map[string]groundtruth.File{}, Preds: map[string]prediction.File{}}
	roots := []string{root}
	if overlay != "" {
		roots = append(roots, overlay)
	}
	for _, e := range in.Manifest.Entries {
		for _, base := range roots {
			file := filepath.Join(base, "groundtruth", string(e.Suite), e.ID+".yaml")
			data, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			if f, ok := groundtruth.Parse(file, data, r); ok {
				in.Truth[e.ID] = f
			}
		}
	}
	dir := filepath.Join(predDir, tool)
	names, err := os.ReadDir(dir)
	if err != nil {
		r.Errorf(dir, "", "%v", err)
		return in, r
	}
	for _, d := range names {
		if d.IsDir() || filepath.Ext(d.Name()) != ".json" {
			continue
		}
		file := filepath.Join(dir, d.Name())
		data, err := os.ReadFile(file)
		if err != nil {
			r.Errorf(file, "", "%v", err)
			continue
		}
		f, ok := prediction.Parse(file, data, r)
		if !ok {
			continue
		}
		stem := strings.TrimSuffix(d.Name(), ".json")
		switch _, known := in.Manifest.Lookup(f.Entry); {
		case f.Tool.Name != tool:
			r.Errorf(file, "tool.name", "%q does not match --tool %q", f.Tool.Name, tool)
		case f.Entry != stem:
			r.Errorf(file, "entry", "%q does not match the file name %q", f.Entry, stem)
		case !known:
			r.Errorf(file, "entry", "%q is not a manifest entry", f.Entry)
		default:
			in.Preds[f.Entry] = f
		}
	}
	return in, r
}
