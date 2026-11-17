# Go pilot: tools compared

crypto-finder finds all 202 Go assets with no false positive and passes every Go target cell.

Suite A, 16 Go entries (`fixtures/go/*`, one per library, the standard library included), 202 assets in 152 objects, scored on 2026-11-17 against the ground truth of 2026-11-08. crypto-finder is v0.54.0 with remote ruleset v1.48.0, run on 2026-11-17 with the same predictions as v0.53.1 on 2026-11-12; sonar-cryptography and cdxgen are scored from their stored predictions.

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| go | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 202 / 0 / 0 | 177 / 202 (0.876) | 144 / 152 (0.947) |
| go | sonar-cryptography | 1.000 | 0.243 | 0.390 | 1.000 | 0.276 | 0.433 | 49 / 0 / 153 | 40 / 49 (0.816) | 40 / 42 (0.952) |
| go | cdxgen | 0.533 | 0.243 | 0.333 | 0.506 | 0.289 | 0.368 | 49 / 43 / 153 | 35 / 49 (0.714) | 31 / 44 (0.705) |

`op` counts every annotated operation; `obj` counts each crypto object once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed means the match also names the right algorithm and parameters.

## crypto-finder against the targets

Precision >= 0.90 and recall >= 0.85 per cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). 5 of 5 Go cells pass at object level.

| View | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| object (headline) | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| operation | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |

## What crypto-finder still gets wrong

None: no false positive and no false negative. 25 of 202 matches are not attributed ([score.md](crypto-finder/go-pilot/score.md), "Attribution").

## Other tools

- **sonar-cryptography** (object level 1.000 / 0.276) makes no false positive but reports algorithms only, and nothing in 6 of the 16 Go entries.
- **cdxgen** runs golem through `evinse -l go` (object level 0.506 / 0.289). It finds algorithms with locations; most of its 43 false positives are keys, passwords and tokens it derives from variable names. It reports nothing in 5 of the 16 entries.

## History of pinned crypto-finder runs

Go only, every run against the ground truth of 2026-11-08 and the same matching rules. All languages together: [README.md](README.md#history-of-pinned-runs).

| crypto-finder | Ruleset | Scored | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-08 | 142 / 25 / 60 | 0.850 / 0.703 | 0.805 / 0.678 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 186 / 17 / 16 | 0.916 / 0.921 | 0.892 / 0.928 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 188 / 10 / 14 | 0.949 / 0.931 | 0.935 / 0.941 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 199 / 2 / 3 | 0.990 / 0.985 | 0.987 / 0.980 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 199 / 0 / 3 | 1.000 / 0.985 | 1.000 / 0.980 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 199 / 0 / 3 | 1.000 / 0.985 | 1.000 / 0.980 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 202 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 202 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |

## Run records

Exact versions, commands and the full score report of each tool: [crypto-finder](crypto-finder/go-pilot/README.md), [sonar-cryptography](sonar-cryptography/go-pilot/README.md), [cdxgen](cdxgen/go-pilot/README.md).
