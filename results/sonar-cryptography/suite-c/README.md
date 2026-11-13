# sonar-cryptography, suite C first score

sonar-cryptography 1.7.0, run on 2026-11-13 over the six suite C entries in the languages it supports (Go, Java, Python) and scored on the `files` scope: the 142 sampled files of the ten repositories, of which 90 files and 125 assets are in these six. This is the [first score](../../../METHODOLOGY.md#first-score-freeze), published as is. [suite-c.md](../../suite-c.md) compares the tools.

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 1.000 | 0.224 | 0.366 | 28 | 0 | 97 |
| object | 1.000 | 0.228 | 0.371 | 28 | 0 | 95 |

`score.md` is the full `cbench score` report, `predictions/sonar-cryptography/` the adapted predictions, `commands/` the docker command `scripts/run-sonar-cryptography.sh` recorded for each entry and `build/` the two Java build helpers below. 158 predictions fall outside the sampled files and are ignored.

## Run record

- **sonar-cryptography:** 1.7.0, embedded in CBOMkit-action v2.3.0 through cbomkit-lib 1.3.0 (the script checked the three versions on every run). Image `ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b`, Docker 29.1.3, no network during the scan, `CBOMKIT_EXCLUDE` matching nothing so test files are scanned.
- **Checkouts:** each repository at the commit pinned in `corpus/manifest.yaml`, fetched with `cbench fetch` into `work/corpus/ID`, scanned whole.
- **Not run:** `rust-c-*` and `javascript-c-*`. sonar-cryptography 1.7.0 has no Rust or JavaScript support, as in suite A.
- **Go and Python,** as recorded for suite A:

```sh
scripts/run-sonar-cryptography.sh work/corpus/ID work/sonar-cryptography/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0 --entry ID \
  --input work/sonar-cryptography/ID.cbom.json \
  --out results/sonar-cryptography/suite-c/predictions/sonar-cryptography/ID.json
```

Both Python repositories have `pyproject.toml` at the root, so no `setup.cfg` marker was added; both Go repositories have `go.mod` at the root.

- **Java,** built in a copy with Temurin JDK 21.0.12, then scanned whole:

```sh
B=$PWD/results/sonar-cryptography/suite-c/build
cp -a work/corpus/java-c-mall work/build/java-c-mall && (cd work/build/java-c-mall && \
  mvn -q -s $B/mall-settings.xml -Ddocker.skip=true package dependency:copy-dependencies)
cp -a work/corpus/java-c-jadx work/build/java-c-jadx && (cd work/build/java-c-jadx && \
  ./gradlew --no-daemon -I $B/copydeps.init.gradle assemble testClasses copyDeps)
scripts/run-sonar-cryptography.sh work/build/ID work/sonar-cryptography/ID.cbom.json   # with deviation 1
```

## Deviations from the recorded configuration

| # | Deviation | Why | Effect |
|---|---|---|---|
| 1 | The runner's check that a Java entry has `src/` at its root was removed for the two Java entries; nothing else in the script changed (same image, environment and versions). | Both repositories are multi-module builds: `src/` lives in each module. | cbomkit scanned every module: mall 7 projects, all 524 `.java` files; jadx 18 projects, 2,003 files. |
| 2 | mall: `-s build/mall-settings.xml` mirrors the pom's `aliyunmaven` repository to Maven Central, and `-Ddocker.skip=true` skips the docker-maven-plugin bound to `package`. | The original repository hung on a TLS read for over 12 minutes; the plugin tries to reach a Docker daemon at an address in the pom. | The build is the recorded `mvn package dependency:copy-dependencies`; its tests ran and passed. |
| 3 | jadx: Gradle has no `dependency:copy-dependencies`, so `build/copydeps.init.gradle`, an init script outside the entry, copies each project's runtime and test runtime jars into `<project>/build/dependency`. Built with jadx's own Gradle wrapper (9.8.2), `assemble testClasses`, tests not run. | The scanner resolves types from classes and jars inside the entry. | 8,491 `.class` files and 730 dependency jars. jadx's 37 Kotlin files are not analyzed by the Java frontend. |

The adapter read every CBOM without a note: no component without a location and no dropped location.

## What it reports

- **Go (27 TP of 61, no FP).** Standard library hashes in non-test and test files alike, including `crypto/sha256`. It misses `rand.Read` draws, the `tls.NewListener` session in caveman, ollama's Ed25519 key generation and SSH key serialization, and 25 of the 33 SHA-256 hashes in `_test.go` files.
- **Java (1 of 3).** It finds jadx's MD5 file hash and misses apksig's `ApkVerifier.verify()` and Spring Security's `BCryptPasswordEncoder`.
- **Python (0 of 61).** It scanned 687 `.py` files in graphify and 1,124 in yt-dlp and reported no component in either. The Python assets here are `hashlib`, `hmac`, `secrets` and `ssl` calls and yt-dlp's own AES module; its suite A Python matches were in other libraries.
