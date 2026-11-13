# Suite C pool: assets only one tool found

`cbench unique` reports, for each compared run, the pool assets it matched and those no other compared run matched, by language and asset type. Every count is over the 993 adjudicated pool assets, so it is relative to the pool, not to every asset in the repositories. [suite-c.md](../suite-c.md#assets-only-one-tool-found) reads them.

| Report | Runs compared |
|---|---|
| [first-score.md](first-score.md) | the three first scores: crypto-finder v0.53.1, sonar-cryptography 1.7.0, cdxgen 13.6.0 |
| [later-run.md](later-run.md) | the later crypto-finder `--include-tests` run in place of its first score |
| [all-runs.md](all-runs.md) | all four runs; "found by none" is then the assets only the grep found |

```sh
CF=results/crypto-finder/suite-c-pool/predictions/crypto-finder
CFT=results/crypto-finder/suite-c-pool-include-tests/predictions/crypto-finder
SONAR=results/sonar-cryptography/suite-c-pool/predictions/sonar-cryptography
CDX=results/cdxgen/suite-c-pool/predictions/cdxgen
go run ./cmd/cbench unique --run "crypto-finder v0.53.1=$CF" --run "sonar-cryptography 1.7.0=$SONAR" \
  --run "cdxgen 13.6.0=$CDX" --out results/suite-c-pool-unique/first-score.md
go run ./cmd/cbench unique --run "crypto-finder --include-tests (later run)=$CFT" \
  --run "sonar-cryptography 1.7.0=$SONAR" --run "cdxgen 13.6.0=$CDX" --out results/suite-c-pool-unique/later-run.md
go run ./cmd/cbench unique --run "crypto-finder v0.53.1=$CF" --run "crypto-finder --include-tests (later run)=$CFT" \
  --run "sonar-cryptography 1.7.0=$SONAR" --run "cdxgen 13.6.0=$CDX" --out results/suite-c-pool-unique/all-runs.md
```
