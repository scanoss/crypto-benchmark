# Score: crypto-finder

- Tool: crypto-finder 0.54.0
- Config: ruleset_checksum=a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6, ruleset_source=remote, ruleset_version=v1.48.0

Entries scored (90): go-b-direct-reachable, go-b-direct-unreachable, go-b-transitive-reachable, go-b-transitive-unreachable, go-b-depth3-reachable, go-b-depth3-unreachable, go-b-wrapper-reachable, go-b-wrapper-unreachable, go-b-dispatch-reachable, go-b-dispatch-unreachable, go-b-deadfn-reachable, go-b-deadfn-unreachable, go-b-unused-import-reachable, go-b-unused-import-unreachable, python-b-direct-reachable, python-b-direct-unreachable, python-b-transitive-reachable, python-b-transitive-unreachable, python-b-depth3-reachable, python-b-depth3-unreachable, python-b-wrapper-reachable, python-b-wrapper-unreachable, python-b-dispatch-reachable, python-b-dispatch-unreachable, python-b-deadfn-reachable, python-b-deadfn-unreachable, python-b-unused-import-reachable, python-b-unused-import-unreachable, java-b-direct-reachable, java-b-direct-unreachable, java-b-transitive-reachable, java-b-transitive-unreachable, java-b-depth3-reachable, java-b-depth3-unreachable, java-b-wrapper-reachable, java-b-wrapper-unreachable, java-b-dispatch-reachable, java-b-dispatch-unreachable, java-b-deadfn-reachable, java-b-deadfn-unreachable, java-b-unused-import-reachable, java-b-unused-import-unreachable, javascript-b-direct-reachable, javascript-b-direct-unreachable, javascript-b-transitive-reachable, javascript-b-transitive-unreachable, javascript-b-depth3-reachable, javascript-b-depth3-unreachable, javascript-b-wrapper-reachable, javascript-b-wrapper-unreachable, javascript-b-dispatch-reachable, javascript-b-dispatch-unreachable, javascript-b-deadfn-reachable, javascript-b-deadfn-unreachable, javascript-b-unused-import-reachable, javascript-b-unused-import-unreachable, rust-b-direct-reachable, rust-b-direct-unreachable, rust-b-transitive-reachable, rust-b-transitive-unreachable, rust-b-depth3-reachable, rust-b-depth3-unreachable, rust-b-wrapper-reachable, rust-b-wrapper-unreachable, rust-b-dispatch-reachable, rust-b-dispatch-unreachable, rust-b-deadfn-reachable, rust-b-deadfn-unreachable, rust-b-unused-import-reachable, rust-b-unused-import-unreachable, go-b-handler-nethttp-reachable, go-b-handler-nethttp-unreachable, go-b-handler-gin-reachable, go-b-handler-gin-unreachable, go-b-callback-reachable, go-b-callback-unreachable, python-b-handler-flask-reachable, python-b-handler-flask-unreachable, python-b-handler-fastapi-reachable, python-b-handler-fastapi-unreachable, python-b-callback-reachable, python-b-callback-unreachable, java-b-handler-spring-reachable, java-b-handler-spring-unreachable, javascript-b-handler-express-reachable, javascript-b-handler-express-unreachable, rust-b-handler-axum-reachable, rust-b-handler-axum-unreachable, rust-b-handler-actix-reachable, rust-b-handler-actix-unreachable

Not run, no prediction file (99): python-stdlib, python-cryptography, python-pyjwt, python-truststore, python-rsa, python-pyopenssl, python-pynacl, python-bcrypt, python-paramiko, python-joserfc, python-pycryptodome, python-argon2-cffi, python-ecdsa, python-python-jose, python-passlib, python-pyotp, javascript-stdlib, javascript-jose, javascript-noble-hashes, javascript-jwa, javascript-jsonwebtoken, javascript-noble-ciphers, javascript-node-forge, javascript-fast-sha256, javascript-tweetnacl, javascript-noble-curves, javascript-crypto-random-string, javascript-blake3-wasm, javascript-hash-js, javascript-selfsigned, javascript-peculiar-x509, javascript-iron-webcrypto, go-stdlib, go-x-crypto-bcrypt, go-dgrijalva-jwt-go, go-go-ethereum-crypto, go-x-crypto-ssh, go-x-crypto-sha3, go-golang-jwt-v5, go-x-crypto-pbkdf2, go-x-crypto-ripemd160, go-x-crypto-chacha20poly1305, go-x-crypto-hkdf, go-x-crypto-blake2b, go-x-crypto-scrypt, go-x-crypto-ssh-agent, go-x-crypto-acme-autocert, go-x-crypto-curve25519, rust-aes, rust-chacha20, rust-curve25519-dalek, rust-ecdsa, rust-hkdf, rust-hmac, rust-md-5, rust-openssl, rust-ring, rust-rustls-webpki, rust-rustls, rust-sha1, rust-sha2, rust-siphasher, rust-tokio-rustls, java-stdlib, java-bcprov-jdk15on, java-bcpkix-jdk15on, java-bcpg-jdk15on, java-jsch, java-spring-security-crypto, java-eddsa, java-jasypt, java-jjwt-api, java-jbcrypt, java-nimbus-jose-jwt, java-jose4j, java-scrypt, java-bcrypt, java-cryptography-core-jvm, java-xmlsec, java-c-mall, java-c-jadx, python-c-yt-dlp, python-c-graphify, go-c-ollama, go-c-caveman, rust-c-clash-verge-rev, rust-c-cc-switch, javascript-c-freecodecamp, javascript-c-firecrawl, java-c-mall-pool, java-c-jadx-pool, python-c-yt-dlp-pool, python-c-graphify-pool, go-c-ollama-pool, go-c-caveman-pool, rust-c-clash-verge-rev-pool, rust-c-cc-switch-pool, javascript-c-freecodecamp-pool, javascript-c-firecrawl-pool

Predictions outside scope.files or scope.components, ignored: 1394

## Detection

Candidate match: same origin and component, same file and asset type, line ranges that overlap (LINE_TOLERANCE = 0). Ties prefer an attributed pair.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| all | 1.000 | 1.000 | 1.000 | 104 | 0 | 0 | 8 |

### By suite

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| B | 1.000 | 1.000 | 1.000 | 104 | 0 | 0 | 8 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 4 |
| java | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 0 |
| javascript | 1.000 | 1.000 | 1.000 | 18 | 0 | 0 | 0 |
| python | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 0 |
| rust | 1.000 | 1.000 | 1.000 | 20 | 0 | 0 | 4 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 1.000 | 1.000 | 100 | 0 | 0 | 8 |
| related-crypto-material | 1.000 | 1.000 | 1.000 | 4 | 0 | 0 | 0 |

### By language × asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go / algorithm | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 4 |
| java / algorithm | 1.000 | 1.000 | 1.000 | 18 | 0 | 0 | 0 |
| java / related-crypto-material | 1.000 | 1.000 | 1.000 | 4 | 0 | 0 | 0 |
| javascript / algorithm | 1.000 | 1.000 | 1.000 | 18 | 0 | 0 | 0 |
| python / algorithm | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 0 |
| rust / algorithm | 1.000 | 1.000 | 1.000 | 20 | 0 | 0 | 4 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates |
|---|---|---|---|---|---|---|---|
| go-b-callback-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-callback-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 |
| go-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 |
| go-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 |
| go-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 |
| go-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-handler-gin-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-handler-gin-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-handler-nethttp-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-handler-nethttp-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| go-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 3 | 0 | 0 | 0 |
| java-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 3 | 0 | 0 | 0 |
| java-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| java-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| java-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-handler-spring-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-handler-spring-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| java-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| javascript-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| javascript-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-handler-express-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-handler-express-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| javascript-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-callback-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-callback-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| python-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 |
| python-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-handler-fastapi-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-handler-fastapi-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-handler-flask-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-handler-flask-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| python-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 |
| rust-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 |
| rust-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 |
| rust-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 |
| rust-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-handler-actix-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-handler-actix-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-handler-axum-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-handler-axum-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |
| rust-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 |

## Attribution

A matched pair is attributed when the canonical algorithm and every ground-truth parameter also agree (`internal/algo`).

### Overall

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| all | 104 / 104 (1.000) | 1.000 | 1.000 |

### By suite

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| B | 104 / 104 (1.000) | 1.000 | 1.000 |

### By language

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go | 22 / 22 (1.000) | 1.000 | 1.000 |
| java | 22 / 22 (1.000) | 1.000 | 1.000 |
| javascript | 18 / 18 (1.000) | 1.000 | 1.000 |
| python | 22 / 22 (1.000) | 1.000 | 1.000 |
| rust | 20 / 20 (1.000) | 1.000 | 1.000 |

### By asset type

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| algorithm | 100 / 100 (1.000) | 1.000 | 1.000 |
| related-crypto-material | 4 / 4 (1.000) | 1.000 | 1.000 |

### By entry

| Group | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|
| go-b-callback-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-callback-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-deadfn-reachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| go-b-deadfn-unreachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| go-b-depth3-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-depth3-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-direct-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-direct-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-dispatch-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-dispatch-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-gin-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-gin-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-nethttp-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-nethttp-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-transitive-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-transitive-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-unused-import-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-unused-import-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-wrapper-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-wrapper-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-deadfn-reachable | 3 / 3 (1.000) | 1.000 | 1.000 |
| java-b-deadfn-unreachable | 3 / 3 (1.000) | 1.000 | 1.000 |
| java-b-depth3-reachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| java-b-depth3-unreachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| java-b-direct-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-direct-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-dispatch-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-dispatch-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-handler-spring-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-handler-spring-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-transitive-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-transitive-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-unused-import-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-unused-import-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-wrapper-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-wrapper-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-deadfn-reachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| javascript-b-deadfn-unreachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| javascript-b-depth3-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-depth3-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-direct-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-direct-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-dispatch-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-dispatch-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-handler-express-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-handler-express-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-transitive-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-transitive-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-unused-import-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-unused-import-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-wrapper-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-wrapper-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-callback-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-callback-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-deadfn-reachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| python-b-deadfn-unreachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| python-b-depth3-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-depth3-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-direct-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-direct-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-dispatch-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-dispatch-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-fastapi-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-fastapi-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-flask-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-flask-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-transitive-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-transitive-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-unused-import-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-unused-import-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-wrapper-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-wrapper-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-deadfn-reachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| rust-b-deadfn-unreachable | 2 / 2 (1.000) | 1.000 | 1.000 |
| rust-b-depth3-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-depth3-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-direct-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-direct-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-dispatch-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-dispatch-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-actix-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-actix-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-axum-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-axum-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-transitive-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-transitive-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-unused-import-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-unused-import-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-wrapper-reachable | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-wrapper-unreachable | 1 / 1 (1.000) | 1.000 | 1.000 |

## Reachability

Over matched pairs whose ground truth reachable is not unknown, in 90 of 90 scored entries.

### Overall

| Ground truth \ Predicted | reachable | unreachable | no answer |
|---|---|---|---|
| reachable | 47 | 0 | 0 |
| unreachable | 0 | 57 | 0 |

Reachable precision 1.000, reachable recall 1.000.

### Ground-truth depth 0

| Ground truth \ Predicted | reachable | unreachable | no answer |
|---|---|---|---|
| reachable | 0 | 0 | 0 |
| unreachable | 0 | 0 | 0 |

Reachable precision n/a, reachable recall n/a.

### Ground-truth depth 1

| Ground truth \ Predicted | reachable | unreachable | no answer |
|---|---|---|---|
| reachable | 5 | 0 | 0 |
| unreachable | 0 | 5 | 0 |

Reachable precision 1.000, reachable recall 1.000.

### Ground-truth depth 2+

| Ground truth \ Predicted | reachable | unreachable | no answer |
|---|---|---|---|
| reachable | 42 | 0 | 0 |
| unreachable | 0 | 52 | 0 |

Reachable precision 1.000, reachable recall 1.000.

## Suite B pair discrimination

Correct pairs: 45 / 45 (1.000)

| Reachable variant | Unreachable variant | Result | Why |
|---|---|---|---|
| go-b-direct-reachable | go-b-direct-unreachable | correct |  |
| go-b-transitive-reachable | go-b-transitive-unreachable | correct |  |
| go-b-depth3-reachable | go-b-depth3-unreachable | correct |  |
| go-b-wrapper-reachable | go-b-wrapper-unreachable | correct |  |
| go-b-dispatch-reachable | go-b-dispatch-unreachable | correct |  |
| go-b-deadfn-reachable | go-b-deadfn-unreachable | correct |  |
| go-b-unused-import-reachable | go-b-unused-import-unreachable | correct |  |
| python-b-direct-reachable | python-b-direct-unreachable | correct |  |
| python-b-transitive-reachable | python-b-transitive-unreachable | correct |  |
| python-b-depth3-reachable | python-b-depth3-unreachable | correct |  |
| python-b-wrapper-reachable | python-b-wrapper-unreachable | correct |  |
| python-b-dispatch-reachable | python-b-dispatch-unreachable | correct |  |
| python-b-deadfn-reachable | python-b-deadfn-unreachable | correct |  |
| python-b-unused-import-reachable | python-b-unused-import-unreachable | correct |  |
| java-b-direct-reachable | java-b-direct-unreachable | correct |  |
| java-b-transitive-reachable | java-b-transitive-unreachable | correct |  |
| java-b-depth3-reachable | java-b-depth3-unreachable | correct |  |
| java-b-wrapper-reachable | java-b-wrapper-unreachable | correct |  |
| java-b-dispatch-reachable | java-b-dispatch-unreachable | correct |  |
| java-b-deadfn-reachable | java-b-deadfn-unreachable | correct |  |
| java-b-unused-import-reachable | java-b-unused-import-unreachable | correct |  |
| javascript-b-direct-reachable | javascript-b-direct-unreachable | correct |  |
| javascript-b-transitive-reachable | javascript-b-transitive-unreachable | correct |  |
| javascript-b-depth3-reachable | javascript-b-depth3-unreachable | correct |  |
| javascript-b-wrapper-reachable | javascript-b-wrapper-unreachable | correct |  |
| javascript-b-dispatch-reachable | javascript-b-dispatch-unreachable | correct |  |
| javascript-b-deadfn-reachable | javascript-b-deadfn-unreachable | correct |  |
| javascript-b-unused-import-reachable | javascript-b-unused-import-unreachable | correct |  |
| rust-b-direct-reachable | rust-b-direct-unreachable | correct |  |
| rust-b-transitive-reachable | rust-b-transitive-unreachable | correct |  |
| rust-b-depth3-reachable | rust-b-depth3-unreachable | correct |  |
| rust-b-wrapper-reachable | rust-b-wrapper-unreachable | correct |  |
| rust-b-dispatch-reachable | rust-b-dispatch-unreachable | correct |  |
| rust-b-deadfn-reachable | rust-b-deadfn-unreachable | correct |  |
| rust-b-unused-import-reachable | rust-b-unused-import-unreachable | correct |  |
| go-b-handler-nethttp-reachable | go-b-handler-nethttp-unreachable | correct |  |
| go-b-handler-gin-reachable | go-b-handler-gin-unreachable | correct |  |
| go-b-callback-reachable | go-b-callback-unreachable | correct |  |
| python-b-handler-flask-reachable | python-b-handler-flask-unreachable | correct |  |
| python-b-handler-fastapi-reachable | python-b-handler-fastapi-unreachable | correct |  |
| python-b-callback-reachable | python-b-callback-unreachable | correct |  |
| java-b-handler-spring-reachable | java-b-handler-spring-unreachable | correct |  |
| javascript-b-handler-express-reachable | javascript-b-handler-express-unreachable | correct |  |
| rust-b-handler-axum-reachable | rust-b-handler-axum-unreachable | correct |  |
| rust-b-handler-actix-reachable | rust-b-handler-actix-unreachable | correct |  |

## Suite B reachability

Over every ground-truth asset of the 90 scored pair entries, matched or not. An asset no prediction matched is `not detected` and counts as a miss for reachable recall. Reachable precision is reachable assets answered reachable over all assets answered reachable. False reachable counts the `reachable: false` assets of unreachable variants answered reachable, over those detected. Depth (0, 1, 2+) and distance (exact) compare matched assets the tool answered with a depth.

### Overall

| Group | Truth | Assets | reachable | unreachable | no answer | not detected | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| all | reachable | 47 | 47 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 52 (0.000) | 104 / 104 (1.000) | 104 / 104 (1.000) | 45 / 45 |
| all | unreachable | 57 | 0 | 57 | 0 | 0 | | | | | | |

### By language

| Group | Truth | Assets | reachable | unreachable | no answer | not detected | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| go | reachable | 10 | 10 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 11 (0.000) | 22 / 22 (1.000) | 22 / 22 (1.000) | 10 / 10 |
| go | unreachable | 12 | 0 | 12 | 0 | 0 | | | | | | |
| java | reachable | 10 | 10 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 11 (0.000) | 22 / 22 (1.000) | 22 / 22 (1.000) | 8 / 8 |
| java | unreachable | 12 | 0 | 12 | 0 | 0 | | | | | | |
| javascript | reachable | 8 | 8 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 9 (0.000) | 18 / 18 (1.000) | 18 / 18 (1.000) | 8 / 8 |
| javascript | unreachable | 10 | 0 | 10 | 0 | 0 | | | | | | |
| python | reachable | 10 | 10 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 11 (0.000) | 22 / 22 (1.000) | 22 / 22 (1.000) | 10 / 10 |
| python | unreachable | 12 | 0 | 12 | 0 | 0 | | | | | | |
| rust | reachable | 9 | 9 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 10 (0.000) | 20 / 20 (1.000) | 20 / 20 (1.000) | 9 / 9 |
| rust | unreachable | 11 | 0 | 11 | 0 | 0 | | | | | | |

### By scenario

| Group | Truth | Assets | reachable | unreachable | no answer | not detected | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| callback | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 2 / 2 |
| callback | unreachable | 2 | 0 | 2 | 0 | 0 | | | | | | |
| deadfn | reachable | 6 | 6 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 11 (0.000) | 22 / 22 (1.000) | 22 / 22 (1.000) | 5 / 5 |
| deadfn | unreachable | 16 | 0 | 16 | 0 | 0 | | | | | | |
| depth3 | reachable | 6 | 6 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 6 (0.000) | 12 / 12 (1.000) | 12 / 12 (1.000) | 5 / 5 |
| depth3 | unreachable | 6 | 0 | 6 | 0 | 0 | | | | | | |
| direct | reachable | 5 | 5 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 5 (0.000) | 10 / 10 (1.000) | 10 / 10 (1.000) | 5 / 5 |
| direct | unreachable | 5 | 0 | 5 | 0 | 0 | | | | | | |
| dispatch | reachable | 5 | 5 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 5 (0.000) | 10 / 10 (1.000) | 10 / 10 (1.000) | 5 / 5 |
| dispatch | unreachable | 5 | 0 | 5 | 0 | 0 | | | | | | |
| handler | reachable | 8 | 8 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 8 (0.000) | 16 / 16 (1.000) | 16 / 16 (1.000) | 8 / 8 |
| handler | unreachable | 8 | 0 | 8 | 0 | 0 | | | | | | |
| transitive | reachable | 5 | 5 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 5 (0.000) | 10 / 10 (1.000) | 10 / 10 (1.000) | 5 / 5 |
| transitive | unreachable | 5 | 0 | 5 | 0 | 0 | | | | | | |
| unused-import | reachable | 5 | 5 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 5 (0.000) | 10 / 10 (1.000) | 10 / 10 (1.000) | 5 / 5 |
| unused-import | unreachable | 5 | 0 | 5 | 0 | 0 | | | | | | |
| wrapper | reachable | 5 | 5 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 5 (0.000) | 10 / 10 (1.000) | 10 / 10 (1.000) | 5 / 5 |
| wrapper | unreachable | 5 | 0 | 5 | 0 | 0 | | | | | | |

### By language and scenario

| Group | Truth | Assets | reachable | unreachable | no answer | not detected | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| go / callback | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / callback | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / deadfn | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 1 / 1 |
| go / deadfn | unreachable | 3 | 0 | 3 | 0 | 0 | | | | | | |
| go / depth3 | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / depth3 | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / direct | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / direct | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / dispatch | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / dispatch | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / handler | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 2 / 2 |
| go / handler | unreachable | 2 | 0 | 2 | 0 | 0 | | | | | | |
| go / transitive | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / transitive | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / unused-import | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / unused-import | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| go / wrapper | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| go / wrapper | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / deadfn | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 3 (0.000) | 6 / 6 (1.000) | 6 / 6 (1.000) | 1 / 1 |
| java / deadfn | unreachable | 4 | 0 | 4 | 0 | 0 | | | | | | |
| java / depth3 | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 1 / 1 |
| java / depth3 | unreachable | 2 | 0 | 2 | 0 | 0 | | | | | | |
| java / direct | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / direct | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / dispatch | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / dispatch | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / handler | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / handler | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / transitive | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / transitive | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / unused-import | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / unused-import | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| java / wrapper | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| java / wrapper | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / deadfn | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 1 / 1 |
| javascript / deadfn | unreachable | 3 | 0 | 3 | 0 | 0 | | | | | | |
| javascript / depth3 | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / depth3 | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / direct | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / direct | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / dispatch | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / dispatch | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / handler | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / handler | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / transitive | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / transitive | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / unused-import | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / unused-import | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| javascript / wrapper | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| javascript / wrapper | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / callback | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / callback | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / deadfn | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 1 / 1 |
| python / deadfn | unreachable | 3 | 0 | 3 | 0 | 0 | | | | | | |
| python / depth3 | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / depth3 | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / direct | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / direct | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / dispatch | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / dispatch | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / handler | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 2 / 2 |
| python / handler | unreachable | 2 | 0 | 2 | 0 | 0 | | | | | | |
| python / transitive | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / transitive | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / unused-import | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / unused-import | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| python / wrapper | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| python / wrapper | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / deadfn | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 1 / 1 |
| rust / deadfn | unreachable | 3 | 0 | 3 | 0 | 0 | | | | | | |
| rust / depth3 | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / depth3 | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / direct | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / direct | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / dispatch | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / dispatch | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / handler | reachable | 2 | 2 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 2 (0.000) | 4 / 4 (1.000) | 4 / 4 (1.000) | 2 / 2 |
| rust / handler | unreachable | 2 | 0 | 2 | 0 | 0 | | | | | | |
| rust / transitive | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / transitive | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / unused-import | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / unused-import | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |
| rust / wrapper | reachable | 1 | 1 | 0 | 0 | 0 | 1.000 | 1.000 | 0 / 1 (0.000) | 2 / 2 (1.000) | 2 / 2 (1.000) | 1 / 1 |
| rust / wrapper | unreachable | 1 | 0 | 1 | 0 | 0 | | | | | | |

## Object-level view

Each ground-truth object counts once: the assets that share an `object` id, or one asset without one. An object is detected when any of its assets is matched above, and attributed when any of its matched assets is attributed. A prediction matched to an object already detected is a duplicate; FPs are the same as above.

### Overall

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| all | 1.000 | 1.000 | 1.000 | 104 | 0 | 0 | 8 | 104 / 104 (1.000) | 1.000 | 1.000 |

### By language

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 4 | 22 / 22 (1.000) | 1.000 | 1.000 |
| java | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 0 | 22 / 22 (1.000) | 1.000 | 1.000 |
| javascript | 1.000 | 1.000 | 1.000 | 18 | 0 | 0 | 0 | 18 / 18 (1.000) | 1.000 | 1.000 |
| python | 1.000 | 1.000 | 1.000 | 22 | 0 | 0 | 0 | 22 / 22 (1.000) | 1.000 | 1.000 |
| rust | 1.000 | 1.000 | 1.000 | 20 | 0 | 0 | 4 | 20 / 20 (1.000) | 1.000 | 1.000 |

### By asset type

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| algorithm | 1.000 | 1.000 | 1.000 | 100 | 0 | 0 | 8 | 100 / 100 (1.000) | 1.000 | 1.000 |
| related-crypto-material | 1.000 | 1.000 | 1.000 | 4 | 0 | 0 | 0 | 4 / 4 (1.000) | 1.000 | 1.000 |

### By entry

| Group | Precision | Recall | F1 | TP | FP | FN | Duplicates | Attributed TP / TP | Attributed precision | Attributed recall |
|---|---|---|---|---|---|---|---|---|---|---|
| go-b-callback-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-callback-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 | 2 / 2 (1.000) | 1.000 | 1.000 |
| go-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 | 2 / 2 (1.000) | 1.000 | 1.000 |
| go-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-gin-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-gin-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-nethttp-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-handler-nethttp-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| go-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 3 | 0 | 0 | 0 | 3 / 3 (1.000) | 1.000 | 1.000 |
| java-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 3 | 0 | 0 | 0 | 3 / 3 (1.000) | 1.000 | 1.000 |
| java-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| java-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| java-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-handler-spring-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-handler-spring-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| java-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| javascript-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| javascript-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-handler-express-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-handler-express-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| javascript-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-callback-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-callback-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| python-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 0 | 2 / 2 (1.000) | 1.000 | 1.000 |
| python-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-fastapi-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-fastapi-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-flask-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-handler-flask-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| python-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-deadfn-reachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 | 2 / 2 (1.000) | 1.000 | 1.000 |
| rust-b-deadfn-unreachable | 1.000 | 1.000 | 1.000 | 2 | 0 | 0 | 1 | 2 / 2 (1.000) | 1.000 | 1.000 |
| rust-b-depth3-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-depth3-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 1 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-direct-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-direct-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-dispatch-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-dispatch-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-actix-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-actix-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-axum-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-handler-axum-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-transitive-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-transitive-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-unused-import-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-unused-import-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-wrapper-reachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
| rust-b-wrapper-unreachable | 1.000 | 1.000 | 1.000 | 1 | 0 | 0 | 0 | 1 / 1 (1.000) | 1.000 | 1.000 |
