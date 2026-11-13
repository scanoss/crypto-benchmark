# sonar-cryptography, suite C first `pool` score

**sonar-cryptography 1.7.0 matches 169 of the 590 adjudicated pool assets in its three languages with 3 false positives: precision 0.983, recall 0.286 relative to the pool.** These are the frozen predictions of its [first score](../suite-c/README.md), re-keyed to the `<id>-pool` entries. Nothing was rerun. [suite-c.md](../../suite-c.md#pool-scope-first-score) compares the tools.

| Level | Precision | Recall, relative to pool | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.983 | 0.286 | 0.444 | 169 | 3 | 421 | 16 | 166 / 169 (0.982) |
| object | 0.983 | 0.300 | 0.460 | 169 | 3 | 394 | 16 | 166 / 169 (0.982) |

- **Scope:** `pool`, six entries (Go, Java, Python). The Rust and JavaScript entries are not run: sonar-cryptography 1.7.0 has no support for them. Recall counts only the assets some tool or the grep found.
- **Dropped outside the pool frame:** none.
- **By language:** Go 168 / 0 / 199 (1.000 / 0.458), Java 1 / 3 / 10 (0.250 / 0.091), Python 0 / 0 / 212. The 3 FP are jadx's `getThumbPrint(x509cert, "MD5")`, `"SHA-1"` and `"SHA-256"` calls (`CertificateManager.java` 83-85): the algorithm names passed to a project helper, reported where they are passed instead of at the `MessageDigest.getInstance` call inside it (line 158), which it does not report.

`score.md` is the full `cbench score` report. `predictions/sonar-cryptography/` holds the prediction files.

## Run record

- **Predictions:** `results/sonar-cryptography/suite-c/predictions/sonar-cryptography/ID.json`, unchanged, with one line changed: `"entry": "ID"` became `"entry": "ID-pool"`. Every asset is byte-identical.
- **Tool run:** none. The scan is recorded in [suite-c/README.md](../suite-c/README.md): 1.7.0 in CBOMkit-action v2.3.0, test files scanned.
- **Commands,** per entry `ID`:

```sh
mkdir -p results/sonar-cryptography/suite-c-pool/predictions/sonar-cryptography
sed 's/^  "entry": "ID",$/  "entry": "ID-pool",/' \
  results/sonar-cryptography/suite-c/predictions/sonar-cryptography/ID.json \
  > results/sonar-cryptography/suite-c-pool/predictions/sonar-cryptography/ID-pool.json
go run ./cmd/cbench score --tool sonar-cryptography \
  --predictions results/sonar-cryptography/suite-c-pool/predictions \
  --out results/sonar-cryptography/suite-c-pool/score.md
```
