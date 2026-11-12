# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (16): go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519

Not run, no prediction file (153): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 0.533 | 0.243 | 0.333 | 49 | 43 | 153 | 3 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| A | 0.533 | 0.243 | 0.333 | 49 | 43 | 153 | 3 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go | 0.533 | 0.243 | 0.333 | 49 | 43 | 153 | 3 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.188 | 0.316 | 25 | 0 | 108 | 1 |
| certificate | 0.000 | 0.000 | 0.000 | 0 | 3 | 4 | 0 |
| protocol | 0.500 | 0.100 | 0.167 | 1 | 1 | 9 | 0 |
| related-crypto-material | 0.371 | 0.418 | 0.393 | 23 | 39 | 32 | 2 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm | 1.000 | 0.188 | 0.316 | 25 | 0 | 108 | 1 |
| go / certificate | 0.000 | 0.000 | 0.000 | 0 | 3 | 4 | 0 |
| go / protocol | 0.500 | 0.100 | 0.167 | 1 | 1 | 9 | 0 |
| go / related-crypto-material | 0.371 | 0.418 | 0.393 | 23 | 39 | 32 | 2 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-dgrijalva-jwt-go | 0.429 | 0.375 | 0.400 | 3 | 4 | 5 | 1 |
| go-go-ethereum-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 13 | 0 |
| go-golang-jwt-v5 | 0.500 | 0.312 | 0.385 | 5 | 5 | 11 | 1 |
| go-stdlib | 0.545 | 0.316 | 0.400 | 30 | 25 | 65 | 0 |
| go-x-crypto-acme-autocert | 0.667 | 0.667 | 0.667 | 2 | 1 | 1 | 0 |
| go-x-crypto-bcrypt | 0.000 | 0.000 | 0.000 | 0 | 1 | 2 | 0 |
| go-x-crypto-blake2b | n/a | 0.000 | 0.000 | 0 | 0 | 5 | 0 |
| go-x-crypto-chacha20poly1305 | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 |
| go-x-crypto-curve25519 | 0.600 | 0.375 | 0.462 | 3 | 2 | 5 | 0 |
| go-x-crypto-hkdf | 1.000 | 0.111 | 0.200 | 1 | 0 | 8 | 0 |
| go-x-crypto-pbkdf2 | 0.400 | 0.400 | 0.400 | 2 | 3 | 3 | 1 |
| go-x-crypto-ripemd160 | 0.500 | 0.500 | 0.500 | 1 | 1 | 1 | 0 |
| go-x-crypto-scrypt | 0.000 | 0.000 | 0.000 | 0 | 1 | 3 | 0 |
| go-x-crypto-sha3 | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| go-x-crypto-ssh | 1.000 | 0.118 | 0.211 | 2 | 0 | 15 | 0 |
| go-x-crypto-ssh-agent | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 35 / 49 (0.714) | 0.380 | 0.173 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| A | 35 / 49 (0.714) | 0.380 | 0.173 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go | 35 / 49 (0.714) | 0.380 | 0.173 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 11 / 25 (0.440) | 0.440 | 0.083 |
| certificate | 0 / 0 (n/a) | 0.000 | 0.000 |
| protocol | 1 / 1 (1.000) | 0.500 | 0.100 |
| related-crypto-material | 23 / 23 (1.000) | 0.371 | 0.418 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-dgrijalva-jwt-go | 2 / 3 (0.667) | 0.286 | 0.250 |
| go-go-ethereum-crypto | 0 / 0 (n/a) | n/a | 0.000 |
| go-golang-jwt-v5 | 3 / 5 (0.600) | 0.300 | 0.188 |
| go-stdlib | 24 / 30 (0.800) | 0.436 | 0.253 |
| go-x-crypto-acme-autocert | 1 / 2 (0.500) | 0.333 | 0.333 |
| go-x-crypto-bcrypt | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-x-crypto-blake2b | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-chacha20poly1305 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-curve25519 | 2 / 3 (0.667) | 0.400 | 0.250 |
| go-x-crypto-hkdf | 0 / 1 (0.000) | 0.000 | 0.000 |
| go-x-crypto-pbkdf2 | 1 / 2 (0.500) | 0.200 | 0.200 |
| go-x-crypto-ripemd160 | 1 / 1 (1.000) | 0.500 | 0.500 |
| go-x-crypto-scrypt | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-x-crypto-sha3 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-ssh | 1 / 2 (0.500) | 0.500 | 0.059 |
| go-x-crypto-ssh-agent | 0 / 0 (n/a) | n/a | 0.000 |

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
| all | 0.506 | 0.289 | 0.368 | 44 | 43 | 108 | 8 | 31 / 44 (0.705) | 0.356 | 0.204 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go | 0.506 | 0.289 | 0.368 | 44 | 43 | 108 | 8 | 31 / 44 (0.705) | 0.356 | 0.204 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.258 | 0.410 | 24 | 0 | 69 | 2 | 11 / 24 (0.458) | 0.458 | 0.118 |
| certificate | 0.000 | 0.000 | 0.000 | 0 | 3 | 2 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| protocol | 0.500 | 0.100 | 0.167 | 1 | 1 | 9 | 0 | 1 / 1 (1.000) | 0.500 | 0.100 |
| related-crypto-material | 0.328 | 0.404 | 0.362 | 19 | 39 | 28 | 6 | 19 / 19 (1.000) | 0.328 | 0.404 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-dgrijalva-jwt-go | 0.429 | 0.750 | 0.545 | 3 | 4 | 1 | 1 | 2 / 3 (0.667) | 0.286 | 0.500 |
| go-go-ethereum-crypto | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-golang-jwt-v5 | 0.500 | 0.625 | 0.556 | 5 | 5 | 3 | 1 | 3 / 5 (0.600) | 0.300 | 0.375 |
| go-stdlib | 0.500 | 0.357 | 0.417 | 25 | 25 | 45 | 5 | 20 / 25 (0.800) | 0.400 | 0.286 |
| go-x-crypto-acme-autocert | 0.667 | 0.667 | 0.667 | 2 | 1 | 1 | 0 | 1 / 2 (0.500) | 0.333 | 0.333 |
| go-x-crypto-bcrypt | 0.000 | 0.000 | 0.000 | 0 | 1 | 2 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-x-crypto-blake2b | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-chacha20poly1305 | n/a | 0.000 | 0.000 | 0 | 0 | 8 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-curve25519 | 0.600 | 0.500 | 0.545 | 3 | 2 | 3 | 0 | 2 / 3 (0.667) | 0.400 | 0.333 |
| go-x-crypto-hkdf | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 0 | 0 / 1 (0.000) | 0.000 | 0.000 |
| go-x-crypto-pbkdf2 | 0.400 | 0.400 | 0.400 | 2 | 3 | 3 | 1 | 1 / 2 (0.500) | 0.200 | 0.200 |
| go-x-crypto-ripemd160 | 0.500 | 0.500 | 0.500 | 1 | 1 | 1 | 0 | 1 / 1 (1.000) | 0.500 | 0.500 |
| go-x-crypto-scrypt | 0.000 | 0.000 | 0.000 | 0 | 1 | 3 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-x-crypto-sha3 | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-x-crypto-ssh | 1.000 | 0.154 | 0.267 | 2 | 0 | 11 | 0 | 1 / 2 (0.500) | 0.500 | 0.077 |
| go-x-crypto-ssh-agent | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
