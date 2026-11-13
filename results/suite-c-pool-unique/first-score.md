# Assets only one run found

Pool entries (10): java-c-mall-pool, java-c-jadx-pool, python-c-yt-dlp-pool, python-c-graphify-pool, go-c-ollama-pool, go-c-caveman-pool, rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool

Each run is matched against the pool ground truth by the score's rules. An asset counts for a run when the run matches it, and as found only by that run when no other run in this comparison matches it. Recall is relative to the pool: 993 adjudicated assets, not every asset in the repositories.

| Run | Pool assets found | Found only by this run | Entries not run |
|---|---|---|---|
| crypto-finder v0.53.1 | 444 | 284 | none |
| sonar-cryptography 1.7.0 | 169 | 71 | rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool |
| cdxgen 13.6.0 | 76 | 9 | none |
| found by none of these runs | | 469 | |

## Found only by one run, by language

| language | crypto-finder v0.53.1 | sonar-cryptography 1.7.0 | cdxgen 13.6.0 | none of these runs |
|---|---|---|---|---|
| go | 42 | 71 | 1 | 155 |
| java | 7 | 0 | 0 | 3 |
| javascript | 104 | 0 | 0 | 193 |
| python | 101 | 0 | 0 | 111 |
| rust | 30 | 0 | 8 | 7 |

## Found only by one run, by asset type

| asset type | crypto-finder v0.53.1 | sonar-cryptography 1.7.0 | cdxgen 13.6.0 | none of these runs |
|---|---|---|---|---|
| algorithm | 254 | 71 | 8 | 335 |
| certificate | 8 | 0 | 0 | 33 |
| protocol | 9 | 0 | 0 | 37 |
| related-crypto-material | 13 | 0 | 1 | 64 |

