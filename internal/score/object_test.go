package score

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

// show renders counts as one literal for assertions.
func show(c Counts) string {
	return fmt.Sprintf("TP=%d FP=%d FN=%d dup=%d attributed=%d P=%s R=%s F1=%s",
		c.TP, c.FP, c.FN, c.Duplicates, c.AttributedTP, c.Precision(), c.Recall(), c.F1())
}

func onObject(a groundtruth.Asset, object string) groundtruth.Asset {
	a.Object = object
	return a
}

func scoreOne(gt groundtruth.File, preds ...prediction.Asset) Report {
	return Score("t", manifest(entry("e", corpus.Python, corpus.SuiteA)), map[string]groundtruth.File{"e": gt},
		map[string]prediction.File{"e": {Tool: prediction.Tool{Name: "t"}, Entry: "e", Assets: preds}})
}

// keyObject is an RSA key generated on line 5, then used to encrypt on line
// 10 and decrypt on line 15, plus a one-shot hash on line 20 with no object.
func keyObject() groundtruth.File {
	return exhaustive(
		onObject(truthAsset("gt-1", "main.py", 5, 5), "main/key"),
		onObject(truthAsset("gt-2", "main.py", 10, 10), "main/key"),
		onObject(truthAsset("gt-3", "main.py", 15, 15), "main/key"),
		truthAsset("gt-4", "main.py", 20, 20),
	)
}

func TestPerObjectToolGetsFullObjectRecall(t *testing.T) {
	rep := scoreOne(keyObject(), predAsset("main.py", 5, 5), predAsset("main.py", 20, 20))
	if got, want := show(rep.Overall), "TP=2 FP=0 FN=2 dup=0 attributed=2 P=1.000 R=0.500 F1=0.667"; got != want {
		t.Errorf("operation level: got %s, want %s", got, want)
	}
	if got, want := show(rep.Objects.Overall), "TP=2 FP=0 FN=0 dup=0 attributed=2 P=1.000 R=1.000 F1=1.000"; got != want {
		t.Errorf("object level: got %s, want %s", got, want)
	}
}

func TestAnyMatchedAssetDetectsTheObject(t *testing.T) {
	// Only the decrypt is found: the object is detected, the hash is not.
	rep := scoreOne(keyObject(), predAsset("main.py", 15, 15))
	if got, want := show(rep.Objects.Overall), "TP=1 FP=0 FN=1 dup=0 attributed=1 P=1.000 R=0.500 F1=0.667"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestExtraPredictionsOnOneObjectAreDuplicates(t *testing.T) {
	// One prediction per operation, plus a second one on the key line that
	// is already an operation-level duplicate, plus one far from everything.
	rep := scoreOne(keyObject(), predAsset("main.py", 5, 5), predAsset("main.py", 5, 5),
		predAsset("main.py", 10, 10), predAsset("main.py", 15, 15), predAsset("main.py", 40, 40))
	if got, want := show(rep.Overall), "TP=3 FP=1 FN=1 dup=1 attributed=3 P=0.750 R=0.750 F1=0.750"; got != want {
		t.Errorf("operation level: got %s, want %s", got, want)
	}
	if got, want := show(rep.Objects.Overall), "TP=1 FP=1 FN=1 dup=3 attributed=1 P=0.500 R=0.500 F1=0.500"; got != want {
		t.Errorf("object level: got %s, want %s", got, want)
	}
}

func TestObjectIsAttributedWhenAnyMatchedAssetIs(t *testing.T) {
	wrong := func(start int) prediction.Asset {
		p := predAsset("main.py", start, start)
		p.Algorithm = "ChaCha20"
		return p
	}
	cases := []struct {
		name  string
		preds []prediction.Asset
		want  string
	}{
		{"one of two matches attributed", []prediction.Asset{wrong(5), predAsset("main.py", 10, 10)}, "TP=1 FP=0 FN=1 dup=1 attributed=1"},
		{"no match attributed", []prediction.Asset{wrong(5), wrong(10)}, "TP=1 FP=0 FN=1 dup=1 attributed=0"},
		{"unmatched asset does not attribute", []prediction.Asset{wrong(5)}, "TP=1 FP=0 FN=1 dup=0 attributed=0"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _, _ := strings.Cut(show(scoreOne(keyObject(), c.preds...).Objects.Overall), " P=")
			if got != c.want {
				t.Fatalf("got %s, want %s", got, c.want)
			}
		})
	}
}

func TestObjectsAreScopedToTheirEntry(t *testing.T) {
	// The same object id in two entries names two objects.
	m := manifest(entry("a", corpus.Python, corpus.SuiteA), entry("b", corpus.Java, corpus.SuiteA))
	truth := map[string]groundtruth.File{
		"a": exhaustive(onObject(truthAsset("gt-1", "x.py", 5, 5), "k"), onObject(truthAsset("gt-2", "x.py", 9, 9), "k")),
		"b": exhaustive(onObject(truthAsset("gt-1", "x.py", 5, 5), "k"), onObject(truthAsset("gt-2", "x.py", 9, 9), "k")),
	}
	preds := map[string]prediction.File{
		"a": {Tool: prediction.Tool{Name: "t"}, Entry: "a", Assets: []prediction.Asset{predAsset("x.py", 5, 5)}},
		"b": {Tool: prediction.Tool{Name: "t"}, Entry: "b"},
	}
	o := Score("t", m, truth, preds).Objects
	if got, want := show(o.Overall), "TP=1 FP=0 FN=1 dup=0 attributed=1 P=1.000 R=0.500 F1=0.667"; got != want {
		t.Errorf("overall: got %s, want %s", got, want)
	}
	var rows []string
	for _, g := range append(o.ByEntry, o.ByLanguage...) {
		rows = append(rows, g.Key+": "+show(g.Counts))
	}
	want := []string{
		"a: TP=1 FP=0 FN=0 dup=0 attributed=1 P=1.000 R=1.000 F1=1.000",
		"b: TP=0 FP=0 FN=1 dup=0 attributed=0 P=n/a R=0.000 F1=0.000",
		"java: TP=0 FP=0 FN=1 dup=0 attributed=0 P=n/a R=0.000 F1=0.000",
		"python: TP=1 FP=0 FN=0 dup=0 attributed=1 P=1.000 R=1.000 F1=1.000",
	}
	if strings.Join(rows, "\n") != strings.Join(want, "\n") {
		t.Fatalf("groups:\n%s\nwant:\n%s", strings.Join(rows, "\n"), strings.Join(want, "\n"))
	}
}

func TestObjectsBreakDownByAssetType(t *testing.T) {
	cert := func(id string, line int) groundtruth.Asset {
		a := onObject(truthAsset(id, "main.py", line, line), "main/cert")
		a.Type = groundtruth.Certificate
		return a
	}
	gt := exhaustive(cert("gt-1", 5), cert("gt-2", 30), truthAsset("gt-3", "main.py", 50, 50))
	proto := predAsset("main.py", 70, 70)
	proto.Type = groundtruth.Protocol
	certPred := predAsset("main.py", 30, 30)
	certPred.Type = groundtruth.Certificate
	var rows []string
	for _, g := range scoreOne(gt, certPred, proto).Objects.ByType {
		rows = append(rows, g.Key+": "+show(g.Counts))
	}
	want := []string{
		"algorithm: TP=0 FP=0 FN=1 dup=0 attributed=0 P=n/a R=0.000 F1=0.000",
		"certificate: TP=1 FP=0 FN=0 dup=0 attributed=1 P=1.000 R=1.000 F1=1.000",
		"protocol: TP=0 FP=1 FN=0 dup=0 attributed=0 P=0.000 R=n/a F1=0.000",
	}
	if strings.Join(rows, "\n") != strings.Join(want, "\n") {
		t.Fatalf("by type:\n%s\nwant:\n%s", strings.Join(rows, "\n"), strings.Join(want, "\n"))
	}
}

func TestObjectViewRendersAfterTheOperationView(t *testing.T) {
	rep := scoreOne(keyObject(), predAsset("main.py", 5, 5), predAsset("main.py", 20, 20))
	md := Markdown(rep)
	ops, objects, found := strings.Cut(md, "## Object-level view")
	if !found || !strings.Contains(ops, "## Suite B pair discrimination") {
		t.Fatalf("object-level view missing or not last:\n%s", md)
	}
	for _, want := range []string{
		"| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |",
		"| all | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |",
		"| python | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |",
		"| e | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |",
	} {
		if !strings.Contains(objects, want) {
			t.Errorf("object-level view lacks %q", want)
		}
	}
	if !strings.Contains(ops, "| all | 1.000 | 0.500 | 0.667 | 2 | 0 | 2 | 0 |") {
		t.Errorf("operation view changed:\n%s", ops)
	}
	data, err := JSON(rep)
	if err != nil {
		t.Fatal(err)
	}
	var got struct {
		ObjectLevel struct {
			Overall struct {
				TP     int     `json:"tp"`
				FN     int     `json:"fn"`
				Recall float64 `json:"recall"`
			} `json:"overall"`
			ByEntry []struct {
				Group string `json:"group"`
			} `json:"by_entry"`
		} `json:"object_level"`
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if o := got.ObjectLevel; len(o.ByEntry) != 1 || fmt.Sprintf("%d %d %g %s", o.Overall.TP, o.Overall.FN, o.Overall.Recall, o.ByEntry[0].Group) != "2 0 1 e" {
		t.Fatalf("object_level JSON: %+v\n%s", o, data)
	}
}
