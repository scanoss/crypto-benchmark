# crypto-finder, suite A Rust pilot

crypto-finder v0.54.0 with remote ruleset v1.48.0, run on 2026-11-17 over the Rust suite A entries and scored against the ground truth of 2026-11-08. `score.md` is the full `cbench score` report; `predictions/crypto-finder/` holds the adapted findings. What it still gets wrong, and the history of earlier runs, are on the [Rust pilot page](../../rust-pilot.md).

| Level | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.985 | 0.985 | 0.985 | 194 | 3 | 3 | 9 | 166 / 194 (0.856) |
| object | 0.980 | 0.987 | 0.983 | 149 | 3 | 2 | 54 | 134 / 149 (0.899) |

## Run record

- **crypto-finder:** v0.54.0, the released `crypto-finder_0.54.0_linux_amd64.tar.gz` (sha256 `06255e9dace89710ee81308aedecd9e8e11b9ea37199b89afe9a81962971a1d0`, checked against the release's checksums file). The binary reports version `0.54.0`, commit `1bf3cb12ebf3731f40917e8dac4aea54c40f4567`. Its only change from v0.53.1 is how it downloads the ruleset: from `/v3/cryptography/rulesets` as a bundle with a signed lease, kept encrypted in the cache. The run on 2026-11-12 with v0.53.1 gave the same predictions, asset for asset; only `tool.version` differs.
- **Ruleset:** remote, v1.48.0, the latest release on 2026-11-17, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, as the interim report records it. The prediction files carry the same values in `tool.config`.
- **Commands,** per entry `ID` with directory `fixtures/rust/CRATE`:

```sh
crypto-finder scan --ruleset-version v1.48.0 --include-tests --export-callgraph-project-reachability \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json fixtures/rust/CRATE
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json --out results/crypto-finder/rust-pilot/predictions/crypto-finder/ID.json
go run ./cmd/cbench score --tool crypto-finder --predictions results/crypto-finder/rust-pilot/predictions \
  --out results/crypto-finder/rust-pilot/score.md
```

`--include-tests` is the current configuration; suite A has no test files, so it changes nothing here. `--export-callgraph-project-reachability` classifies reachability against the program's own code: every suite A program is an application with a `main`. There is no `--scan-dependencies`, because suite A ground truth covers project code only.

The score report lists the entries of the other languages as not run, since each pilot covers one language.
