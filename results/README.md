# Suite A results

crypto-finder v0.54.0 with remote ruleset v1.48.0, scored on 2026-11-17, finds 1,413 of the 1,416 suite A assets with 9 false positives, and passes 28 of 30 target cells. sonar-cryptography 1.7.0 and cdxgen 13.6.0 pass none.

| Tool | Languages | TP / FP / FN op | P / R op | P / R obj | Target cells, obj |
|---|---|---|---|---|---|
| crypto-finder v0.54.0, ruleset v1.48.0 | 5 | 1,413 / 9 / 3 | 0.994 / 0.998 | 0.991 / 0.998 | 28 of 30 |
| sonar-cryptography 1.7.0 | 3 (no JavaScript or Rust support) | 150 / 43 / 696 | 0.777 / 0.177 | 0.728 / 0.197 | 0 of 20 |
| cdxgen 13.6.0 | 5 | 112 / 77 / 1,304 | 0.593 / 0.079 | 0.560 / 0.097 | 0 of 30 |

- **Ground truth:** the one of 2026-11-08, 1,416 assets in 1,015 objects (Python 295 / 212, JavaScript 373 / 279, Rust 197 / 151, Go 202 / 152, Java 349 / 221), 75 of them with anchors ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules)).
- **Matching:** strict. Line ranges must overlap (`LINE_TOLERANCE = 0`), ties prefer an attributed pair, and an algorithm asset also matches on its anchors, the call that names its algorithm for the same object in the same function.
- **Per language:** [Python](python-pilot.md), [JavaScript](javascript-pilot.md), [Rust](rust-pilot.md), [Go](go-pilot.md), [Java](java-pilot.md). Each page compares the tools and links to every run record.
- **Suite B**, transitive reachability: [suite-b.md](suite-b.md). crypto-finder discriminates all 45 pairs.
- **Suite C**, real repositories: [suite-c.md](suite-c.md). Status: complete, `files` and `pool` scored; reachability is phase 2. First score on the `files` scope (142 sampled files, 214 assets): crypto-finder 108 / 8 / 106 (0.931 / 0.505), sonar-cryptography 28 / 0 / 97 in three languages, cdxgen 18 / 33 / 196. First score on the `pool` scope (993 adjudicated assets, recall relative to the pool): crypto-finder 444 / 24 / 549 (0.949 / 0.447), sonar-cryptography 169 / 3 / 421 (0.983 / 0.286) in three languages, cdxgen 76 / 689 / 917 (0.099 / 0.077); assets only that tool found: 284, 71 and 9. Most of crypto-finder's misses are test files its default left out; a later run with `--include-tests` is shown beside it.

## Tools

| Tool | Version and configuration | Languages run |
|---|---|---|
| crypto-finder | v0.54.0, the released linux_amd64 binary (tarball sha256 `06255e9dace89710ee81308aedecd9e8e11b9ea37199b89afe9a81962971a1d0`, checked against the release's checksums file), which downloads the ruleset from `/v3/cryptography/rulesets` with a signed lease. Remote ruleset v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, pinned with `--ruleset-version`. | all five |
| sonar-cryptography | 1.7.0, headless in CBOMkit-action v2.3.0 with cbomkit-lib 1.3.0 | Python, Go, Java; it has no JavaScript or Rust support |
| cdxgen | 13.6.0: `cbom` for Python, JavaScript and Java, golem through `evinse -l go` for Go, rusi through `evinse -l rust` for Rust. For Python and Java it inventories libraries, keystores and certificates but not algorithm use in source, so it scores 0 there. | all five |

sonar-cryptography and cdxgen were run once; the corpus has not changed since, so their scores come from their stored predictions, scored against the current ground truth. [METHODOLOGY.md](../METHODOLOGY.md#competitor-configuration) has the exact configuration of every tool.

## Targets

Pre-registered on 2026-11-05: precision >= 0.90 and recall >= 0.85 in every language × asset type cell, in each language over all asset types (`all types`) and in each asset type over all languages (`all languages`). The object level is the headline; the operation level follows it ([METHODOLOGY.md](../METHODOLOGY.md#targets)). Each cell is `precision / recall` and its verdict; `n/a` has neither ground truth nor predictions.

### crypto-finder, object level (headline)

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | 0.972 / 1.000 pass | 1.000 / 1.000 pass | 0.800 / 1.000 fail | 1.000 / 1.000 pass | 0.972 / 1.000 pass |
| javascript | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| rust | 0.988 / 1.000 pass | 1.000 / 1.000 pass | 0.882 / 1.000 fail | 1.000 / 0.958 pass | 0.980 / 0.987 pass |
| go | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| java | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| all languages | 0.992 / 1.000 pass | 1.000 / 1.000 pass | 0.905 / 1.000 pass | 1.000 / 0.994 pass | 0.991 / 0.998 pass |

28 of 30 cells pass.

### crypto-finder, operation level

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | 0.981 / 1.000 pass | 1.000 / 1.000 pass | 0.800 / 1.000 fail | 1.000 / 1.000 pass | 0.980 / 1.000 pass |
| javascript | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| rust | 0.991 / 0.991 pass | 1.000 / 1.000 pass | 0.882 / 1.000 fail | 1.000 / 0.964 pass | 0.985 / 0.985 pass |
| go | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| java | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| all languages | 0.995 / 0.999 pass | 1.000 / 1.000 pass | 0.905 / 1.000 pass | 1.000 / 0.995 pass | 0.994 / 0.998 pass |

28 of 30 cells pass.

### sonar-cryptography, object level (headline)

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | 1.000 / 0.158 fail | n/a / 0.000 fail | n/a / 0.000 fail | 0.786 / 0.177 fail | 0.917 / 0.156 fail |
| go | 1.000 / 0.452 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | 1.000 / 0.276 fail |
| java | 0.950 / 0.270 fail | n/a / 0.000 fail | 1.000 / 0.500 fail | 0.026 / 0.014 fail | 0.500 / 0.181 fail |
| all languages | 0.981 / 0.273 fail | n/a / 0.000 fail | 1.000 / 0.050 fail | 0.226 / 0.066 fail | 0.728 / 0.197 fail |

0 of 20 cells pass.

### sonar-cryptography, operation level

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | 1.000 / 0.119 fail | n/a / 0.000 fail | n/a / 0.000 fail | 0.786 / 0.141 fail | 0.921 / 0.119 fail |
| go | 1.000 / 0.368 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | 1.000 / 0.243 fail |
| java | 0.970 / 0.260 fail | n/a / 0.000 fail | 1.000 / 0.500 fail | 0.026 / 0.011 fail | 0.623 / 0.189 fail |
| all languages | 0.986 / 0.236 fail | n/a / 0.000 fail | 1.000 / 0.050 fail | 0.226 / 0.054 fail | 0.777 / 0.177 fail |

0 of 20 cells pass.

### cdxgen, object level (headline)

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail |
| javascript | 1.000 / 0.137 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | 1.000 / 0.090 fail |
| rust | 1.000 / 0.051 fail | n/a / 0.000 fail | n/a / 0.000 fail | 0.424 / 0.521 fail | 0.460 / 0.192 fail |
| go | 1.000 / 0.258 fail | 0.000 / 0.000 fail | 0.500 / 0.100 fail | 0.328 / 0.404 fail | 0.506 / 0.289 fail |
| java | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail |
| all languages | 1.000 / 0.084 fail | 0.000 / 0.000 fail | 0.500 / 0.026 fail | 0.376 / 0.139 fail | 0.560 / 0.097 fail |

0 of 30 cells pass.

### cdxgen, operation level

| Language | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| python | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail |
| javascript | 1.000 / 0.122 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | 1.000 / 0.086 fail |
| rust | 1.000 / 0.036 fail | n/a / 0.000 fail | n/a / 0.000 fail | 0.443 / 0.482 fail | 0.477 / 0.157 fail |
| go | 1.000 / 0.188 fail | 0.000 / 0.000 fail | 0.500 / 0.100 fail | 0.371 / 0.418 fail | 0.533 / 0.243 fail |
| java | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail | n/a / 0.000 fail |
| all languages | 1.000 / 0.064 fail | 0.000 / 0.000 fail | 0.500 / 0.026 fail | 0.407 / 0.134 fail | 0.593 / 0.079 fail |

0 of 30 cells pass.

## Known gaps

What crypto-finder v0.54.0 with ruleset v1.48.0 still gets wrong in suite A: 9 false positives and 3 false negatives at operation level, and one family of matches that is not attributed. Go, Java and JavaScript have none.

| Gap | Count | Where | What happens |
|---|---|---|---|
| TLS configuration setters counted as protocols | 4 FP | Python `ctx.minimum_version` (python-stdlib `tls-context/main.py:6`, python-truststore `inject/main.py:10`); Rust openssl `set_min_proto_version` and `set_cipher_list` (rust-openssl `tls_client.rs:5,6`) | crypto-finder reports each setter as a TLS protocol asset of its own. The ground truth treats it as configuration of the TLS context it is called on, which is already an asset. These four are the only reason the Python protocol (0.800 / 1.000) and Rust protocol (0.882 / 1.000) cells fail. |
| Hash class passed to a signing call | 3 FP | python-cryptography `x509-cert/main.py:22`; python-pyopenssl `cert-load/main.py:22`, `x509-cert/main.py:21` | In `.sign(key, hashes.SHA256())` the `hashes.SHA256()` argument is reported as a SHA-256 hash operation. The ground truth counts it as the signature's digest parameter. |
| Fernet `encrypt` | 1 FP | python-cryptography `fernet/main.py:8` | `f.encrypt(...)` is reported as an operation. The ground truth counts it as a follow-up on the `Fernet(key)` object. |
| Random generator construction | 1 FP | rust-curve25519-dalek `ristretto_dh.rs:7` | Creating the random number generator is reported as a CSPRNG use. The ground truth annotates the key generations that draw from it. |
| ECDSA verify inside a Rust macro | 1 FN | rust-openssl `ecdsa_p256.rs:12` | The verify call sits inside `assert!(...)` and is not reported. |
| ring PKCS#8 private keys | 2 FN | rust-ring `ecdsa_p256.rs:6`, `ed25519.rs:6` | `generate_pkcs8` is reported once, as the PKCS#8 encoding it returns; the generated private key is not reported as related crypto material of its own. |
| Unattributed ristretto255 curve | 4 matched, not attributed | rust-curve25519-dalek `ristretto_dh.rs:8,9,14,15` | The ristretto255 key generations and key agreements match and are reported as ECDH on the ristretto255 group. They stay unattributed because the benchmark's curve table (`internal/algo`) has no ristretto255, so the adapter does not carry the curve. This is a limit of the benchmark, not a tool error. |

## History of pinned runs

crypto-finder only, all five languages, every run against the ground truth of 2026-11-08 and the same matching rules. Each run pins the ruleset with `--ruleset-version`; every release is a published binary. The v0.54.0 run gave the same predictions as v0.53.1 with the same ruleset, asset for asset: v0.54.0 changes only how the ruleset is downloaded.

| crypto-finder | Ruleset | Run | TP / FP / FN op | P / R op | P / R obj | Target cells obj / op | Suite B pairs |
|---|---|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-06, rescored 2026-11-08 | 835 / 113 / 581 | 0.881 / 0.590 | 0.840 / 0.585 | 1 / 2 | 27 of 45 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 1,190 / 69 / 226 | 0.945 / 0.840 | 0.928 / 0.870 | 12 / 12 | 40 of 45 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 1,231 / 40 / 185 | 0.969 / 0.869 | 0.958 / 0.900 | 20 / 17 | 40 of 45 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 1,313 / 20 / 103 | 0.985 / 0.927 | 0.980 / 0.948 | 28 / 28 | 40 of 45 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 1,397 / 9 / 19 | 0.994 / 0.987 | 0.991 / 0.989 | 28 / 28 | 40 of 45 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 1,402 / 9 / 14 | 0.994 / 0.990 | 0.991 / 0.991 | 28 / 28 | 45 of 45 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 1,413 / 9 / 3 | 0.994 / 0.998 | 0.991 / 0.998 | 28 / 28 | 45 of 45 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 1,413 / 9 / 3 | 0.994 / 0.998 | 0.991 / 0.998 | 28 / 28 | 45 of 45 |

From v1.44.0 on, no run lost a true positive or added a false positive or false negative; every change removed one. Runs before 2026-11-08 were scored against an earlier ground truth and are not comparable. Each language page has the same table for its language.

## Comparison with the other tools

- **Coverage is the main difference.** crypto-finder finds 954 of 955 algorithm assets across the five languages; sonar-cryptography finds 24% in its three languages and cdxgen 6%. On related crypto material (IVs, nonces, salts, tags, shared secrets, seeds, private keys) F1 is 0.997 for crypto-finder, 0.201 for cdxgen and 0.087 for sonar-cryptography.
- **Precision on what each tool reports is closer.** Algorithm precision is 0.995 for crypto-finder (5 FP over five languages), 0.986 for sonar-cryptography (2 FP over three) and 1.000 for cdxgen (0 FP). sonar-cryptography attributes 84% of its algorithm matches, crypto-finder 80% and cdxgen 36%.
- **Best competitor per language, object-level F1:** sonar-cryptography in Python (0.266), Go (0.433) and Java (0.266); cdxgen in JavaScript (0.164) and Rust (0.271).
- **Granularity.** sonar-cryptography reports a crypto object once, where it is created. The object-level view exists so that this is not counted as missed crypto; its object-level recall is still 0.197.
- **Reachability is not measured in suite A.** Every suite A asset is reachable, and crypto-finder answers `reachable` on every match. sonar-cryptography and cdxgen report no reachability. Suite B scores it.

## How the scoring changed

The rules and ground truth changed four times before the pinned runs above, each time for every tool alike, and every tool was re-scored from its stored predictions. [METHODOLOGY.md](../METHODOLOGY.md#change-log) has the reasons.

| Date | Change | Effect on the scores at the time |
|---|---|---|
| 2026-11-04 | Related crypto material and the `hash` parameter enter the ground truth; the object-level view is added | Python pilot re-scored under both views |
| 2026-11-04 | Strict matching: `LINE_TOLERANCE` 2 -> 0, attributed tie-break | 111 pairs matched only through the tolerance, none the same operation on a neighbouring line ([strict-matching.md](../docs/strict-matching.md)). TP fell from 688 to 635 for crypto-finder, 156 to 115 for sonar-cryptography and 105 to 88 for cdxgen. |
| 2026-11-05 | Anchors on algorithm assets; Fernet `decrypt` becomes its own asset | 49 of the 57 pairs broken by strict matching return. TP rose from 635 to 671 for crypto-finder, 115 to 139 for sonar-cryptography and 88 to 89 for cdxgen. |
| 2026-11-08 | A generated private key is related crypto material | 150 suite A assets added. Every tool's recall over all types fell, because no tool reported all of them; no target verdict changed. |

## All tables

The output of `go run ./scripts/scoreboard`, which scores every `results/<tool>/<lang>-pilot/predictions` directory. `op` is the operation level, the primary view: every annotated operation counts. `obj` is the object level: each crypto object counts once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed is attributed TP over TP: the matches that also name the right algorithm and parameters.

### Language × tool

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| python | crypto-finder | 0.980 | 1.000 | 0.990 | 0.972 | 1.000 | 0.986 | 295 / 6 / 0 | 265 / 295 (0.898) | 199 / 212 (0.939) |
| python | sonar-cryptography | 0.921 | 0.119 | 0.210 | 0.917 | 0.156 | 0.266 | 35 / 3 / 260 | 31 / 35 (0.886) | 30 / 33 (0.909) |
| python | cdxgen | n/a | 0.000 | 0.000 | n/a | 0.000 | 0.000 | 0 / 0 / 295 | 0 / 0 (n/a) | 0 / 0 (n/a) |
| javascript | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 373 / 0 / 0 | 321 / 373 (0.861) | 259 / 279 (0.928) |
| javascript | cdxgen | 1.000 | 0.086 | 0.158 | 1.000 | 0.090 | 0.164 | 32 / 0 / 341 | 7 / 32 (0.219) | 7 / 25 (0.280) |
| rust | crypto-finder | 0.985 | 0.985 | 0.985 | 0.980 | 0.987 | 0.983 | 194 / 3 / 3 | 166 / 194 (0.856) | 134 / 149 (0.899) |
| rust | cdxgen | 0.477 | 0.157 | 0.237 | 0.460 | 0.192 | 0.271 | 31 / 34 / 166 | 18 / 31 (0.581) | 16 / 29 (0.552) |
| go | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 202 / 0 / 0 | 177 / 202 (0.876) | 144 / 152 (0.947) |
| go | sonar-cryptography | 1.000 | 0.243 | 0.390 | 1.000 | 0.276 | 0.433 | 49 / 0 / 153 | 40 / 49 (0.816) | 40 / 42 (0.952) |
| go | cdxgen | 0.533 | 0.243 | 0.333 | 0.506 | 0.289 | 0.368 | 49 / 43 / 153 | 35 / 49 (0.714) | 31 / 44 (0.705) |
| java | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 349 / 0 / 0 | 297 / 349 (0.851) | 210 / 221 (0.950) |
| java | sonar-cryptography | 0.623 | 0.189 | 0.290 | 0.500 | 0.181 | 0.266 | 66 / 40 / 283 | 57 / 66 (0.864) | 37 / 40 (0.925) |
| java | cdxgen | n/a | 0.000 | 0.000 | n/a | 0.000 | 0.000 | 0 / 0 / 349 | 0 / 0 (n/a) | 0 / 0 (n/a) |

### Asset type × tool, all languages

Each tool's rows sum the pilots it ran on.

| Asset type | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op |
|---|---|---:|---:|---:|---:|---:|---:|---|---|
| algorithm | crypto-finder | 0.995 | 0.999 | 0.997 | 0.992 | 1.000 | 0.996 | 954 / 5 / 1 | 767 / 954 (0.804) |
| algorithm | sonar-cryptography | 0.986 | 0.236 | 0.381 | 0.981 | 0.273 | 0.428 | 137 / 2 / 444 | 115 / 137 (0.839) |
| algorithm | cdxgen | 1.000 | 0.064 | 0.120 | 1.000 | 0.084 | 0.154 | 61 / 0 / 894 | 22 / 61 (0.361) |
| certificate | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 49 / 0 / 0 | 49 / 49 (1.000) |
| certificate | sonar-cryptography | n/a | 0.000 | 0.000 | n/a | 0.000 | 0.000 | 0 / 0 / 21 | 0 / 0 (n/a) |
| certificate | cdxgen | 0.000 | 0.000 | 0.000 | 0.000 | 0.000 | 0.000 | 0 / 3 / 49 | 0 / 0 (n/a) |
| protocol | crypto-finder | 0.905 | 1.000 | 0.950 | 0.905 | 1.000 | 0.950 | 38 / 4 / 0 | 38 / 38 (1.000) |
| protocol | sonar-cryptography | 1.000 | 0.050 | 0.095 | 1.000 | 0.050 | 0.095 | 1 / 0 / 19 | 1 / 1 (1.000) |
| protocol | cdxgen | 0.500 | 0.026 | 0.050 | 0.500 | 0.026 | 0.050 | 1 / 1 / 37 | 1 / 1 (1.000) |
| related-crypto-material | crypto-finder | 1.000 | 0.995 | 0.997 | 1.000 | 0.994 | 0.997 | 372 / 0 / 2 | 372 / 372 (1.000) |
| related-crypto-material | sonar-cryptography | 0.226 | 0.054 | 0.087 | 0.226 | 0.066 | 0.102 | 12 / 41 / 212 | 12 / 12 (1.000) |
| related-crypto-material | cdxgen | 0.407 | 0.134 | 0.201 | 0.376 | 0.139 | 0.203 | 50 / 73 / 324 | 37 / 50 (0.740) |

### crypto-finder, language × asset type

| Language | Asset type | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op |
|---|---|---:|---:|---:|---:|---:|---:|---|---|
| python | algorithm | 0.981 | 1.000 | 0.990 | 0.972 | 1.000 | 0.986 | 202 / 4 / 0 | 172 / 202 (0.851) |
| python | certificate | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 7 / 0 / 0 | 7 / 7 (1.000) |
| python | protocol | 0.800 | 1.000 | 0.889 | 0.800 | 1.000 | 0.889 | 8 / 2 / 0 | 8 / 8 (1.000) |
| python | related-crypto-material | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 78 / 0 / 0 | 78 / 78 (1.000) |
| javascript | algorithm | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 262 / 0 / 0 | 210 / 262 (0.802) |
| javascript | certificate | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 14 / 0 / 0 | 14 / 14 (1.000) |
| javascript | protocol | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 3 / 0 / 0 | 3 / 3 (1.000) |
| javascript | related-crypto-material | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 94 / 0 / 0 | 94 / 94 (1.000) |
| rust | algorithm | 0.991 | 0.991 | 0.991 | 0.988 | 1.000 | 0.994 | 111 / 1 / 1 | 83 / 111 (0.748) |
| rust | certificate | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 14 / 0 / 0 | 14 / 14 (1.000) |
| rust | protocol | 0.882 | 1.000 | 0.938 | 0.882 | 1.000 | 0.938 | 15 / 2 / 0 | 15 / 15 (1.000) |
| rust | related-crypto-material | 1.000 | 0.964 | 0.982 | 1.000 | 0.958 | 0.979 | 54 / 0 / 2 | 54 / 54 (1.000) |
| go | algorithm | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 133 / 0 / 0 | 108 / 133 (0.812) |
| go | certificate | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 4 / 0 / 0 | 4 / 4 (1.000) |
| go | protocol | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 10 / 0 / 0 | 10 / 10 (1.000) |
| go | related-crypto-material | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 55 / 0 / 0 | 55 / 55 (1.000) |
| java | algorithm | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 246 / 0 / 0 | 194 / 246 (0.789) |
| java | certificate | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 10 / 0 / 0 | 10 / 10 (1.000) |
| java | protocol | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 2 / 0 / 0 | 2 / 2 (1.000) |
| java | related-crypto-material | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 91 / 0 / 0 | 91 / 91 (1.000) |
