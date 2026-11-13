# crypto-finder, suite C `pool` score of the later `--include-tests` run

**Later run, 2026-11-13, configuration changed after the first score was seen.** It matches 880 of the 993 adjudicated pool assets with 45 false positives: precision 0.951, recall 0.886 relative to the pool. It sits beside the [first pool score](../suite-c-pool/README.md) (444 / 24 / 549, 0.949 / 0.447) and never replaces it ([METHODOLOGY.md](../../../METHODOLOGY.md#first-score-freeze)).

| Level | Precision | Recall, relative to pool | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.951 | 0.886 | 0.918 | 880 | 45 | 113 | 19 | 861 / 880 (0.978) |
| object | 0.949 | 0.896 | 0.922 | 830 | 45 | 96 | 69 | 824 / 830 (0.993) |

- **Pooled, so not independent of the pool.** This run's findings are one of the pool's sources ([METHODOLOGY.md](../../../METHODOLOGY.md#pool-scope)). Pool recall favors the tools that built the pool, and this run brought in 365 assets no other source run found.
- **Dropped outside the pool frame:** 208 predictions, all in files of another language than the entry's.
- **Against the first pool score:** 436 FN become TP, all in test-path files (Go `_test.go` 207, JavaScript `__tests__` 180, Python `test/` 37 and `tests/` 12), and 21 FP are added, all in test-path files. Every FN and FP of the first score is still in this run, or became a TP.

`score.md` is the full `cbench score` report. `predictions/crypto-finder/` holds the prediction files.

## Run record

- **Predictions:** `results/crypto-finder/suite-c-include-tests/predictions/crypto-finder/ID.json`, unchanged, with one line changed: `"entry": "ID"` became `"entry": "ID-pool"`. Every asset is byte-identical.
- **Tool run:** none. The scan is recorded in [suite-c-include-tests/README.md](../suite-c-include-tests/README.md): the first score's binary, ruleset and checkouts, plus `--include-tests`.
- **Commands,** per entry `ID`:

```sh
mkdir -p results/crypto-finder/suite-c-pool-include-tests/predictions/crypto-finder
sed 's/^  "entry": "ID",$/  "entry": "ID-pool",/' \
  results/crypto-finder/suite-c-include-tests/predictions/crypto-finder/ID.json \
  > results/crypto-finder/suite-c-pool-include-tests/predictions/crypto-finder/ID-pool.json
go run ./cmd/cbench score --tool crypto-finder \
  --predictions results/crypto-finder/suite-c-pool-include-tests/predictions \
  --out results/crypto-finder/suite-c-pool-include-tests/score.md
```
