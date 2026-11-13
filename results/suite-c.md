# Suite C: real repositories, first score

**On the 142 sampled files of ten real repositories, crypto-finder v0.53.1 with ruleset v1.48.0 finds 108 of 214 assets with 8 false positives (precision 0.931, recall 0.505).** sonar-cryptography 1.7.0 finds 28 of the 125 assets in its three languages with none, and cdxgen 13.6.0 finds 18 of 214 with 33. This is the first score, taken at the versions frozen before any tool ran on suite C, and published as is ([METHODOLOGY.md](../METHODOLOGY.md#first-score-freeze)).

**Most of crypto-finder's misses are test files it did not scan.** Its recorded configuration ran with crypto-finder's default, which leaves test sources out (`--include-tests` adds them). Suite A has no test files, so the default never mattered there; the competitors were configured not to skip tests. In this run crypto-finder has no prediction in any test-path file of the sample: Go `_test.go` files, Python `test/`, `tests/` and `devscripts/` directories and JavaScript `__tests__` directories. Those files hold 90 of the 214 assets, and they are 90 of its 106 FN. A [later run with `--include-tests`](#later-run-crypto-finder-with---include-tests) scores 184 / 14 / 30 (precision 0.929, recall 0.860). It is shown separately and does not replace the first score.

| Tool | Languages run | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|
| crypto-finder v0.53.1, ruleset v1.48.0 | 5 | 108 / 8 / 106 | 0.931 / 0.505 | 0.924 / 0.485 |
| sonar-cryptography 1.7.0 | 3 (Go, Java, Python; no JavaScript or Rust support) | 28 / 0 / 97 | 1.000 / 0.224 | 1.000 / 0.228 |
| cdxgen 13.6.0 | 5 | 18 / 33 / 196 | 0.353 / 0.084 | 0.353 / 0.090 |

- **Scope:** `files`. Every asset in every sampled file is annotated, so precision and recall are both exact within the 142-file sample. Predictions in other files are ignored and only counted (crypto-finder 508, cdxgen 714, sonar-cryptography 158). Recall here is not relative to a pool.
- **Ground truth:** the merged `files` ground truth of 2026-11-13, 214 assets in 200 objects over 142 files (Go 61, Java 3, JavaScript 57, Python 61, Rust 32), annotated twice independently before any tool ran on suite C ([annotation log](../docs/annotation-log.md#suite-c-files-scope-2026-11-13)).
- **Matching:** the suite A rules, strict (`LINE_TOLERANCE = 0`). Suite C ground truth has no anchors.
- **Reachability** is not scored in suite C; it is phase 2 ([METHODOLOGY.md](../METHODOLOGY.md#suite-c)).
- **No targets** are set for suite C, and the suite A targets are not applied to it.
- **Run records:** [crypto-finder](crypto-finder/suite-c/README.md), [sonar-cryptography](sonar-cryptography/suite-c/README.md), [cdxgen](cdxgen/suite-c/README.md). Each holds the predictions, `score.md`, the commands and every deviation.

## What suite C measures

Whether the suite A scores hold on real code no tool was tuned on: ten open-source applications, two per language, chosen by a rule written before any candidate was retrieved, and fetched at a pinned commit ([METHODOLOGY.md](../METHODOLOGY.md#suite-c-selection)). The `files` scope draws 15 files per repository among those that import a crypto module, seeded by the pinned commit, and annotates every asset in them. A signal file with no crypto asset still counts: it tests precision.

| Entry | Repository | Sampled files | Assets |
|---|---|---|---|
| go-c-caveman | JuliusBrussee/caveman | 15 | 30 |
| go-c-ollama | ollama/ollama | 15 | 31 |
| java-c-jadx | skylot/jadx | 15 | 2 |
| java-c-mall | macrozheng/mall | 15 | 1 |
| javascript-c-firecrawl | firecrawl/firecrawl | 15 | 36 |
| javascript-c-freecodecamp | freeCodeCamp/freeCodeCamp | 15 | 21 |
| python-c-graphify | Graphify-Labs/graphify | 15 | 16 |
| python-c-yt-dlp | yt-dlp/yt-dlp | 15 | 45 |
| rust-c-cc-switch | farion1231/cc-switch | 15 | 23 |
| rust-c-clash-verge-rev | clash-verge-rev/clash-verge-rev | 7 | 9 |

## First score, operation level

Each cell is `TP / FP / FN (precision / recall)`. sonar-cryptography is `not run` where it has no language support.

### By language

| Language | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| go | 26 / 0 / 35 (1.000 / 0.426) | 27 / 0 / 34 (1.000 / 0.443) | 2 / 6 / 59 (0.250 / 0.033) |
| java | 2 / 0 / 1 (1.000 / 0.667) | 1 / 0 / 2 (1.000 / 0.333) | 0 / 0 / 3 (n/a / 0.000) |
| javascript | 31 / 2 / 26 (0.939 / 0.544) | not run | 8 / 0 / 49 (1.000 / 0.140) |
| python | 26 / 5 / 35 (0.839 / 0.426) | 0 / 0 / 61 (n/a / 0.000) | 0 / 0 / 61 (n/a / 0.000) |
| rust | 23 / 1 / 9 (0.958 / 0.719) | not run | 8 / 27 / 24 (0.229 / 0.250) |

On the three languages both run, crypto-finder scores 54 / 5 / 71 (0.915 / 0.432) and sonar-cryptography 28 / 0 / 97 (1.000 / 0.224). In Go, where most assets are in test files, sonar-cryptography finds one more than crypto-finder.

### By asset type

sonar-cryptography's rows cover its three languages only.

| Asset type | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| algorithm | 97 / 2 / 71 (0.980 / 0.577) | 28 / 0 / 66 (1.000 / 0.298) | 17 / 0 / 151 (1.000 / 0.101) |
| certificate | 2 / 1 / 9 (0.667 / 0.182) | 0 / 0 / 9 (n/a / 0.000) | 0 / 0 / 11 (n/a / 0.000) |
| protocol | 6 / 4 / 12 (0.600 / 0.333) | 0 / 0 / 12 (n/a / 0.000) | 0 / 0 / 18 (n/a / 0.000) |
| related-crypto-material | 3 / 1 / 14 (0.750 / 0.176) | 0 / 0 / 10 (n/a / 0.000) | 1 / 33 / 16 (0.029 / 0.059) |

### By entry

| Entry | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| go-c-caveman | 12 / 0 / 18 (1.000 / 0.400) | 15 / 0 / 15 (1.000 / 0.500) | 0 / 4 / 30 (0.000 / 0.000) |
| go-c-ollama | 14 / 0 / 17 (1.000 / 0.452) | 12 / 0 / 19 (1.000 / 0.387) | 2 / 2 / 29 (0.500 / 0.065) |
| java-c-jadx | 1 / 0 / 1 (1.000 / 0.500) | 1 / 0 / 1 (1.000 / 0.500) | 0 / 0 / 2 (n/a / 0.000) |
| java-c-mall | 1 / 0 / 0 (1.000 / 1.000) | 0 / 0 / 1 (n/a / 0.000) | 0 / 0 / 1 (n/a / 0.000) |
| javascript-c-firecrawl | 11 / 0 / 25 (1.000 / 0.306) | not run | 5 / 0 / 31 (1.000 / 0.139) |
| javascript-c-freecodecamp | 20 / 2 / 1 (0.909 / 0.952) | not run | 3 / 0 / 18 (1.000 / 0.143) |
| python-c-graphify | 12 / 0 / 4 (1.000 / 0.750) | 0 / 0 / 16 (n/a / 0.000) | 0 / 0 / 16 (n/a / 0.000) |
| python-c-yt-dlp | 14 / 5 / 31 (0.737 / 0.311) | 0 / 0 / 45 (n/a / 0.000) | 0 / 0 / 45 (n/a / 0.000) |
| rust-c-cc-switch | 15 / 1 / 8 (0.938 / 0.652) | not run | 8 / 23 / 15 (0.258 / 0.348) |
| rust-c-clash-verge-rev | 8 / 0 / 1 (1.000 / 0.889) | not run | 0 / 4 / 9 (0.000 / 0.000) |

**Attribution.** crypto-finder attributes 104 of its 108 matches (0.963), sonar-cryptography 28 of 28 and cdxgen 17 of 18. crypto-finder's four unattributed matches are freeCodeCamp's `jwt.verify(token, secret)` calls: it reports each as a signature with no algorithm, where the secret makes it HMAC (`jwt.sign` with the same secret is attributed).

## First score, object level

Each crypto object counts once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Suite C has 200 objects for its 214 assets, so the two views are close.

| Group | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| all | 97 / 8 / 103 (0.924 / 0.485) | 28 / 0 / 95 (1.000 / 0.228) | 18 / 33 / 182 (0.353 / 0.090) |
| go | 25 / 0 / 35 (1.000 / 0.417) | 27 / 0 / 33 (1.000 / 0.450) | 2 / 6 / 58 (0.250 / 0.033) |
| java | 2 / 0 / 1 (1.000 / 0.667) | 1 / 0 / 2 (1.000 / 0.333) | 0 / 0 / 3 (n/a / 0.000) |
| javascript | 22 / 2 / 25 (0.917 / 0.468) | not run | 8 / 0 / 39 (1.000 / 0.170) |
| python | 26 / 5 / 34 (0.839 / 0.433) | 0 / 0 / 60 (n/a / 0.000) | 0 / 0 / 60 (n/a / 0.000) |
| rust | 22 / 1 / 8 (0.957 / 0.733) | not run | 8 / 27 / 22 (0.229 / 0.267) |
| algorithm | 87 / 2 / 69 (0.978 / 0.558) | 28 / 0 / 65 (1.000 / 0.301) | 17 / 0 / 139 (1.000 / 0.109) |
| certificate | 2 / 1 / 8 (0.667 / 0.200) | 0 / 0 / 9 (n/a / 0.000) | 0 / 0 / 10 (n/a / 0.000) |
| protocol | 6 / 4 / 12 (0.600 / 0.333) | 0 / 0 / 12 (n/a / 0.000) | 0 / 0 / 18 (n/a / 0.000) |
| related-crypto-material | 2 / 1 / 14 (0.667 / 0.125) | 0 / 0 / 9 (n/a / 0.000) | 1 / 33 / 15 (0.029 / 0.062) |

By entry, the object level differs most in freeCodeCamp, where one JWT secret signs and verifies in several files (crypto-finder 11 / 2 / 1 at object level against 20 / 2 / 1, cdxgen 3 / 0 / 9 against 3 / 0 / 18); elsewhere it moves by one object at most. Every `score.md` has the full object-level tables.

## The test-file finding

| | Go | Python | JavaScript | Rust | Java |
|---|---|---|---|---|---|
| Test-path files in the sample | 16 `_test.go` | 5 (`test/`, `tests/`, `devscripts/`) | 6 in `__tests__`, 4 `*.test.ts` elsewhere | none (inline `#[cfg(test)]` modules only) | 6 under `src/test/` |
| Assets in them | 35 | 30 | 25 in `__tests__`, 9 elsewhere | 2 in inline test modules | 0 |
| crypto-finder predictions in them, first score | 0 | 0 | 0 in `__tests__`; 9 TP in `*.test.ts` elsewhere | scanned | 0 |

- crypto-finder's default leaves out `_test.go` files, `test/` and `tests/` directories, `devscripts/` and `__tests__`; it still scanned JavaScript `*.test.ts` files outside `__tests__` and Rust inline test modules. The 90 assets under the default exclusion are all FN in the first score.
- The [competitor configuration](../METHODOLOGY.md#competitor-configuration) sets sonar-cryptography's exclusion to match nothing, because "the benchmark scores every file", and suite C's [decisions](../METHODOLOGY.md#suite-c-decisions) keep test files in scope for the same reason. crypto-finder's recorded configuration did not set `--include-tests`. That was an omission in the run configuration, carried over from suite A where no test file exists; it is reported here, not corrected in the first score.
- The competitors scan test files but find little in them: sonar-cryptography 8 of the 35 Go test assets and none of the Python ones; cdxgen none.

## Later run: crypto-finder with `--include-tests`

**Dated 2026-11-13. Configuration changed after the first score was seen.** Same binary, ruleset, checkouts and commands, plus `--include-tests`; nothing else changed ([run record](crypto-finder/suite-c-include-tests/README.md)). Outside test paths every finding is identical to the first score. This run sits beside the first score and never replaces it.

| Language | First score TP / FP / FN | P / R | Later run TP / FP / FN | P / R |
|---|---|---|---|---|
| go | 26 / 0 / 35 | 1.000 / 0.426 | 60 / 0 / 1 | 1.000 / 0.984 |
| java | 2 / 0 / 1 | 1.000 / 0.667 | 2 / 0 / 1 | 1.000 / 0.667 |
| javascript | 31 / 2 / 26 | 0.939 / 0.544 | 55 / 5 / 2 | 0.917 / 0.965 |
| python | 26 / 5 / 35 | 0.839 / 0.426 | 44 / 8 / 17 | 0.846 / 0.721 |
| rust | 23 / 1 / 9 | 0.958 / 0.719 | 23 / 1 / 9 | 0.958 / 0.719 |
| **all, operation** | **108 / 8 / 106** | **0.931 / 0.505** | **184 / 14 / 30** | **0.929 / 0.860** |
| all, object | 97 / 8 / 103 | 0.924 / 0.485 | 172 / 14 / 28 | 0.925 / 0.860 |

| Asset type | Later run TP / FP / FN | P / R |
|---|---|---|
| algorithm | 154 / 5 / 14 | 0.969 / 0.917 |
| certificate | 8 / 1 / 3 | 0.889 / 0.727 |
| protocol | 14 / 7 / 4 | 0.667 / 0.778 |
| related-crypto-material | 8 / 1 / 9 | 0.889 / 0.471 |

76 FN become TP and 6 FP are added, all in test-path files. Attribution: 179 of 184 matches. The run record lists what the test files add and what is left in them.

## What crypto-finder misses

The first score's 106 FN and 8 FP, grouped by cause. Counts are operation level. "Left in the later run" is the same miss in the `--include-tests` run.

| Cause | FN | FP | Where | Left in the later run |
|---|---|---|---|---|
| Test sources left out by the default configuration | 90 | 0 | Go `_test.go` (35), Python `test/`, `tests/`, `devscripts/` (30), JavaScript `__tests__` (25) | 14 FN and 6 FP in those files, by the causes below and the run record |
| A hash call inside a Rust macro argument is not reported | 5 | 0 | cc-switch `Sha256::digest(...)` inside `format!(...)` (`pi_config/mod.rs:494`, `kimi_oauth/account.rs:137`, `session_usage_dsh.rs:547`) and inside `assert_eq!(...)` in inline test modules (`claude_account/layout.rs:167`, `dsh_account.rs:739`) | all 5 |
| TLS client built with rustls, tokio-rustls and hyper-rustls | 3 | 1 | cc-switch `hyper_client.rs`: the two `TlsConnector::connect(...)` handshakes (431, 537) and the certificates parsed by `RootCertStore::add_parsable_certificates` (623) are not reported; the `HttpsConnectorBuilder` chain (62) is reported as a certificate | all 4 |
| Python `ssl` trust stores and client keys | 3 | 0 | yt-dlp `networking/_helper.py`: `load_verify_locations(cafile=...)` (25) and `load_verify_locations(cadata=...)` (50) are not reported as certificates; `load_cert_chain(...)` (146) is reported as a certificate but its private key is not reported as key material | all 3, and 5 more `load_cert_chain` keys in `test/` |
| TLS configuration setters reported as protocols | 0 | 4 | yt-dlp `networking/_helper.py`: `set_alpn_protocols` (119), `set_ciphers` (125, 140), `minimum_version` (142). The ground truth counts them as configuration of the context, which is already an asset; suite A has the same gap | all 4, and 3 more setters in `test/` |
| An HMAC created with only a key and fed through `update()` | 1 | 1 | yt-dlp `extractor/abematv.py:105`: `hmac.new(key, digestmod=hashlib.sha256)` is not reported, and the `h.update(nonce)` on the next line is reported as key material. The `hmac.new(key, msg, digestmod=...)` in the same file is found | both |
| A project's own cipher implementation | 1 | 0 | yt-dlp `abematv.py:81` `aes_ecb_decrypt(...)`, from yt-dlp's own `aes` module | it, and 6 assets in `devscripts/generate_aes_testdata.py` (yt-dlp's own AES functions, an IV and an `openssl` command line) |
| JWT decode without verification reported as an algorithm | 0 | 2 | freeCodeCamp `api/src/plugins/auth.ts:98,164`, `jwt.decode(...)` | both |
| Nonce drawn into a buffer | 1 | 0 | clash-verge-rev `config/encrypt.rs:23`, `getrandom::fill(&mut nonce)`: the draw is reported as a CSPRNG use, the nonce it fills is not reported as material | it |
| Browser crypto reached through `window` | 1 | 0 | freeCodeCamp `client/src/utils/growthbook-cookie.ts:7`, `window.crypto.randomUUID()` behind an optional-chaining guard | it |
| A signature library with no coverage | 1 | 0 | jadx `ApkSignatureNode.java:176`, apksig `ApkVerifier.verify()`; sonar-cryptography and cdxgen miss it too | it |
| **Total** | **106** | **8** | | **30 FN, 14 FP** |

Two of these are detectable by other tools on the same lines: cdxgen reports three of the five `Sha256::digest` calls inside `format!`.

## What the other tools miss

- **sonar-cryptography.** Python: no component in either repository (61 FN); it scanned 687 and 1,124 `.py` files. Go: 34 FN, of which 25 are SHA-256 hashes in `_test.go` files; outside tests it misses `rand.Read` draws, caveman's `tls.NewListener` session and ollama's Ed25519 key generation and SSH key serialization. Java: apksig's `ApkVerifier.verify()` and Spring Security's `BCryptPasswordEncoder`. No false positive. JavaScript and Rust are not supported, so 89 assets are out of its reach.
- **cdxgen.** Python and Java: no algorithm analysis in source in 13.6.0, so all 64 assets are FN. Go: golem reports 2 of 61, none in `_test.go` files. JavaScript: `createHash`, `createHmac` and `crypto.subtle.digest` only; no `jsonwebtoken`, no random draws, nothing in `__tests__`. Rust: `Sha256::digest` calls only. All 33 of its FP are `related-crypto-material` named after variables (`input_tokens`, `refresh_token`, `password`, `key`) on lines that compute no crypto.

Details and deviations are on each tool's [run record](#suite-c-real-repositories-first-score).

## Limitations

- The sample is 142 files. One repository's assets can dominate a row: yt-dlp holds 45 of the 61 Python assets, firecrawl 36 of the 57 JavaScript ones, and the two Java repositories only 3 assets between them.
- Most signal files import crypto for a hash or a random draw: 168 of the 214 assets are algorithms, and SHA-256 alone is most of Go. Suite C measures what real applications do with crypto, not library breadth; suite A does that.
- The `pool` scope, which covers crypto outside the sampled files, is not scored yet.
