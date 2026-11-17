# crypto-finder, suite C later run with `--include-tests`

**Later run, 2026-11-13: crypto-finder with `--include-tests`, configuration changed after the first score was seen.** Under the [first score freeze](../../../METHODOLOGY.md#first-score-freeze) this is a separate, dated run. It does not replace or edit the [first score](../suite-c/README.md), which stays the suite C result for crypto-finder v0.53.1.

| Level | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.929 | 0.860 | 0.893 | 184 | 14 | 30 | 1 | 179 / 184 (0.973) |
| object | 0.925 | 0.860 | 0.891 | 172 | 14 | 28 | 13 | 172 / 172 (1.000) |

First score, for comparison: operation 108 / 8 / 106 (0.931 / 0.505), object 97 / 8 / 103 (0.924 / 0.485).

## What changed

Only `--include-tests`. The binary, ruleset, checkouts, commands and adapter are those of the first score.

| | First score | This run |
|---|---|---|
| Flag | default (test sources left out) | `--include-tests` |
| Why it changed | | The first score showed no prediction in any test-path file of the sample. The benchmark scores test files ([METHODOLOGY.md](../../../METHODOLOGY.md#suite-c-decisions), "Test files"), and the competitors were configured not to skip them. Suite A has no test files, so the default had never mattered. |
| Detection outside test paths | | Identical in every entry: the same findings with the same fields. Only crypto-finder's own `reachable`, depth and distance answers moved on some of them, and suite C does not score reachability. |

76 FN become TP, all in test-path files; 6 FP are added, all in test-path files (below). The prediction files do not record the flag in `tool.config`; this page does.

**Reproduced with v0.54.0 on 2026-11-17.** The same scan with crypto-finder v0.54.0 (tarball sha256 `06255e9dace89710ee81308aedecd9e8e11b9ea37199b89afe9a81962971a1d0`), which downloads ruleset v1.48.0 from `/v3/cryptography/rulesets` with a signed lease, gave the same predictions on all ten entries, asset for asset. Nothing here changed, and the pool gained no item.

## Run record

- **crypto-finder:** v0.53.1, the same released binary as the first score (tarball sha256 `3a7db29d52a911fe8d02dec1bfc36aece2f64e66ab0d70f63a58ea3fd93bd8e8`), commit `20722e3aadc1dc23020e07a1050e9c7562554cf6`.
- **Ruleset:** remote, v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, as every interim report records it.
- **Commands,** per entry `ID`, on the same checkouts:

```sh
crypto-finder scan --ruleset-version v1.48.0 --export-callgraph-project-reachability --include-tests \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json work/corpus/ID
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json \
  --out results/crypto-finder/suite-c-include-tests/predictions/crypto-finder/ID.json
go run ./cmd/cbench score --tool crypto-finder \
  --predictions results/crypto-finder/suite-c-include-tests/predictions \
  --out results/crypto-finder/suite-c-include-tests/score.md
```

Each scan took 13 s to 120 s, all exited 0, and the adapter accepted every entry.

## By language, operation level

| Language | First score TP / FP / FN | P / R | This run TP / FP / FN | P / R |
|---|---|---|---|---|
| go | 26 / 0 / 35 | 1.000 / 0.426 | 60 / 0 / 1 | 1.000 / 0.984 |
| java | 2 / 0 / 1 | 1.000 / 0.667 | 2 / 0 / 1 | 1.000 / 0.667 |
| javascript | 31 / 2 / 26 | 0.939 / 0.544 | 55 / 5 / 2 | 0.917 / 0.965 |
| python | 26 / 5 / 35 | 0.839 / 0.426 | 44 / 8 / 17 | 0.846 / 0.721 |
| rust | 23 / 1 / 9 | 0.958 / 0.719 | 23 / 1 / 9 | 0.958 / 0.719 |
| all | 108 / 8 / 106 | 0.931 / 0.505 | 184 / 14 / 30 | 0.929 / 0.860 |

Rust and Java do not move: the sample has no Rust test-path file and no Java ground truth in one, and crypto-finder already scanned Rust inline `#[cfg(test)]` modules and JavaScript `*.test.ts` files outside `__tests__` by default.

## What the test files add

- **Go (34 FN become TP).** Every `sha256.Sum256` and `sha256.New()` in the sampled `_test.go` files, and a `rand.Read`. One FN is left: caveman's `httptest.NewTLSServer` (a TLS server session).
- **JavaScript (24 FN become TP, 3 FP added).** firecrawl's `crypto.randomUUID` draws and jose RSA key generations in `__tests__`. Left: one `exportJWK(publicKey)` reported as an RS256 signature rather than key material. Added FP: the JWK object literals around `exportJWK` and a call to the test's own `sign({ alg: "RS512", ... })` helper, reported as RSA signatures.
- **Python (18 FN become TP, 3 FP added).** graphify's `hashlib.sha256` in `tests/`; yt-dlp's TLS server contexts and certificate loads in `test/`. Left: the six AES and IV assets of `devscripts/generate_aes_testdata.py` (yt-dlp's own AES functions and an `openssl` command line), one `wrap_bio` TLS session, and the five private keys loaded by `load_cert_chain`. Added FP: three TLS configuration setters (`verify_mode`, `maximum_version`, `set_ciphers`) reported as protocols.

The 16 FN and 8 FP outside test paths are those of the first score; [suite-c.md](../../suite-c.md#what-crypto-finder-misses) lists them by cause.
