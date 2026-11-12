# cdxgen, suite A Python pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Python pilot page](../../python-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | n/a | 0.000 | 0.000 | 0 | 0 | 295 |
| object | n/a | 0.000 | 0.000 | 0 | 0 | 212 |

Run on 2026-11-04 over the same 16 Python suite A entries as [crypto-finder](../../crypto-finder/python-pilot/README.md), and run again the same day with the same image, script and commands. `score.md` is the `cbench score` report. `predictions/cdxgen/` holds the adapted predictions. `commands/` holds the exact commands that `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative.

## Run record

- **cdxgen:** 13.6.0, run through `npm exec`, with OpenJDK 21.0.12 for atom.
- **Commands,** per entry `ID` with directory `fixtures/python/LIB`:

```sh
scripts/run-cdxgen.sh fixtures/python/LIB work/cdxgen/ID.cbom.json     # runs: cbom -o OUT DIR
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath fixtures/python/LIB)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/python-pilot/predictions/cdxgen/ID.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/python-pilot/predictions \
  --out results/cdxgen/python-pilot/score.md
```

The second run used the same cdxgen version, script and adapter call. Its prediction files and recorded commands are byte-identical to the first run's. For the third score, the second run's CBOMs were kept, and re-adapting them reproduced the prediction files byte for byte. The current score is against the ground truth at commit `d12fe34`.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

Every entry scores the same: no predictions, so every ground-truth asset is an FN, and precision and attribution are n/a in every row, overall, per asset type and per library (`score.md`). Each CBOM lists the packages that `requirements.txt` resolves to as `library` components, and holds no `cryptographic-asset` component. The python-stdlib CBOM is empty.

This matches [METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration): for Python, cdxgen 13.6.0 inventories crypto libraries, keystores and certificates, but not algorithm use in source. Its atom analyzer also failed on 15 of the 16 entries ("Failed to enhance CPG: java.util.NoSuchElementException", then an OverflowDB storage error), and wrote no reachables or usages slice. python-stdlib has no third-party package, and atom did not report a failure there. The atom failure is not the cause of the zero: by the methodology above, cdxgen has no Python algorithm analysis for those slices to feed.

The FN have one diagnosis, so they are not listed one by one: there is no Python source analysis for algorithms.
