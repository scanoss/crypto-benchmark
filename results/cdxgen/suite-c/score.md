# Score: cdxgen

- Tool: cdxgen 13.6.0
- Config: cdxgen_locations=true

Entries scored (10): java-c-mall, java-c-jadx, python-c-yt-dlp, python-c-graphify, go-c-ollama, go-c-caveman, rust-c-clash-verge-rev, rust-c-cc-switch, javascript-c-freecodecamp, javascript-c-firecrawl

Not run, no prediction file (169): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

Predictions outside scope.files or scope.components, ignored: 714

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 0.353 | 0.084 | 0.136 | 18 | 33 | 196 | 0 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| C | 0.353 | 0.084 | 0.136 | 18 | 33 | 196 | 0 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go | 0.250 | 0.033 | 0.058 | 2 | 6 | 59 | 0 |
| java | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| javascript | 1.000 | 0.140 | 0.246 | 8 | 0 | 49 | 0 |
| python | n/a | 0.000 | 0.000 | 0 | 0 | 61 | 0 |
| rust | 0.229 | 0.250 | 0.239 | 8 | 27 | 24 | 0 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.101 | 0.184 | 17 | 0 | 151 | 0 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 11 | 0 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 18 | 0 |
| related-crypto-material | 0.029 | 0.059 | 0.039 | 1 | 33 | 16 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm | 1.000 | 0.018 | 0.035 | 1 | 0 | 55 | 0 |
| go / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| go / related-crypto-material | 0.143 | 0.333 | 0.200 | 1 | 6 | 2 | 0 |
| java / algorithm | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 |
| javascript / algorithm | 1.000 | 0.157 | 0.271 | 8 | 0 | 43 | 0 |
| javascript / related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| python / algorithm | n/a | 0.000 | 0.000 | 0 | 0 | 35 | 0 |
| python / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 9 | 0 |
| python / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 |
| python / related-crypto-material | n/a | 0.000 | 0.000 | 0 | 0 | 7 | 0 |
| rust / algorithm | 1.000 | 0.348 | 0.516 | 8 | 0 | 15 | 0 |
| rust / certificate | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| rust / protocol | n/a | 0.000 | 0.000 | 0 | 0 | 6 | 0 |
| rust / related-crypto-material | 0.000 | 0.000 | 0.000 | 0 | 27 | 1 | 0 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-c-caveman | 0.000 | 0.000 | 0.000 | 0 | 4 | 30 | 0 |
| go-c-ollama | 0.500 | 0.065 | 0.114 | 2 | 2 | 29 | 0 |
| java-c-jadx | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 |
| java-c-mall | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 |
| javascript-c-firecrawl | 1.000 | 0.139 | 0.244 | 5 | 0 | 31 | 0 |
| javascript-c-freecodecamp | 1.000 | 0.143 | 0.250 | 3 | 0 | 18 | 0 |
| python-c-graphify | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 |
| python-c-yt-dlp | n/a | 0.000 | 0.000 | 0 | 0 | 45 | 0 |
| rust-c-cc-switch | 0.258 | 0.348 | 0.296 | 8 | 23 | 15 | 0 |
| rust-c-clash-verge-rev | 0.000 | 0.000 | 0.000 | 0 | 4 | 9 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 17 / 18 (0.944) | 0.333 | 0.079 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| C | 17 / 18 (0.944) | 0.333 | 0.079 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go | 2 / 2 (1.000) | 0.250 | 0.033 |
| java | 0 / 0 (n/a) | n/a | 0.000 |
| javascript | 7 / 8 (0.875) | 0.875 | 0.123 |
| python | 0 / 0 (n/a) | n/a | 0.000 |
| rust | 8 / 8 (1.000) | 0.229 | 0.250 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 16 / 17 (0.941) | 0.941 | 0.095 |
| certificate | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 1 / 1 (1.000) | 0.029 | 0.059 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-c-caveman | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-c-ollama | 2 / 2 (1.000) | 0.500 | 0.065 |
| java-c-jadx | 0 / 0 (n/a) | n/a | 0.000 |
| java-c-mall | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-c-firecrawl | 4 / 5 (0.800) | 0.800 | 0.111 |
| javascript-c-freecodecamp | 3 / 3 (1.000) | 1.000 | 0.143 |
| python-c-graphify | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp | 0 / 0 (n/a) | n/a | 0.000 |
| rust-c-cc-switch | 8 / 8 (1.000) | 0.258 | 0.348 |
| rust-c-clash-verge-rev | 0 / 0 (n/a) | 0.000 | 0.000 |

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
| all | 0.353 | 0.090 | 0.143 | 18 | 33 | 182 | 0 | 17 / 18 (0.944) | 0.333 | 0.085 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go | 0.250 | 0.033 | 0.059 | 2 | 6 | 58 | 0 | 2 / 2 (1.000) | 0.250 | 0.033 |
| java | n/a | 0.000 | 0.000 | 0 | 0 | 3 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript | 1.000 | 0.170 | 0.291 | 8 | 0 | 39 | 0 | 7 / 8 (0.875) | 0.875 | 0.149 |
| python | n/a | 0.000 | 0.000 | 0 | 0 | 60 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust | 0.229 | 0.267 | 0.246 | 8 | 27 | 22 | 0 | 8 / 8 (1.000) | 0.229 | 0.267 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 0.109 | 0.197 | 17 | 0 | 139 | 0 | 16 / 17 (0.941) | 0.941 | 0.103 |
| certificate | n/a | 0.000 | 0.000 | 0 | 0 | 10 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| protocol | n/a | 0.000 | 0.000 | 0 | 0 | 18 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| related-crypto-material | 0.029 | 0.062 | 0.040 | 1 | 33 | 15 | 0 | 1 / 1 (1.000) | 0.029 | 0.062 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-c-caveman | 0.000 | 0.000 | 0.000 | 0 | 4 | 30 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
| go-c-ollama | 0.500 | 0.067 | 0.118 | 2 | 2 | 28 | 0 | 2 / 2 (1.000) | 0.500 | 0.067 |
| java-c-jadx | n/a | 0.000 | 0.000 | 0 | 0 | 2 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| java-c-mall | n/a | 0.000 | 0.000 | 0 | 0 | 1 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| javascript-c-firecrawl | 1.000 | 0.143 | 0.250 | 5 | 0 | 30 | 0 | 4 / 5 (0.800) | 0.800 | 0.114 |
| javascript-c-freecodecamp | 1.000 | 0.250 | 0.400 | 3 | 0 | 9 | 0 | 3 / 3 (1.000) | 1.000 | 0.250 |
| python-c-graphify | n/a | 0.000 | 0.000 | 0 | 0 | 16 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| python-c-yt-dlp | n/a | 0.000 | 0.000 | 0 | 0 | 44 | 0 | 0 / 0 (n/a) | n/a | 0.000 |
| rust-c-cc-switch | 0.258 | 0.364 | 0.302 | 8 | 23 | 14 | 0 | 8 / 8 (1.000) | 0.258 | 0.364 |
| rust-c-clash-verge-rev | 0.000 | 0.000 | 0.000 | 0 | 4 | 8 | 0 | 0 / 0 (n/a) | 0.000 | 0.000 |
