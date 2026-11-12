# Strict matching: the pairs that matched only through the line tolerance

**Update 2026-11-05.** Class A is resolved by multi-anchor assets ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules) change log): an `algorithm` asset may carry the call that names its algorithm for the same object in the same function as an anchor, and a prediction on it matches. 49 of the 57 class A pairs below match again through an anchor; the table is kept as it was, the record of the change to `LINE_TOLERANCE = 0`.

On 2026-11-04 the matching rules changed from `LINE_TOLERANCE = 2` to `LINE_TOLERANCE = 0` ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules), change log). Before the change, 111 matched pairs across the published pilots had line ranges that did not overlap: the prediction was 1 or 2 lines away from the ground-truth asset. This page lists them, from `internal/score.MatchEntry` on the stored predictions under the old rule, filtered to pairs whose unwidened ranges do not overlap.

Each pair was read against the program source and put in one of these classes:

- **Same operation on a neighbouring line.** The tool reports the asset's own expression at another line, such as the first line of a multi-line chained call. The fix for this class is the ground-truth range, not a tolerance. No pair was in this class.
- **A. Setup call.** The call that names the algorithm or builds the generator, 1 or 2 lines before the call that performs the operation: `KeyGenerator.getInstance("AES")` before `generateKey()`, `jwt.NewWithClaims(jwt.SigningMethodHS256, ...)` before `SignedString`, `UnparsedPublicKey::new(&ED25519, ...)` before `verify`, `SystemRandom::new()` before `fill`, a JWS header before `sign`. 57 pairs.
- **B. Material where it is used or declared.** An IV, nonce or salt reported where the program passes it on (`modes.CBC(iv)`, `new IvParameterSpec(iv)`) or where it declares the buffer (`let mut nonce = [0u8; 12];`), not where the draw creates it, which is where the [annotation guide](annotation-guide.md#related-crypto-material) puts it. 9 pairs.
- **C. Another call or asset.** A different operation near the asset: a key generation next to the key's serialization, a sign next to a verify, a secret key next to a nonce, a digest argument next to a random draw. 45 pairs.

The separate check of every ground-truth range against its source found no single-line range for a multi-line expression: every range that a bracket count flags ends at the call's closing parenthesis, with a follow-up call (`.map_err`, `.getPayload()`) or an `if` block after it. So no ground-truth range changed.

| Tool | Class A | Class B | Class C | Total |
|---|---:|---:|---:|---:|
| crypto-finder | 31 | 2 | 20 | 53 |
| sonar-cryptography | 25 | 1 | 15 | 41 |
| cdxgen | 1 | 6 | 10 | 17 |
| all | 57 | 9 | 45 | 111 |

By pilot: crypto-finder Python 4, JavaScript 5, Rust 14, Go 6, Java 24; sonar-cryptography Python 9, Go 4, Java 28; cdxgen Rust 8, Go 9, and none in Python, JavaScript or Java.

Under the new rule each of these pairs is an FN and an FP, unless another prediction on the asset's own lines takes the asset.

| Tool | Entry | File | Ground truth (lines: asset) | Prediction (lines: name) | Class |
|---|---|---|---|---|---|
| crypto-finder | python-cryptography | `rsa-oaep/main.py` | 9: algorithm RSA-OAEP | 7: RSA-OAEP | A |
| crypto-finder | python-cryptography | `aes-cbc/main.py` | 9: related-crypto-material initialization-vector | 11: initialization-vector | B |
| crypto-finder | python-ecdsa | `key-load/main.py` | 8: related-crypto-material key encoding | 7: private-key | C |
| crypto-finder | python-ecdsa | `key-load/main.py` | 15: related-crypto-material key encoding | 14: signature | C |
| crypto-finder | javascript-jose | `jwk/main.js` | 9: related-crypto-material key encoding | 10: digest | C |
| crypto-finder | javascript-jose | `jwk/main.js` | 10: algorithm SHA-256 | 9: RSASSA-PSS-SHA-256 | C |
| crypto-finder | javascript-node-forge | `hmac-sha256/main.js` | 6: algorithm HMAC | 5: HMAC | A |
| crypto-finder | javascript-tweetnacl | `box-before/main.js` | 6: related-crypto-material shared-secret | 5: key | C |
| crypto-finder | javascript-tweetnacl | `box/main.js` | 6: related-crypto-material nonce | 5: key | C |
| crypto-finder | rust-curve25519-dalek | `src/bin/x25519_dh.rs` | 12: algorithm X25519 | 10: Curve25519 | C |
| crypto-finder | rust-curve25519-dalek | `src/bin/ristretto_dh.rs` | 14: algorithm ECDH | 12: Curve25519 | C |
| crypto-finder | rust-ring | `src/bin/ed25519.rs` | 13: algorithm Ed25519 | 12: Ed25519 | A |
| crypto-finder | rust-ring | `src/bin/ecdsa_p256.rs` | 14: algorithm ECDSA | 13: ECDSA | A |
| crypto-finder | rust-ring | `src/bin/rsa_pkcs1.rs` | 14: algorithm RSA | 13: RSA-PKCS1-v1.5 | A |
| crypto-finder | rust-ring | `src/bin/pbkdf2.rs` | 8: algorithm random draw | 6: CSPRNG | A |
| crypto-finder | rust-ring | `src/bin/aead_aes256gcm.rs` | 7: algorithm random draw | 5: CSPRNG | A |
| crypto-finder | rust-ring | `src/bin/x25519.rs` | 12: algorithm X25519 | 10: X25519 | A |
| crypto-finder | rust-ring | `src/bin/random.rs` | 6: algorithm random draw | 4: CSPRNG | A |
| crypto-finder | rust-rustls | `src/bin/handshake.rs` | 9: certificate key encoding | 10: (no name) | C |
| crypto-finder | rust-rustls | `src/bin/handshake.rs` | 16-17: certificate key encoding | 19-21: (no name) | C |
| crypto-finder | rust-tokio-rustls | `src/bin/connector.rs` | 9-11: protocol key encoding | 12: (no name) | C |
| crypto-finder | rust-tokio-rustls | `src/bin/handshake.rs` | 14-16: protocol key encoding | 17: (no name) | C |
| crypto-finder | rust-tokio-rustls | `src/bin/handshake.rs` | 21-23: protocol key encoding | 24: (no name) | C |
| crypto-finder | go-dgrijalva-jwt-go | `hs256/main.go` | 13: algorithm HMAC | 12: HMAC-SHA-256 | A |
| crypto-finder | go-dgrijalva-jwt-go | `rs256/main.go` | 25: algorithm RSA | 24: RSASSA-PKCS1-SHA256 | A |
| crypto-finder | go-golang-jwt-v5 | `hs256/main.go` | 13: algorithm HMAC | 12: HMAC-SHA-256 | A |
| crypto-finder | go-golang-jwt-v5 | `rs256/main.go` | 27: algorithm RSA | 26: RSASSA-PKCS1-SHA256 | A |
| crypto-finder | go-golang-jwt-v5 | `es256/main.go` | 18: algorithm ECDSA | 17: ECDSA-SHA256 | A |
| crypto-finder | go-golang-jwt-v5 | `eddsa/main.go` | 17: algorithm Ed25519 | 16: Ed25519 | A |
| crypto-finder | java-stdlib | `src/main/java/bench/AesCbc.java` | 17: related-crypto-material initialization-vector | 18: initialization-vector | B |
| crypto-finder | java-stdlib | `src/main/java/bench/HmacSha256.java` | 15: algorithm HMAC | 14: HMAC-SHA256 | A |
| crypto-finder | java-stdlib | `src/main/java/bench/Pbkdf2.java` | 17: algorithm PBKDF2 | 16: PBKDF2 | A |
| crypto-finder | java-stdlib | `src/main/java/bench/AesGcm.java` | 14: algorithm AES | 16: (no name) | C |
| crypto-finder | java-stdlib | `src/main/java/bench/RandomBytes.java` | 10: algorithm random draw | 8: (no name) | A |
| crypto-finder | java-bcprov-jdk15on | `src/main/java/bench/Ed25519Sign.java` | 14: algorithm Ed25519 | 13: (no name) | A |
| crypto-finder | java-bcprov-jdk15on | `src/main/java/bench/RsaOaep.java` | 17: algorithm RSA | 16: (no name) | A |
| crypto-finder | java-bcprov-jdk15on | `src/main/java/bench/AesCbc.java` | 16: algorithm random draw | 14: (no name) | A |
| crypto-finder | java-bcprov-jdk15on | `src/main/java/bench/AesGcm.java` | 14: algorithm random draw | 12: (no name) | A |
| crypto-finder | java-bcpg-jdk15on | `src/main/java/bench/PgpEncrypt.java` | 46: algorithm RSA | 44: (no name) | C |
| crypto-finder | java-bcpg-jdk15on | `src/main/java/bench/PgpPassword.java` | 40: algorithm OpenPGP-S2K | 38: (no name) | C |
| crypto-finder | java-jsch | `src/main/java/bench/SessionConfig.java` | 9: protocol key encoding | 8: (no name) | A |
| crypto-finder | java-jsch | `src/main/java/bench/AddIdentity.java` | 12: related-crypto-material key encoding | 10: private-key | C |
| crypto-finder | java-jsch | `src/main/java/bench/GenerateKeyPair.java` | 12: related-crypto-material key encoding | 10: private-key | C |
| crypto-finder | java-nimbus-jose-jwt | `src/main/java/bench/JwsHs256.java` | 21: algorithm HMAC | 20: HMAC-SHA256 | A |
| crypto-finder | java-nimbus-jose-jwt | `src/main/java/bench/JweDirect.java` | 22: algorithm AES | 20: AES-GCM-256 | A |
| crypto-finder | java-nimbus-jose-jwt | `src/main/java/bench/JwsEs256.java` | 26: algorithm ECDSA | 24: ECDSA-SHA256 | A |
| crypto-finder | java-nimbus-jose-jwt | `src/main/java/bench/JwsHs256.java` | 17: algorithm random draw | 15: (no name) | A |
| crypto-finder | java-nimbus-jose-jwt | `src/main/java/bench/JwtRs256.java` | 31: algorithm RSA | 29: RSASSA-PKCS1-SHA256 | A |
| crypto-finder | java-jose4j | `src/main/java/bench/JwtEs256.java` | 31: algorithm ECDSA | 30: ECDSA-P256-SHA256 | A |
| crypto-finder | java-jose4j | `src/main/java/bench/JweRsaOaep.java` | 20: algorithm RSA-OAEP | 18: AES-256-GCM | C |
| crypto-finder | java-jose4j | `src/main/java/bench/JwkJson.java` | 15: related-crypto-material key encoding | 13: key | C |
| crypto-finder | java-jose4j | `src/main/java/bench/JwsHs256.java` | 21: algorithm HMAC | 19: HMAC-SHA256 | A |
| crypto-finder | java-jose4j | `src/main/java/bench/JwtEs256.java` | 42: algorithm ECDSA | 40: ECDSA-P256-SHA256 | A |
| sonar-cryptography | python-cryptography | `chacha20-poly1305/main.py` | 8: algorithm ChaCha20-Poly1305 | 7: Poly1305 | C |
| sonar-cryptography | python-cryptography | `ecdsa-p256/main.py` | 9: algorithm ECDSA | 8: SHA-256 | C |
| sonar-cryptography | python-cryptography | `fernet/main.py` | 6: algorithm Fernet | 5: Fernet | C |
| sonar-cryptography | python-cryptography | `key-load/main.py` | 8-12: related-crypto-material key encoding | 7: private-key | C |
| sonar-cryptography | python-cryptography | `ed25519/main.py` | 7: algorithm Ed25519 | 5: SHA-512 | C |
| sonar-cryptography | python-cryptography | `hmac-sha256/main.py` | 7: algorithm random draw | 9: SHA-256 | C |
| sonar-cryptography | python-cryptography | `aes-gcm/main.py` | 9: related-crypto-material nonce | 7: secret-key | C |
| sonar-cryptography | python-cryptography | `chacha20-poly1305/main.py` | 9: related-crypto-material nonce | 7: secret-key | C |
| sonar-cryptography | python-python-jose | `jwt-rs256/main.py` | 8-12: related-crypto-material key encoding | 7: private-key | C |
| sonar-cryptography | go-stdlib | `ecdh-x25519/main.go` | 22: algorithm X25519 | 21: ECDH | A |
| sonar-cryptography | go-stdlib | `pbkdf2/main.go` | 13: algorithm random draw | 14: SHA-256 | C |
| sonar-cryptography | go-x-crypto-pbkdf2 | `derive/main.go` | 14: algorithm random draw | 15: SHA-256 | C |
| sonar-cryptography | go-x-crypto-hkdf | `derive/main.go` | 16: algorithm random draw | 17: SHA-256 | C |
| sonar-cryptography | java-stdlib | `src/main/java/bench/AesCbc.java` | 17: related-crypto-material initialization-vector | 18: iv | B |
| sonar-cryptography | java-stdlib | `src/main/java/bench/EcdsaP256.java` | 13: algorithm EC | 12: EC-secp256r1 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/Ed25519Signature.java` | 11: algorithm Ed25519 | 10: Ed25519 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/HmacSha256.java` | 15: algorithm HMAC | 14: HMAC-SHA-256 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/KeySerialization.java` | 16: algorithm EC | 15: EC-secp256r1 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/KeySerialization.java` | 22: related-crypto-material key encoding | 21: private-key | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/X25519Agreement.java` | 11: algorithm X25519 | 10: x25519 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/AesCbc.java` | 14: algorithm AES | 12: AES-256 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/AesGcm.java` | 14: algorithm AES | 12: AES-256 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/ChaCha20Poly1305.java` | 14: algorithm ChaCha20 | 12: ChaCha20 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/ChaCha20Poly1305.java` | 17: algorithm random draw | 19: Poly1305 | C |
| sonar-cryptography | java-stdlib | `src/main/java/bench/Pbkdf2.java` | 14: related-crypto-material salt | 16: password | C |
| sonar-cryptography | java-stdlib | `src/main/java/bench/RsaOaep.java` | 12: algorithm RSA | 10: RSA-3072 | A |
| sonar-cryptography | java-stdlib | `src/main/java/bench/RsaSignature.java` | 12: algorithm RSA | 10: RSA-3072 | A |
| sonar-cryptography | java-bcprov-jdk15on | `src/main/java/bench/Pbkdf2.java` | 15: algorithm random draw | 17: SHA-256 | C |
| sonar-cryptography | java-bcprov-jdk15on | `src/main/java/bench/ProviderAesGcm.java` | 18: algorithm AES | 16: AES-256 | A |
| sonar-cryptography | java-bcprov-jdk15on | `src/main/java/bench/RsaOaep.java` | 17: algorithm RSA | 15: RSA-3072 | A |
| sonar-cryptography | java-bcpkix-jdk15on | `src/main/java/bench/Csr.java` | 21: algorithm EC | 20: EC-secp256r1 | A |
| sonar-cryptography | java-bcpkix-jdk15on | `src/main/java/bench/PemRoundTrip.java` | 30: algorithm EC | 29: EC-secp256r1 | A |
| sonar-cryptography | java-bcpkix-jdk15on | `src/main/java/bench/CmsSignedData.java` | 31: algorithm RSA | 29: RSA-3072 | A |
| sonar-cryptography | java-bcpkix-jdk15on | `src/main/java/bench/SelfSignedCert.java` | 25: algorithm RSA | 23: RSA-3072 | A |
| sonar-cryptography | java-bcpg-jdk15on | `src/main/java/bench/PgpEncrypt.java` | 32: algorithm RSA | 30: RSA-3072 | A |
| sonar-cryptography | java-bcpg-jdk15on | `src/main/java/bench/PgpKeyRing.java` | 30: algorithm RSA | 28: RSA-3072 | A |
| sonar-cryptography | java-bcpg-jdk15on | `src/main/java/bench/PgpSign.java` | 25: algorithm RSA | 23: RSA-3072 | A |
| sonar-cryptography | java-nimbus-jose-jwt | `src/main/java/bench/JweDirect.java` | 17: algorithm AES | 15: AES-256 | A |
| sonar-cryptography | java-xmlsec | `src/main/java/bench/XmlEncryption.java` | 27: algorithm AES | 25: AES-256 | A |
| sonar-cryptography | java-xmlsec | `src/main/java/bench/XmlEncryption.java` | 31: algorithm RSA | 29: RSA-3072 | A |
| sonar-cryptography | java-xmlsec | `src/main/java/bench/XmlSignature.java` | 24: algorithm RSA | 22: RSA-3072 | A |
| cdxgen | rust-chacha20 | `src/bin/chacha20.rs` | 8: related-crypto-material nonce | 7: nonce | B |
| cdxgen | rust-curve25519-dalek | `src/bin/ristretto_dh.rs` | 11: related-crypto-material key encoding | 9: bob_secret | C |
| cdxgen | rust-openssl | `src/bin/aes_cbc.rs` | 9: related-crypto-material initialization-vector | 8: iv | B |
| cdxgen | rust-openssl | `src/bin/aes_gcm.rs` | 8: related-crypto-material nonce | 7: nonce | B |
| cdxgen | rust-openssl | `src/bin/pbkdf2.rs` | 7: related-crypto-material salt | 6: salt | B |
| cdxgen | rust-ring | `src/bin/pbkdf2.rs` | 8: related-crypto-material salt | 7: salt | B |
| cdxgen | rust-ring | `src/bin/aead_aes256gcm.rs` | 11: related-crypto-material nonce | 10: nonce_bytes | B |
| cdxgen | rust-ring | `src/bin/x25519.rs` | 12: related-crypto-material shared-secret | 10: peer_public_key | C |
| cdxgen | go-stdlib | `ecdh-x25519/main.go` | 22: algorithm X25519 | 21: curveX25519 | A |
| cdxgen | go-stdlib | `pbkdf2/main.go` | 13: related-crypto-material salt | 14: key | C |
| cdxgen | go-dgrijalva-jwt-go | `rs256/main.go` | 19: related-crypto-material key encoding | 18: public-key | C |
| cdxgen | go-golang-jwt-v5 | `rs256/main.go` | 20: related-crypto-material key encoding | 19: public-key | C |
| cdxgen | go-x-crypto-pbkdf2 | `derive/main.go` | 14: algorithm random draw | 15: sha-256 | C |
| cdxgen | go-x-crypto-pbkdf2 | `derive/main.go` | 14: related-crypto-material salt | 15: key | C |
| cdxgen | go-x-crypto-pbkdf2 | `wpa-psk/main.go` | 13: related-crypto-material salt | 14: key | C |
| cdxgen | go-x-crypto-scrypt | `derive/main.go` | 13: related-crypto-material salt | 11: password | C |
| cdxgen | go-x-crypto-acme-autocert | `manager/main.go` | 29: protocol key encoding | 31: TLS | C |
