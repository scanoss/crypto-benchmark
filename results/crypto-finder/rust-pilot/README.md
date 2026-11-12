# crypto-finder, suite A Rust pilot

crypto-finder v0.53.1 with remote ruleset v1.48.0, run on 2026-11-12 over the Rust suite A entries and scored against the ground truth of 2026-11-08. `score.md` is the full `cbench score` report; `predictions/crypto-finder/` holds the adapted findings. What it still gets wrong, and the history of earlier runs, are on the [Rust pilot page](../../rust-pilot.md).

| Level | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.985 | 0.985 | 0.985 | 194 | 3 | 3 | 9 | 166 / 194 (0.856) |
| object | 0.980 | 0.987 | 0.983 | 149 | 3 | 2 | 54 | 134 / 149 (0.899) |

## Run record

- **crypto-finder:** v0.53.1, the released `crypto-finder_0.53.1_linux_amd64.tar.gz` (sha256 `3a7db29d52a911fe8d02dec1bfc36aece2f64e66ab0d70f63a58ea3fd93bd8e8`, checked against the release's checksums file). The binary reports version `0.53.1`.
- **Ruleset:** remote, v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, as the interim report records it. The prediction files carry the same values in `tool.config`.
- **Commands,** per entry `ID` with directory `fixtures/rust/CRATE`:

```sh
crypto-finder scan --ruleset-version v1.48.0 --export-callgraph-project-reachability \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json fixtures/rust/CRATE
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json --out results/crypto-finder/rust-pilot/predictions/crypto-finder/ID.json
go run ./cmd/cbench score --tool crypto-finder --predictions results/crypto-finder/rust-pilot/predictions \
  --out results/crypto-finder/rust-pilot/score.md
```

`--export-callgraph-project-reachability` classifies reachability against the program's own code: every suite A program is an application with a `main`. There is no `--scan-dependencies`, because suite A ground truth covers project code only.

The score report lists the entries of the other languages as not run, since each pilot covers one language.
