# sonar-cryptography, suite A Python pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Python pilot page](../../python-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 0.921 | 0.119 | 0.210 | 35 | 3 | 260 |
| object | 0.917 | 0.156 | 0.266 | 33 | 3 | 179 |

Run on 2026-11-04 over the same 16 Python suite A entries as [crypto-finder](../../crypto-finder/python-pilot/README.md) and [cdxgen](../../cdxgen/python-pilot/README.md), first scored against the ground truth at commit `3aff652`, and run and scored again the same day after methodology changes (below). The current score is against the ground truth at commit `d12fe34`. `score.md` is the `cbench score` report. `predictions/sonar-cryptography/` holds the adapted predictions. `commands/` holds the exact commands that `scripts/run-sonar-cryptography.sh` recorded for each entry.

## Run record

- **sonar-cryptography:** 1.7.0, embedded in CBOMkit-action v2.3.0 through cbomkit-lib 1.3.0, with the SonarQube Python frontend 5.29.0.35837.
- **Image:** `ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b`, Docker 29.1.3, no network during the scan.
- **Commands,** per entry `ID` with directory `fixtures/python/LIB`:

```sh
scripts/run-sonar-cryptography.sh fixtures/python/LIB work/sonar-cryptography/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0 \
  --entry ID --input work/sonar-cryptography/ID.cbom.json \
  --out results/sonar-cryptography/python-pilot/predictions/sonar-cryptography/ID.json
go run ./cmd/cbench score --tool sonar-cryptography --predictions results/sonar-cryptography/python-pilot/predictions \
  --out results/sonar-cryptography/python-pilot/score.md
```

The adapter read every CBOM without a note: no component without a location and no dropped location. Rerunning gives the same `score.md`; only the generated key names (`private-key@<uuid>`) in the predictions change.

The second run used the same image, script and commands; the recorded commands are byte-identical to the first run's. It was needed because the first run's CBOMs were not kept and the adapter now reads `material_type`. The new predictions equal the first ones asset for asset, apart from the generated key names and a `material_type` (`private-key` or `secret-key`) on the 15 key components.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

Operation level (op) and object level (obj). FP is the same in both views.

**The 12 entries with no prediction use unsupported libraries.** sonar-cryptography's Python rules cover pyca/cryptography only, and none of these fixtures imports it. The imports are `argon2`, `bcrypt`, `ecdsa` and `hashlib`, `joserfc`, `paramiko`, `passlib`, `Crypto` (pycryptodome), `nacl`, `pyotp` and `hashlib`, `rsa`, `hashlib`/`hmac`/`secrets`/`ssl`, and `truststore`/`ssl`. The standard library is not covered either. It is not a runner problem. Every `.py` file of every entry was scanned (the scanner's file count equals the fixture's), and the same run finds 38 assets in python-cryptography.

**pyjwt, pyopenssl and python-jose score only on pyca/cryptography calls.** Those fixtures generate their keys with pyca/cryptography, and the key generation lines are the only matches. The library's own calls (`jwt.encode`, `OpenSSL.SSL.Context`, `jose.jwt.encode`) produce nothing.

**All 14 FP are key material at key generation sites.** For every generated key, sonar-cryptography writes a `related-crypto-material` component (`private-key@<uuid>` or `secret-key@<uuid>`) at the generation line, next to the algorithm. The ground truth records the algorithm on those lines and no key material, so the key components count as FP. Under the old tolerance 4 of the 15 key components matched assets one or two lines away (two nonces, two key serializations); under strict matching none does. Whether a generated key is an asset of its own is an annotation guide question, not a detection error.

**python-cryptography misses come from where the tool reports.** sonar-cryptography reports a crypto object once, at the call that creates it, and attaches the operations to that object. The ground truth also annotates each later operation call, such as `encrypt` and `decrypt` on an AEAD or RSA-OAEP object, `verify` after `sign`, and each X25519 exchange. Many of its FN are such calls, which is what the object-level view corrects for. The three certificates in `x509-cert/main.py` are also missed: no certificate component is written. Of 24 TP, 4 are not attributed. For example, RSA-PSS signing is reported as `SHA-256` and the RSA-PSS key generation as `RSA-PKCS1-1.5-SHA-256`. The 10 duplicates are extra components on a line that already matched, such as `SHA-256` beside `HMAC-SHA-256` or `SHA-512` beside `Ed25519`.

There is no reachability: CBOMs carry none, so every prediction answers `unknown`.
