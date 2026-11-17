# JavaScript pilot: tools compared

crypto-finder finds all 373 JavaScript assets with no false positive and passes every JavaScript target cell.


## Summary

| | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| JavaScript crypto found | **100%** (373 of 373) | not supported | 9% (32 of 373) |
| Findings that are right | **100%** | not supported | 100% |

**Found**: of all the crypto in the code, the share the tool reported (recall). **Right**: of what the tool reported, the share that is real crypto (precision).

## Scores in detail

Suite A, 16 JavaScript entries (`fixtures/javascript/*`, one per library, the standard library included), 373 assets in 279 objects, scored on 2026-11-17 against the ground truth of 2026-11-08. crypto-finder is v0.54.0 with remote ruleset v1.48.0, run on 2026-11-17 with the same predictions as v0.53.1 on 2026-11-12; sonar-cryptography and cdxgen are scored from their stored predictions.

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| javascript | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 373 / 0 / 0 | 321 / 373 (0.861) | 259 / 279 (0.928) |
| javascript | cdxgen | 1.000 | 0.086 | 0.158 | 1.000 | 0.090 | 0.164 | 32 / 0 / 341 | 7 / 32 (0.219) | 7 / 25 (0.280) |

`op` counts every annotated operation; `obj` counts each crypto object once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed means the match also names the right algorithm and parameters.

## crypto-finder against the targets

Precision >= 0.90 and recall >= 0.85 per cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). 5 of 5 JavaScript cells pass at object level.

| View | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| object (headline) | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| operation | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |

## What crypto-finder still gets wrong

None: no false positive and no false negative. 52 of 373 matches are not attributed: the asset is found, but the algorithm name or a parameter differs from the ground truth ([score.md](crypto-finder/javascript-pilot/score.md), "Attribution").

## Other tools

- **cdxgen** (object level 1.000 / 0.090) has no false positives, but finds the algorithm strings the code passes rather than the operations, so only 7 of its 32 matches attribute. It reports nothing in 11 of the 16 entries.
- **sonar-cryptography** has no JavaScript support and was not run.

## History of pinned crypto-finder runs

JavaScript only, every run against the ground truth of 2026-11-08 and the same matching rules. All languages together: [README.md](README.md#history-of-pinned-runs).

| crypto-finder | Ruleset | Scored | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-08 | 232 / 10 / 141 | 0.959 / 0.622 | 0.941 / 0.570 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 319 / 6 / 54 | 0.982 / 0.855 | 0.975 / 0.853 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 345 / 6 / 28 | 0.983 / 0.925 | 0.977 / 0.928 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 345 / 6 / 28 | 0.983 / 0.925 | 0.977 / 0.928 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 372 / 0 / 1 | 1.000 / 0.997 | 1.000 / 1.000 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 372 / 0 / 1 | 1.000 / 0.997 | 1.000 / 1.000 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 373 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 373 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |

## Run records

Exact versions, commands and the full score report of each tool: [crypto-finder](crypto-finder/javascript-pilot/README.md), [sonar-cryptography](sonar-cryptography/javascript-pilot/README.md) (not run), [cdxgen](cdxgen/javascript-pilot/README.md).
