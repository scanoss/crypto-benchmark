# Suite C: real repositories

**Both scopes are scored. On the 142 sampled files (`files`, the headline), crypto-finder v0.53.1 with ruleset v1.48.0 finds 108 of 214 assets with 8 false positives (precision 0.931, recall 0.505).** sonar-cryptography 1.7.0 finds 28 of the 125 assets in its three languages with none, and cdxgen 13.6.0 finds 18 of 214 with 33.

**Over whole repositories (`pool`), crypto-finder matches 444 of the 993 adjudicated pool assets with 24 false positives: precision 0.949, recall 0.447 relative to the pool.** sonar-cryptography matches 169 of the 590 pool assets in its languages with 3 (0.983 / 0.286 relative to the pool), and cdxgen 76 of 993 with 689 (0.099 / 0.077 relative to the pool). Pool recall counts only the assets some tool or the grep found, so it overstates true recall and is never comparable to `files` or suite A recall ([`pool` scope](#pool-scope-first-score)). 284 pool assets were found by crypto-finder's first score and by neither competitor, 71 only by sonar-cryptography and 9 only by cdxgen.

These are the first scores, taken at the versions frozen before any tool ran on suite C, and published as is ([METHODOLOGY.md](../METHODOLOGY.md#first-score-freeze)).

**Most of crypto-finder's misses, in both scopes, are test files it did not scan.** Its recorded configuration ran with crypto-finder's default, which leaves test sources out (`--include-tests` adds them). Suite A has no test files, so the default never mattered there; the competitors were configured not to skip tests. In the `files` scope those files hold 90 of its 106 FN; in the pool, 436 of its 549. A [later run with `--include-tests`](#later-run-crypto-finder-with---include-tests) scores 184 / 14 / 30 on `files` (precision 0.929, recall 0.860) and 880 / 45 / 113 on the pool (0.951 / 0.886 relative to the pool). It is shown separately and does not replace the first score.

## `files` scope, first score

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

## `pool` scope, first score

Recall in this table is relative to the pool: the denominator is the 993 assets found by some tool or the grep, not every asset in the repositories.

| Tool | Languages run | TP / FP / FN op | P / R op, R relative to pool | P / R obj, R relative to pool | Found only by this tool | Predictions outside the pool frame |
|---|---|---|---|---|---|---|
| crypto-finder v0.53.1, ruleset v1.48.0 | 5 | 444 / 24 / 549 | 0.949 / 0.447 | 0.944 / 0.435 | 284 | 143 dropped |
| sonar-cryptography 1.7.0 | 3 (Go, Java, Python) | 169 / 3 / 421 | 0.983 / 0.286 | 0.983 / 0.300 | 71 | 0 |
| cdxgen 13.6.0 | 5 | 76 / 689 / 917 | 0.099 / 0.077 | 0.098 / 0.081 | 9 | 0 |

- **Scope:** `pool`. Every finding of the three tools, of the later crypto-finder run and of a grep for crypto terms, on the whole repository, was adjudicated by two annotators: 10,655 items, 993 assets ([annotation log](../docs/annotation-log.md#suite-c-pool-scope-2026-11-13)). Precision is exact.
- **Predictions:** the same frozen prediction files as the `files` scope, re-keyed to the `<id>-pool` entries; no tool was rerun. Predictions in files of another language than the entry's are dropped before matching and counted.
- **Run records:** [crypto-finder](crypto-finder/suite-c-pool/README.md), [sonar-cryptography](sonar-cryptography/suite-c-pool/README.md), [cdxgen](cdxgen/suite-c-pool/README.md), and the later run's [pool score](crypto-finder/suite-c-pool-include-tests/README.md).

## What suite C measures

Whether the suite A scores hold on real code no tool was tuned on: ten open-source applications, two per language, chosen by a rule written before any candidate was retrieved, and fetched at a pinned commit ([METHODOLOGY.md](../METHODOLOGY.md#suite-c-selection)). The `files` scope draws 15 files per repository among those that import a crypto module, seeded by the pinned commit, and annotates every asset in them. A signal file with no crypto asset still counts: it tests precision. The `pool` scope covers the whole repository instead, but only the assets some tool or the grep found.

| Entry | Repository | Sampled files | `files` assets | `pool` assets |
|---|---|---|---|---|
| go-c-caveman | JuliusBrussee/caveman | 15 | 30 | 258 |
| go-c-ollama | ollama/ollama | 15 | 31 | 109 |
| java-c-jadx | skylot/jadx | 15 | 2 | 7 |
| java-c-mall | macrozheng/mall | 15 | 1 | 4 |
| javascript-c-firecrawl | firecrawl/firecrawl | 15 | 36 | 297 |
| javascript-c-freecodecamp | freeCodeCamp/freeCodeCamp | 15 | 21 | 50 |
| python-c-graphify | Graphify-Labs/graphify | 15 | 16 | 58 |
| python-c-yt-dlp | yt-dlp/yt-dlp | 15 | 45 | 154 |
| rust-c-cc-switch | farion1231/cc-switch | 15 | 23 | 43 |
| rust-c-clash-verge-rev | clash-verge-rev/clash-verge-rev | 7 | 9 | 13 |

## `files` scope: first score, operation level

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

## `files` scope: first score, object level

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

## `files` scope: the test-file finding

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

In the `files` scope: the first score's 106 FN and 8 FP, grouped by cause. The pool's are under [What crypto-finder misses in the pool](#what-crypto-finder-misses-in-the-pool). Counts are operation level. "Left in the later run" is the same miss in the `--include-tests` run.

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

## `files` scope: what the other tools miss

- **sonar-cryptography.** Python: no component in either repository (61 FN); it scanned 687 and 1,124 `.py` files. Go: 34 FN, of which 25 are SHA-256 hashes in `_test.go` files; outside tests it misses `rand.Read` draws, caveman's `tls.NewListener` session and ollama's Ed25519 key generation and SSH key serialization. Java: apksig's `ApkVerifier.verify()` and Spring Security's `BCryptPasswordEncoder`. No false positive. JavaScript and Rust are not supported, so 89 assets are out of its reach.
- **cdxgen.** Python and Java: no algorithm analysis in source in 13.6.0, so all 64 assets are FN. Go: golem reports 2 of 61, none in `_test.go` files. JavaScript: `createHash`, `createHmac` and `crypto.subtle.digest` only; no `jsonwebtoken`, no random draws, nothing in `__tests__`. Rust: `Sha256::digest` calls only. All 33 of its FP are `related-crypto-material` named after variables (`input_tokens`, `refresh_token`, `password`, `key`) on lines that compute no crypto.

Details and deviations are on each tool's run record (linked [above](#files-scope-first-score)).

## `pool` scope: first score in detail

**Every recall in this section is relative to the pool**, 993 assets that some tool or the grep found, and is not comparable to `files` or suite A recall. Each cell is `TP / FP / FN (precision / recall relative to pool)`. sonar-cryptography is `not run` where it has no language support.

### What the pool is, and what it cannot see

- **Sources.** Every finding of crypto-finder's first score, of the later `--include-tests` run, of sonar-cryptography and of cdxgen, plus every line a grep for crypto terms printed (`cbench frame --grep`), in source files of the entry's language outside excluded paths ([METHODOLOGY.md](../METHODOLOGY.md#pool-scope)). Every item was adjudicated with no tool attribution, so precision is exact for every tool.
- **It favors the tools that built it.** An asset no source found is not in the denominator. 104 of the 993 assets were found by the grep alone, by no tool run.
- **Known blind spots.** The annotators listed 25 assets that are in the repositories but no pool item covers, so no tool is scored on them ([annotation log](../docs/annotation-log.md#known-misses-of-the-pool)): 12 `httptest.NewTLSServer` test servers and one `StartTLS` in the two Go repositories, certificate parses (`CertPool.AppendCertsFromPEM` ×2, `load_verify_locations(cafile=...)`, rustls `add_parsable_certificates`), three IVs created from literals in yt-dlp, a `wrap_socket` TLS session in graphify, signature verifications through `AlipaySignature.rsaCheckV1` (mall) and `windows.WinVerifyTrustEx` (ollama), a crypto/rand draw through `io.ReadFull` and an ECDSA public-key serialization. TLS in test code and key material are what every source missed most, so pool recall is most generous on those types.
- **Dropped outside the pool frame.** crypto-finder's first-score files hold 143 predictions in files of another language than the entry's (TypeScript, JavaScript and Python files in the Go and Rust repositories; Rust and Go files in firecrawl), the later run's 208. They are dropped before matching and counted, never scored. sonar-cryptography and cdxgen have none.

### By language

| Language | Pool assets | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|---|
| go | 367 | 140 / 3 / 227 (0.979 / 0.381) | 168 / 0 / 199 (1.000 / 0.458) | 7 / 245 / 360 (0.028 / 0.019) |
| java | 11 | 8 / 0 / 3 (1.000 / 0.727) | 1 / 3 / 10 (0.250 / 0.091) | 0 / 0 / 11 (n/a / 0.000) |
| javascript | 347 | 154 / 11 / 193 (0.933 / 0.444) | not run | 50 / 0 / 297 (1.000 / 0.144) |
| python | 212 | 101 / 5 / 111 (0.953 / 0.476) | 0 / 0 / 212 (n/a / 0.000) | 0 / 0 / 212 (n/a / 0.000) |
| rust | 56 | 41 / 5 / 15 (0.891 / 0.732) | not run | 19 / 444 / 37 (0.041 / 0.339) |
| **all** | **993** | **444 / 24 / 549 (0.949 / 0.447)** | **169 / 3 / 421 (0.983 / 0.286)** | **76 / 689 / 917 (0.099 / 0.077)** |

On the three languages both run, crypto-finder scores 249 / 8 / 341 (0.969 / 0.422 relative to pool) and sonar-cryptography 169 / 3 / 421 (0.983 / 0.286). In Go, sonar-cryptography finds more: 220 of crypto-finder's 227 Go FN are in `_test.go` files its first-score configuration did not scan.

### By asset type

sonar-cryptography's rows cover its three languages only.

| Asset type | Pool assets | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|---|
| algorithm | 827 | 413 / 10 / 414 (0.976 / 0.499) | 169 / 3 / 286 (0.983 / 0.371) | 74 / 8 / 753 (0.902 / 0.089) |
| certificate | 41 | 8 / 3 / 33 (0.727 / 0.195) | 0 / 0 / 39 (n/a / 0.000) | 0 / 1 / 41 (0.000 / 0.000) |
| protocol | 46 | 9 / 8 / 37 (0.529 / 0.196) | 0 / 0 / 39 (n/a / 0.000) | 0 / 1 / 46 (0.000 / 0.000) |
| related-crypto-material | 79 | 14 / 3 / 65 (0.824 / 0.177) | 0 / 0 / 57 (n/a / 0.000) | 2 / 679 / 77 (0.003 / 0.025) |

### By entry

| Entry | Pool assets | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|---|
| go-c-caveman-pool | 258 | 100 / 3 / 158 (0.971 / 0.388) | 110 / 0 / 148 (1.000 / 0.426) | 4 / 139 / 254 (0.028 / 0.016) |
| go-c-ollama-pool | 109 | 40 / 0 / 69 (1.000 / 0.367) | 58 / 0 / 51 (1.000 / 0.532) | 3 / 106 / 106 (0.028 / 0.028) |
| java-c-jadx-pool | 7 | 6 / 0 / 1 (1.000 / 0.857) | 1 / 3 / 6 (0.250 / 0.143) | 0 / 0 / 7 (n/a / 0.000) |
| java-c-mall-pool | 4 | 2 / 0 / 2 (1.000 / 0.500) | 0 / 0 / 4 (n/a / 0.000) | 0 / 0 / 4 (n/a / 0.000) |
| javascript-c-firecrawl-pool | 297 | 109 / 2 / 188 (0.982 / 0.367) | not run | 47 / 0 / 250 (1.000 / 0.158) |
| javascript-c-freecodecamp-pool | 50 | 45 / 9 / 5 (0.833 / 0.900) | not run | 3 / 0 / 47 (1.000 / 0.060) |
| python-c-graphify-pool | 58 | 45 / 0 / 13 (1.000 / 0.776) | 0 / 0 / 58 (n/a / 0.000) | 0 / 0 / 58 (n/a / 0.000) |
| python-c-yt-dlp-pool | 154 | 56 / 5 / 98 (0.918 / 0.364) | 0 / 0 / 154 (n/a / 0.000) | 0 / 0 / 154 (n/a / 0.000) |
| rust-c-cc-switch-pool | 43 | 30 / 5 / 13 (0.857 / 0.698) | not run | 19 / 399 / 24 (0.045 / 0.442) |
| rust-c-clash-verge-rev-pool | 13 | 11 / 0 / 2 (1.000 / 0.846) | not run | 0 / 45 / 13 (0.000 / 0.000) |

**Attribution.** crypto-finder attributes 429 of its 444 matches (0.966), sonar-cryptography 166 of 169 and cdxgen 63 of 76. **Object level** (recall relative to pool): crypto-finder 403 / 24 / 523 (0.944 / 0.435), sonar-cryptography 169 / 3 / 394 (0.983 / 0.300), cdxgen 75 / 689 / 851 (0.098 / 0.081). Every `score.md` has the full tables.

### Assets only one tool found

An asset counts for a tool when its prediction matches it by the score's rules, and as found only by that tool when no other tool in the comparison matches it (`cbench unique`; reports in [suite-c-pool-unique](suite-c-pool-unique/)). The first comparison is the three first scores; the second puts the later crypto-finder run in place of the first.

| Tool | Pool assets matched | Only this tool, first scores | Only this tool, with the later crypto-finder run |
|---|---|---|---|
| crypto-finder | 444 first score; 880 later run | 284 | 649 |
| sonar-cryptography 1.7.0 | 169 | 71 | 0 |
| cdxgen 13.6.0 | 76 | 9 | 9 |
| none of the three | | 469 | 104 |

| Language | Only crypto-finder, first score | Only sonar-cryptography | Only cdxgen | None of the three first scores |
|---|---|---|---|---|
| go | 42 | 71 | 1 | 155 |
| java | 7 | 0 | 0 | 3 |
| javascript | 104 | 0 | 0 | 193 |
| python | 101 | 0 | 0 | 111 |
| rust | 30 | 0 | 8 | 7 |

- **sonar-cryptography's 71** are all in Go `_test.go` files, which crypto-finder's first configuration left out. Against the later run it has none of its own.
- **cdxgen's 9** are eight cc-switch `Sha256::digest(...)` calls inside `format!` macro arguments, which crypto-finder does not report, and the key half of caveman's `tls.LoadX509KeyPair` (`proxy/internal/identity/tls.go:63`).
- **The 469 found by none of the three first scores** split into 365 found only by the later crypto-finder run, all in test-path files, and 104 found only by the grep.

## Later run in the pool: crypto-finder with `--include-tests`

**Dated 2026-11-13, configuration changed after the first score was seen; beside the first pool score, never in its place.** Same prediction files as the later `files` run, re-keyed ([run record](crypto-finder/suite-c-pool-include-tests/README.md)). This run is one of the pool's sources, so its pool recall is the least independent of the four.

| Language | First score TP / FP / FN | P / R, R relative to pool | Later run TP / FP / FN | P / R, R relative to pool |
|---|---|---|---|---|
| go | 140 / 3 / 227 | 0.979 / 0.381 | 347 / 8 / 20 | 0.977 / 0.946 |
| java | 8 / 0 / 3 | 1.000 / 0.727 | 8 / 0 / 3 | 1.000 / 0.727 |
| javascript | 154 / 11 / 193 | 0.933 / 0.444 | 334 / 18 / 13 | 0.949 / 0.963 |
| python | 101 / 5 / 111 | 0.953 / 0.476 | 150 / 14 / 62 | 0.915 / 0.708 |
| rust | 41 / 5 / 15 | 0.891 / 0.732 | 41 / 5 / 15 | 0.891 / 0.732 |
| **all, operation** | **444 / 24 / 549** | **0.949 / 0.447** | **880 / 45 / 113** | **0.951 / 0.886** |
| all, object | 403 / 24 / 523 | 0.944 / 0.435 | 830 / 45 / 96 | 0.949 / 0.896 |

| Asset type | Later run TP / FP / FN | P / R, R relative to pool |
|---|---|---|
| algorithm | 761 / 19 / 66 | 0.976 / 0.920 |
| certificate | 37 / 6 / 4 | 0.860 / 0.902 |
| protocol | 35 / 17 / 11 | 0.673 / 0.761 |
| related-crypto-material | 47 / 3 / 32 | 0.940 / 0.595 |

436 FN become TP, all in test-path files (Go `_test.go` 207, JavaScript `__tests__` 180, Python `test/` 37 and `tests/` 12), and 21 FP are added, all in test-path files. Every other miss of the first pool score is still there. Predictions dropped outside the pool frame: 208.

### What crypto-finder misses in the pool

The first pool score's 549 FN and 24 FP, grouped by cause, with the later run's FP beside. The 113 FN below the first row are missed by both runs; 53 of them are in test-path files. Counts are operation level; where a cause also appears in the `files` table [above](#what-crypto-finder-misses), the pool counts include those lines.

| Cause | FN | FP, first score | FP, later run | Where |
|---|---|---|---|---|
| Test sources left out by the default configuration | 436 | 0 | | Go `_test.go` (207), JavaScript `__tests__` (180), Python `test/` (37) and `tests/` (12). The later run finds all 436 |
| Calls into a project's own cipher implementation | 25 | 0 | 0 | yt-dlp's pure-Python AES (`yt_dlp.aes`) called from 5 extractors, `test/test_aes.py` (12) and `devscripts/generate_aes_testdata.py` (5, with an `openssl enc` command line); yt-dlp's own textbook RSA (`adn.py`, `ohdave_rsa_encrypt`) and RC4 (`VideaIE.rc4`) |
| Certificates and private keys loaded into a TLS context | 17 | 0 | 0 | Python `load_cert_chain` key half (10, 9 of them in `test/`) and `load_verify_locations(cadata=...)`; Go `tls.LoadX509KeyPair` and `tls.X509KeyPair` key half (3) and `CertPool.AppendCertsFromPEM`; Node `https.createServer({ key, cert })` key and certificate (firecrawl test) |
| Library or OS calls whose algorithm the pinned library fixes but the call does not name | 13 | 0 | 0 | Hutool `JWTUtil.createToken` and `verify` (mall, HS256), `@fastify/cookie` `sign` (freeCodeCamp, HMAC-SHA-256, 3), Gatsby `createContentDigest` (MD5), tauri-plugin-updater `Update::verify` (Ed25519), apksig `ApkVerifier.verify`, `x/crypto/ssh` `Signer.Sign` and `PublicKey.Verify` on an Ed25519 key (ollama), pypdf `PdfWriter.encrypt(algorithm="RC4-128")`, Windows `CryptQueryObject` (certificate) and `CryptUnprotectData` (DPAPI) |
| pycryptodome reached through yt-dlp's own `dependencies.Cryptodome` module | 11 | 0 | 0 | `Cryptodome.AES.new`, `RSA.importKey`, `RSA.generate`, `PKCS1_v1_5.new`, `PKCS1_OAEP.new`, `CMAC.new` and a key export, called as attributes of the module yt-dlp imports instead of `Crypto`/`Cryptodome` (`aes.py`, `bilibili.py`, `ivi.py`, `tarangplus.py`, `wrestleuniverse.py`), and the `AES.new` probe in that module |
| A hash call inside a Rust macro argument | 11 | 0 | 0 | cc-switch `Sha256::digest(...)` inside `format!(...)` (9) and `assert_eq!(...)` (2) |
| TLS contexts and sessions | 11 | 0 | 0 | Go `&tls.Config{...}` literals passed or assigned in test code (5) and an `httptest.NewTLSServer`; tokio-rustls `TlsConnector::connect` (2, cc-switch); Python `wrap_bio` (test) and `websockets.sync.client.connect(ssl=...)`; Node `https.createServer` |
| Salt, IV and nonce material | 8 | 0 | 0 | literal salts passed to yt-dlp's `pbkdf2_sha1` helper (3), two IVs (a default argument and a `struct.pack`), firecrawl's zero IV `Buffer.alloc(16)`, a salt filled by `rand.Read` (caveman), a nonce filled by `getrandom::fill` (clash-verge-rev) |
| Key parse and serialization | 7 | 1 | 8 | Go `x509.MarshalECPrivateKey` (2) and `ecdsa.ParseUncompressedPublicKey` are reported as ECDSA signatures, not as key material; ECDSA `PublicKey.Bytes()` is not reported; jose `exportJWK` (3) is reported as an RSA signature, and so are the JWK object literals around it |
| Python standard-library call forms | 6 | 1 | 1 | `hmac.HMAC(...)` (2), `hmac.new(key, digestmod=...)` fed through `update()`, whose `update()` is reported as key material (FP), `hashlib.file_digest` (2), `hashlib.sha256` passed by reference to a project function |
| Certificate chain verification | 2 | 2 | 4 | caveman `x509.Certificate.Verify(opts)`: reported as a certificate on the call and its `VerifyOptions` literal, not as a signature verification |
| Crypto in a guarded or mocked expression | 2 | 0 | 0 | `window?.crypto?.randomUUID` (freeCodeCamp), `createHash` inside a `vi.mock` factory (firecrawl test) |
| TLS configuration reported as an asset of its own | 0 | 9 | 19 | Python `ssl` setters (`set_ciphers`, `set_alpn_protocols`, `minimum_version`, `maximum_version`, `verify_mode`: 4 first, 12 later); `rustls::crypto::ring::default_provider().install_default()` as a protocol (4, cc-switch); hyper-rustls `HttpsConnectorBuilder` as a certificate; later run only, a Go `tls.VersionTLS11` constant and Python `load_default_certs()` |
| JWT decode without verification reported as an algorithm | 0 | 9 | 9 | freeCodeCamp `jwt.decode(...)`: `auth.ts` (2) and `user.test.ts` (7), a test file the default configuration scans |
| JWKS set construction reported as public-key material | 0 | 2 | 2 | firecrawl jose `createLocalJWKSet`, `createRemoteJWKSet` |
| A test's own `sign()` helper reported as an RSA signature | 0 | 0 | 2 | firecrawl `__tests__` |
| **Total** | **549** | **24** | **45** | |

### What the other tools miss in the pool

- **sonar-cryptography.** No certificate, protocol or key-material asset in any language (135 FN of those types). Python: nothing in either repository (212 FN). Java: 1 of 11; its 3 FP are jadx's `getThumbPrint(x509cert, "MD5" | "SHA-1" | "SHA-256")` calls, where the algorithm name is passed to a project helper; the helper's `MessageDigest.getInstance` (line 158) is the asset, and it is not reported. Go: 168 of 367, precision 1.000.
- **cdxgen.** 679 of its 689 FP are `related-crypto-material`, 570 of them `key` and 50 `token`, all in Go (245) and Rust (444): lines where an identifier is named like key material (`input_tokens`, `output_key`, `secret := ""`, `SchemaStructuralKeys`) and no crypto runs. Every one was adjudicated as a pool item; the reported lines are correct, they hold no asset. Its 76 TP are JavaScript `createHash`, `createHmac` and `crypto.subtle.digest` (50), Rust `Sha256::digest` (19) and Go (7). Python and Java: nothing, as in the `files` scope.

## Limitations

- The sample is 142 files. One repository's assets can dominate a row: yt-dlp holds 45 of the 61 Python assets, firecrawl 36 of the 57 JavaScript ones, and the two Java repositories only 3 assets between them.
- Most signal files import crypto for a hash or a random draw: 168 of the 214 assets are algorithms, and SHA-256 alone is most of Go. Suite C measures what real applications do with crypto, not library breadth; suite A does that.
- Pool recall is relative to the pool and overstates true recall: it cannot see the assets no source found, and it favors the tools whose findings built the pool. Pool precision is exact.
- Reachability is not scored in suite C. It is phase 2 ([METHODOLOGY.md](../METHODOLOGY.md#suite-c-decisions)).
