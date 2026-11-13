# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (10): java-c-mall-pool, java-c-jadx-pool, python-c-yt-dlp-pool, python-c-graphify-pool, go-c-ollama-pool, go-c-caveman-pool, rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool

Not run, no prediction file (179): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable, java-c-mall, java-c-jadx, python-c-yt-dlp, python-c-graphify, go-c-ollama, go-c-caveman, rust-c-clash-verge-rev, rust-c-cc-switch, javascript-c-freecodecamp, javascript-c-firecrawl

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.
Rows marked † include entries with scope.kind pool; their recall is relative to pool.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all † | 0.099 | 0.077 | 0.086 | 76 | 689 | 917 | 0 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| C † | 0.099 | 0.077 | 0.086 | 76 | 689 | 917 | 0 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go † | 0.028 | 0.019 | 0.023 | 7 | 245 | 360 | 0 |
| java † | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 |
| javascript † | 1.000 | 0.144 | 0.252 | 50 | 0 | 297 | 0 |
| python † | n/a | 0.000 | 0.000 | 0 | 0 | 212 | 0 |
| rust † | 0.041 | 0.339 | 0.073 | 19 | 444 | 37 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm † | 0.902 | 0.089 | 0.163 | 74 | 8 | 753 | 0 |
| certificate † | 0.000 | 0.000 | 0.000 | 0 | 1 | 41 | 0 |
| protocol † | 0.000 | 0.000 | 0.000 | 0 | 1 | 46 | 0 |
| related-crypto-material † | 0.003 | 0.025 | 0.005 | 2 | 679 | 77 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm † | 1.000 | 0.017 | 0.034 | 5 | 0 | 286 | 0 |
| go / certificate † | 0.000 | 0.000 | 0.000 | 0 | 1 | 23 | 0 |
| go / protocol † | 0.000 | 0.000 | 0.000 | 0 | 1 | 14 | 0 |
| go / related-crypto-material † | 0.008 | 0.051 | 0.014 | 2 | 243 | 37 | 0 |
| java / algorithm † | n/a | 0.000 | 0.000 | 0 | 0 | 9 | 0 |
| java / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| javascript / algorithm † | 1.000 | 0.154 | 0.267 | 50 | 0 | 274 | 0 |
| javascript / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| javascript / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| javascript / related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 21 | 0 |
| python / algorithm † | n/a | 0.000 | 0.000 | 0 | 0 | 155 | 0 |
| python / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 14 | 0 |
| python / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 25 | 0 |
| python / related-crypto-material † | n/a | 0.000 | 0.000 | 0 | 0 | 18 | 0 |
| rust / algorithm † | 0.704 | 0.396 | 0.507 | 19 | 8 | 29 | 0 |
| rust / certificate † | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| rust / protocol † | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| rust / related-crypto-material † | 0.000 | 0.000 | 0.000 | 0 | 436 | 1 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-c-caveman-pool † | 0.028 | 0.016 | 0.020 | 4 | 139 | 254 | 0 |
| go-c-ollama-pool † | 0.028 | 0.028 | 0.028 | 3 | 106 | 106 | 0 |
| java-c-jadx-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 7 | 0 |
| java-c-mall-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 4 | 0 |
| javascript-c-firecrawl-pool † | 1.000 | 0.158 | 0.273 | 47 | 0 | 250 | 0 |
| javascript-c-freecodecamp-pool † | 1.000 | 0.060 | 0.113 | 3 | 0 | 47 | 0 |
| python-c-graphify-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 58 | 0 |
| python-c-yt-dlp-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 154 | 0 |
| rust-c-cc-switch-pool † | 0.045 | 0.442 | 0.082 | 19 | 399 | 24 | 0 |
| rust-c-clash-verge-rev-pool † | 0.000 | 0.000 | 0.000 | 0 | 45 | 13 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all † | 63 / 76 (0.829) | 0.082 | 0.063 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| C † | 63 / 76 (0.829) | 0.082 | 0.063 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go † | 6 / 7 (0.857) | 0.024 | 0.016 |
| java † | 0 / 0 (n/a) | n/a | 0.000 |
| javascript † | 38 / 50 (0.760) | 0.760 | 0.110 |
| python † | 0 / 0 (n/a) | n/a | 0.000 |
| rust † | 19 / 19 (1.000) | 0.041 | 0.339 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm † | 61 / 74 (0.824) | 0.744 | 0.074 |
| certificate † | 0 / 0 (n/a) | 0.000 | 0.000 |
| protocol † | 0 / 0 (n/a) | 0.000 | 0.000 |
| related-crypto-material † | 2 / 2 (1.000) | 0.003 | 0.025 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-c-caveman-pool † | 3 / 4 (0.750) | 0.021 | 0.012 |
| go-c-ollama-pool † | 3 / 3 (1.000) | 0.028 | 0.028 |
| java-c-jadx-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| java-c-mall-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-c-firecrawl-pool † | 35 / 47 (0.745) | 0.745 | 0.118 |
| javascript-c-freecodecamp-pool † | 3 / 3 (1.000) | 1.000 | 0.060 |
| python-c-graphify-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp-pool † | 0 / 0 (n/a) | n/a | 0.000 |
| rust-c-cc-switch-pool † | 19 / 19 (1.000) | 0.045 | 0.442 |
| rust-c-clash-verge-rev-pool † | 0 / 0 (n/a) | 0.000 | 0.000 |

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
| all † | 0.098 | 0.081 | 0.089 | 75 | 689 | 851 | 1 | 63 / 75 (0.840) | 0.082 | 0.068 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go † | 0.028 | 0.020 | 0.023 | 7 | 245 | 344 | 0 | 6 / 7 (0.857) | 0.024 | 0.017 |
| java † | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript † | 1.000 | 0.159 | 0.275 | 49 | 0 | 259 | 1 | 38 / 49 (0.776) | 0.776 | 0.123 |
| python † | n/a | 0.000 | 0.000 | 0 | 0 | 202 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust † | 0.041 | 0.345 | 0.073 | 19 | 444 | 36 | 0 | 19 / 19 (1.000) | 0.041 | 0.345 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm † | 0.901 | 0.095 | 0.171 | 73 | 8 | 698 | 1 | 61 / 73 (0.836) | 0.753 | 0.079 |
| certificate † | 0.000 | 0.000 | 0.000 | 0 | 1 | 36 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| protocol † | 0.000 | 0.000 | 0.000 | 0 | 1 | 46 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| related-crypto-material † | 0.003 | 0.027 | 0.005 | 2 | 679 | 71 | 0 | 2 / 2 (1.000) | 0.003 | 0.027 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-c-caveman-pool † | 0.028 | 0.016 | 0.020 | 4 | 139 | 247 | 0 | 3 / 4 (0.750) | 0.021 | 0.012 |
| go-c-ollama-pool † | 0.028 | 0.030 | 0.029 | 3 | 106 | 97 | 0 | 3 / 3 (1.000) | 0.028 | 0.030 |
| java-c-jadx-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 7 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-c-mall-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-c-firecrawl-pool † | 1.000 | 0.159 | 0.274 | 46 | 0 | 244 | 1 | 35 / 46 (0.761) | 0.761 | 0.121 |
| javascript-c-freecodecamp-pool † | 1.000 | 0.167 | 0.286 | 3 | 0 | 15 | 0 | 3 / 3 (1.000) | 1.000 | 0.167 |
| python-c-graphify-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 58 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp-pool † | n/a | 0.000 | 0.000 | 0 | 0 | 144 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust-c-cc-switch-pool † | 0.045 | 0.442 | 0.082 | 19 | 399 | 24 | 0 | 19 / 19 (1.000) | 0.045 | 0.442 |
| rust-c-clash-verge-rev-pool † | 0.000 | 0.000 | 0.000 | 0 | 45 | 12 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
