# crypto-finder, suite C first score

crypto-finder v0.53.1 with remote ruleset v1.48.0, run on 2026-11-13 over the ten suite C entries and scored on the `files` scope: 142 sampled files, 214 assets. This is the [first score](../../../METHODOLOGY.md#first-score-freeze), published as is. [suite-c.md](../../suite-c.md) compares the tools and lists every miss.

| Level | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.931 | 0.505 | 0.655 | 108 | 8 | 106 | 1 | 104 / 108 (0.963) |
| object | 0.924 | 0.485 | 0.636 | 97 | 8 | 103 | 12 | 97 / 97 (1.000) |

**Test sources were not scanned.** This run used crypto-finder's default, which leaves test sources out (the `--include-tests` flag adds them). Suite A has no test files, so the default never mattered there; the competitors were configured not to skip tests. 90 of the 106 FN are in `_test.go` files, Python `test/`, `tests/` and `devscripts/` directories and JavaScript `__tests__` directories, where this run has no prediction at all. A later run with `--include-tests` is recorded separately in [suite-c-include-tests](../suite-c-include-tests/README.md); it does not replace this score.

`score.md` is the full `cbench score` report and `predictions/crypto-finder/` the adapted findings. 508 predictions fall outside the sampled files and are ignored.

## Run record

- **crypto-finder:** v0.53.1, the released `crypto-finder_0.53.1_linux_amd64.tar.gz` (sha256 `3a7db29d52a911fe8d02dec1bfc36aece2f64e66ab0d70f63a58ea3fd93bd8e8`, checked against the release's checksums file). The binary reports version `0.53.1`, commit `20722e3aadc1dc23020e07a1050e9c7562554cf6`.
- **Ruleset:** remote, v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, as every interim report records it. The prediction files carry the same values in `tool.config`.
- **Checkouts:** each repository at its pinned commit, fetched with `cbench fetch` into `work/corpus/ID`, scanned whole and unbuilt. Java entries ran with `JAVA_HOME` set to Temurin JDK 21.0.12.
- **Commands,** per entry `ID`, as recorded for suite A:

```sh
crypto-finder scan --ruleset-version v1.48.0 --export-callgraph-project-reachability \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json work/corpus/ID
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json --out results/crypto-finder/suite-c/predictions/crypto-finder/ID.json
go run ./cmd/cbench score --tool crypto-finder --predictions results/crypto-finder/suite-c/predictions \
  --out results/crypto-finder/suite-c/score.md
```

No `--scan-dependencies`, default exclusions, default `--timeout 10m`. Each scan took 13 s to 74 s. No deviation from the recorded configuration, and the adapter accepted every entry.

Reachability is not scored in suite C ([METHODOLOGY.md](../../../METHODOLOGY.md#suite-c)); the `reachable` answers in the prediction files are crypto-finder's own and every ground-truth asset is `unknown`.
