# Rust pilot: tools compared

crypto-finder finds 194 of the 197 Rust assets with 3 false positives. Rust protocol is its one target cell that fails, on two TLS configuration setters.

Suite A, 15 Rust entries (`fixtures/rust/*`, one per crate), 197 assets in 151 objects, scored on 2026-11-12 against the ground truth of 2026-11-08. crypto-finder is v0.53.1 with remote ruleset v1.48.0; sonar-cryptography and cdxgen are scored from their stored predictions.

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| rust | crypto-finder | 0.985 | 0.985 | 0.985 | 0.980 | 0.987 | 0.983 | 194 / 3 / 3 | 166 / 194 (0.856) | 134 / 149 (0.899) |
| rust | cdxgen | 0.477 | 0.157 | 0.237 | 0.460 | 0.192 | 0.271 | 31 / 34 / 166 | 18 / 31 (0.581) | 16 / 29 (0.552) |

`op` counts every annotated operation; `obj` counts each crypto object once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed means the match also names the right algorithm and parameters.

## crypto-finder against the targets

Precision >= 0.90 and recall >= 0.85 per cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). 4 of 5 Rust cells pass at object level.

| View | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| object (headline) | 0.988 / 1.000 pass | 1.000 / 1.000 pass | 0.882 / 1.000 fail | 1.000 / 0.958 pass | 0.980 / 0.987 pass |
| operation | 0.991 / 0.991 pass | 1.000 / 1.000 pass | 0.882 / 1.000 fail | 1.000 / 0.964 pass | 0.985 / 0.985 pass |

## What crypto-finder still gets wrong

| Gap | Count | Where |
|---|---|---|
| openssl TLS configuration setters `set_min_proto_version` and `set_cipher_list` reported as protocols of their own | 2 FP | rust-openssl `tls_client.rs:5,6` |
| Creating the random number generator reported as a CSPRNG use; the ground truth annotates the key generations that draw from it | 1 FP | rust-curve25519-dalek `ristretto_dh.rs:7` |
| ECDSA verify inside `assert!(...)`, a call inside a Rust macro, not reported | 1 FN | rust-openssl `ecdsa_p256.rs:12` |
| ring `generate_pkcs8` reported once, as the PKCS#8 encoding it returns, not also as the generated private key | 2 FN | rust-ring `ecdsa_p256.rs:6`, `ed25519.rs:6` |
| ristretto255 key generations and key agreements match but are not attributed, because the benchmark's curve table has no ristretto255 | 4 unattributed | rust-curve25519-dalek `ristretto_dh.rs:8,9,14,15` |

## Other tools

- **cdxgen** (object level 0.460 / 0.192) runs the rusi analyzer, which reports variables by name. That finds part of the related crypto material (object level 0.424 / 0.521) at the cost of false positives on variables that hold keys or buffers. It reports nothing in 2 of the 15 entries.
- **sonar-cryptography** has no Rust support and was not run.

## History of pinned crypto-finder runs

Rust only, every run against the ground truth of 2026-11-08 and the same matching rules. All languages together: [README.md](README.md#history-of-pinned-runs).

| crypto-finder | Ruleset | Scored | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-08 | 129 / 14 / 68 | 0.902 / 0.655 | 0.873 / 0.636 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 172 / 14 / 25 | 0.925 / 0.873 | 0.906 / 0.894 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 172 / 4 / 25 | 0.977 / 0.873 | 0.971 / 0.894 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 172 / 4 / 25 | 0.977 / 0.873 | 0.971 / 0.894 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 189 / 3 / 8 | 0.984 / 0.959 | 0.980 / 0.967 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 193 / 3 / 4 | 0.985 / 0.980 | 0.980 / 0.980 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 194 / 3 / 3 | 0.985 / 0.985 | 0.980 / 0.987 |

## Run records

Exact versions, commands and the full score report of each tool: [crypto-finder](crypto-finder/rust-pilot/README.md), [sonar-cryptography](sonar-cryptography/rust-pilot/README.md) (not run), [cdxgen](cdxgen/rust-pilot/README.md).
