# Java pilot: tools compared

crypto-finder finds all 349 Java assets with no false positive and passes every Java target cell.


## Summary

| | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| Java crypto found | **100%** (349 of 349) | 19% (66 of 349) | 0% (0 of 349) |
| Findings that are right | **100%** | 62% | n/a, it reports nothing |

**Found**: of all the crypto in the code, the share the tool reported (recall). **Right**: of what the tool reported, the share that is real crypto (precision).

## Scores in detail

Suite A, 16 Java entries (`fixtures/java/*`, one per library, the standard library included), 349 assets in 221 objects, scored on 2026-11-17 against the ground truth of 2026-11-08. crypto-finder is v0.54.0 with remote ruleset v1.48.0, run on 2026-11-17 with the same predictions as v0.53.1 on 2026-11-12; sonar-cryptography and cdxgen are scored from their stored predictions.

| Language | Tool | P op | R op | F1 op | P obj | R obj | F1 obj | TP / FP / FN op | Attributed op | Attributed obj |
|---|---|---:|---:|---:|---:|---:|---:|---|---|---|
| java | crypto-finder | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 1.000 | 349 / 0 / 0 | 297 / 349 (0.851) | 210 / 221 (0.950) |
| java | sonar-cryptography | 0.623 | 0.189 | 0.290 | 0.500 | 0.181 | 0.266 | 66 / 40 / 283 | 57 / 66 (0.864) | 37 / 40 (0.925) |
| java | cdxgen | n/a | 0.000 | 0.000 | n/a | 0.000 | 0.000 | 0 / 0 / 349 | 0 / 0 (n/a) | 0 / 0 (n/a) |

`op` counts every annotated operation; `obj` counts each crypto object once ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)). Attributed means the match also names the right algorithm and parameters.

## crypto-finder against the targets

Precision >= 0.90 and recall >= 0.85 per cell ([METHODOLOGY.md](../METHODOLOGY.md#targets)). 5 of 5 Java cells pass at object level.

| View | algorithm | certificate | protocol | related-crypto-material | all types |
|---|---|---|---|---|---|
| object (headline) | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |
| operation | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass | 1.000 / 1.000 pass |

## What crypto-finder still gets wrong

None: no false positive and no false negative. 52 of 349 matches are not attributed ([score.md](crypto-finder/java-pilot/score.md), "Attribution").

## Other tools

- **sonar-cryptography** (object level 0.500 / 0.181) covers JCA and Bouncy Castle and reports nothing in 9 of the 16 Java entries. Its algorithm precision is high (0.950 object level); most of its 40 false positives are key material (object level 0.026 / 0.014).
- **cdxgen** 13.6.0 inventories libraries, keystores and certificates for Java but does not analyze algorithm use in source, so it scores 0 and reports nothing in all 16 entries.

## History of pinned crypto-finder runs

Java only, every run against the ground truth of 2026-11-08 and the same matching rules. All languages together: [README.md](README.md#history-of-pinned-runs).

| crypto-finder | Ruleset | Scored | TP / FP / FN op | P / R op | P / R obj |
|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-08 | 179 / 35 / 170 | 0.836 / 0.513 | 0.778 / 0.557 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 275 / 15 / 74 | 0.948 / 0.788 | 0.926 / 0.846 |
| v0.51.0 | v1.44.0 | 2026-11-10 | 288 / 3 / 61 | 0.990 / 0.825 | 0.985 / 0.878 |
| v0.51.0 | v1.45.0 | 2026-11-11 | 317 / 2 / 32 | 0.994 / 0.908 | 0.991 / 0.950 |
| v0.51.0 | v1.46.0 | 2026-11-11 | 344 / 0 / 5 | 1.000 / 0.986 | 1.000 / 0.986 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 345 / 0 / 4 | 1.000 / 0.989 | 1.000 / 0.986 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 349 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 349 / 0 / 0 | 1.000 / 1.000 | 1.000 / 1.000 |

## Run records

Exact versions, commands and the full score report of each tool: [crypto-finder](crypto-finder/java-pilot/README.md), [sonar-cryptography](sonar-cryptography/java-pilot/README.md), [cdxgen](cdxgen/java-pilot/README.md).
