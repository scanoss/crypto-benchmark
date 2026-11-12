# cdxgen, suite A Rust pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Rust pilot page](../../rust-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 0.477 | 0.157 | 0.237 | 31 | 34 | 166 |
| object | 0.460 | 0.192 | 0.271 | 29 | 34 | 122 |

Run on 2026-11-04 over the same 15 Rust suite A entries as [crypto-finder](../../crypto-finder/rust-pilot/README.md). `score.md` is the `cbench score` report. `predictions/cdxgen/` holds the adapted predictions. `commands/` holds the exact commands that `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative.

## Run record

- **cdxgen:** 13.6.0, run through `npm exec` on Node 24.19.0. A directory with `Cargo.toml` goes through `cdxgen -t rust` and then `evinse -l rust`, which runs cdxgen's rusi source analyzer.
- **Commands,** per entry `ID` with directory `fixtures/rust/CRATE`:

```sh
scripts/run-cdxgen.sh fixtures/rust/CRATE work/cdxgen/ID.cbom.json   # runs: cdxgen -t rust, then evinse -l rust
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath fixtures/rust/CRATE)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/rust-pilot/predictions/cdxgen/ID.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/rust-pilot/predictions \
  --out results/cdxgen/rust-pilot/score.md
```

Every component in the 15 CBOMs carries a source location, so the adapter dropped none. CBOMs carry no reachability, so reachability is n/a.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

Unlike the Python pilot, where cdxgen reported nothing, rusi analyzes Rust source and finds something in 13 of the 15 entries. Almost all of it is related crypto material, found by variable name:

- **Material by variable name.** 61 of its 65 predictions are `related-crypto-material` components named after a variable (`key`, `nonce`, `salt`, `iv`, `key_pair`), placed on the line that binds it. Where that line is the one the ground truth annotates (a key loaded with `from_pem_slice`, a literal nonce or salt), it matches: 17 TPs, recall 0.415 on related crypto material. Where the variable is a key the program generates or draws, or an IV, nonce or salt buffer declared one line before the draw that fills it, the guide annotates no material there, so it is an FP: 44 FPs, 38 of material type `key` and 6 typed `nonce`.
- **Material type.** 3 of the 5 matched nonces and salts are not attributed, because rusi types every salt as `nonce`. It matches no IV and no shared secret.
- **Algorithms.** It reports 4 algorithm assets, all matched and attributed: openssl `pbkdf2_hmac`, the certificate fingerprint `digest(MessageDigest::sha256())`, `Sha1::digest` and `Sha512::digest`. Algorithm recall is 0.036. It reports no certificate and no protocol asset.
- **Nothing** in rust-md-5 and rust-siphasher.

The object view lowers precision (0.323 to 0.302) because the material FPs still count once each, and raises recall a little (0.115 to 0.140), because the 112 algorithm assets fold into fewer objects while each matched material is its own object.
