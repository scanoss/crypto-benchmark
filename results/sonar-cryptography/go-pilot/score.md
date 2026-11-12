# Score: sonar-cryptography

- Tool: sonar-cryptography 1.7.0
- Config: (none)

Entries scored (16): go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519

Not run, no prediction file (153): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 1.000 | 0.243 | 0.390 | 49 | 0 | 153 | 20 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| A | 1.000 | 0.243 | 0.390 | 49 | 0 | 153 | 20 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go | 1.000 | 0.243 | 0.390 | 49 | 0 | 153 | 20 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.368 | 0.538 | 49 | 0 | 84 | 20 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 |
| related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 55 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm | 1.000 | 0.368 | 0.538 | 49 | 0 | 84 | 20 |
| go / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| go / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 |
| go / related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 55 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-dgrijalva-jwt-go | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 0 |
| go-go-ethereum-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 13 | 0 |
| go-golang-jwt-v5 | 1.000 | 0.188 | 0.316 | 3 | 0 | 13 | 3 |
| go-stdlib | 1.000 | 0.358 | 0.527 | 34 | 0 | 61 | 7 |
| go-x-crypto-acme-autocert | 1.000 | 0.333 | 0.500 | 1 | 0 | 2 | 1 |
| go-x-crypto-bcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| go-x-crypto-blake2b | n/a | 0.000 | 0.000 | 0 | 0 | 5 | 0 |
| go-x-crypto-chacha20poly1305 | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 |
| go-x-crypto-curve25519 | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 |
| go-x-crypto-hkdf | 1.000 | 0.111 | 0.200 | 1 | 0 | 8 | 1 |
| go-x-crypto-pbkdf2 | 1.000 | 0.400 | 0.571 | 2 | 0 | 3 | 2 |
| go-x-crypto-ripemd160 | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |
| go-x-crypto-scrypt | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| go-x-crypto-sha3 | 1.000 | 0.500 | 0.667 | 2 | 0 | 2 | 0 |
| go-x-crypto-ssh | 1.000 | 0.176 | 0.300 | 3 | 0 | 14 | 4 |
| go-x-crypto-ssh-agent | 1.000 | 0.250 | 0.400 | 1 | 0 | 3 | 2 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 40 / 49 (0.816) | 0.816 | 0.198 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| A | 40 / 49 (0.816) | 0.816 | 0.198 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go | 40 / 49 (0.816) | 0.816 | 0.198 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 40 / 49 (0.816) | 0.816 | 0.301 |
| certificate | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-dgrijalva-jwt-go | 1 / 1 (1.000) | 1.000 | 0.125 |
| go-go-ethereum-crypto | 0 / 0 (n/a) | n/a | 0.000 |
| go-golang-jwt-v5 | 3 / 3 (1.000) | 1.000 | 0.188 |
| go-stdlib | 26 / 34 (0.765) | 0.765 | 0.274 |
| go-x-crypto-acme-autocert | 1 / 1 (1.000) | 1.000 | 0.333 |
| go-x-crypto-bcrypt | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-blake2b | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-chacha20poly1305 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-curve25519 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-hkdf | 0 / 1 (0.000) | 0.000 | 0.000 |
| go-x-crypto-pbkdf2 | 2 / 2 (1.000) | 1.000 | 0.400 |
| go-x-crypto-ripemd160 | 1 / 1 (1.000) | 1.000 | 0.500 |
| go-x-crypto-scrypt | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-sha3 | 2 / 2 (1.000) | 1.000 | 0.500 |
| go-x-crypto-ssh | 3 / 3 (1.000) | 1.000 | 0.176 |
| go-x-crypto-ssh-agent | 1 / 1 (1.000) | 1.000 | 0.250 |

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
| all | 1.000 | 0.276 | 0.433 | 42 | 0 | 110 | 27 | 40 / 42 (0.952) | 0.952 | 0.263 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go | 1.000 | 0.276 | 0.433 | 42 | 0 | 110 | 27 | 40 / 42 (0.952) | 0.952 | 0.263 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.452 | 0.622 | 42 | 0 | 51 | 27 | 40 / 42 (0.952) | 0.952 | 0.430 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 47 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-dgrijalva-jwt-go | 1.000 | 0.250 | 0.400 | 1 | 0 | 3 | 0 | 1 / 1 (1.000) | 1.000 | 0.250 |
| go-go-ethereum-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-golang-jwt-v5 | 1.000 | 0.375 | 0.545 | 3 | 0 | 5 | 3 | 3 / 3 (1.000) | 1.000 | 0.375 |
| go-stdlib | 1.000 | 0.386 | 0.557 | 27 | 0 | 43 | 14 | 26 / 27 (0.963) | 0.963 | 0.371 |
| go-x-crypto-acme-autocert | 1.000 | 0.333 | 0.500 | 1 | 0 | 2 | 1 | 1 / 1 (1.000) | 1.000 | 0.333 |
| go-x-crypto-bcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-blake2b | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-chacha20poly1305 | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-curve25519 | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-hkdf | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 1 | 0 / 1 (0.000) | 0.000 | 0.000 |
| go-x-crypto-pbkdf2 | 1.000 | 0.400 | 0.571 | 2 | 0 | 3 | 2 | 2 / 2 (1.000) | 1.000 | 0.400 |
| go-x-crypto-ripemd160 | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 | 1 / 1 (1.000) | 1.000 | 0.500 |
| go-x-crypto-scrypt | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-sha3 | 1.000 | 0.500 | 0.667 | 2 | 0 | 2 | 0 | 2 / 2 (1.000) | 1.000 | 0.500 |
| go-x-crypto-ssh | 1.000 | 0.231 | 0.375 | 3 | 0 | 10 | 4 | 3 / 3 (1.000) | 1.000 | 0.231 |
| go-x-crypto-ssh-agent | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 2 | 1 / 1 (1.000) | 1.000 | 0.500 |
