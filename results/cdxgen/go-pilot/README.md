# cdxgen, suite A Go pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Go pilot page](../../go-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 0.533 | 0.243 | 0.333 | 49 | 43 | 153 |
| object | 0.506 | 0.289 | 0.368 | 44 | 43 | 108 |

Run on 2026-11-04 over the same 16 Go suite A entries as [crypto-finder](../../crypto-finder/go-pilot/README.md), scored against the ground truth at commit `12bf5cd`: 180 assets in 130 objects. `score.md` is the `cbench score` report, `predictions/cdxgen/` the adapted predictions and `commands/` the commands `scripts/run-cdxgen.sh` recorded for each entry, with paths made relative.

## Run record

- **cdxgen:** 13.6.0, run through `npm exec`. Each Go entry has `go.mod` at its root, so the script ran `cdxgen -t go` and then `evinse -l go`, cdxgen's golem analyzer, on its default package pattern `./...`.
- **Commands,** per entry `ID` with directory `fixtures/go/LIB`:

```sh
scripts/run-cdxgen.sh fixtures/go/LIB work/cdxgen/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen \
  --source-root "$(realpath fixtures/go/LIB)" --entry ID --input work/cdxgen/ID.cbom.json \
  --out results/cdxgen/go-pilot/predictions/cdxgen/ID.json
go run ./cmd/cbench score --tool cdxgen --predictions results/cdxgen/go-pilot/predictions \
  --out results/cdxgen/go-pilot/score.md
```

Every component had a location, so the adapter dropped none. The CBOMs carry no reachability.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

Unlike Python, where cdxgen 13.6.0 analyzes no source for algorithms, golem reads Go source and reports algorithms, keys and protocols with locations. It finds 35 of 180 assets, and 57 of its 92 predictions are FPs.

- **What it finds.** RSA and ECDSA key generation, X25519 in x/crypto/curve25519, stdlib hashes, HMAC, AES, PBKDF2, HKDF under its digest's name, the `tls.Config` literal, and some key serializations and parses.
- **What makes its FPs.** Almost all are `related-crypto-material` it derives from variable names and types: a `private-key` next to every key generation, a `public-key` on each line of an `x509.Certificate` template that mentions the key, `token` and `secret` for JWT values, `password` for a password variable, and `shared-secret` on the X25519 public-key derivations. The guide annotates none of these. It also reports the `x509.MarshalPKCS1PrivateKey` line in key-load and in both JWT RS256 programs as an X.509 certificate (3), and the RC4 key literal as a key. Two FPs are one or two lines from the asset they name: `curveX25519` on `ecdh.X25519()`, the line before the key generation, and TLS on a read of `server.TLSConfig`, two lines after the autocert `manager.TLSConfig()`.
- **Its related-crypto-material recall is 0.303.** 10 of its matches are material: `keyPEM`, `public-key` and `PKCS#1 private key` labels on eight key serializations and parses, and the two X25519 shared secrets. All 10 are attributed. Its `key`, `password` and `public-key` labels next to a salt or a key serialization match nothing and are FPs.
- **What it misses.** No go-ethereum, blake2b, chacha20poly1305, sha3 or ssh/agent call, no IV, nonce or salt, no certificate, no TLS or SSH session, and no autocert `TLSConfig` call.
- **Attribution.** 22 of 35 matches are attributed. Its names are often lower case or carry no parameters (`rsa` without the key size the call states, `ecdsa` without the curve, `hmacSHA`, `curveX25519`), and it names HKDF by its digest (`sha-256`).
