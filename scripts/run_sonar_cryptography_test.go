package scripts

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// fakeDocker answers the runner's version probe with the given lines and, for
// the scan, writes a CBOM into the directory mounted at /scan/out and logs
// the files it was handed in /scan/ws.
const fakeDocker = `#!/usr/bin/env bash
case " $* " in
*" --entrypoint sh "*)
	echo "$FAKE_VERSIONS" | tr '|' '\n'
	exit 0 ;;
esac
while [ $# -gt 0 ]; do
	case "$1" in
	-v)
		case "$2" in
		*:/scan/out) out=${2%%:*} ;;
		*:/scan/ws:ro) ws=${2%%:*} ;;
		esac
		shift ;;
	esac
	shift
done
ls -A "$ws" >"$FAKE_LOG"
echo '{"bomFormat":"CycloneDX","specVersion":"1.6","components":[]}' >"$out/cbom.json"
`

const goodVersions = "com.ibm/sonar-cryptography-plugin 1.7.0|org.pqca/cbomkit-lib 1.3.0|org.pqca/CBOMkit-action v2.3.0"

func runRunner(t *testing.T, versions string) (out string, stderr string, err error) {
	t.Helper()
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not available")
	}
	dir := t.TempDir()
	docker := filepath.Join(dir, "docker")
	if err := os.WriteFile(docker, []byte(fakeDocker), 0o755); err != nil {
		t.Fatal(err)
	}
	entry := filepath.Join(dir, "entry")
	if err := os.MkdirAll(filepath.Join(entry, "a"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(entry, "a", "main.py"), []byte("import hashlib\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out = filepath.Join(dir, "res", "entry.cbom.json")
	cmd := exec.Command("bash", "run-sonar-cryptography.sh", entry, out)
	cmd.Env = append(os.Environ(), "DOCKER="+docker, "FAKE_VERSIONS="+versions,
		"FAKE_LOG="+filepath.Join(dir, "ws.txt"))
	var errBuf strings.Builder
	cmd.Stderr = &errBuf
	err = cmd.Run()
	return out, errBuf.String(), err
}

func TestSonarCryptographyRunnerRefusesWrongPluginVersion(t *testing.T) {
	wrong := strings.Replace(goodVersions, "plugin 1.7.0", "plugin 1.6.1", 1)
	out, stderr, err := runRunner(t, wrong)
	if err == nil {
		t.Fatal("runner accepted sonar-cryptography-plugin 1.6.1")
	}
	if !strings.Contains(stderr, "com.ibm/sonar-cryptography-plugin 1.6.1") {
		t.Errorf("stderr does not name the version found:\n%s", stderr)
	}
	if _, err := os.Stat(out); err == nil {
		t.Error("runner wrote OUT despite the wrong version")
	}
}

func TestSonarCryptographyRunnerRecordsAndMarksPythonRoot(t *testing.T) {
	out, stderr, err := runRunner(t, goodVersions)
	if err != nil {
		t.Fatalf("runner failed: %v\n%s", err, stderr)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("OUT not written: %v", err)
	}
	record, err := os.ReadFile(out + ".command")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"sonar-cryptography-plugin 1.7.0", "touch WS/setup.cfg", "CBOMKIT_LANGUAGES=python", "-w /scan/ws"} {
		if !strings.Contains(string(record), want) {
			t.Errorf("OUT.command lacks %q:\n%s", want, record)
		}
	}
	ws, err := os.ReadFile(filepath.Join(filepath.Dir(filepath.Dir(out)), "ws.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Fields(string(ws)); strings.Join(got, " ") != "a setup.cfg" {
		t.Errorf("scanned workspace holds %v, want [a setup.cfg]", got)
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(filepath.Dir(out)), "entry", "setup.cfg")); err == nil {
		t.Error("runner wrote the marker into ENTRY_DIR instead of its copy")
	}
}
