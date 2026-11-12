# Rust suite A fixtures

One binary crate per library in the frozen Rust list (`corpus/libraries.yaml`), with one binary per primitive under `src/bin/`. Each crate pins its dependencies with `=` in `Cargo.toml` and commits its `Cargo.lock`. Every binary builds and runs with Rust 1.98:

```sh
CARGO_TARGET_DIR=/tmp/cbench-rust-target cargo run --manifest-path fixtures/rust/ring/Cargo.toml --bin ed25519
```

Keep `target/` out of the repository by setting `CARGO_TARGET_DIR`.

## Limits

- **ecdsa** is a generic crate. Programs reach it through the p256 and k256 curve crates, as its documentation shows; `key_bytes.rs` names the `ecdsa` types directly.
- **rustls, tokio-rustls and rustls-webpki** run without a network. `handshake.rs` completes a TLS 1.3 handshake in memory (rustls) or over a `tokio::io::duplex` pipe (tokio-rustls). `client_config.rs` and `connector.rs` build a client against the webpki-roots trust store and stop before connecting.
- **rustls-webpki** verifies a leaf certificate against a test CA and checks a signature made with the leaf's key. The test PKI in each `certs/` directory is an ECDSA P-256 CA and a `localhost` leaf valid for 100 years, made with this OpenSSL 3 script (kept here, outside the entries, so no scanner reads it as part of an entry):

  ```sh
  openssl req -x509 -newkey ec -pkeyopt ec_paramgen_curve:P-256 -nodes -keyout ca.key -out ca.pem \
    -days 36500 -subj "/CN=cbench test CA" -addext basicConstraints=critical,CA:TRUE -addext keyUsage=critical,keyCertSign
  openssl req -newkey ec -pkeyopt ec_paramgen_curve:P-256 -nodes -keyout localhost.key -out localhost.csr -subj "/CN=localhost"
  printf 'subjectAltName=DNS:localhost\nbasicConstraints=CA:FALSE\nextendedKeyUsage=serverAuth\nkeyUsage=critical,digitalSignature\n' > leaf.ext
  openssl x509 -req -in localhost.csr -CA ca.pem -CAkey ca.key -CAcreateserial -days 36500 -sha256 -extfile leaf.ext -out localhost.pem
  openssl pkcs8 -topk8 -nocrypt -in localhost.key -out localhost.key.pk8 && mv localhost.key.pk8 localhost.key
  ```

- **ring** cannot generate RSA keys, so `rsa_pkcs1.rs` loads a 2048-bit test key from `keys/rsa-2048.pk8` (`openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048`, converted to PKCS#8 DER). The keys are test data and protect nothing.
- **siphasher** is a keyed hash for hash tables, not a general-purpose MAC. The programs use it with explicit keys, which is how it gets its keyed-PRF property.
- **curve25519-dalek** is group arithmetic. The programs do X25519 with `MontgomeryPoint`, Diffie-Hellman over ristretto255 (the key agreement is the `*` operator), and hash-to-scalar and hash-to-group with SHA-512.
