package scripts

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// fakeCdxgen stands in for cdxgen, evinse and cbom: it answers the version
// probe and writes an empty BOM to the path after -o.
const fakeCdxgen = `#!/usr/bin/env bash
if [ "${1:-}" = --version ]; then echo "CycloneDX Generator 13.6.0"; exit 0; fi
while [ $# -gt 0 ]; do
	if [ "$1" = -o ]; then echo '{}' >"$2"; fi
	shift
done
`

func TestCdxgenRunnerRecordsCommandsWithoutTrailingSpace(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not available")
	}
	dir := t.TempDir()
	bin := filepath.Join(dir, "bin")
	if err := os.MkdirAll(bin, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"cdxgen", "evinse", "cbom"} {
		if err := os.WriteFile(filepath.Join(bin, name), []byte(fakeCdxgen), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	entry := filepath.Join(dir, "entry")
	if err := os.MkdirAll(entry, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(entry, "go.mod"), []byte("module x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "res", "entry.cbom.json")
	cmd := exec.Command("bash", "run-cdxgen.sh", entry, out)
	cmd.Env = append(os.Environ(), "CDXGEN_BIN="+bin)
	var errBuf strings.Builder
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		t.Fatalf("runner failed: %v\n%s", err, errBuf.String())
	}
	record, err := os.ReadFile(out + ".command")
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(record), "\n"), "\n")
	if len(lines) != 3 || lines[0] != "# CycloneDX Generator 13.6.0" ||
		!strings.HasPrefix(lines[1], "cdxgen -t go -o ") || !strings.HasSuffix(lines[1], " "+entry) ||
		!strings.HasPrefix(lines[2], "evinse -i ") || !strings.HasSuffix(lines[2], " -l go "+entry) {
		t.Fatalf("OUT.command:\n%q", record)
	}
}
