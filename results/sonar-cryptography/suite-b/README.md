# sonar-cryptography, suite B (transitive reachability)

Run on 2026-11-05 over the 42 suite B entries of the languages it supports: Go, Python and Java. JavaScript and Rust have no sonar-cryptography rules and were not run, as in suite A. `score.md` is the `cbench score` report.

**sonar-cryptography has no dependency analysis and no reachability.** It analyzes the source of the scanned project only ([METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration)), and every suite B asset is in a dependency. It detected none of the 52 assets in those languages and reported nothing else, so it is shown as having no reachability, and its detection recall on suite B is 0 by design, not by a missed rule.

## Run record

- **sonar-cryptography 1.7.0** headless in CBOMkit-action v2.3.0 through `scripts/run-sonar-cryptography.sh ENTRY_DIR OUT`, on each entry built by `scripts/prepare-suite-b.sh`.
- Go and Python: the built application directory (`work/suite-b/<entry>/app`). The Python libraries are installed in a virtual environment outside it, as a user's would be; the Go libraries are the modules its `go.mod` replaces.
- Java: the built application with its dependencies copied in (`mvn package dependency:copy-dependencies` against the entry's Maven repository), so `target/dependency` holds the library jars the scanner resolves types from. The library sources are not part of the project, as they are not in a real one.
- **Adapter:** `cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0`.

| Language | Entries | Assets | Detected |
|---|---|---|---|
| Go | 14 | 16 | 0 |
| Python | 14 | 16 | 0 |
| Java | 14 | 20 | 0 |
