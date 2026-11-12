# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (15): rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls

Not run, no prediction file (154): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 0.477 | 0.157 | 0.237 | 31 | 34 | 166 | 0 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| A | 0.477 | 0.157 | 0.237 | 31 | 34 | 166 | 0 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| rust | 0.477 | 0.157 | 0.237 | 31 | 34 | 166 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.036 | 0.069 | 4 | 0 | 108 | 0 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 15 | 0 |
| related-crypto-material | 0.443 | 0.482 | 0.462 | 27 | 34 | 29 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| rust / algorithm | 1.000 | 0.036 | 0.069 | 4 | 0 | 108 | 0 |
| rust / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| rust / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 15 | 0 |
| rust / related-crypto-material | 0.443 | 0.482 | 0.462 | 27 | 34 | 29 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| rust-aes | 0.000 | 0.000 | 0.000 | 0 | 2 | 4 | 0 |
| rust-chacha20 | 0.250 | 0.143 | 0.182 | 1 | 3 | 6 | 0 |
| rust-curve25519-dalek | 0.500 | 0.105 | 0.174 | 2 | 2 | 17 | 0 |
| rust-ecdsa | 0.857 | 0.375 | 0.522 | 6 | 1 | 10 | 0 |
| rust-hkdf | 0.500 | 0.500 | 0.500 | 2 | 2 | 2 | 0 |
| rust-hmac | 0.000 | 0.000 | 0.000 | 0 | 1 | 4 | 0 |
| rust-md-5 | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| rust-openssl | 0.500 | 0.130 | 0.206 | 7 | 7 | 47 | 0 |
| rust-ring | 0.304 | 0.156 | 0.206 | 7 | 16 | 38 | 0 |
| rust-rustls | 1.000 | 0.154 | 0.267 | 2 | 0 | 11 | 0 |
| rust-rustls-webpki | 1.000 | 0.091 | 0.167 | 1 | 0 | 10 | 0 |
| rust-sha1 | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 |
| rust-sha2 | 1.000 | 0.250 | 0.400 | 1 | 0 | 3 | 0 |
| rust-siphasher | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| rust-tokio-rustls | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 18 / 31 (0.581) | 0.277 | 0.091 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| A | 18 / 31 (0.581) | 0.277 | 0.091 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| rust | 18 / 31 (0.581) | 0.277 | 0.091 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 4 / 4 (1.000) | 1.000 | 0.036 |
| certificate | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 14 / 27 (0.519) | 0.230 | 0.250 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| rust-aes | 0 / 0 (n/a) | 0.000 | 0.000 |
| rust-chacha20 | 1 / 1 (1.000) | 0.250 | 0.143 |
| rust-curve25519-dalek | 0 / 2 (0.000) | 0.000 | 0.000 |
| rust-ecdsa | 4 / 6 (0.667) | 0.571 | 0.250 |
| rust-hkdf | 0 / 2 (0.000) | 0.000 | 0.000 |
| rust-hmac | 0 / 0 (n/a) | 0.000 | 0.000 |
| rust-md-5 | 0 / 0 (n/a) | n/a | 0.000 |
| rust-openssl | 3 / 7 (0.429) | 0.214 | 0.056 |
| rust-ring | 4 / 7 (0.571) | 0.174 | 0.089 |
| rust-rustls | 2 / 2 (1.000) | 1.000 | 0.154 |
| rust-rustls-webpki | 1 / 1 (1.000) | 1.000 | 0.091 |
| rust-sha1 | 1 / 1 (1.000) | 1.000 | 0.500 |
| rust-sha2 | 1 / 1 (1.000) | 1.000 | 0.250 |
| rust-siphasher | 0 / 0 (n/a) | n/a | 0.000 |
| rust-tokio-rustls | 1 / 1 (1.000) | 1.000 | 0.125 |

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
| all | 0.460 | 0.192 | 0.271 | 29 | 34 | 122 | 2 | 16 / 29 (0.552) | 0.254 | 0.106 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| rust | 0.460 | 0.192 | 0.271 | 29 | 34 | 122 | 2 | 16 / 29 (0.552) | 0.254 | 0.106 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.051 | 0.096 | 4 | 0 | 75 | 0 | 4 / 4 (1.000) | 1.000 | 0.051 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 9 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 15 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 0.424 | 0.521 | 0.467 | 25 | 34 | 23 | 2 | 12 / 25 (0.480) | 0.203 | 0.250 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| rust-aes | 0.000 | 0.000 | 0.000 | 0 | 2 | 4 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| rust-chacha20 | 0.250 | 0.167 | 0.200 | 1 | 3 | 5 | 0 | 1 / 1 (1.000) | 0.250 | 0.167 |
| rust-curve25519-dalek | 0.500 | 0.118 | 0.190 | 2 | 2 | 15 | 0 | 0 / 2 (0.000) | 0.000 | 0.000 |
| rust-ecdsa | 0.800 | 0.500 | 0.615 | 4 | 1 | 4 | 2 | 2 / 4 (0.500) | 0.400 | 0.250 |
| rust-hkdf | 0.500 | 0.500 | 0.500 | 2 | 2 | 2 | 0 | 0 / 2 (0.000) | 0.000 | 0.000 |
| rust-hmac | 0.000 | 0.000 | 0.000 | 0 | 1 | 3 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| rust-md-5 | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust-openssl | 0.500 | 0.184 | 0.269 | 7 | 7 | 31 | 0 | 3 / 7 (0.429) | 0.214 | 0.079 |
| rust-ring | 0.304 | 0.212 | 0.250 | 7 | 16 | 26 | 0 | 4 / 7 (0.571) | 0.174 | 0.121 |
| rust-rustls | 1.000 | 0.154 | 0.267 | 2 | 0 | 11 | 0 | 2 / 2 (1.000) | 1.000 | 0.154 |
| rust-rustls-webpki | 1.000 | 0.167 | 0.286 | 1 | 0 | 5 | 0 | 1 / 1 (1.000) | 1.000 | 0.167 |
| rust-sha1 | 1.000 | 0.500 | 0.667 | 1 | 0 | 1 | 0 | 1 / 1 (1.000) | 1.000 | 0.500 |
| rust-sha2 | 1.000 | 0.250 | 0.400 | 1 | 0 | 3 | 0 | 1 / 1 (1.000) | 1.000 | 0.250 |
| rust-siphasher | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust-tokio-rustls | 1.000 | 0.125 | 0.222 | 1 | 0 | 7 | 0 | 1 / 1 (1.000) | 1.000 | 0.125 |
