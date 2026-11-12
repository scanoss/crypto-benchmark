# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (16): javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto

Not run, no prediction file (153): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 1.000 | 0.086 | 0.158 | 32 | 0 | 341 | 0 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| A | 1.000 | 0.086 | 0.158 | 32 | 0 | 341 | 0 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| javascript | 1.000 | 0.086 | 0.158 | 32 | 0 | 341 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.122 | 0.218 | 32 | 0 | 230 | 0 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 94 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| javascript / algorithm | 1.000 | 0.122 | 0.218 | 32 | 0 | 230 | 0 |
| javascript / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| javascript / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| javascript / related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 94 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| javascript-blake3-wasm | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| javascript-crypto-random-string | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| javascript-fast-sha256 | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 |
| javascript-hash-js | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| javascript-iron-webcrypto | n/a | 0.000 | 0.000 | 0 | 0 | 17 | 0 |
| javascript-jose | 1.000 | 0.200 | 0.333 | 6 | 0 | 24 | 0 |
| javascript-jsonwebtoken | 1.000 | 0.400 | 0.571 | 4 | 0 | 6 | 0 |
| javascript-jwa | 1.000 | 0.182 | 0.308 | 2 | 0 | 9 | 0 |
| javascript-noble-ciphers | n/a | 0.000 | 0.000 | 0 | 0 | 25 | 0 |
| javascript-noble-curves | n/a | 0.000 | 0.000 | 0 | 0 | 28 | 0 |
| javascript-noble-hashes | n/a | 0.000 | 0.000 | 0 | 0 | 24 | 0 |
| javascript-node-forge | n/a | 0.000 | 0.000 | 0 | 0 | 49 | 0 |
| javascript-peculiar-x509 | 1.000 | 0.200 | 0.333 | 3 | 0 | 12 | 0 |
| javascript-selfsigned | n/a | 0.000 | 0.000 | 0 | 0 | 7 | 0 |
| javascript-stdlib | 1.000 | 0.175 | 0.298 | 17 | 0 | 80 | 0 |
| javascript-tweetnacl | n/a | 0.000 | 0.000 | 0 | 0 | 34 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 7 / 32 (0.219) | 0.219 | 0.019 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| A | 7 / 32 (0.219) | 0.219 | 0.019 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| javascript | 7 / 32 (0.219) | 0.219 | 0.019 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 7 / 32 (0.219) | 0.219 | 0.027 |
| certificate | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| javascript-blake3-wasm | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-crypto-random-string | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-fast-sha256 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-hash-js | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-iron-webcrypto | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-jose | 3 / 6 (0.500) | 0.500 | 0.100 |
| javascript-jsonwebtoken | 0 / 4 (0.000) | 0.000 | 0.000 |
| javascript-jwa | 0 / 2 (0.000) | 0.000 | 0.000 |
| javascript-noble-ciphers | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-noble-curves | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-noble-hashes | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-node-forge | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-peculiar-x509 | 0 / 3 (0.000) | 0.000 | 0.000 |
| javascript-selfsigned | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-stdlib | 4 / 17 (0.235) | 0.235 | 0.041 |
| javascript-tweetnacl | 0 / 0 (n/a) | n/a | 0.000 |

## Reachability

n/a: no scored entry's predictions report reachability.

## Suite B pair discrimination

n/a: no scored entry's predictions report reachability.

## Suite B reachability

n/a: no suite B pair entry was scored.

## Object-level view

Each ground-truth object counts once: the assets that share an `object` id, or one asset without one. An object is detected when any of its assets is matched above, and attributed when any of its matched assets is attributed. A prediction matched to an object already detected is a duplicate; FPs are the same as above.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| all | 1.000 | 0.090 | 0.164 | 25 | 0 | 254 | 7 | 7 / 25 (0.280) | 0.280 | 0.025 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| javascript | 1.000 | 0.090 | 0.164 | 25 | 0 | 254 | 7 | 7 / 25 (0.280) | 0.280 | 0.025 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.137 | 0.242 | 25 | 0 | 157 | 7 | 7 / 25 (0.280) | 0.280 | 0.038 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 7 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 87 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| javascript-blake3-wasm | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-crypto-random-string | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-fast-sha256 | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-hash-js | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-iron-webcrypto | n/a | 0.000 | 0.000 | 0 | 0 | 17 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-jose | 1.000 | 0.400 | 0.571 | 6 | 0 | 9 | 0 | 3 / 6 (0.500) | 0.500 | 0.200 |
| javascript-jsonwebtoken | 1.000 | 0.600 | 0.750 | 3 | 0 | 2 | 1 | 0 / 3 (0.000) | 0.000 | 0.000 |
| javascript-jwa | 1.000 | 0.286 | 0.444 | 2 | 0 | 5 | 0 | 0 / 2 (0.000) | 0.000 | 0.000 |
| javascript-noble-ciphers | n/a | 0.000 | 0.000 | 0 | 0 | 22 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-noble-curves | n/a | 0.000 | 0.000 | 0 | 0 | 18 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-noble-hashes | n/a | 0.000 | 0.000 | 0 | 0 | 23 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-node-forge | n/a | 0.000 | 0.000 | 0 | 0 | 34 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-peculiar-x509 | 1.000 | 0.300 | 0.462 | 3 | 0 | 7 | 0 | 0 / 3 (0.000) | 0.000 | 0.000 |
| javascript-selfsigned | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-stdlib | 1.000 | 0.155 | 0.268 | 11 | 0 | 60 | 6 | 4 / 11 (0.364) | 0.364 | 0.056 |
| javascript-tweetnacl | n/a | 0.000 | 0.000 | 0 | 0 | 25 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
