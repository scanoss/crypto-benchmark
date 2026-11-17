# Python pilot: tools compared

crypto-finder finds all 295 Python assets with 6 false positives. Python protocol is its one target cell that fails, on two TLS configuration setters.


## Summary

| | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| Python crypto found | **100%** (295 of 295) | 12% (35 of 295) | 0% (0 of 295) |
| Findings that are right | **98%** | 92% | n/a, it reports nothing |

**Found**: of all the crypto in the code, the share the tool reported (recall). **Right**: of what the tool reported, the share that is real crypto (precision).

## Scores in detail

Suite A, 16 Python entries (`fixtures/python/*`, one per library, the standard library included), 295 assets in 212 objects, scored on 2026-11-17 against the ground truth of 2026-11-08. crypto-finder is v0.54.0 with remote ruleset v1.48.0, run on 2026-11-17 with the same predictions as v0.53.1 on 2026-11-12; sonar-cryptography and cdxgen are scored from their stored predictions.

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| python | crypto-finder | 0.980 | 1.000 | 0.990 | 0.972 | 1.000 | 0.986 | 295 / 6 / 0 | 265 / 295 (0.898) | 199 / 212 (0.939) |
| python | sonar-cryptography | 0.921 | 0.119 | 0.210 | 0.917 | 0.156 | 0.266 | 35 / 3 / 260 | 31 / 35 (0.886) | 30 / 33 (0.909) |
| python | cdxgen | n/a | 0.000 | 0.000 | n/a | 0.000 | 0.000 | 0 / 0 / 295 | 0 / 0 (n/a) | 0 / 0 (n/a) |

`op` counts every annotated operation; `obj` counts each crypto object once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed means the match also names the right algorithm and parameters.

## crypto-finder against the targets

Precision >= 0.90 and recall >= 0.85 per cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). 4 of 5 Python cells pass at object level.

| View | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| object (headline) | 0.972 / 1.000 pass | 1.000 / 1.000 pass | 0.800 / 1.000 fail | 1.000 / 1.000 pass | 0.972 / 1.000 pass |
| operation | 0.981 / 1.000 pass | 1.000 / 1.000 pass | 0.800 / 1.000 fail | 1.000 / 1.000 pass | 0.980 / 1.000 pass |

## What crypto-finder still gets wrong

| Gap | Count | Where |
|---|---|---|
| TLS configuration setter `ctx.minimum_version` reported as a protocol of its own | 2 FP | python-stdlib `tls-context/main.py:6`, python-truststore `inject/main.py:10` |
| Hash class in `.sign(key, hashes.SHA256())` reported as a SHA-256 hash operation; the ground truth counts it as the signature's digest parameter | 3 FP | python-cryptography `x509-cert/main.py:22`, python-pyopenssl `cert-load/main.py:22` and `x509-cert/main.py:21` |
| Fernet `f.encrypt(...)` reported as an operation; the ground truth counts it as a follow-up on the `Fernet(key)` object | 1 FP | python-cryptography `fernet/main.py:8` |

No false negatives.

## Other tools

- **sonar-cryptography** (object level 0.917 / 0.156) reports algorithms and some key material from pyca/cryptography and finds nothing in 12 of the 16 Python entries.
- **cdxgen** 13.6.0 inventories libraries, keystores and certificates for Python but does not analyze algorithm use in source, so it scores 0 and reports nothing in all 16 entries.

## History of pinned crypto-finder runs

Python only, every run against the ground truth of 2026-11-08 and the same matching rules. All languages together: [README.md](README.md#history-of-pinned-runs).

| crypto-finder | Ruleset | Scored | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-08 | 153 / 29 / 142 | 0.841 / 0.519 | 0.796 / 0.533 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 238 / 17 / 57 | 0.933 / 0.807 | 0.915 / 0.858 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 238 / 17 / 57 | 0.933 / 0.807 | 0.915 / 0.858 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 280 / 6 / 15 | 0.979 / 0.949 | 0.972 / 0.986 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 293 / 6 / 2 | 0.980 / 0.993 | 0.972 / 1.000 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 293 / 6 / 2 | 0.980 / 0.993 | 0.972 / 1.000 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 295 / 6 / 0 | 0.980 / 1.000 | 0.972 / 1.000 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 295 / 6 / 0 | 0.980 / 1.000 | 0.972 / 1.000 |

## Run records

Exact versions, commands and the full score report of each tool: [crypto-finder](crypto-finder/python-pilot/README.md), [sonar-cryptography](sonar-cryptography/python-pilot/README.md), [cdxgen](cdxgen/python-pilot/README.md).
