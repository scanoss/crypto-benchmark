package score

import (
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
	"github.com/scanoss/crypto-benchmark/internal/groundtruth"
	"github.com/scanoss/crypto-benchmark/internal/prediction"
)

func poolPreds(entry string, assets ...prediction.Asset) prediction.File {
	return prediction.File{Entry: entry, Assets: assets}
}

func TestUniqueCountsAssetsOnlyOneRunMatched(t *testing.T) {
	cert := truthAsset("gt-3", "c.go", 30, 30)
	cert.Type = groundtruth.Certificate
	pool := groundtruth.File{Entry: "go-app-pool", Scope: groundtruth.Scope{Kind: groundtruth.Pool}, Assets: []groundtruth.Asset{
		truthAsset("gt-1", "a.go", 10, 10), truthAsset("gt-2", "b.go", 20, 20), cert,
	}}
	files := groundtruth.File{Entry: "go-app", Scope: groundtruth.Scope{Kind: groundtruth.Files, Files: []string{"a.go"}},
		Assets: []groundtruth.Asset{truthAsset("gt-1", "a.go", 10, 10)}}
	m := manifest(entry("go-app", corpus.Go, corpus.SuiteC), entry("go-app-pool", corpus.Go, corpus.SuiteC), entry("go-lib-pool", corpus.Go, corpus.SuiteC))
	truth := map[string]groundtruth.File{"go-app": files, "go-app-pool": pool, "go-lib-pool": {Entry: "go-lib-pool", Scope: groundtruth.Scope{Kind: groundtruth.Pool}}}
	runs := []Run{
		{"A", map[string]prediction.File{
			"go-app":      poolPreds("go-app", predAsset("a.go", 10, 10)),
			"go-app-pool": poolPreds("go-app-pool", predAsset("a.go", 10, 10), predAsset("b.go", 20, 20)),
		}},
		{"B", map[string]prediction.File{"go-app-pool": poolPreds("go-app-pool", predAsset("b.go", 20, 20))}},
		{"C", map[string]prediction.File{}},
	}
	r := Unique(m, truth, runs)
	if r.Assets != 3 || strings.Join(r.Entries, ",") != "go-app-pool,go-lib-pool" {
		t.Fatalf("pool entries %v with %d assets, want go-app-pool,go-lib-pool with 3; files entries stay out", r.Entries, r.Assets)
	}
	want := []struct{ found, only, notRun int }{{2, 1, 1}, {1, 0, 1}, {0, 0, 2}}
	for i, w := range want {
		row := r.Runs[i]
		if row.Found != w.found || row.Only != w.only || len(row.EntriesNoRun) != w.notRun {
			t.Errorf("%s: found %d only %d not run %v, want %d %d %d", row.Label, row.Found, row.Only, row.EntriesNoRun, w.found, w.only, w.notRun)
		}
	}
	if r.Runs[0].OnlyByLang["go"] != 1 || r.Runs[0].OnlyByType["algorithm"] != 1 {
		t.Errorf("A only by language %v, by type %v", r.Runs[0].OnlyByLang, r.Runs[0].OnlyByType)
	}
	if r.None != 1 || r.NoneByType["certificate"] != 1 {
		t.Errorf("found by none %d %v, want the certificate", r.None, r.NoneByType)
	}
	md := UniqueMarkdown(r)
	for _, s := range []string{"| A | 2 | 1 | go-lib-pool |", "| found by none of these runs | | 1 | |", "| certificate | 0 | 0 | 0 | 1 |", "relative to the pool"} {
		if !strings.Contains(md, s) {
			t.Errorf("markdown lacks %q:\n%s", s, md)
		}
	}
}
