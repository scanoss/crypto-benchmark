// Package fetch materializes a remote corpus entry at its pinned commit.
package fetch

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/scanoss/crypto-benchmark/internal/corpus"
)

type Outcome int

const (
	Fetched Outcome = iota
	AlreadyPresent
)

// Fetch checks out src at src.SHA into <destRoot>/<id>. A checkout already at
// src.SHA is left alone. Any other existing directory is an error and is never
// removed. The clone is built in a temporary sibling and renamed into place,
// so an interrupted run never leaves a half-made <destRoot>/<id>.
func Fetch(ctx context.Context, id string, src corpus.Remote, destRoot string) (Outcome, error) {
	final := filepath.Join(destRoot, id)
	if _, err := os.Stat(final); err == nil {
		head, err := git(ctx, "", "--git-dir", filepath.Join(final, ".git"), "rev-parse", "HEAD")
		if err != nil {
			return 0, fmt.Errorf("fetch: %s exists but is not a git checkout with a HEAD; remove it to refetch: %w", final, err)
		}
		if head != src.SHA {
			return 0, fmt.Errorf("fetch: %s is at %s, want %s; remove it to refetch", final, head, src.SHA)
		}
		return AlreadyPresent, nil
	} else if !os.IsNotExist(err) {
		return 0, fmt.Errorf("fetch: %w", err)
	}

	if err := os.MkdirAll(destRoot, 0o755); err != nil {
		return 0, fmt.Errorf("fetch: %w", err)
	}
	tmp, err := os.MkdirTemp(destRoot, "."+id+".partial-")
	if err != nil {
		return 0, fmt.Errorf("fetch: %w", err)
	}
	defer os.RemoveAll(tmp)

	steps := [][]string{
		{"init", "-q"},
		{"remote", "add", "origin", src.URL},
		{"fetch", "-q", "--depth", "1", "origin", src.SHA},
		{"checkout", "-q", "--detach", "FETCH_HEAD"},
	}
	for _, args := range steps {
		if _, err := git(ctx, tmp, args...); err != nil {
			return 0, fmt.Errorf("fetch: %s: %w", id, err)
		}
	}
	head, err := git(ctx, tmp, "rev-parse", "HEAD")
	if err != nil {
		return 0, fmt.Errorf("fetch: %s: %w", id, err)
	}
	if head != src.SHA {
		return 0, fmt.Errorf("fetch: %s: checked out %s, want %s", id, head, src.SHA)
	}
	if err := os.Rename(tmp, final); err != nil {
		return 0, fmt.Errorf("fetch: %w", err)
	}
	return Fetched, nil
}

func git(ctx context.Context, dir string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
