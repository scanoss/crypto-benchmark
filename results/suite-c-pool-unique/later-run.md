# Assets only one run found

Pool entries (10): java-c-mall-pool, java-c-jadx-pool, python-c-yt-dlp-pool, python-c-graphify-pool, go-c-ollama-pool, go-c-caveman-pool, rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool

Each run is matched against the pool ground truth by the score's rules. An asset counts for a run when the run matches it, and as found only by that run when no other run in this comparison matches it. Recall is relative to the pool: 993 adjudicated assets, not every asset in the repositories.

| Run | Pool assets found | Found only by this run | Entries not run |
|---|---|---|---|
| crypto-finder --include-tests (later run) | 880 | 649 | none |
| sonar-cryptography 1.7.0 | 169 | 0 | rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool |
| cdxgen 13.6.0 | 76 | 9 | none |
| found by none of these runs | | 104 | |

## Found only by one run, by language

| language | crypto-finder --include-tests (later run) | sonar-cryptography 1.7.0 | cdxgen 13.6.0 | none of these runs |
|---|---|---|---|---|
| go | 178 | 0 | 1 | 19 |
| java | 7 | 0 | 0 | 3 |
| javascript | 284 | 0 | 0 | 13 |
| python | 150 | 0 | 0 | 62 |
| rust | 30 | 0 | 8 | 7 |

## Found only by one run, by asset type

| asset type | crypto-finder --include-tests (later run) | sonar-cryptography 1.7.0 | cdxgen 13.6.0 | none of these runs |
|---|---|---|---|---|
| algorithm | 531 | 0 | 8 | 58 |
| certificate | 37 | 0 | 0 | 4 |
| protocol | 35 | 0 | 0 | 11 |
| related-crypto-material | 46 | 0 | 1 | 31 |

