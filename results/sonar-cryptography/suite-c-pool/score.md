# Score: sonar-cryptography

- Tool: sonar-cryptography 1.7.0
- Config: (none)

Entries scored (6): java-c-mall-pool, java-c-jadx-pool, python-c-yt-dlp-pool, python-c-graphify-pool, go-c-ollama-pool, go-c-caveman-pool

Not run, no prediction file (183): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable, java-c-mall, java-c-jadx, python-c-yt-dlp, python-c-graphify, go-c-ollama, go-c-caveman, rust-c-clash-verge-rev, rust-c-cc-switch, javascript-c-freecodecamp, javascript-c-firecrawl, rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.
Rows marked † include entries with scope.kind pool; their recall is relative to pool.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all † | 0.983 | 0.286 | 0.444 | 169 | 3 | 421 | 16 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| C † | 0.983 | 0.286 | 0.444 | 169 | 3 | 421 | 16 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go † | 1.000 | 0.458 | 0.628 | 168 | 0 | 199 | 16 |
| java † | 0.250 | 0.091 | 0.133 | 1 | 3 | 10 | 0 |
| python † | n/a | 0.000 | 0.000 | 0 | 0 | 212 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm † | 0.983 | 0.371 | 0.539 | 169 | 3 | 286 | 16 |
| certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 39 | 0 |
| protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 39 | 0 |
| related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 57 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm † | 1.000 | 0.577 | 0.732 | 168 | 0 | 123 | 16 |
| go / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 23 | 0 |
| go / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| go / related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 39 | 0 |
| java / algorithm † | 0.250 | 0.111 | 0.154 | 1 | 3 | 8 | 0 |
| java / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| python / algorithm † | n/a | 0.000 | 0.000 | 0 | 0 | 155 | 0 |
| python / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| python / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 25 | 0 |
| python / related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 18 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-c-caveman-pool † | 1.000 | 0.426 | 0.598 | 110 | 0 | 148 | 12 |
| go-c-ollama-pool † | 1.000 | 0.532 | 0.695 | 58 | 0 | 51 | 4 |
| java-c-jadx-pool † | 0.250 | 0.143 | 0.182 | 1 | 3 | 6 | 0 |
| java-c-mall-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| python-c-graphify-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 58 | 0 |
| python-c-yt-dlp-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 154 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all † | 166 / 169 (0.982) | 0.965 | 0.281 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| C † | 166 / 169 (0.982) | 0.965 | 0.281 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go † | 165 / 168 (0.982) | 0.982 | 0.450 |
| java † | 1 / 1 (1.000) | 0.250 | 0.091 |
| python † | 0 / 0 (n/a) | n/a | 0.000 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm † | 166 / 169 (0.982) | 0.965 | 0.365 |
| certificate † | 0 / 0 (n/a) | n/a | 0.000 |
| protocol † | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material † | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-c-caveman-pool † | 107 / 110 (0.973) | 0.973 | 0.415 |
| go-c-ollama-pool † | 58 / 58 (1.000) | 1.000 | 0.532 |
| java-c-jadx-pool † | 1 / 1 (1.000) | 0.250 | 0.143 |
| java-c-mall-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-graphify-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp-pool † | 0 / 0 (n/a) | n/a | 0.000 |

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
| all † | 0.983 | 0.300 | 0.460 | 169 | 3 | 394 | 16 | 166 / 169 (0.982) | 0.965 | 0.295 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go † | 1.000 | 0.479 | 0.647 | 168 | 0 | 183 | 16 | 165 / 168 (0.982) | 0.982 | 0.470 |
| java † | 0.250 | 0.100 | 0.143 | 1 | 3 | 9 | 0 | 1 / 1 (1.000) | 0.250 | 0.100 |
| python † | n/a | 0.000 | 0.000 | 0 | 0 | 202 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm † | 0.983 | 0.385 | 0.553 | 169 | 3 | 270 | 16 | 166 / 169 (0.982) | 0.965 | 0.378 |
| certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 34 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 39 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 51 | 0 | 0 / 0 (n/a) | n/a | 0.000 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-c-caveman-pool † | 1.000 | 0.438 | 0.609 | 110 | 0 | 141 | 12 | 107 / 110 (0.973) | 0.973 | 0.426 |
| go-c-ollama-pool † | 1.000 | 0.580 | 0.734 | 58 | 0 | 42 | 4 | 58 / 58 (1.000) | 1.000 | 0.580 |
| java-c-jadx-pool † | 0.250 | 0.143 | 0.182 | 1 | 3 | 6 | 0 | 1 / 1 (1.000) | 0.250 | 0.143 |
| java-c-mall-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-graphify-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 58 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 144 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
