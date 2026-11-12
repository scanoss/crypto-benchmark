# Score: sonar-cryptography

- Tool: sonar-cryptography 1.7.0
- Config: (none)

Entries scored (16): java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec

Not run, no prediction file (153): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 0.623 | 0.189 | 0.290 | 66 | 40 | 283 | 34 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| A | 0.623 | 0.189 | 0.290 | 66 | 40 | 283 | 34 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| java | 0.623 | 0.189 | 0.290 | 66 | 40 | 283 | 34 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 0.970 | 0.260 | 0.410 | 64 | 2 | 182 | 34 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 |
| protocol | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |
| related-crypto-material | 0.026 | 0.011 | 0.015 | 1 | 38 | 90 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| java / algorithm | 0.970 | 0.260 | 0.410 | 64 | 2 | 182 | 34 |
| java / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 |
| java / protocol | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |
| java / related-crypto-material | 0.026 | 0.011 | 0.015 | 1 | 38 | 90 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| java-bcpg-jdk15on | 0.500 | 0.125 | 0.200 | 3 | 3 | 21 | 0 |
| java-bcpkix-jdk15on | 0.400 | 0.143 | 0.211 | 4 | 6 | 24 | 2 |
| java-bcprov-jdk15on | 0.864 | 0.365 | 0.514 | 19 | 3 | 33 | 12 |
| java-bcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 |
| java-cryptography-core-jvm | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 |
| java-eddsa | 1.000 | 0.133 | 0.235 | 2 | 0 | 13 | 0 |
| java-jasypt | n/a | 0.000 | 0.000 | 0 | 0 | 9 | 0 |
| java-jbcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 5 | 0 |
| java-jjwt-api | n/a | 0.000 | 0.000 | 0 | 0 | 24 | 0 |
| java-jose4j | n/a | 0.000 | 0.000 | 0 | 0 | 26 | 0 |
| java-jsch | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 |
| java-nimbus-jose-jwt | 0.500 | 0.023 | 0.044 | 1 | 1 | 42 | 0 |
| java-scrypt | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| java-spring-security-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| java-stdlib | 0.586 | 0.557 | 0.571 | 34 | 24 | 27 | 20 |
| java-xmlsec | 0.500 | 0.250 | 0.333 | 3 | 3 | 9 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 57 / 66 (0.864) | 0.538 | 0.163 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| A | 57 / 66 (0.864) | 0.538 | 0.163 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| java | 57 / 66 (0.864) | 0.538 | 0.163 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 55 / 64 (0.859) | 0.833 | 0.224 |
| certificate | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 1 / 1 (1.000) | 1.000 | 0.500 |
| related-crypto-material | 1 / 1 (1.000) | 0.026 | 0.011 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| java-bcpg-jdk15on | 3 / 3 (1.000) | 0.500 | 0.125 |
| java-bcpkix-jdk15on | 4 / 4 (1.000) | 0.400 | 0.143 |
| java-bcprov-jdk15on | 16 / 19 (0.842) | 0.727 | 0.308 |
| java-bcrypt | 0 / 0 (n/a) | n/a | 0.000 |
| java-cryptography-core-jvm | 0 / 0 (n/a) | n/a | 0.000 |
| java-eddsa | 0 / 2 (0.000) | 0.000 | 0.000 |
| java-jasypt | 0 / 0 (n/a) | n/a | 0.000 |
| java-jbcrypt | 0 / 0 (n/a) | n/a | 0.000 |
| java-jjwt-api | 0 / 0 (n/a) | n/a | 0.000 |
| java-jose4j | 0 / 0 (n/a) | n/a | 0.000 |
| java-jsch | 0 / 0 (n/a) | n/a | 0.000 |
| java-nimbus-jose-jwt | 1 / 1 (1.000) | 0.500 | 0.023 |
| java-scrypt | 0 / 0 (n/a) | n/a | 0.000 |
| java-spring-security-crypto | 0 / 0 (n/a) | n/a | 0.000 |
| java-stdlib | 30 / 34 (0.882) | 0.517 | 0.492 |
| java-xmlsec | 3 / 3 (1.000) | 0.500 | 0.250 |

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
| all | 0.500 | 0.181 | 0.266 | 40 | 40 | 181 | 60 | 37 / 40 (0.925) | 0.463 | 0.167 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| java | 0.500 | 0.181 | 0.266 | 40 | 40 | 181 | 60 | 37 / 40 (0.925) | 0.463 | 0.167 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 0.950 | 0.270 | 0.420 | 38 | 2 | 103 | 60 | 35 / 38 (0.921) | 0.875 | 0.248 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 5 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 | 1 / 1 (1.000) | 1.000 | 0.500 |
| related-crypto-material | 0.026 | 0.014 | 0.018 | 1 | 38 | 72 | 0 | 1 / 1 (1.000) | 0.026 | 0.014 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| java-bcpg-jdk15on | 0.500 | 0.176 | 0.261 | 3 | 3 | 14 | 0 | 3 / 3 (1.000) | 0.500 | 0.176 |
| java-bcpkix-jdk15on | 0.400 | 0.308 | 0.348 | 4 | 6 | 9 | 2 | 4 / 4 (1.000) | 0.400 | 0.308 |
| java-bcprov-jdk15on | 0.786 | 0.275 | 0.407 | 11 | 3 | 29 | 20 | 9 / 11 (0.818) | 0.643 | 0.225 |
| java-bcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-cryptography-core-jvm | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-eddsa | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 1 | 0 / 1 (0.000) | 0.000 | 0.000 |
| java-jasypt | n/a | 0.000 | 0.000 | 0 | 0 | 9 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-jbcrypt | n/a | 0.000 | 0.000 | 0 | 0 | 5 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-jjwt-api | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-jose4j | n/a | 0.000 | 0.000 | 0 | 0 | 15 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-jsch | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-nimbus-jose-jwt | 0.500 | 0.062 | 0.111 | 1 | 1 | 15 | 0 | 1 / 1 (1.000) | 0.500 | 0.062 |
| java-scrypt | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-spring-security-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-stdlib | 0.415 | 0.436 | 0.425 | 17 | 24 | 22 | 37 | 17 / 17 (1.000) | 0.415 | 0.436 |
| java-xmlsec | 0.500 | 0.500 | 0.500 | 3 | 3 | 3 | 0 | 3 / 3 (1.000) | 0.500 | 0.500 |
