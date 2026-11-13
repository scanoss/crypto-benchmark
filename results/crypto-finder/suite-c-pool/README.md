# crypto-finder, suite C first `pool` score

**crypto-finder v0.53.1 with remote ruleset v1.48.0 matches 444 of the 993 adjudicated pool assets with 24 false positives: precision 0.949, recall 0.447 relative to the pool.** These are the frozen predictions of the [first score](../suite-c/README.md), re-keyed to the `<id>-pool` entries. Nothing was rerun. [suite-c.md](../../suite-c.md#pool-scope-first-score) compares the tools.

| Level | Precision | Recall, relative to pool | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.949 | 0.447 | 0.608 | 444 | 24 | 549 | 14 | 429 / 444 (0.966) |
| object | 0.944 | 0.435 | 0.596 | 403 | 24 | 523 | 55 | 398 / 403 (0.988) |

- **Scope:** `pool`, the whole repository of each entry: source files of the entry's language outside excluded paths ([METHODOLOGY.md](../../../METHODOLOGY.md#pool-scope)). Precision is exact. Recall counts only the assets some tool or the grep found, so it overstates true recall.
- **Dropped outside the pool frame:** 143 predictions, all in files of another language than the entry's (TypeScript, JavaScript and Python files in the Go and Rust repositories, Rust and Go files in firecrawl). They are not scored.
- **Test sources were not scanned,** as in the `files` first score: this configuration used crypto-finder's default. 436 of the 549 FN are in test-path files the [later `--include-tests` run](../suite-c-pool-include-tests/README.md) finds.

`score.md` is the full `cbench score` report. `predictions/crypto-finder/` holds the prediction files.

## Run record

- **Predictions:** `results/crypto-finder/suite-c/predictions/crypto-finder/ID.json`, unchanged, with one line changed: `"entry": "ID"` became `"entry": "ID-pool"`. The tool, version, configuration and every asset are byte-identical. Those files cover the whole repository; the `files` score only ignored the predictions outside its sample.
- **Tool run:** none. The run record of the scan is [suite-c/README.md](../suite-c/README.md): crypto-finder v0.53.1 (commit `20722e3aadc1dc23020e07a1050e9c7562554cf6`), ruleset v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, default configuration.
- **Ground truth:** `groundtruth/C/<id>-pool.yaml`, 993 assets, with the verdict of every pool item in `<id>-pool.verdicts.tsv` ([annotation log](../../../docs/annotation-log.md#suite-c-pool-scope-2026-11-13)).
- **Commands,** per entry `ID`:

```sh
mkdir -p results/crypto-finder/suite-c-pool/predictions/crypto-finder
sed 's/^  "entry": "ID",$/  "entry": "ID-pool",/' \
  results/crypto-finder/suite-c/predictions/crypto-finder/ID.json \
  > results/crypto-finder/suite-c-pool/predictions/crypto-finder/ID-pool.json
go run ./cmd/cbench score --tool crypto-finder --predictions results/crypto-finder/suite-c-pool/predictions \
  --out results/crypto-finder/suite-c-pool/score.md
```

Reachability is not scored in suite C; the `reachable` answers are crypto-finder's own and every ground-truth asset is `unknown`.
