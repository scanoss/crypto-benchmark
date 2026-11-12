# sonar-cryptography, suite A Java pilot

Current score against the ground truth of 2026-11-08, from the stored predictions ([Java pilot page](../../java-pilot.md) compares the tools):

| Level | Precision | Recall | F1 | TP | FP | FN |
|---|---|---|---|---|---|---|
| operation | 0.623 | 0.189 | 0.290 | 66 | 40 | 283 |
| object | 0.500 | 0.181 | 0.266 | 40 | 40 | 181 |

Run on 2026-11-04 over the same 16 Java suite A entries as [crypto-finder](../../crypto-finder/java-pilot/README.md) and [cdxgen](../../cdxgen/java-pilot/README.md) (`fixtures/java/*`, one Maven project per library in the frozen list plus JCA/JCE). `score.md` is the `cbench score` report, with the operation-level and object-level views. `predictions/sonar-cryptography/` holds the adapted predictions. `commands/` holds the exact commands that `scripts/run-sonar-cryptography.sh` recorded for each entry, with paths made relative.

## Run record

- **sonar-cryptography:** 1.7.0, embedded in CBOMkit-action v2.3.0 through cbomkit-lib 1.3.0, with the SonarQube Java frontend. The script checked the three embedded versions.
- **Image:** `ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b`, Docker 29.1.3, no network during the scan.
- **Build:** the scanner needs compiled classes, so each entry was copied to a scratch directory and built there with Maven 3.8.7 on Temurin JDK 21.0.12 (`mvn -q package dependency:copy-dependencies`), which puts `target/classes` and the dependency jars inside the copy. The repository never holds `target/`.
- **Commands,** per entry `ID` with directory `fixtures/java/LIB`:

```sh
cp -r fixtures/java/LIB work/build/LIB && (cd work/build/LIB && mvn -q package dependency:copy-dependencies)
scripts/run-sonar-cryptography.sh work/build/LIB work/sonar-cryptography/ID.cbom.json
go run ./cmd/cbench adapt cbom --tool sonar-cryptography --tool-version 1.7.0 \
  --entry ID --input work/sonar-cryptography/ID.cbom.json \
  --out results/sonar-cryptography/java-pilot/predictions/sonar-cryptography/ID.json
go run ./cmd/cbench score --tool sonar-cryptography --predictions results/sonar-cryptography/java-pilot/predictions \
  --out results/sonar-cryptography/java-pilot/score.md
```

The adapter read every CBOM without a note: no component without a location and no dropped location. The generated key names (`key@<uuid>`) change between runs; nothing else does.

## What it reports

Written after the 2026-11-04 run, against the ground truth of that day; counts in this section are of that ground truth, not of the current score above.

sonar-cryptography's Java rules cover JCA/JCE and the Bouncy Castle provider and lightweight API. It reports nothing for the other libraries, it reports a crypto object once where it is created, and it names a generated key as key material.

**104 FN are libraries it does not cover.** bcrypt, cryptography-core-jvm, jasypt, jbcrypt, jjwt-api, jose4j, jsch, scrypt and spring-security-crypto have no prediction at all. nimbus-jose-jwt gets no match: the JCA `KeyGenerator` the JWE program calls itself is reported at `KeyGenerator.getInstance`, two lines before `generateKey()`, and its 37 other assets (JWS and JWE signers, verifiers, encrypters, JWK key generation and serialization) are missed. eddsa gets two matches, the JCA `MessageDigest.getInstance("SHA-512")` nested in each `EdDSAEngine`, which pair with the Ed25519 signer assets but name SHA-512, so they do not attribute; the i2p key generation and key encoding are missed.

**In the libraries it covers, the misses are key generations reported at their setup call, draws, material, certificates and operations on objects.**

- Key generations (stdlib 10, bcprov 2, bcpkix 4, bcpg 3, xmlsec 3, nimbus-jose-jwt 1): each is reported, but at the generator's `getInstance`, `initialize` or constructor, one or two lines before the `generateKey()` or `generateKeyPair()` the guide annotates, so it no longer matches (above).
- Random draws and the material they become (stdlib 13, bcprov 17): no `SecureRandom.nextBytes` is reported. The IVs and salts it does report sit on the `GCMParameterSpec`, `IvParameterSpec` and `PBEKeySpec` lines that pass them on, which the guide does not annotate (FP, below).
- Certificates and key bytes (stdlib 6, bcpkix 11, bcprov 3, bcpg 2): no certificate parse or build, no `getEncoded`, no `KeyFactory` parse (the private key it reports at `KeyFactory.getInstance` is a line early) and no PEM read or write is reported.
- OpenPGP and XML security (bcpg 15, xmlsec 7): only the JCA key generations in these programs are reported, at their setup call. Bouncy Castle's OpenPGP operator builders and Santuario's `XMLSignature` and `XMLCipher` produce nothing.
- CMS and certificate signing in bcpkix (9): `JcaContentSignerBuilder` and the verify calls produce nothing.
- Lightweight API calls in bcprov (8): the four X25519 key agreements, the Argon2 and scrypt KDFs, and the Ed25519 and EC key pair generations. In stdlib, Bob's X25519 key generation, the second from the same generator, is missed (1).

**All 69 FP are setup calls, material or parameters the guide does not annotate.** The algorithm of a key generation at its setup call (`getInstance`, `initialize(new ECGenParameterSpec(...))` or a Bouncy Castle generator constructor), one or two lines before the generating call: 27 FP, two for each EC generator. The `SHA-512` and `SHA-256` components it nests in the Ed25519 and HmacSHA256 generators: 2 FP. For every key generation it also writes a `related-crypto-material` key component (`key@<uuid>` or `secret-key@<uuid>`) at the generator line, next to the algorithm: 27 FP. IVs and nonces at the `GCMParameterSpec` or `IvParameterSpec` that passes them: 8 FP. The password and salt components on `PBEKeySpec` and the private key at `KeyFactory.getInstance`: 3 FP. `KeyFactory.getInstance("EC")`, which computes nothing, as an EC algorithm, and `new SecretKeySpec(derived, "AES")` as an AES-128 algorithm where the ground truth has a key parse: 2 FP.

**Attribution (34 of 43).** The unattributed matches are names `internal/algo` does not split or a parameter it reports wrongly: `RSA-PKCS1-1.5-SHA-256` for `SHA256withRSA` (2), `ChaCha20` for ChaCha20-Poly1305 (2), `PBES2` for the bcprov PBKDF2 (1), `SHA-512` for the Ed25519 signer (2) and bare `ECDSA` with no digest for `ECDSASigner` (2).

Reachability: sonar-cryptography reports none, so it is n/a.
