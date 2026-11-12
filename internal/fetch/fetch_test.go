package fetch

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

func run(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-c", "user.name=t", "-c", "user.email=t@example.com"}, args...)...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

// upstream creates a repository with two commits and returns its file:// URL
// and both commit SHAs, oldest first.
func upstream(t *testing.T) (url, first, second string) {
	t.Helper()
	dir := t.TempDir()
	run(t, dir, "init", "-q")
	for _, name := range []string{"first.txt", "second.txt"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o644); err != nil {
			t.Fatal(err)
		}
		run(t, dir, "add", name)
		run(t, dir, "commit", "-q", "-m", name)
		if first == "" {
			first = run(t, dir, "rev-parse", "HEAD")
		}
	}
	return "file://" + dir, first, run(t, dir, "rev-parse", "HEAD")
}

func TestFetchPinsCommitAndIsIdempotent(t *testing.T) {
	url, first, _ := upstream(t)
	dest := t.TempDir()
	ctx := context.Background()

	out, err := Fetch(ctx, "demo", corpus.Remote{URL: url, SHA: first}, dest)
	if err != nil || out != Fetched {
		t.Fatalf("first fetch: outcome %v, err %v", out, err)
	}
	checkout := filepath.Join(dest, "demo")
	if got := run(t, checkout, "rev-parse", "HEAD"); got != first {
		t.Fatalf("HEAD = %s, want %s", got, first)
	}
	if _, err := os.Stat(filepath.Join(checkout, "first.txt")); err != nil {
		t.Fatalf("first.txt missing at the pinned commit: %v", err)
	}
	if _, err := os.Stat(filepath.Join(checkout, "second.txt")); !os.IsNotExist(err) {
		t.Fatalf("second.txt exists, so the checkout is not at the pinned (older) commit: %v", err)
	}
	if entries, _ := os.ReadDir(dest); len(entries) != 1 {
		t.Fatalf("dest holds %d entries, want only demo (no leftover partial clone)", len(entries))
	}

	out, err = Fetch(ctx, "demo", corpus.Remote{URL: url, SHA: first}, dest)
	if err != nil || out != AlreadyPresent {
		t.Fatalf("rerun: outcome %v, err %v, want AlreadyPresent", out, err)
	}
}

func TestFetchRefusesCheckoutAtAnotherCommit(t *testing.T) {
	url, first, second := upstream(t)
	dest := t.TempDir()
	ctx := context.Background()
	if _, err := Fetch(ctx, "demo", corpus.Remote{URL: url, SHA: first}, dest); err != nil {
		t.Fatal(err)
	}

	_, err := Fetch(ctx, "demo", corpus.Remote{URL: url, SHA: second}, dest)
	if err == nil || !strings.Contains(err.Error(), "is at "+first+", want "+second) {
		t.Fatalf("err = %v, want refusal naming both commits", err)
	}
	if got := run(t, filepath.Join(dest, "demo"), "rev-parse", "HEAD"); got != first {
		t.Fatalf("existing checkout changed to %s, want it left at %s", got, first)
	}
}

func TestFetchRefusesNonCheckoutDirectoryInsideAnotherRepo(t *testing.T) {
	url, first, _ := upstream(t)
	outer := t.TempDir()
	run(t, outer, "init", "-q")
	run(t, outer, "commit", "-q", "--allow-empty", "-m", "outer")
	dest := filepath.Join(outer, "work")
	keep := filepath.Join(dest, "demo", "keep.txt")
	if err := os.MkdirAll(filepath.Dir(keep), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keep, []byte("mine"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := Fetch(context.Background(), "demo", corpus.Remote{URL: url, SHA: first}, dest)
	if err == nil || !strings.Contains(err.Error(), "is not a git checkout") {
		t.Fatalf("err = %v, want refusal for a directory that is not a checkout", err)
	}
	if b, err := os.ReadFile(keep); err != nil || string(b) != "mine" {
		t.Fatalf("existing directory was modified: %q, %v", b, err)
	}
}

func TestFetchUnknownCommitLeavesNothingBehind(t *testing.T) {
	url, _, _ := upstream(t)
	dest := t.TempDir()
	_, err := Fetch(context.Background(), "demo", corpus.Remote{URL: url, SHA: strings.Repeat("ab", 20)}, dest)
	if err == nil {
		t.Fatal("fetching a commit the remote does not have succeeded")
	}
	if entries, _ := os.ReadDir(dest); len(entries) != 0 {
		t.Fatalf("dest holds %d entries after a failed fetch, want none", len(entries))
	}
}
