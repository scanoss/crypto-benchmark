package score

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

func truthAsset(id, file string, start, end int) groundtruth.Asset {
	return groundtruth.Asset{
		ID: id, File: file, Lines: groundtruth.Lines{Start: start, End: end},
		Type: groundtruth.Algorithm, Algorithm: "AES", Origin: groundtruth.Project{},
		Reachability: groundtruth.Known{Reachable: true, Depth: 0},
	}
}

func predAsset(file string, start, end int) prediction.Asset {
	return prediction.Asset{
		File: file, Lines: groundtruth.Lines{Start: start, End: end}, Type: groundtruth.Algorithm,
		Algorithm: "AES", Origin: groundtruth.Project{}, Reachable: prediction.NoAnswer,
	}
}

func exhaustive(assets ...groundtruth.Asset) groundtruth.File {
	return groundtruth.File{Entry: "e", Scope: groundtruth.Scope{Kind: groundtruth.Exhaustive}, Assets: assets}
}

// summary renders a match as "TP=.. FP=.. FN=.. dup=.. ignored=.." plus the
// matched pairs as truth-id@pred-start, so assertions compare one literal.
func summary(m Match) string {
	var pairs []string
	for _, p := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s@%d", p.Truth.ID, p.Pred.Lines.Start))
	}
	return fmt.Sprintf("TP=%d FP=%d FN=%d dup=%d ignored=%d %v", len(m.Pairs), len(m.FP), len(m.FN), len(m.Duplicates), m.Ignored, pairs)
}

func TestLinesMustOverlap(t *testing.T) {
	gt := exhaustive(truthAsset("gt-1", "./src/Crypto.java", 10, 12))
	cases := []struct {
		name       string
		start, end int
		want       string
	}{
		{"inside the range", 11, 11, "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@11]"},
		{"on the last line of a multi-line range", 12, 12, "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@12]"},
		{"range overlapping the first line", 3, 10, "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@3]"},
		{"one line after", 13, 13, "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
		{"one line before", 9, 9, "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
		{"range ending one line before", 3, 9, "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := summary(MatchEntry(gt, []prediction.Asset{predAsset("src/Crypto.java", c.start, c.end)})); got != c.want {
				t.Fatalf("got %s, want %s", got, c.want)
			}
		})
	}
}

func withAnchors(a groundtruth.Asset, ranges ...[2]int) groundtruth.Asset {
	for _, r := range ranges {
		a.Anchors = append(a.Anchors, groundtruth.Lines{Start: r[0], End: r[1]})
	}
	return a
}

func TestAnAnchorMatchesLikeTheLines(t *testing.T) {
	// generateKeyPair on line 14, KeyPairGenerator.getInstance on line 11.
	gt := exhaustive(withAnchors(truthAsset("gt-1", "K.java", 14, 14), [2]int{11, 11}))
	cases := []struct {
		name  string
		lines [][2]int
		want  string
	}{
		{"on the operation call", [][2]int{{14, 14}}, "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@14]"},
		{"on the anchor", [][2]int{{11, 11}}, "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@11]"},
		{"between them", [][2]int{{12, 13}}, "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
		{"on both: one asset, one prediction", [][2]int{{11, 11}, {14, 14}}, "TP=1 FP=0 FN=0 dup=1 ignored=0 [gt-1@11]"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var preds []prediction.Asset
			for _, l := range c.lines {
				preds = append(preds, predAsset("K.java", l[0], l[1]))
			}
			if got := summary(MatchEntry(gt, preds)); got != c.want {
				t.Fatalf("got %s, want %s", got, c.want)
			}
		})
	}
}

func TestAPredictionOnASharedAnchorConsumesOneAsset(t *testing.T) {
	// One new SecureRandom() on line 5 anchors two draws on lines 7 and 9.
	gt := exhaustive(withAnchors(truthAsset("gt-1", "R.java", 7, 7), [2]int{5, 5}), withAnchors(truthAsset("gt-2", "R.java", 9, 9), [2]int{5, 5}))
	if got, want := summary(MatchEntry(gt, []prediction.Asset{predAsset("R.java", 5, 5)})), "TP=1 FP=0 FN=1 dup=0 ignored=0 [gt-1@5]"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestDuplicatesAreNeitherTPNorFP(t *testing.T) {
	gt := exhaustive(truthAsset("gt-1", "a.go", 10, 11), truthAsset("gt-2", "a.go", 30, 30))
	preds := []prediction.Asset{predAsset("a.go", 11, 11), predAsset("a.go", 10, 10), predAsset("a.go", 10, 11), predAsset("a.go", 50, 50)}
	m := MatchEntry(gt, preds)
	if got, want := summary(m), "TP=1 FP=1 FN=1 dup=2 ignored=0 [gt-1@10]"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	rep := Score("t", manifest(entry("e", corpus.Go, corpus.SuiteA)), map[string]groundtruth.File{"e": gt},
		map[string]prediction.File{"e": {Tool: prediction.Tool{Name: "t"}, Entry: "e", Assets: preds}})
	if got, want := fmt.Sprint(rep.Overall.Precision(), rep.Overall.Recall()), "0.500 0.500"; got != want {
		t.Fatalf("precision recall = %s, want %s (duplicates must not count)", got, want)
	}
}

func withAlgorithm[T groundtruth.Asset | prediction.Asset](a T, name string) T {
	switch a := any(&a).(type) {
	case *groundtruth.Asset:
		a.Algorithm = name
	case *prediction.Asset:
		a.Algorithm = name
	}
	return a
}

// matchedNames renders each matched pair as truth-algorithm=prediction-algorithm.
func matchedNames(m Match) []string {
	var out []string
	for _, p := range m.Pairs {
		out = append(out, p.Truth.Algorithm+"="+p.Pred.Algorithm)
	}
	return out
}

func TestTiesPreferAnAttributedPrediction(t *testing.T) {
	// A generic and a specific name for one call. "AES" sorts before
	// "AES-GCM", but only the specific one attributes against AES/GCM.
	truth := withAlgorithm(truthAsset("gt-1", "a.java", 20, 20), "AES")
	truth.Parameters.Mode = "GCM"
	preds := []prediction.Asset{withAlgorithm(predAsset("a.java", 20, 20), "AES"), withAlgorithm(predAsset("a.java", 20, 20), "AES-GCM")}
	m := MatchEntry(exhaustive(truth), preds)
	if got, want := fmt.Sprintf("%v %d %s", matchedNames(m), len(m.Duplicates), m.Duplicates[0].Algorithm), "[AES=AES-GCM] 1 AES"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestTiesPreferTheAttributedPairingAcrossTruthAssets(t *testing.T) {
	// Two algorithms on one call and two findings on its lines. Ground-truth
	// id and prediction order alone would pair them crosswise.
	gt := exhaustive(withAlgorithm(truthAsset("gt-1", "a.java", 9, 13), "RSA-OAEP"), withAlgorithm(truthAsset("gt-2", "a.java", 9, 13), "AES"))
	preds := []prediction.Asset{withAlgorithm(predAsset("a.java", 9, 9), "AES"), withAlgorithm(predAsset("a.java", 9, 9), "RSA-OAEP")}
	m := MatchEntry(gt, preds)
	if got, want := fmt.Sprint(matchedNames(m)), "[RSA-OAEP=RSA-OAEP AES=AES]"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestTiesWithoutAttributionFallBackToPredictionOrder(t *testing.T) {
	truth := withAlgorithm(truthAsset("gt-1", "a.java", 20, 20), "SHA-256")
	preds := []prediction.Asset{withAlgorithm(predAsset("a.java", 20, 20), "MD5"), withAlgorithm(predAsset("a.java", 20, 20), "AES")}
	m := MatchEntry(exhaustive(truth), preds)
	if got, want := fmt.Sprint(matchedNames(m), len(m.Duplicates)), "[SHA-256=AES] 1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestAttributionNeverWinsAcrossALineGap(t *testing.T) {
	// The attributed finding is on the next line, so it is no candidate at
	// all: the tie-break orders candidates, it does not widen them.
	truth := withAlgorithm(truthAsset("gt-1", "a.java", 20, 20), "AES")
	truth.Parameters.Mode = "GCM"
	preds := []prediction.Asset{withAlgorithm(predAsset("a.java", 20, 20), "AES"), withAlgorithm(predAsset("a.java", 21, 21), "AES-GCM")}
	m := MatchEntry(exhaustive(truth), preds)
	if got, want := fmt.Sprint(matchedNames(m), len(m.Duplicates), len(m.FP)), "[AES=AES] 0 1"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestGreedyBreaksDistanceTiesByTruthID(t *testing.T) {
	// Two annotations on one line, listed out of id order: the lower id takes the prediction.
	gt := exhaustive(truthAsset("gt-2", "a.go", 20, 20), truthAsset("gt-1", "a.go", 20, 20))
	m := MatchEntry(gt, []prediction.Asset{predAsset("a.go", 20, 20)})
	if got, want := summary(m), "TP=1 FP=0 FN=1 dup=0 ignored=0 [gt-1@20]"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestOriginAndComponentMustAgree(t *testing.T) {
	dep := truthAsset("gt-1", "org/x/Cipher.java", 5, 5)
	dep.Origin = groundtruth.Dependency{Component: "pkg:maven/org.x/lib@1.0"}
	gt := exhaustive(dep)
	withOrigin := func(o groundtruth.Origin) prediction.Asset {
		p := predAsset("org/x/Cipher.java", 5, 5)
		p.Origin = o
		return p
	}
	cases := []struct {
		name string
		pred prediction.Asset
		want string
	}{
		{"same purl", withOrigin(groundtruth.Dependency{Component: "pkg:maven/org.x/lib@1.0"}), "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@5]"},
		{"qualifiers and subpath dropped", withOrigin(groundtruth.Dependency{Component: "pkg:maven/org.x/lib@1.0?type=jar#src"}), "TP=1 FP=0 FN=0 dup=0 ignored=0 [gt-1@5]"},
		{"other version", withOrigin(groundtruth.Dependency{Component: "pkg:maven/org.x/lib@2.0"}), "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
		{"project origin", withOrigin(groundtruth.Project{}), "TP=0 FP=1 FN=1 dup=0 ignored=0 []"},
	}
	if got, want := summary(MatchEntry(exhaustive(truthAsset("gt-1", "org/x/Cipher.java", 5, 5)), []prediction.Asset{withOrigin(dep.Origin)})), "TP=0 FP=1 FN=1 dup=0 ignored=0 []"; got != want {
		t.Fatalf("dependency prediction for project truth: got %s, want %s", got, want)
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := summary(MatchEntry(gt, []prediction.Asset{c.pred})); got != c.want {
				t.Fatalf("got %s, want %s", got, c.want)
			}
		})
	}
}

func TestAssetTypeMustAgree(t *testing.T) {
	p := predAsset("a.go", 3, 3)
	p.Type = groundtruth.Protocol
	if got, want := summary(MatchEntry(exhaustive(truthAsset("gt-1", "a.go", 3, 3)), []prediction.Asset{p})), "TP=0 FP=1 FN=1 dup=0 ignored=0 []"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestPredictionsOutsideScopeFilesAreIgnored(t *testing.T) {
	gt := groundtruth.File{Entry: "e", Scope: groundtruth.Scope{Kind: groundtruth.Files, Files: []string{"crypto/a.go"}},
		Assets: []groundtruth.Asset{truthAsset("gt-1", "crypto/a.go", 3, 3)}}
	m := MatchEntry(gt, []prediction.Asset{predAsset("crypto/a.go", 3, 3), predAsset("crypto/a.go", 40, 40), predAsset("other/b.go", 3, 3)})
	if got, want := summary(m), "TP=1 FP=1 FN=0 dup=0 ignored=1 [gt-1@3]"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestAttributedIsStricterThanDetected(t *testing.T) {
	gtAES := truthAsset("gt-1", "a.go", 3, 3)
	gtAES.Algorithm, gtAES.Parameters = "AES", groundtruth.Parameters{Mode: "GCM", KeySize: 256}
	gtSHA := truthAsset("gt-2", "a.go", 9, 9)
	gtSHA.Algorithm = "SHA-256"
	gtAny := truthAsset("gt-3", "a.go", 20, 20)
	gtAny.Algorithm = ""
	aes := predAsset("a.go", 3, 3)
	aes.Algorithm, aes.Params = "aes", groundtruth.Parameters{Mode: "gcm", KeySize: 128}
	sha := predAsset("a.go", 9, 9)
	sha.Algorithm = "sha_256"
	anyAlg := predAsset("a.go", 20, 20)
	anyAlg.Algorithm = "whatever"
	m := MatchEntry(exhaustive(gtAES, gtSHA, gtAny), []prediction.Asset{aes, sha, anyAlg})
	var got []string
	for _, p := range m.Pairs {
		ok, _ := Attributed(p)
		got = append(got, fmt.Sprintf("%s=%v", p.Truth.ID, ok))
	}
	slices.Sort(got)
	if strings.Join(got, " ") != "gt-1=false gt-2=true gt-3=true" {
		t.Fatalf("attributed: %v", got)
	}
	aes.Params.KeySize = 256
	rep := Score("t", manifest(entry("e", corpus.Go, corpus.SuiteA)), map[string]groundtruth.File{"e": exhaustive(gtAES, gtSHA)},
		map[string]prediction.File{"e": {Tool: prediction.Tool{Name: "t"}, Entry: "e", Assets: []prediction.Asset{aes, predAsset("a.go", 9, 9), predAsset("b.go", 1, 1)}}})
	c := rep.Overall
	if got, want := fmt.Sprint(c.TP, c.AttributedTP, c.AttributedShare(), c.AttributedPrecision(), c.AttributedRecall(), c.Precision()), "2 1 0.500 0.333 0.500 0.667"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestAttributionCanonicalizesBothSides(t *testing.T) {
	cases := []struct {
		name      string
		truthAlg  string
		truth     groundtruth.Parameters
		predAlg   string
		pred      groundtruth.Parameters
		want      bool
		conflicts string
	}{
		{"split truth, compact prediction", "AES", groundtruth.Parameters{KeySize: 256, Mode: "GCM"}, "AES-256-GCM", groundtruth.Parameters{}, true, ""},
		{"compact truth, split prediction", "AES-256-GCM", groundtruth.Parameters{}, "aes", groundtruth.Parameters{KeySize: 256, Mode: "gcm"}, true, ""},
		{"name parameter required of the prediction", "AES-256-GCM", groundtruth.Parameters{}, "AES-GCM", groundtruth.Parameters{}, false, ""},
		{"prediction may say more than the truth", "AES", groundtruth.Parameters{}, "AES-128-CBC", groundtruth.Parameters{}, true, ""},
		{"curve alias", "ECDSA", groundtruth.Parameters{Curve: "P-256"}, "ECDSA-prime256v1", groundtruth.Parameters{}, true, ""},
		{"specific name against its family", "SHA-256", groundtruth.Parameters{}, "SHA-2", groundtruth.Parameters{}, false, ""},
		{"3DES synonym", "3DES", groundtruth.Parameters{}, "DESede", groundtruth.Parameters{}, true, ""},
		{"unresolved prediction is absent", "SHA-256", groundtruth.Parameters{}, "SHA-$bits", groundtruth.Parameters{}, false, ""},
		{"HMAC name against split truth", "HMAC", groundtruth.Parameters{Hash: "SHA-256"}, "HMAC-SHA256", groundtruth.Parameters{}, true, ""},
		{"bare HMAC lacks the truth's hash", "HMAC", groundtruth.Parameters{Hash: "SHA-256"}, "HMAC", groundtruth.Parameters{}, false, ""},
		{"compact HMAC truth needs the hash", "HMAC-SHA-256", groundtruth.Parameters{}, "HMAC", groundtruth.Parameters{}, false, ""},
		{"wrong hash", "ECDSA", groundtruth.Parameters{Hash: "SHA-256"}, "ECDSA-SHA384", groundtruth.Parameters{}, false, ""},
		{"hash from an explicit field", "ECDSA", groundtruth.Parameters{Hash: "SHA-256"}, "ECDSA", groundtruth.Parameters{Hash: "SHA256"}, true, ""},
		{"truth without hash accepts any", "ECDSA", groundtruth.Parameters{}, "ECDSA-SHA-256", groundtruth.Parameters{}, true, ""},
		{"JCA name against split truth", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-256"}, "SHA256withRSA", groundtruth.Parameters{}, true, ""},
		{"RSASSA-PKCS1 with digest", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-256"}, "RSASSA-PKCS1-SHA-256", groundtruth.Parameters{}, true, ""},
		{"JCA RSA is not PSS", "RSA", groundtruth.Parameters{Padding: "PSS", Hash: "SHA-256"}, "SHA256withRSA", groundtruth.Parameters{}, false, ""},
		{"RSA-OAEP with digest against the guide's RSA-OAEP", "RSA-OAEP", groundtruth.Parameters{}, "RSA-OAEP-SHA-256", groundtruth.Parameters{}, true, ""},
		{"PBKDF2 with digest against bare PBKDF2", "PBKDF2", groundtruth.Parameters{}, "PBKDF2-SHA256", groundtruth.Parameters{}, true, ""},
		{"JCA PBKDF2 against bare PBKDF2", "PBKDF2", groundtruth.Parameters{}, "PBKDF2WithHmacSHA256", groundtruth.Parameters{}, true, ""},
		{"prefixed RS256 against the guide's RS256", "RSA", groundtruth.Parameters{Padding: "PKCS1v15", Hash: "SHA-256"}, "RSA-RS256", groundtruth.Parameters{Hash: "SHA-256"}, true, ""},
		{"ES256 carries its curve", "ECDSA", groundtruth.Parameters{Curve: "P-256", Hash: "SHA-256"}, "ES256", groundtruth.Parameters{}, true, ""},
		{"ECDSA-SHA-256 lacks ES256's curve", "ECDSA", groundtruth.Parameters{Curve: "P-256", Hash: "SHA-256"}, "ECDSA-SHA-256", groundtruth.Parameters{}, false, ""},
		{"PS256 against the guide's RSA-PSS", "RSA-PSS", groundtruth.Parameters{Hash: "SHA-256"}, "PS256", groundtruth.Parameters{}, true, ""},
		{"EdDSA is not Ed25519", "Ed25519", groundtruth.Parameters{}, "EdDSA", groundtruth.Parameters{}, false, ""},
		{"conflicting explicit field wins", "AES", groundtruth.Parameters{KeySize: 256, Mode: "GCM"}, "AES-256-GCM", groundtruth.Parameters{KeySize: 128}, false,
			`gt-1 prediction: name "AES-256-GCM" implies key_size 256, explicit key_size 128 wins`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tr := truthAsset("gt-1", "a.go", 3, 3)
			tr.Algorithm, tr.Parameters = c.truthAlg, c.truth
			pr := predAsset("a.go", 3, 3)
			pr.Algorithm, pr.Params = c.predAlg, c.pred
			ok, conflicts := Attributed(Pair{tr, pr})
			if ok != c.want {
				t.Fatalf("attributed = %v, want %v", ok, c.want)
			}
			if got := strings.Join(conflicts, "; "); got != c.conflicts {
				t.Fatalf("conflicts %q, want %q", got, c.conflicts)
			}
		})
	}
}

func TestMaterialTypeIsAttributedOnlyWhenTheTruthSetsIt(t *testing.T) {
	cases := []struct {
		name        string
		truth, pred groundtruth.MaterialType
		want        bool
	}{
		{"same type", "nonce", "nonce", true},
		{"different type", "nonce", "initialization-vector", false},
		{"prediction without type", "salt", "", false},
		{"truth without type", "", "private-key", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tr := truthAsset("gt-1", "a.go", 3, 3)
			tr.Type, tr.Algorithm, tr.Primitive, tr.MaterialType = groundtruth.RelatedCryptoMaterial, "", "", c.truth
			pr := predAsset("a.go", 3, 3)
			pr.Type, pr.Algorithm, pr.MaterialType = groundtruth.RelatedCryptoMaterial, "", c.pred
			if ok, _ := Attributed(Pair{tr, pr}); ok != c.want {
				t.Fatalf("attributed = %v, want %v", ok, c.want)
			}
		})
	}
}

func TestNameConflictsAreReported(t *testing.T) {
	tr := truthAsset("gt-1", "a.go", 3, 3)
	tr.Parameters.KeySize = 256
	pr := predAsset("a.go", 3, 3)
	pr.Algorithm, pr.Params.KeySize = "AES-256-GCM", 128
	rep := Score("t", manifest(entry("e", corpus.Go, corpus.SuiteA)), map[string]groundtruth.File{"e": exhaustive(tr)},
		map[string]prediction.File{"e": {Tool: prediction.Tool{Name: "t"}, Entry: "e", Assets: []prediction.Asset{pr}}})
	if rep.Overall.AttributedTP != 0 {
		t.Fatalf("attributed a pair whose explicit key_size 128 contradicts the truth")
	}
	want := `Name and field disagree in e, field used: gt-1 prediction: name "AES-256-GCM" implies key_size 256, explicit key_size 128 wins`
	if md := Markdown(rep); !strings.Contains(md, want) {
		t.Fatalf("markdown lacks %q:\n%s", want, md)
	}
	js, err := JSON(rep)
	if err != nil || !strings.Contains(string(js), `"name_conflicts": [`) || !strings.Contains(string(js), `explicit key_size 128 wins`) {
		t.Fatalf("json lacks the conflict (err %v):\n%s", err, js)
	}
}

func entry(id string, lang corpus.Language, suite corpus.Suite) corpus.Entry {
	return corpus.Entry{ID: id, Language: lang, Suite: suite}
}

func manifest(entries ...corpus.Entry) corpus.Manifest { return corpus.Manifest{Entries: entries} }

func TestEntriesWithoutPredictionsAreNotRunAndWithoutTruthSkipped(t *testing.T) {
	m := manifest(entry("ran", corpus.Go, corpus.SuiteA), entry("idle", corpus.Go, corpus.SuiteA), entry("bare", corpus.Go, corpus.SuiteA))
	truth := map[string]groundtruth.File{
		"ran":  exhaustive(truthAsset("gt-1", "a.go", 3, 3)),
		"idle": exhaustive(truthAsset("gt-1", "a.go", 3, 3), truthAsset("gt-2", "a.go", 9, 9)),
	}
	preds := map[string]prediction.File{
		"ran":  {Tool: prediction.Tool{Name: "t"}, Entry: "ran", Assets: []prediction.Asset{predAsset("a.go", 3, 3)}},
		"bare": {Tool: prediction.Tool{Name: "t"}, Entry: "bare", Assets: []prediction.Asset{predAsset("a.go", 3, 3)}},
	}
	rep := Score("t", m, truth, preds)
	got := fmt.Sprintf("scored=%v notrun=%v skipped=%v TP=%d FN=%d FP=%d", rep.Scored, rep.NotRun, rep.Skipped, rep.Overall.TP, rep.Overall.FN, rep.Overall.FP)
	if want := "scored=[ran] notrun=[idle] skipped=[{bare no ground truth}] TP=1 FN=0 FP=0"; got != want {
		t.Fatalf("got %s\nwant %s", got, want)
	}
}

func TestReachabilityIsNAForToolWithoutIt(t *testing.T) {
	m := manifest(entry("e", corpus.Java, corpus.SuiteB))
	truth := map[string]groundtruth.File{"e": exhaustive(truthAsset("gt-1", "a.go", 3, 3))}
	pf := prediction.File{Tool: prediction.Tool{Name: "t"}, Entry: "e", Assets: []prediction.Asset{predAsset("a.go", 3, 3)}}
	rep := Score("t", m, truth, map[string]prediction.File{"e": pf})
	if rep.Reachability != nil || rep.Pairs != nil {
		t.Fatalf("reachability = %+v, pairs = %+v; want both n/a", rep.Reachability, rep.Pairs)
	}
	md := Markdown(rep)
	if !strings.Contains(md, "## Reachability\n\nn/a: no scored entry's predictions report reachability.") {
		t.Fatalf("markdown does not print reachability as n/a:\n%s", md)
	}

	pf.ReachabilityReported = true
	rep = Score("t", m, truth, map[string]prediction.File{"e": pf})
	if got, want := fmt.Sprint(rep.Reachability.Overall[true][prediction.NoAnswer], rep.Reachability.Overall.ReachableRecall()), "1 0.000"; got != want {
		t.Fatalf("no-answer cell, reachable recall = %s, want %s", got, want)
	}
}

func TestReachabilityConfusionByDepth(t *testing.T) {
	known := func(a groundtruth.Asset, reachable bool, d groundtruth.Depth) groundtruth.Asset {
		a.Reachability = groundtruth.Known{Reachable: reachable, Depth: d}
		return a
	}
	unknown := truthAsset("gt-4", "a.go", 40, 40)
	unknown.Reachability = groundtruth.Unknown{}
	truth := exhaustive(known(truthAsset("gt-1", "a.go", 3, 3), true, 0), known(truthAsset("gt-2", "a.go", 10, 10), true, 2),
		known(truthAsset("gt-3", "a.go", 20, 20), false, 2), unknown)
	answer := func(p prediction.Asset, a prediction.Answer) prediction.Asset { p.Reachable = a; return p }
	preds := []prediction.Asset{answer(predAsset("a.go", 3, 3), prediction.Reachable), answer(predAsset("a.go", 10, 10), prediction.Unreachable),
		answer(predAsset("a.go", 20, 20), prediction.Reachable), answer(predAsset("a.go", 40, 40), prediction.Reachable)}
	rep := Score("t", manifest(entry("e", corpus.Go, corpus.SuiteB)), map[string]groundtruth.File{"e": truth},
		map[string]prediction.File{"e": {Tool: prediction.Tool{Name: "t"}, Entry: "e", ReachabilityReported: true, Assets: preds}})
	rc := rep.Reachability
	got := fmt.Sprint(rc.Overall[true], rc.Overall[false], rc.ByDepth["0"][true], rc.ByDepth["2+"][true], rc.ByDepth["2+"][false], len(rc.ByDepth["1"]),
		rc.Overall.ReachablePrecision(), rc.Overall.ReachableRecall())
	want := "map[false:1 true:1] map[true:1] map[true:1] map[false:1] map[true:1] 0 0.500 0.500"
	if got != want {
		t.Fatalf("got  %s\nwant %s", got, want)
	}
}

func TestPairDiscrimination(t *testing.T) {
	r := corpus.Entry{ID: "p-r", Language: corpus.Java, Suite: corpus.SuiteB, Pair: &corpus.Pair{Partner: "p-u", Variant: corpus.Reachable}}
	u := corpus.Entry{ID: "p-u", Language: corpus.Java, Suite: corpus.SuiteB, Pair: &corpus.Pair{Partner: "p-r", Variant: corpus.Unreachable}}
	appHash := truthAsset("gt-app", "App.java", 30, 30) // reachable in both variants; it does not make the pair
	libR := truthAsset("gt-lib", "Lib.java", 5, 5)
	libU := libR
	libU.Reachability = groundtruth.Known{Reachable: false, Depth: 2}
	truth := map[string]groundtruth.File{"p-r": exhaustive(libR, appHash), "p-u": exhaustive(libU, appHash)}
	pred := func(entry string, lib prediction.Answer) prediction.File {
		l, a := predAsset("Lib.java", 5, 5), predAsset("App.java", 30, 30)
		l.Reachable, a.Reachable = lib, prediction.Reachable
		return prediction.File{Tool: prediction.Tool{Name: "t"}, Entry: entry, ReachabilityReported: true, Assets: []prediction.Asset{l, a}}
	}
	cases := []struct {
		name         string
		inR, inU     prediction.Answer
		correct, why string
	}{
		{"discriminates", prediction.Reachable, prediction.Unreachable, "1/1", ""},
		{"says reachable in both", prediction.Reachable, prediction.Reachable, "0/1", "p-u: gt-lib reported reachable"},
		{"misses the reachable variant", prediction.NoAnswer, prediction.Unreachable, "0/1", "p-r: gt-lib not found reachable"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rep := Score("t", manifest(r, u), truth, map[string]prediction.File{"p-r": pred("p-r", c.inR), "p-u": pred("p-u", c.inU)})
			got := fmt.Sprintf("%d/%d", rep.Pairs.Correct(), len(rep.Pairs.Results))
			if got != c.correct || rep.Pairs.Results[0].Why != c.why {
				t.Fatalf("got %s %q, want %s %q", got, rep.Pairs.Results[0].Why, c.correct, c.why)
			}
		})
	}
	rep := Score("t", manifest(r, u), truth, map[string]prediction.File{"p-r": pred("p-r", prediction.Reachable)})
	if got, want := fmt.Sprint(len(rep.Pairs.Results), rep.Pairs.Skipped), "0 [{p-r / p-u p-u was not run}]"; got != want {
		t.Fatalf("partner not run: got %s, want %s", got, want)
	}
}

func TestMetricsPrintNAOnZeroDenominators(t *testing.T) {
	c := Counts{FN: 2}
	if got, want := fmt.Sprint(c.Precision(), c.Recall(), c.F1(), Counts{}.F1()), "n/a 0.000 0.000 n/a"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestBreakdownsAndPoolLabel(t *testing.T) {
	m := manifest(entry("j", corpus.Java, corpus.SuiteA), entry("g", corpus.Go, corpus.SuiteC))
	proto := truthAsset("gt-2", "b.go", 5, 5)
	proto.Type = groundtruth.Protocol
	pooled := groundtruth.File{Entry: "g", Scope: groundtruth.Scope{Kind: groundtruth.Pool}, Assets: []groundtruth.Asset{truthAsset("gt-1", "a.go", 3, 3), proto}}
	truth := map[string]groundtruth.File{"j": exhaustive(truthAsset("gt-1", "A.java", 3, 3)), "g": pooled}
	preds := map[string]prediction.File{
		"j": {Tool: prediction.Tool{Name: "t"}, Entry: "j", Assets: []prediction.Asset{predAsset("A.java", 3, 3)}},
		"g": {Tool: prediction.Tool{Name: "t"}, Entry: "g", Assets: []prediction.Asset{predAsset("a.go", 3, 3)}},
	}
	md := Markdown(Score("t", m, truth, preds))
	for _, want := range []string{
		"| all † | 1.000 | 0.667 | 0.800 | 2 | 0 | 1 | 0 |",
		"| C † | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |",
		"| A | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |",
		"| go / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |",
		"| java / algorithm | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |",
		"| g † | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |",
		"| j | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |",
		"| j | 1 / 1 (1.000) | 1.000 | 1.000 |",
		"Rows marked † include entries with scope.kind pool; their recall is relative to pool.",
	} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown lacks %q", want)
		}
	}
	if t.Failed() {
		t.Log(md)
	}
}
