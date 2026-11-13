# cdxgen, suite C first `pool` score

**cdxgen 13.6.0 matches 76 of the 993 adjudicated pool assets with 689 false positives: precision 0.099, recall 0.077 relative to the pool.** These are the frozen predictions of its [first score](../suite-c/README.md), re-keyed to the `<id>-pool` entries. Nothing was rerun. [suite-c.md](../../suite-c.md#pool-scope-first-score) compares the tools.

| Level | Precision | Recall, relative to pool | F1 | TP | FP | FN | Duplicates | Attributed TP / TP |
|---|---|---|---|---|---|---|---|---|
| operation | 0.099 | 0.077 | 0.086 | 76 | 689 | 917 | 0 | 63 / 76 (0.829) |
| object | 0.098 | 0.081 | 0.089 | 75 | 689 | 851 | 1 | 63 / 75 (0.840) |

- **Scope:** `pool`, all ten entries. Recall counts only the assets some tool or the grep found.
- **Dropped outside the pool frame:** none.
- **The false positives:** 679 of the 689 are `related-crypto-material`, 570 of them `key` and 50 `token`, all in Go (245) and Rust (444) code. Every one was adjudicated as a pool item: they are lines where an identifier is named like key material (`input_tokens`, `output_key`, `secret := ""`, `SchemaStructuralKeys`) and no crypto is computed. The reported lines are correct; the lines hold no asset.
- **What it finds:** JavaScript `createHash`, `createHmac` and `crypto.subtle.digest` (50 TP), Rust `Sha256::digest` (19 TP), Go 7. Python and Java: none, as in the `files` score (cdxgen 13.6.0 has no algorithm analysis in their source).

`score.md` is the full `cbench score` report. `predictions/cdxgen/` holds the prediction files.

## Run record

- **Predictions:** `results/cdxgen/suite-c/predictions/cdxgen/ID.json`, unchanged, with one line changed: `"entry": "ID"` became `"entry": "ID-pool"`. Every asset is byte-identical, including the `src-tauri/` prefix of cc-switch's paths ([deviation 1](../suite-c/README.md#deviations-and-notes)).
- **Tool run:** none. The scan is recorded in [suite-c/README.md](../suite-c/README.md).
- **Commands,** per entry `ID`:

```sh
mkdir -p results/cdxgen/suite-c-pool/predictions/cdxgen
sed 's/^  "entry": "ID",$/  "entry": "ID-pool",/' \
  results/cdxgen/suite-c/predictions/cdxgen/ID.json \
  > results/cdxgen/suite-c-pool/predictions/cdxgen/ID-pool.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/suite-c-pool/predictions \
  --out results/cdxgen/suite-c-pool/score.md
```
