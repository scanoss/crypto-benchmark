# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (14): go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable

Not run, no prediction file (135): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable

Predictions outside scope.files or scope.components, ignored: 22

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| B | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-b-deadfn-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| go-b-deadfn-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| go-b-depth3-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-depth3-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-direct-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-direct-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-dispatch-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-dispatch-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-transitive-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-transitive-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-unused-import-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-unused-import-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-wrapper-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| go-b-wrapper-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 0 / 0 (n/a) | n/a | 0.000 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| B | 0 / 0 (n/a) | n/a | 0.000 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go | 0 / 0 (n/a) | n/a | 0.000 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-b-deadfn-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-deadfn-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-depth3-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-depth3-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-direct-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-direct-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-dispatch-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-dispatch-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-transitive-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-transitive-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-unused-import-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-unused-import-unreachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-wrapper-reachable | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-wrapper-unreachable | 0 / 0 (n/a) | n/a | 0.000 |

## Reachability

n/a: no scored entry's predictions report reachability.

## Suite B pair discrimination

n/a: no scored entry's predictions report reachability.

## Suite B reachability

No reachability: none of the 14 scored pair entries' predictions report reachability. The tool is not scored on reachability, depth or pairs; the outcome table below only shows which assets it detected.

### Overall

| Group | Truth | Assets | reachable | unreachable | no answer | not detected |
|---|---|---|---|---|---|---|
| all | reachable | 7 | 0 | 0 | 0 | 7 |
| all | unreachable | 9 | 0 | 0 | 0 | 9 |

### By language

| Group | Truth | Assets | reachable | unreachable | no answer | not detected |
|---|---|---|---|---|---|---|
| go | reachable | 7 | 0 | 0 | 0 | 7 |
| go | unreachable | 9 | 0 | 0 | 0 | 9 |

### By scenario

| Group | Truth | Assets | reachable | unreachable | no answer | not detected |
|---|---|---|---|---|---|---|
| deadfn | reachable | 1 | 0 | 0 | 0 | 1 |
| deadfn | unreachable | 3 | 0 | 0 | 0 | 3 |
| depth3 | reachable | 1 | 0 | 0 | 0 | 1 |
| depth3 | unreachable | 1 | 0 | 0 | 0 | 1 |
| direct | reachable | 1 | 0 | 0 | 0 | 1 |
| direct | unreachable | 1 | 0 | 0 | 0 | 1 |
| dispatch | reachable | 1 | 0 | 0 | 0 | 1 |
| dispatch | unreachable | 1 | 0 | 0 | 0 | 1 |
| transitive | reachable | 1 | 0 | 0 | 0 | 1 |
| transitive | unreachable | 1 | 0 | 0 | 0 | 1 |
| unused-import | reachable | 1 | 0 | 0 | 0 | 1 |
| unused-import | unreachable | 1 | 0 | 0 | 0 | 1 |
| wrapper | reachable | 1 | 0 | 0 | 0 | 1 |
| wrapper | unreachable | 1 | 0 | 0 | 0 | 1 |

### By language and scenario

| Group | Truth | Assets | reachable | unreachable | no answer | not detected |
|---|---|---|---|---|---|---|
| go / deadfn | reachable | 1 | 0 | 0 | 0 | 1 |
| go / deadfn | unreachable | 3 | 0 | 0 | 0 | 3 |
| go / depth3 | reachable | 1 | 0 | 0 | 0 | 1 |
| go / depth3 | unreachable | 1 | 0 | 0 | 0 | 1 |
| go / direct | reachable | 1 | 0 | 0 | 0 | 1 |
| go / direct | unreachable | 1 | 0 | 0 | 0 | 1 |
| go / dispatch | reachable | 1 | 0 | 0 | 0 | 1 |
| go / dispatch | unreachable | 1 | 0 | 0 | 0 | 1 |
| go / transitive | reachable | 1 | 0 | 0 | 0 | 1 |
| go / transitive | unreachable | 1 | 0 | 0 | 0 | 1 |
| go / unused-import | reachable | 1 | 0 | 0 | 0 | 1 |
| go / unused-import | unreachable | 1 | 0 | 0 | 0 | 1 |
| go / wrapper | reachable | 1 | 0 | 0 | 0 | 1 |
| go / wrapper | unreachable | 1 | 0 | 0 | 0 | 1 |

## Object-level view

Each ground-truth object counts once: the assets that share an `object` id, or one asset without one. An object is detected when any of its assets is matched above, and attributed when any of its matched assets is attributed. A prediction matched to an object already detected is a duplicate; FPs are the same as above.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| all | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-b-deadfn-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-deadfn-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-depth3-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-depth3-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-direct-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-direct-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-dispatch-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-dispatch-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-transitive-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-transitive-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-unused-import-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-unused-import-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-wrapper-reachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| go-b-wrapper-unreachable | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
