# sonar-cryptography, suite A Go pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Go pilot page](../../go-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 1.000 | 0.243 | 0.390 | 49 | 0 | 153 |
| object | 1.000 | 0.276 | 0.433 | 42 | 0 | 110 |

Run on 2026-11-04 over the same 16 Go suite A entries as [crypto-finder](../../crypto-finder/go-pilot/README.md), scored against the ground truth at commit `12bf5cd`: 180 assets in 130 objects. `score.md` is the `cbench score` report, `predictions/sonar-cryptography/` the adapted predictions and `commands/` the commands `scripts/run-sonar-cryptography.sh` recorded for each entry.

## Run record

- **sonar-cryptography** 1.7.0 in the CBOMkit-action v2.3.0 container, pinned by digest, with cbomkit-lib 1.3.0 and the SonarQube Go frontend embedded there. The script checks the three embedded versions and refused nothing.
- **Language:** each Go entry has `go.mod` at its root, so the script picked `go` by itself (`CBOMKIT_LANGUAGES=go`). The runner needed no change for Go.
- **Commands,** per entry `ID` with directory `fixtures/go/LIB`:

```sh
scripts/run-sonar-cryptography.sh fixtures/go/LIB work/sonar-cryptography/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0 \
  --entry ID --input work/sonar-cryptography/ID.cbom.json \
  --out results/sonar-cryptography/go-pilot/predictions/sonar-cryptography/ID.json
go run ./cmd/cbench score --tool sonar-cryptography --predictions results/sonar-cryptography/go-pilot/predictions \
  --out results/sonar-cryptography/go-pilot/score.md
```

Every component in the 16 CBOMs had a location, so the adapter dropped none. The CBOMs carry no reachability, so the predictions report none.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

sonar-cryptography makes one false positive here, and it finds 48 of the 133 algorithm assets. Everything it finds is an algorithm.

- **What it covers.** Go's standard library MD5, SHA-1, SHA-256, SHA-512, SHA3-256 and SHA3-512, HMAC, AES, 3DES, RC4, RSA key generation and OAEP, ECDSA and Ed25519 key generation, ML-KEM-768 and PBKDF2, and x/crypto's pbkdf2, hkdf and sha3 (SHA3-256 and SHA3-512). Its one FP is X25519: it reports `ECDH` on `ecdh.X25519()`, one line before the key generation the guide annotates. Its object recall is 0.315, higher than the 0.267 at operation level, because it reports a key once where it is generated and not each sign or verify on it.
- **What it does not cover.** No certificate, protocol or related-crypto-material asset: no `tls.Config`, TLS or SSH session, `x509.CreateCertificate` or `ParseCertificate`, key serialization, IV, nonce, salt or shared secret, and no random draw. No go-ethereum, bcrypt, blake2b, chacha20poly1305, curve25519 or scrypt call, no JWT call, no SSH signer or agent call, no signature or key agreement after the key generation (except the RSA signatures), and no Keccak-256 or SHAKE256.
- **Granularity.** On a key generation it often adds a `NATIVEPRNG` and a SHA-512 component at the same location (for Ed25519), which become duplicates and count as neither TP nor FP.
- **Attribution.** 40 of 48 matches are attributed. The misses are `RSA-PKCS1-1.5` for the PSS and PKCS#1 v1.5 signatures, which drops the digest and, for PSS, names the wrong padding (4); the SHA-256 that HMAC and RSA-OAEP use reported in place of the HMAC and RSA-OAEP call (3); and `HKDF-SHA-256`, which methodology rule 6 does not split (1).
