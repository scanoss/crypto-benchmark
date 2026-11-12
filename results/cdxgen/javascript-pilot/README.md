# cdxgen, suite A JavaScript pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([JavaScript pilot page](../../javascript-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 1.000 | 0.086 | 0.158 | 32 | 0 | 341 |
| object | 1.000 | 0.090 | 0.164 | 25 | 0 | 254 |

Run on 2026-11-04 over the same 16 JavaScript suite A entries as [crypto-finder](../../crypto-finder/javascript-pilot/README.md). `score.md` is the `cbench score` report. `predictions/cdxgen/` holds the adapted predictions. `commands/` holds the exact commands that `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative.

## Run record

- **cdxgen:** 13.6.0, from an npm install of `@cdxgen/cdxgen@13.6.0` passed as `CDXGEN_BIN`, on Node.js 24.19.0. The script checked the version.
- **Commands,** per entry `ID` with directory `fixtures/javascript/LIB`:

```sh
CDXGEN_BIN=.../node_modules/.bin scripts/run-cdxgen.sh fixtures/javascript/LIB work/cdxgen/ID.cbom.json   # runs: cbom -o OUT DIR
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath fixtures/javascript/LIB)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/javascript-pilot/predictions/cdxgen/ID.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/javascript-pilot/predictions \
  --out results/cdxgen/javascript-pilot/score.md
```

cdxgen's atom analyzer writes a `.chen` directory of graph fragments into the scanned directory. `fixtures/javascript/.gitignore` keeps them out of the repository; the Python fixtures still carry the ones an earlier run committed.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

cdxgen's JavaScript source analysis finds algorithm names in five entries: the JWS and JWE `alg` strings in jose, jsonwebtoken and jwa, the `'sha256'`, `'rsa'`, `'ed25519'` and `'x25519'` strings of node:crypto calls, and `SHA-256` in @peculiar/x509's WebCrypto algorithm objects. Each CBOM component carries every line where its name occurs, so one `sha-256` component lands on the HMAC, HKDF, PBKDF2, RSA and ECDSA calls that pass `'sha256'`. Those lines are real crypto calls, which is why precision is 1.000, but the name says SHA-256 where the call is an HMAC or a signature, so only 7 of 32 matches attribute (the node:crypto Ed25519 and X25519 key generations, `createHash('sha256')`, the jose Ed25519 signature, and the jose `dir` and RSA-OAEP JWE A256GCM encryptions). It names JWS algorithms by their X.509 OID names (`sha256WithRSAEncryption`, `ecdsaWithSHA256`, `hmacWithSHA256`), which `internal/algo` does not map.

The other 11 entries get only `library` components: cdxgen does not recognize the noble packages, node-forge, tweetnacl, hash.js, fast-sha256, blake3-wasm, crypto-random-string, selfsigned or iron-webcrypto as crypto APIs. It reports no certificate, protocol or related-crypto-material asset, and no reachability. The 300 FN have those two diagnoses, so they are not listed one by one.
