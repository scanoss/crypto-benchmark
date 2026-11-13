# cdxgen, suite C first score

cdxgen 13.6.0, run on 2026-11-13 over the ten suite C entries and scored on the `files` scope: 142 sampled files, 214 assets. This is the [first score](../../../METHODOLOGY.md#first-score-freeze), published as is. [suite-c.md](../../suite-c.md) compares the tools.

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 0.353 | 0.084 | 0.136 | 18 | 33 | 196 |
| object | 0.353 | 0.090 | 0.143 | 18 | 33 | 182 |

`score.md` is the full `cbench score` report, `predictions/cdxgen/` the adapted predictions and `commands/` the commands `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative. 714 predictions fall outside the sampled files and are ignored.

## Run record

- **cdxgen:** 13.6.0 from npm (the script checked `CycloneDX Generator 13.6.0` on every run), Node.js 24.19.0, OpenJDK 21.0.12 on `PATH` for atom. Analyzers bundled with it: golem (Go), rusi (Rust), atom (JavaScript and TypeScript slices).
- **Checkouts:** each repository at its pinned commit, fetched with `cbench fetch` into `work/corpus/ID`. cdxgen and atom write files into the scanned directory, so each run scanned a fresh copy.
- **Mode,** chosen by the script from the root, as [METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration) records: `go.mod` gives `cdxgen -t go` then `evinse -l go` (both Go entries); `Cargo.toml` gives `-t rust` and `evinse -l rust` (both Rust entries, see deviation 1); anything else runs `cbom` (JavaScript, Python and Java). For Python and Java, cdxgen 13.6.0 inventories libraries, keystores and certificates but not algorithm use in source.
- **Commands,** per entry `ID`:

```sh
cp -a work/corpus/ID work/cdx-src/ID
scripts/run-cdxgen.sh work/cdx-src/ID work/cdxgen/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath work/cdx-src/ID)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/suite-c/predictions/cdxgen/ID.json
```

Wall time ranged from 13 s (cc-switch) to 41 minutes (yt-dlp, mostly atom slicing).

## Deviations and notes

| # | Entry | What happened | Effect |
|---|---|---|---|
| 1 | `rust-c-cc-switch` | The repository has no `Cargo.toml` at its root; its only crate is `src-tauri/`. Unmodified, the script would run `cbom` (JavaScript mode), not the Rust configuration. Forcing Rust mode on the root made rusi exit 3 with no report (0 assets). The scored run is the unmodified script on `work/cdx-src/rust-c-cc-switch/src-tauri`. rusi writes locations relative to the scanned directory, so the 418 `cdx:rusi:crypto:sourceLocation` values were prefixed with `src-tauri/` before adapting with `--source-root work/cdx-src/rust-c-cc-switch`. Every resulting path exists in the checkout. | The Rust configuration ran on the crate. The repository's TypeScript front end is not part of a Rust entry. |
| 2 | `javascript-c-freecodecamp` | atom's reachables slice hit its default timeout (19.5 minutes per slice) and its callsites slice failed; the usages slice completed. The recorded default was kept for the scored run. | 3 TP, 0 FP. A separate run with a 2-hour atom timeout gave the same 3 / 0 / 18 on this entry; it is not the prediction file. |
| 3 | `python-c-yt-dlp` | Both atom slices timed out (41 minutes in all). | None on the score: cdxgen 13.6.0 has no Python algorithm analysis. 17 certificate and key files under `test/` came out as components with a file but no line; the adapter dropped them (`locations_dropped: 17`). |
| 4 | `java-c-mall`, `java-c-jadx` | cdxgen's own dependency resolution failed (mall's Maven repository, jadx's Gradle `buildSrc`). | None on the score: the CBOMs are build-file inventories with no located crypto, as in the suite A Java pilot. |

## What it reports

- **JavaScript (8 TP of 57, no FP).** `createHash`, `createHmac` and `crypto.subtle.digest` calls. It reports no `jsonwebtoken` sign or verify, no `randomUUID` or `randomBytes` draw and nothing in the `__tests__` directories of firecrawl.
- **Rust (8 TP of 32, 27 FP).** Its TPs are `Sha256::digest` calls, including three inside `format!(...)`. Almost every FP is `related-crypto-material` named after a variable: `input_tokens`, `cache_read_tokens`, `refresh_token`, `password`, `key`, `secret`, `nonce` on lines that compute no crypto. It misses `Sha256::new()` hashers, the rustls client configuration and handshakes, and clash-verge-rev's AES-256-GCM cipher and nonce draw.
- **Go (2 TP of 61, 6 FP).** golem on its default package pattern reports no prediction in a `_test.go` file and few in the rest: it misses the `sha256.Sum256`, `md5.Sum`, `sha256.New()` and `rand.Read` calls and ollama's Ed25519 key generation. Its 6 FP are material named after variables (`sessionKey`, `privKeyPath`).
- **Python and Java (0 TP).** No algorithm analysis in source for these languages, as in suite A.
