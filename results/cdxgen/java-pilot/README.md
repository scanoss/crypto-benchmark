# cdxgen, suite A Java pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Java pilot page](../../java-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | n/a | 0.000 | 0.000 | 0 | 0 | 349 |
| object | n/a | 0.000 | 0.000 | 0 | 0 | 221 |

Run on 2026-11-04 over the same 16 Java suite A entries as [crypto-finder](../../crypto-finder/java-pilot/README.md) and [sonar-cryptography](../../sonar-cryptography/java-pilot/README.md). `score.md` is the `cbench score` report. `predictions/cdxgen/` holds the adapted predictions, all of them empty. `commands/` holds the exact commands that `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative.

## Run record

- **cdxgen:** 13.6.0, from an npm install of `@cdxgen/cdxgen@13.6.0` passed as `CDXGEN_BIN`, on Node.js 24.19.0. The script checked the version.
- **Input:** a scratch copy of each fixture directory, so the analyzer's output never lands in the repository. The copy is unbuilt, as committed.
- **Commands,** per entry `ID` with directory `fixtures/java/LIB`:

```sh
cp -r fixtures/java/LIB work/cdx-src/LIB
CDXGEN_BIN=.../node_modules/.bin scripts/run-cdxgen.sh work/cdx-src/LIB work/cdxgen/ID.cbom.json   # runs: cbom -o OUT DIR
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath work/cdx-src/LIB)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/java-pilot/predictions/cdxgen/ID.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/java-pilot/predictions \
  --out results/cdxgen/java-pilot/score.md
```

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

cdxgen 13.6.0's `cbom` command has source analysis for JavaScript and TypeScript only. For a Maven project it inventories the dependencies from `pom.xml` as `library` and `framework` components, and for four entries it adds algorithm components that carry no location: `rsa` and `sha-256` for bcpkix-jdk15on, and `aes` for bcprov-jdk15on and xmlsec. The adapter counts those four in each prediction file's `metadata` (`components_without_location`), since an asset with no location cannot be matched. stdlib, which has no dependencies, gets an empty CBOM. This is the configuration [METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration) describes; the 310 FN have that one diagnosis, so they are not listed.
