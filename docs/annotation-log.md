# Annotation log

Agreement between two independent annotations of each entry, and how every disagreement was resolved. The process is in [annotation-guide.md](annotation-guide.md#process). Annotation is agent-only.

In every pass each disagreement was resolved against the guide and the resolution recorded in the asset's `notes`; questions the guide left open are listed with their pass. Passes, oldest first:

- [Suite A, Python pilot (2026-11-04)](#suite-a-python-pilot-2026-11-04)
- [Suite A, Python: related crypto material and digests (2026-11-04)](#suite-a-python-related-crypto-material-and-digests-2026-11-04)
- [Suite A, Python: objects (2026-11-04)](#suite-a-python-objects-2026-11-04)
- [Suite A, JavaScript (2026-11-04)](#suite-a-javascript-2026-11-04)
- [Suite A, Rust (2026-11-04)](#suite-a-rust-2026-11-04)
- [Suite A, Go (2026-11-04)](#suite-a-go-2026-11-04)
- [Suite A, Java (2026-11-04)](#suite-a-java-2026-11-04)
- [Suite A, anchors (2026-11-05)](#suite-a-anchors-2026-11-05)
- [Suite B, all five languages (2026-11-05)](#suite-b-all-five-languages-2026-11-05)
- [Suite B, handler and callback (2026-11-05)](#suite-b-handler-and-callback-2026-11-05)
- [Suite A, generated private keys (2026-11-08)](#suite-a-generated-private-keys-2026-11-08)
- [Suite C, `files` scope (2026-11-13)](#suite-c-files-scope-2026-11-13)

## Suite A, Python pilot (2026-11-04)

The first annotation is the ground truth as it stood before this pass. The program author wrote it, and an earlier review had already corrected it; those corrections are the older `notes` that mention a "second annotator". For this pass, one fresh agent per entry got only the entry's programs, `requirements.txt`, the annotation guide, the schema (`internal/groundtruth`) and the `internal/algo` tables, plus the installed library source. It never saw the ground truth. Each resolution is in the asset's `notes`, prefixed `Re-derivation:`.

How assets are matched:

- **Agreed:** both annotations list an asset in the same file, with the same lines, asset type, primitive, algorithm (compared on its `internal/algo` key) and parameters.
- **Field disagreement:** both list the call, but at least one field differs.
- **Only first / only second:** one annotation lists a call the other does not.
- **Agreement:** agreed over the union of all four counts.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|
| python-argon2-cffi | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 |
| python-bcrypt | 7 | 7 | 5 | 2 | 0 | 0 | 0.714 |
| python-cryptography | 44 | 43 | 42 | 1 | 1 | 0 | 0.955 |
| python-ecdsa | 19 | 21 | 19 | 0 | 0 | 2 | 0.905 |
| python-joserfc | 16 | 16 | 15 | 1 | 0 | 0 | 0.938 |
| python-paramiko | 13 | 12 | 12 | 0 | 1 | 0 | 0.923 |
| python-passlib | 10 | 12 | 10 | 0 | 0 | 2 | 0.833 |
| python-pycryptodome | 42 | 42 | 40 | 2 | 0 | 0 | 0.952 |
| python-pyjwt | 9 | 9 | 9 | 0 | 0 | 0 | 1.000 |
| python-pynacl | 19 | 19 | 16 | 3 | 0 | 0 | 0.842 |
| python-pyopenssl | 11 | 11 | 11 | 0 | 0 | 0 | 1.000 |
| python-pyotp | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| python-python-jose | 11 | 11 | 11 | 0 | 0 | 0 | 1.000 |
| python-rsa | 9 | 9 | 9 | 0 | 0 | 0 | 1.000 |
| python-stdlib | 15 | 15 | 15 | 0 | 0 | 0 | 1.000 |
| python-truststore | 3 | 3 | 0 | 0 | 3 | 3 | 0.000 |
| **Total** | **238** | **240** | **224** | **9** | **5** | **7** | **0.914** |

Across all entries, 233 calls were listed by both annotations at the same file, lines and asset type. On those, the two agreed on the `primitive` in 96.6% of cases. Cohen's kappa for `primitive` is 0.960 (chance agreement 0.139). Kappa is not given for whether a call is an asset, because that needs the calls neither annotation listed, and the guide does not enumerate every call.

python-pyotp's first annotation was written in this pass, by the author of its programs.

### Resolutions

Of the 21 disagreements, 2 changed the ground truth (both python-ecdsa, which also added a guide sentence), and 19 kept the first annotation. Of those 19, 9 led to a new sentence in the annotation guide so that the next annotator decides them the same way.

- **python-bcrypt, `bcrypt.kdf` (2, field `algorithm`).** The second annotation said `bcrypt`. The first said `bcrypt-pbkdf`, and it was kept: the call runs bcrypt_pbkdf, which is a different algorithm from the bcrypt password hash.
- **python-cryptography, `chacha20-poly1305` line 7 (1, field `parameters`).** The second annotation added `key_size: 256`. It was not added: ChaCha20-Poly1305 has only one key size, so the size does not tell variants apart. The guide now says so.
- **python-cryptography, `x509.random_serial_number()` (1, only first).** The second annotation folded it into the certificate builder chain. It stays an asset of its own, because the guide says a random draw nested in another call is still its own asset.
- **python-ecdsa, `generate_sharedsecret_bytes()` (2, only second).** Added to the ground truth. The call performs the key agreement with the key the ECDH object holds, as `private_key.exchange()` does in cryptography. The guide now says so.
- **python-joserfc, `jwt.decode` (1, field `parameters`).** The second annotation added padding `PKCS1v15`. It was not added: the decode call names no algorithm, and its API verifies whichever registered algorithm the token header names. The asset already carried this resolution from the earlier review.
- **python-paramiko, `key.fingerprint` (1, only first).** The second annotation left it out because it is a property, not a call. It was kept: reading the property computes the SHA-256 digest that the program prints. The guide now says that a property that computes a digest counts.
- **python-passlib, `context.verify` (2, only second).** Not added. The program created the `CryptContext`, so its methods are follow-ups, just as `PasswordHasher().verify` is in argon2-cffi, where both annotations agreed. The guide now separates a method on an object the program created from a one-shot function such as passlib's `argon2.verify`.
- **python-pycryptodome, ChaCha20-Poly1305 (2, field `parameters`).** The same `key_size: 256` question as in cryptography, with the same outcome.
- **python-pynacl, XSalsa20-Poly1305 in `Box(...).encrypt`, `Box(...).decrypt` and `SecretBox` (3, field `parameters`).** The same `key_size: 256` question, with the same outcome.
- **python-truststore (3 only first, 3 only second).** The second annotation listed the same three TLS context creations but swapped the two file names. It read both programs in one listing and attributed them wrongly. The first annotation matches the files. With the file names corrected, the two annotations agree on all three assets.

Without the truststore attribution error, overall agreement is 227 of 242 (0.938).

## Suite A, Python: related crypto material and digests (2026-11-04)

Related crypto material came into scope, and a digest now splits off signature and MAC names into a `hash` parameter ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules), change log). The guide gained the [Related crypto material](annotation-guide.md#related-crypto-material) section, the `material_type` field and the `hash` rule. This pass re-annotated the 16 Python entries for both.

### Related crypto material

The first annotation was written by the agent that wrote the new guide section. The second came from one fresh agent that got only the 16 entries' programs and `requirements.txt` files, the updated guide, the schema (`internal/groundtruth`) and `internal/algo`, copied into a directory with no ground truth in it. It had no installed library source and reasoned from each library's documented API. It listed only `related-crypto-material` assets with a `material_type`. Agreement is counted as in the pilot above, with `material_type` as the compared field.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|
| python-argon2-cffi | 1 | 1 | 1 | 0 | 0 | 0 | 1.000 |
| python-bcrypt | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 |
| python-cryptography | 7 | 7 | 7 | 0 | 0 | 0 | 1.000 |
| python-ecdsa | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 |
| python-pycryptodome | 7 | 7 | 7 | 0 | 0 | 0 | 1.000 |
| python-stdlib | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 |
| the other 10 entries | 0 | 0 | 0 | 0 | 0 | 0 | n/a |
| **Total** | **21** | **21** | **21** | **0** | **0** | **0** | **1.000** |

By material type, the 21 assets are 9 salts, 4 nonces, 4 shared secrets, 2 IVs and 2 tags. No program passes a seed to a crypto operation.

There was no disagreement to resolve. The second annotator reported five places where the guide made it reason more than it should have to. Each now has a guide sentence:

- **A salt passed through a helper in the same file** (cryptography `kdf(salt)` for PBKDF2 and scrypt). The exclusion for "a value created outside the entry's files" did not say what happens when the caller is in the same file. The guide now says such a value is annotated where it is created.
- **`bcrypt.gensalt()` against the hash-string exclusion.** The salt also ends up in the hash string that `hashpw` returns. The guide now says why it counts: the program creates it and passes it on.
- **A library call that is two assets.** pycryptodome `AES.new(key, AES.MODE_CBC)` is both the cipher asset and the IV. The guide stated the double count only for random draws, and now states it for library calls too.
- **The material type of a value no one passes at creation** (pycryptodome's drawn IV and nonces). The guide now takes it from the parameter the program later passes the value back as.
- **A shared secret inside a program function call** (`derive(alice.exchange(...))`). The nesting rule already covers it, since `derive` is not an asset. The guide was not changed.

Each new asset's `notes` starts with `Re-derivation: agreed.` and names the guide rule a different annotator might have missed.

### Digests

The `hash` parameter was added by the first annotator only, by the guide's new `hash` rule. It was not re-derived independently. Thirty-one signature and MAC assets gained `hash`. Fifteen `HMAC-SHA-256` or `HMAC-SHA-1` names became `HMAC`, with the digest in `hash`, except one. Three assets needed a decision, and each records it in `notes`:

- **python-joserfc, `jws.deserialize_compact` (`jws-hs256/main.py:8`).** It was `HMAC-SHA-256` and is now `HMAC` with no `hash`. The call names no algorithm and takes HS256 from the token, so only the parameters its API fixes apply, as the pilot already decided for the same library's `jwt.decode`.
- **python-pycryptodome, `DSS.new(key.public_key(), ...)` and `pss.new(key.public_key())` (verify).** Both gained `hash: SHA-256`, from the `SHA256.new(message)` digest that the follow-up `verifier.verify` call receives.

Left without `hash` by the guide's rule: python-ecdsa's sign and verify calls, whose digest is bound to the key when it is created or loaded; `rsa.verify` and paramiko `verify_ssh_sig`, which read the digest from the signature; and joserfc `jwt.decode`. RSA-OAEP, KDFs and password hashes get no `hash`, because the methodology splits digests off signature and MAC names only.

## Suite A, Python: objects (2026-11-04)

The object-level view was added ([METHODOLOGY.md](../METHODOLOGY.md#object-level-view)) and the `object` field ([guide](annotation-guide.md#objects)). This pass put `object` on the 16 Python entries. The first annotation was written by the agent that wrote the guide's Objects section, before it saw the second. The second came from one fresh agent that got the 16 entries' programs, the guide, the schema and `internal/algo` in a directory with no ground truth in it. Unlike earlier passes, it also got each entry's asset list (id, file, lines, asset type, primitive, material type, algorithm and parameters) without `object` or `notes`, because the assets were already agreed and only their grouping was being derived.

An asset is agreed when both annotations put it with exactly the same other assets, or both leave it alone. Object ids are not compared, since only their equality within an entry matters to the score.

| Entry | Assets | Multi-asset objects | Objects | Agreed | Disagreed | Agreement |
|---|---:|---:|---:|---:|---:|---:|
| python-argon2-cffi | 5 | 0 | 5 | 5 | 0 | 1.000 |
| python-bcrypt | 9 | 0 | 9 | 9 | 0 | 1.000 |
| python-cryptography | 51 | 12 | 34 | 51 | 0 | 1.000 |
| python-ecdsa | 23 | 6 | 8 | 23 | 0 | 1.000 |
| python-joserfc | 16 | 5 | 6 | 16 | 0 | 1.000 |
| python-paramiko | 13 | 3 | 9 | 13 | 0 | 1.000 |
| python-passlib | 10 | 0 | 10 | 10 | 0 | 1.000 |
| python-pycryptodome | 49 | 8 | 36 | 49 | 0 | 1.000 |
| python-pyjwt | 9 | 3 | 4 | 9 | 0 | 1.000 |
| python-pynacl | 19 | 4 | 14 | 19 | 0 | 1.000 |
| python-pyopenssl | 11 | 1 | 9 | 11 | 0 | 1.000 |
| python-pyotp | 6 | 0 | 6 | 6 | 0 | 1.000 |
| python-python-jose | 11 | 4 | 6 | 11 | 0 | 1.000 |
| python-rsa | 9 | 3 | 4 | 9 | 0 | 1.000 |
| python-stdlib | 17 | 1 | 16 | 17 | 0 | 1.000 |
| python-truststore | 3 | 0 | 3 | 3 | 0 | 1.000 |
| **Total** | **261** | **50** | **179** | **261** | **0** | **1.000** |

The two annotations formed the same 50 multi-asset objects, so the 261 assets make 179 objects. There was no disagreement to resolve. The second annotator listed thirteen places where it hesitated. It chose as the first annotation did in each, but the guide left the choice to the reader, so each now has a sentence in the guide:

- **PyNaCl `Box(...).encrypt`.** The key agreement goes with alice's key and the XSalsa20-Poly1305 with the box's derived key, which stands alone.
- **joserfc JWE.** The AES-GCM content encryption of `encrypt_compact` and `decrypt_compact` is one object, apart from the RSA key that wraps the content key.
- **Ids.** How to name the second object of one key, and an object with no variable (the JWE content key). The guide now gives a convention and says that spelling never changes a score.
- **Public-key serializations** (python-jose, python-ecdsa, pycryptodome, joserfc). They join the private key's serialization object.
- **Operations on a parsed key** (python-ecdsa `key-load` sign and verify). They act on the same key.
- **The scrypt in pycryptodome `export_key`.** Like the AES-GCM there, it belongs to neither key.
- **A certificate and the key that signs it.** Two objects, because they have different asset types.
- **Two KDF calls given the same password.** Two objects, because a password is not a key.
- **A random draw used as a key by one call.** No object. The guide now states this asymmetry with `AESGCM.generate_key()` on purpose.

The other four hesitations needed no change: the parsed and original `to_string()` serializations in python-ecdsa are one object by the "passed on" rule, a helper called twice is still one asset, `add_server_key` configures a transport and is not an asset, and a transport is its own object.

## Suite A, JavaScript (2026-11-04)

The 16 JavaScript entries were annotated under the current guide: specific algorithm names, `hash`, related crypto material with `material_type`, and `object`. The first annotation was written by the author of the programs. For the assets, one fresh agent per entry got only the entry's programs and `package.json`, the guide, the schema (`internal/groundtruth`), `internal/algo` and the installed library source, in a directory with no ground truth in it. Counting is as in the Python pilot above. The table compares the two annotations as written, before any resolution.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|
| javascript-blake3-wasm | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| javascript-crypto-random-string | 3 | 3 | 3 | 0 | 0 | 0 | 1.000 |
| javascript-fast-sha256 | 11 | 11 | 11 | 0 | 0 | 0 | 1.000 |
| javascript-hash-js | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| javascript-iron-webcrypto | 17 | 11 | 7 | 4 | 6 | 0 | 0.412 |
| javascript-jose | 25 | 25 | 22 | 3 | 0 | 0 | 0.880 |
| javascript-jsonwebtoken | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |
| javascript-jwa | 8 | 11 | 5 | 0 | 3 | 6 | 0.357 |
| javascript-noble-ciphers | 25 | 25 | 25 | 0 | 0 | 0 | 1.000 |
| javascript-noble-curves | 21 | 21 | 21 | 0 | 0 | 0 | 1.000 |
| javascript-noble-hashes | 24 | 24 | 24 | 0 | 0 | 0 | 1.000 |
| javascript-node-forge | 43 | 44 | 42 | 0 | 1 | 2 | 0.933 |
| javascript-peculiar-x509 | 12 | 12 | 10 | 2 | 0 | 0 | 0.833 |
| javascript-selfsigned | 5 | 5 | 3 | 2 | 0 | 0 | 0.600 |
| javascript-stdlib | 89 | 90 | 89 | 0 | 0 | 1 | 0.989 |
| javascript-tweetnacl | 29 | 29 | 29 | 0 | 0 | 0 | 1.000 |
| **Total** | **332** | **331** | **311** | **11** | **10** | **9** | **0.912** |

Both annotations listed 322 calls at the same file, lines and asset type, and agreed on the `primitive` of 320 of them (99.4%). After the resolutions below, the ground truth and the second annotation agree on 322 of 341 (0.944); the 19 left are the kept first-annotation decisions.

### Resolutions

Of the 30 differences, 11 changed the ground truth and 19 kept the first annotation. Each resolution is in the asset's `notes`, prefixed `Re-derivation:`.

- **jose and @peculiar/x509 key generation (5, field `algorithm` and `padding`). Changed.** The second annotation named WebCrypto keys by the scheme they are generated for: jose `generateKeyPair('RS256')` is RSA with padding PKCS1v15, `'PS256'` is RSA-PSS, `'RSA-OAEP-256'` is RSA-OAEP, and @peculiar/x509's `RSASSA-PKCS1-v1_5` keys are RSA with padding PKCS1v15. The first wrote plain RSA, while it already named jose's ES256 key ECDSA by the same rule. A WebCrypto key can only be used with the algorithm it was generated for, so it is a key type tied to one scheme.
- **iron-webcrypto HMAC key size (4, field `parameters`). Changed.** The second annotation added `key_size: 256`: the integrity algorithm `sha256` fixes a 256-bit HMAC key in the library's algorithm table, so the API fixes it.
- **selfsigned key generation primitive (2, field `primitive`). Changed** to `signature`. The program only stores the returned key, but the same `generate` call uses it to sign the certificate it returns.
- **iron-webcrypto PBKDF2 (6, only first). Kept.** The second annotation left PBKDF2 out because no argument names it. Iron stretches a string password with PBKDF2, and the options object passed to each call carries its salt size and iteration count.
- **jwa signer objects (3 only first, 6 only second). Kept.** The second annotation treated `jwa('HS256')` as a specification object, because it only returns closures, and annotated each `sign` and `verify`. The guide lists creating a signer object that names its algorithm as an asset and its methods as follow-ups, as for argon2-cffi `PasswordHasher()`; a specification object is one passed as a parameter to another asset. Where the program writes `jwa('RS256').sign(...)` on its own line, both annotations agreed on one asset per line.
- **node-forge certificate build (1 only first, 2 only second). Kept.** The second annotation put the certificate on `forge.pki.createCertificate()` alone and added an RSA PKCS1v15 SHA-256 signature for `cert.sign(...)`. forge builds a certificate in statements, not one chain; the first annotation spans them with one certificate asset, as the guide counts the whole build, signing included.
- **WebCrypto `importKey('raw', password, 'PBKDF2')` (1 only second). Kept** out: it wraps a password, and the guide does not annotate passwords.

### Objects

The `object` field landed while this pass ran, so the entries were grouped after the asset resolutions, by the first annotator. One more fresh agent got the 16 entries' programs, the guide, the schema and the resolved asset lists without `object` or `notes`, and grouped them independently. Counting is as in the Python objects pass.

| Entries | Assets | Multi-asset objects | Objects | Agreed | Disagreed | Agreement |
|---|---:|---:|---:|---:|---:|---:|
| all 16 | 332 | 61 | 238 | 332 | 0 | 1.000 |

By entry, the multi-asset objects are: jose 8, node-forge 8, noble-curves 7, stdlib 17, tweetnacl 6, peculiar-x509 4, jsonwebtoken 3, jwa 3, noble-ciphers 3, noble-hashes 1, selfsigned 1, and none in the other five. There was no disagreement.

### Questions the guide leaves open

The annotators hesitated in these places. Each was decided as above and recorded in `notes`; each would be settled by one guide sentence, which this pass does not write:

- **A library default the call leaves in place.** Both annotations read it as fixed by the API, as the guide does for pyotp's default SHA-1: PKCS#7 padding for noble, forge and node:crypto AES-CBC, PKCS#1 v1.5 for forge and node:crypto RSA signatures, SHA-256 for noble-curves ECDSA.
- **Node's `iv` parameter for a ChaCha20-Poly1305 nonce.** Both chose `initialization-vector`, by the parameter name.
- **A WebCrypto key's algorithm name,** settled above.
- **A certificate built in statements** (forge), settled above, and whether a CSR is a certificate asset (both said yes).
- **A factory that returns a signer computing nothing** (jwa), settled above.
- **A one-shot call that agrees a key and encrypts** (tweetnacl `nacl.box`). Both annotated two assets on the same lines.
- **Iron's derived keys.** Each call derives fresh keys with PBKDF2, so its AES and HMAC assets are their own objects.

## Suite A, Rust (2026-11-04)

The 15 Rust entries (`fixtures/rust/*`, one crate per row of the frozen Rust list) were annotated under the current guide, with related crypto material, `hash` and `object` from the start. The first annotation was written by the agent that wrote the programs. For the second, one fresh agent per entry got only the entry's `Cargo.toml`, programs and data files, the guide, the schema (`internal/groundtruth`), the `internal/algo` tables and an example file, plus the installed crate sources, in a directory with no ground truth in it. It derived the assets and their objects in one pass. Each resolution is in the asset's `notes`, prefixed `Re-derivation:`.

Assets are matched and counted as in the Python pilot above. Object agreement is counted over the assets both annotations list, as in the Python objects pass: an asset is agreed when both put it with the same other assets, or both leave it alone.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement | Objects agreed |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| rust-aes | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 | 4 / 4 |
| rust-chacha20 | 7 | 7 | 7 | 0 | 0 | 0 | 1.000 | 7 / 7 |
| rust-curve25519-dalek | 17 | 17 | 13 | 4 | 0 | 0 | 0.765 | 17 / 17 |
| rust-ecdsa | 13 | 13 | 7 | 6 | 0 | 0 | 0.538 | 13 / 13 |
| rust-hkdf | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 | 4 / 4 |
| rust-hmac | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 | 4 / 4 |
| rust-md-5 | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 | 2 / 2 |
| rust-openssl | 47 | 48 | 44 | 2 | 1 | 2 | 0.898 | 46 / 46 |
| rust-ring | 39 | 46 | 34 | 1 | 4 | 11 | 0.680 | 33 / 35 |
| rust-rustls | 13 | 13 | 13 | 0 | 0 | 0 | 1.000 | 13 / 13 |
| rust-rustls-webpki | 11 | 13 | 11 | 0 | 0 | 2 | 0.846 | 11 / 11 |
| rust-sha1 | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 | 2 / 2 |
| rust-sha2 | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 | 4 / 4 |
| rust-siphasher | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 | 4 / 4 |
| rust-tokio-rustls | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 | 8 / 8 |
| **Total** | **179** | **189** | **161** | **13** | **5** | **15** | **0.830** | **172 / 174** |

After resolution the ground truth has 182 assets in 136 objects: 112 algorithm, 14 certificate, 15 protocol and 41 related-crypto-material assets, 19 of the last with a `material_type` (5 nonces, 5 salts, 6 shared secrets, 2 tags and 1 IV; no seed).

Ten of the fifteen entries agreed on every asset. Of the 33 disagreements, 16 changed the ground truth to the second annotation, 6 were settled by a reading that neither annotation had (ring `UnparsedPublicKey`, below), 5 were rejected with that same reading, and 6 kept the first annotation. Both object disagreements followed from asset resolutions.

### Resolutions

- **rust-curve25519-dalek, ristretto255 key agreement (4, field `algorithm`).** The first annotation named it `Ristretto255`; the second wrote `ECDH` with curve `ristretto255`. Resolved to the second: it is Diffie-Hellman over the ristretto255 group, written the way the guide writes an EC scheme with its curve.
- **rust-ecdsa, curve on sign and verify (6, field `parameters`).** The second annotation set the curve as well as the hash. Resolved to the second: p256 and k256 fix the curve and the digest in the signing key's type, so both are parameters the API fixes, not values carried from key generation. ring's `EcdsaKeyPair` binds its algorithm when the key is loaded, so its `sign` stays without curve or hash.
- **rust-openssl, PKCS#7 padding on `symm::encrypt` and `symm::decrypt` (2, field `parameters`).** Kept the first annotation: the one-shot functions always pad, so the API fixes it, as in the guide's `aes_cbc_pkcs7_encrypt` example.
- **rust-openssl, `connector.configure()` (1, only first).** Kept: it creates the `Ssl` session for one connection, and both annotations annotated session creation as a protocol asset in rustls and tokio-rustls.
- **rust-openssl, `PKey::hmac(b"...")` (1, only second).** Added: it loads a key from bytes, so related-crypto-material with no `material_type`.
- **rust-openssl, AES decrypt in `private_key_from_pem_passphrase` (1, only second).** Not added: the call names no cipher, and the Python pilot annotated pycryptodome `RSA.import_key(passphrase=...)` as a key parse only.
- **rust-ring, HMAC key size (1, field `parameters`).** Added `key_size` 256: `hmac::Key::generate` always draws a key as long as the digest output.
- **rust-ring, `generate_pkcs8` (2, only second).** Added: the call generates the key pair and returns it serialized as PKCS#8, so it is a key generation and a serialization, and the serialization joins the `from_pkcs8` parse in one object.
- **rust-ring, `from_pkcs8` lines (1 only first, 1 only second).** The first annotation ran to the chained `.map_err` on the next line; the call closes on its own line.
- **rust-ring, `UnparsedPublicKey::new` and `verify` (3 only first, 8 only second).** The first annotation made `UnparsedPublicKey::new(&ALG, bytes)` a verifier object and `verify` a follow-up. The second made it a public-key load and `verify` the asset. Neither stands: the wrapper stores bytes with an algorithm and parses nothing (the type is named unparsed), so it is not an asset, and `verify` is a verify call on a key. So the three verify calls from the second annotation were added, and the first annotation's three verifier assets and the second's five key loads were dropped. The algorithm the wrapper was given is not carried into `verify`, by the guide's rule on digests bound to a key.
- **rust-rustls-webpki, `SystemRandom::new()` and `from_pkcs8` as a signer (2, only second).** Not added. `SystemRandom::new()` draws nothing, and the ring entry's own second annotator decided the same. `from_pkcs8` is a key parse only, as both ring annotations decided.

### Questions the guide leaves open

The guide was outside this pass, so these were left for a later revision of it rather than written into it. Each was decided as above and recorded in the assets' notes:

- a parameter fixed by the type of the key a call runs on is one the API fixes;
- a wrapper that stores key bytes and parses nothing is not an asset;
- a call that generates a key and returns it serialized is two assets;
- creating a TLS connection or session object on a context is a protocol asset of its own;
- a one-shot cipher function that always pads fixes its padding;
- an operator that performs a key agreement (curve25519-dalek `scalar * point`) is a call.

## Suite A, Go (2026-11-04)

The 16 Go entries cover the 15 frozen Go packages and the standard library's crypto (`fixtures/go/*`, 46 programs). The first annotation was written by the author of the programs, under the current guide, with `object` ids. For the second, fresh agents got only an entry's programs and `go.mod`, the guide, the schema (`internal/groundtruth`) and `internal/algo`, copied into a directory with no ground truth in it. They could read the library source in the Go module cache. Seven agents did the work: one each for the standard library, golang.org/x/crypto/ssh, golang-jwt v5, dgrijalva/jwt-go and go-ethereum/crypto, and three for the 11 small x/crypto entries, three or four entries each, written to separate files. Assets are matched as in the Python pilot. The last column compares `object` grouping over the assets both annotations list, as in the Python objects pass.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement | Objects agreed |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go-dgrijalva-jwt-go | 7 | 7 | 6 | 1 | 0 | 0 | 0.857 | 7 / 7 |
| go-go-ethereum-crypto | 11 | 11 | 11 | 0 | 0 | 0 | 1.000 | 11 / 11 |
| go-golang-jwt-v5 | 13 | 13 | 13 | 0 | 0 | 0 | 1.000 | 13 / 13 |
| go-stdlib | 84 | 84 | 82 | 2 | 0 | 0 | 0.976 | 84 / 84 |
| go-x-crypto-acme-autocert | 2 | 3 | 2 | 0 | 0 | 1 | 0.667 | 2 / 2 |
| go-x-crypto-bcrypt | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 | 2 / 2 |
| go-x-crypto-blake2b | 5 | 5 | 5 | 0 | 0 | 0 | 1.000 | 5 / 5 |
| go-x-crypto-chacha20poly1305 | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 | 8 / 8 |
| go-x-crypto-curve25519 | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 | 8 / 8 |
| go-x-crypto-hkdf | 9 | 9 | 9 | 0 | 0 | 0 | 1.000 | 9 / 9 |
| go-x-crypto-pbkdf2 | 5 | 5 | 5 | 0 | 0 | 0 | 1.000 | 5 / 5 |
| go-x-crypto-ripemd160 | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 | 2 / 2 |
| go-x-crypto-scrypt | 3 | 3 | 3 | 0 | 0 | 0 | 1.000 | 3 / 3 |
| go-x-crypto-sha3 | 4 | 4 | 3 | 1 | 0 | 0 | 0.750 | 4 / 4 |
| go-x-crypto-ssh | 14 | 17 | 12 | 0 | 2 | 5 | 0.632 | 12 / 12 |
| go-x-crypto-ssh-agent | 3 | 3 | 3 | 0 | 0 | 0 | 1.000 | 3 / 3 |
| **Total** | **180** | **184** | **174** | **4** | **2** | **6** | **0.935** | **178 / 178** |

These counts correct a line-numbering slip in two second annotations. In go-ethereum `key-hex/main.go` and dgrijalva `rs256/main.go`, every line is 28 higher than the file's, because the agent numbered lines across a concatenated listing, as the truststore annotator did in the Python pilot. The calls, types and fields are otherwise as listed. Without the correction, overall agreement is 165 of 196 (0.842).

### Resolutions

Of the 12 disagreements, 4 changed the ground truth and 8 kept the first annotation. Each resolution is in the asset's `notes`, prefixed `Re-derivation:`.

- **dgrijalva `jwt.Parse` with an RSA key (1, field `parameters`).** The first set padding PKCS1v15, because the program's key function accepts only `*jwt.SigningMethodRSA`. The second left it out. Changed: the restriction is in the program's key function, not fixed by the `Parse` API, as python-joserfc's `jwt.decode` was decided.
- **stdlib `aes.NewCipher` before `cipher.NewGCM` (1, field `primitive`).** The first said `ae`, the primitive of the GCM it is used under; the second said `block-cipher`. Changed: the call creates a block cipher with no mode, and the AEAD is the separate `NewGCM` asset. The score does not compare primitives.
- **SHAKE256 in stdlib `crypto/sha3` and x/crypto/sha3 (2, field `primitive`).** Both second annotators chose `xof`, which the schema lists and the guide does not mention. Changed to `xof`, the CycloneDX primitive for an extendable-output function.
- **x/crypto/ssh config literals (2, only first).** The second left out the `ssh.ServerConfig` and `ssh.ClientConfig` literals. Kept: the stdlib annotator, independently, annotated `tls.Config` literals as TLS contexts, and an SSH config is the same kind of context.
- **x/crypto/ssh `signer.Sign` (1, only second).** The second counted it as a sign call on a key, since `ssh.Signer` is the library's private-key type. Kept out as a follow-up on the signer that `NewSignerFromKey` creates, by the pycryptodome `DSS.new(...).sign` precedent.
- **x/crypto/ssh passphrase protection (4, only second).** The second added bcrypt-pbkdf and AES-256-CTR assets at `MarshalPrivateKeyWithPassphrase` and `ParsePrivateKeyWithPassphrase`, citing pycryptodome `export_key(protection=...)`. Not added: `export_key` names its scheme in the call, while these calls name none, and python-paramiko's password-protected `write_private_key_file` was annotated as the serialization only.
- **autocert `&acme.Client{Key: accountKey}` (1, only second).** The second listed it as a signer object built on the account key. Not added: the client signs nothing in this program, and attaching a key to a client is configuration, like `keyring.Add` in ssh/agent.

### Questions the guide leaves open

Both annotations decided these the same way, but each needed a reading of the guide for Go. They are recorded here and in the assets' `notes`, and they are the guide's open points for the next language:

- **A Go composite literal as an asset.** Go creates a TLS or SSH context with a struct literal, not a call. Both annotate `&tls.Config{...}` as a `protocol` asset on the literal's lines, and a session created from it (`tls.Server`, `tls.Client`, `ssh.NewServerConn`, `ssh.NewClientConn`) as another.
- **`make` then `rand.Read`.** A buffer allocated by `make` and filled by `rand.Read` takes the `rand.Read` line, for the draw and for the IV, nonce or salt it becomes. Passing `rand.Reader` to a call is not a draw.
- **A hash constructor passed as a value.** `sha256.New` passed to `hmac.New`, `hkdf.Key` or `pbkdf2.Key`, and `sha256.New()` nested in `rsa.EncryptOAEP`, are the digest parameter of the enclosing asset, not hash assets.
- **Two-step JWT signing.** `jwt.NewWithClaims(method, claims)` builds a token and computes nothing; `token.SignedString(key)` signs with the key and is the asset, with the parameters the token's method fixes. The two calls sit on adjacent lines in every program, within the line tolerance.
- **A block and its mode.** `aes.NewCipher` and `cipher.NewGCM` or `cipher.NewCBCEncrypter` are separate assets of one object; `Seal`, `Open` and `CryptBlocks` are follow-ups.
- **ML-KEM.** `Encapsulate` and `Decapsulate` return a shared secret, annotated like a key agreement's.
- **Keccak-256.** go-ethereum's `Keccak256` and x/crypto's `NewLegacyKeccak256` use Keccak's original padding, not SHA3-256's, so both annotations wrote `Keccak-256`, which `internal/algo` does not list and so keeps its name.

## Suite A, Java (2026-11-04)

The 16 Java entries (`fixtures/java/*`, one Maven project per library in the frozen list plus JCA/JCE) were annotated under the current guide: specific algorithm names, `hash`, related crypto material with `material_type`, and `object`. The first annotation was written by the authors of the programs. For the second, one fresh agent per entry got only the entry's `pom.xml` and sources, the guide, the schema (`internal/groundtruth`) and `internal/algo`, plus the library jars in the local Maven repository, in a directory with no ground truth in it. It annotated assets and objects in one pass. Counting is as in the Python pilot above. The table compares the two annotations as written, before any resolution.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|
| java-bcpg-jdk15on | 20 | 22 | 12 | 3 | 5 | 7 | 0.444 |
| java-bcpkix-jdk15on | 24 | 24 | 20 | 0 | 4 | 4 | 0.714 |
| java-bcprov-jdk15on | 47 | 46 | 41 | 2 | 4 | 3 | 0.820 |
| java-bcrypt | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |
| java-cryptography-core-jvm | 11 | 11 | 11 | 0 | 0 | 0 | 1.000 |
| java-eddsa | 13 | 17 | 13 | 0 | 0 | 4 | 0.765 |
| java-jasypt | 9 | 10 | 9 | 0 | 0 | 1 | 0.900 |
| java-jbcrypt | 5 | 4 | 3 | 1 | 1 | 0 | 0.600 |
| java-jjwt-api | 20 | 20 | 13 | 0 | 7 | 7 | 0.481 |
| java-jose4j | 22 | 20 | 20 | 0 | 2 | 0 | 0.909 |
| java-jsch | 9 | 9 | 8 | 1 | 0 | 0 | 0.889 |
| java-nimbus-jose-jwt | 38 | 38 | 37 | 0 | 1 | 1 | 0.949 |
| java-scrypt | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| java-spring-security-crypto | 14 | 14 | 14 | 0 | 0 | 0 | 1.000 |
| java-stdlib | 55 | 56 | 51 | 3 | 1 | 2 | 0.895 |
| java-xmlsec | 10 | 11 | 6 | 0 | 4 | 5 | 0.400 |
| **Total** | **311** | **316** | **272** | **10** | **29** | **34** | **0.788** |

Both annotations listed 282 calls at the same file, lines and asset type, and agreed on the `primitive` of 280 of them (99.3%). The largest single source of disagreement is where a JCA key generation sits: 11 key generations (22 of the 73 differences) were put on `KeyPairGenerator.getInstance` or `KeyGenerator.getInstance` by the second annotation and on `generateKeyPair()` or `generateKey()` by the first. After the resolutions below, the ground truth has 310 assets in 182 objects, and it agrees with the second annotation on 281 of 336 (0.836); the differences left are the kept first-annotation decisions.

### Java decisions

The guide's examples are Python and JavaScript. The first annotation applied them to Java with these decisions, which the notes cite as "annotation log, Java decisions":

- **Engine objects that name an algorithm** (`Cipher`, `Mac`, `MessageDigest`, `Signature`, `KeyAgreement`, `SecretKeyFactory`, and their Bouncy Castle lightweight counterparts such as `new GCMBlockCipher(new AESEngine())` or `new HMac(new SHA256Digest())`): the asset is the call that creates the object. `init`, `update`, `doFinal`, `digest`, `sign`, `verify`, `doPhase` and `generateSecret` are follow-ups. A nested engine or digest is part of the asset it is nested in.
- **Key generation** (`KeyGenerator`, `KeyPairGenerator`, Bouncy Castle key pair generators): the asset is the generating call, `generateKey()` or `generateKeyPair()`, by the guide's "generating a key or key pair". Its algorithm comes from the generator's `getInstance`, and `key_size` or `curve` from `init` or `initialize` on the same generator in the same method. The generator's `getInstance` is not a separate asset. The same holds for a generator constructed with no algorithm (Bouncy Castle `Argon2BytesGenerator`): the asset is the call that derives.
- **Transformations.** `AES/CBC/PKCS5Padding` is `AES` with mode `CBC` and padding `PKCS7`. A mode that admits only no padding (GCM, CTR, CCM) gets no `padding`, by the guide's rule for a parameter with one admissible value. `RSA/ECB/PKCS1Padding` is `RSA` with padding `PKCS1v15` and no mode. A key size is not carried from the key generation to the cipher.
- **Random draws.** `new SecureRandom()` and a generator handed to a library draw nothing and are not assets. Each `nextBytes(...)`, `generateSeed(...)` or library random helper is a `drbg` asset, and a draw that becomes an IV, nonce, salt or seed is also the material, on the `nextBytes` line.
- **Material specs.** `new GCMParameterSpec(...)`, `new IvParameterSpec(...)` and `new PBEKeySpec(...)` pass material on and are not assets. Material takes its type from the parameter class: an `IvParameterSpec` value is an `initialization-vector`, also for ChaCha20-Poly1305.
- **Keys and certificates as bytes.** `new SecretKeySpec(bytes, alg)`, Bouncy Castle `new KeyParameter(bytes)`, jose4j `new HmacKey(bytes)` and `KeyFactory` `generatePublic`/`generatePrivate` parse a key. `key.getEncoded()` serializes one, but `getEncoded()` on a KDF's output reads that output and is not an asset. `KeyFactory.getInstance` and `CertificateFactory.getInstance` compute nothing; `generateCertificate` parses a certificate.
- **Key agreement.** `KeyAgreement.getInstance` (or `new X25519Agreement()`) is the `key-agree` asset; `generateSecret()` (or `calculateAgreement`) returns the shared secret and is that material only.
- **TLS and SSH.** `SSLContext.getInstance(...)` and JSch `getSession(...)` are protocol assets; `init`, `createSSLEngine` and `setConfig` configure them.
- **Fluent chains** span from the chain root to the closing parenthesis of the call that performs the operation; trailing accessors such as `.getPayload()` are not part of it.

### Resolutions

Of the 73 differences, 18 changed the ground truth and 55 kept the first annotation. Each resolution is in the asset's `notes`, prefixed `Re-derivation:`.

- **jjwt parse chains (7 only first, 7 only second). Changed.** The first annotation ended the verify and decrypt chains at the trailing `.getPayload().getSubject()`; the second ended them at the parse call, which the guide's line rule supports.
- **xmlsec verify (1 and 1). Changed** from `new XMLSignature(sigElement, "")`, which names no algorithm, to `checkSignatureValue`, which verifies.
- **stdlib PBKDF2 `getEncoded()` (1 only first). Changed:** it reads the KDF output.
- **jsch key generation primitive (1 field). Changed** to none: the program never connects, so the key is only stored and loaded.
- **JCA key generation placement (11 and 11). Kept** on the generating call, as above.
- **Key agreement (4 only second, stdlib and bcprov). Kept:** the second annotation also made `generateSecret()` and `calculateAgreement` key-agree assets, by the python-ecdsa precedent; the creating call already is the key-agree asset.
- **GCM `padding: none` (4 field). Kept out,** by the one-admissible-value rule.
- **eddsa nested `MessageDigest.getInstance("SHA-512")` (4 only second). Kept out:** it is a parameter of the Ed25519 signer it is nested in.
- **Key objects from raw bytes** (bcprov `KeyParameter` 3, jose4j `HmacKey` 1, only first). **Kept,** as the counterparts of `SecretKeySpec`, which both annotations of stdlib list.
- **One-offs kept:** the ChaCha20-Poly1305 `IvParameterSpec` value as an IV; jbcrypt `gensalt` as a draw and a salt (python-bcrypt precedent); jose4j `setGeneratedJwtId()` as a draw; the Argon2id asset on `generateBytes`; jasypt `new RandomIvGenerator()` and bcpg `new SecureRandom()` passed to a builder left out; an OpenPGP key ring serialized as keys, not as a certificate; the extra SHA-256 that xmlsec `checkSignatureValue` recomputes left out; `OpenPGP-S2K` against `S2K` (neither is in `internal/algo`).

### Objects

Objects were compared on the 290 assets paired by location whose whole object is paired in both annotations: the two annotations grouped all 257 of them the same way. The other 33 cannot be compared, because one asset of their object sits on a different line in the other annotation (mostly the key generation placement above). The ground truth has 64 multi-asset objects.

### Questions the guide leaves open

Each was decided as above and recorded in `notes`; each would be settled by one guide sentence, which this pass does not write:

- where a key generation that JCA splits over `getInstance`, `initialize` and `generateKeyPair` sits, and whether `initialize(3072)` states its key size;
- whether a call that returns a key agreement's shared secret is also the agreement;
- whether wrapping raw bytes in a key object (`SecretKeySpec`, `KeyParameter`, `HmacKey`) is parsing a key;
- whether a generator object handed to a library (`new SecureRandom()`, `new RandomIvGenerator()`) is a draw;
- whether an OpenPGP key ring is a certificate;
- cryptography-kotlin's AES-GCM, ECDSA and RSA cannot be called from Java source: their parameters are Kotlin value classes, so the methods compile to mangled names. The fixture covers ChaCha20-Poly1305, AES-CBC, HMAC and SHA-256 instead.

After the tool runs, one ground-truth mistake was fixed: java-bcprov-jdk15on `ProviderAesGcm.java:21` said `nonce` for the value passed to `GCMParameterSpec`, which is an `initialization-vector` by the Java decisions above. No score changed.

## Suite A, anchors (2026-11-05)

Multi-anchor assets were introduced ([annotation guide](annotation-guide.md#anchors), [METHODOLOGY.md](../METHODOLOGY.md#matching-rules) change log). One annotator went through every `algorithm` asset of the five languages that is a key generation, a JWT, JWS or JWE call or a random draw, and a second pass over every `algorithm` asset with a naming call in the eight lines before it added the jose4j JWE decrypt (gt-0004, gt-0005), starting from class A of [strict-matching.md](strict-matching.md), and read the lines before it in the same function for a call that names the asset's algorithm on the same object. No asset's `lines` changed; 75 assets in 14 entries gained an anchor. Python has none: every Python key generation, JWT call and draw names its algorithm on its own call, and the two candidates (a `padding.OAEP(...)` specification object, a literal list of JOSE algorithm names) are arguments, not calls on the object. The candidates left out on purpose are the calls on another object: the `SystemRandom::new()` or `new SecureRandom()` handed to a key generation, the peer's `UnparsedPublicKey::new(&X25519, ...)` in ring's key agreement, `new PBEKeySpec(...)`, and `new JSch()`. jasypt's `setAlgorithm` and Bouncy Castle's `new Argon2BytesGenerator()` also name an algorithm on the same object, but they are a cipher, a digester and a KDF, outside the key generation, JWT and random-draw sweep; where the Argon2 asset sits is still an open question.

The same day Fernet's `decrypt` became its own operation: python-cryptography gains gt-0052, `fernet/main.py:9`, in the object `fernet/key`.

| Entry | Asset | Lines | Anchors (line: call) |
|---|---|---|---|
| go-dgrijalva-jwt-go | gt-0001 | `main.go:13` | 12: `token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClai...` |
| go-dgrijalva-jwt-go | gt-0006 | `main.go:25` | 24: `token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"s...` |
| go-golang-jwt-v5 | gt-0001 | `main.go:13` | 12: `token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"s...` |
| go-golang-jwt-v5 | gt-0006 | `main.go:27` | 26: `token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)` |
| go-golang-jwt-v5 | gt-0009 | `main.go:18` | 17: `token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{"s...` |
| go-golang-jwt-v5 | gt-0012 | `main.go:17` | 16: `token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{"s...` |
| go-stdlib | gt-0013 | `main.go:22` | 21: `curve := ecdh.X25519()` |
| go-stdlib | gt-0014 | `main.go:26` | 21: `curve := ecdh.X25519()` |
| java-bcpg-jdk15on | gt-0001 | `PgpEncrypt.java:32` | 30: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcpg-jdk15on | gt-0006 | `PgpKeyRing.java:30` | 28: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcpg-jdk15on | gt-0007 | `PgpKeyRing.java:31` | 28: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcpg-jdk15on | gt-0018 | `PgpSign.java:25` | 23: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcpkix-jdk15on | gt-0001 | `CmsSignedData.java:31` | 29: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcpkix-jdk15on | gt-0006 | `Csr.java:21` | 19: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC", "B...`; 20: `keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));` |
| java-bcpkix-jdk15on | gt-0010 | `PemRoundTrip.java:30` | 28: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC", "B...`; 29: `keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));` |
| java-bcpkix-jdk15on | gt-0020 | `SelfSignedCert.java:25` | 23: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA", "...` |
| java-bcprov-jdk15on | gt-0001 | `AesCbc.java:16` | 14: `SecureRandom random = new SecureRandom();` |
| java-bcprov-jdk15on | gt-0003 | `AesCbc.java:19` | 14: `SecureRandom random = new SecureRandom();` |
| java-bcprov-jdk15on | gt-0007 | `AesGcm.java:14` | 12: `SecureRandom random = new SecureRandom();` |
| java-bcprov-jdk15on | gt-0009 | `AesGcm.java:17` | 12: `SecureRandom random = new SecureRandom();` |
| java-bcprov-jdk15on | gt-0016 | `EcdsaP256.java:23` | 21: `ECKeyPairGenerator generator = new ECKeyPairGenerator();` |
| java-bcprov-jdk15on | gt-0020 | `Ed25519Sign.java:14` | 12: `Ed25519KeyPairGenerator generator = new Ed25519KeyPairGenerator();` |
| java-bcprov-jdk15on | gt-0029 | `ProviderAesGcm.java:18` | 16: `KeyGenerator keyGen = KeyGenerator.getInstance("AES", "BC");` |
| java-bcprov-jdk15on | gt-0035 | `RsaOaep.java:17` | 15: `RSAKeyPairGenerator generator = new RSAKeyPairGenerator();` |
| java-bcprov-jdk15on | gt-0042 | `X25519Agree.java:15` | 13: `X25519KeyPairGenerator generator = new X25519KeyPairGenerator();` |
| java-bcprov-jdk15on | gt-0043 | `X25519Agree.java:16` | 13: `X25519KeyPairGenerator generator = new X25519KeyPairGenerator();` |
| java-eddsa | gt-0001 | `KeyEncoding.java:14` | 13: `KeyPairGenerator keyPairGen = new KeyPairGenerator();` |
| java-eddsa | gt-0011 | `SignVerify.java:13` | 12: `KeyPairGenerator keyPairGen = new KeyPairGenerator();` |
| java-jose4j | gt-0002 | `JweRsaOaep.java:20` | 17: `jwe.setAlgorithmHeaderValue(KeyManagementAlgorithmIdentifiers.RSA_O...` |
| java-jose4j | gt-0003 | `JweRsaOaep.java:20` | 18: `jwe.setEncryptionMethodHeaderParameter(ContentEncryptionAlgorithmId...` |
| java-jose4j | gt-0004 | `JweRsaOaep.java:27` | 23: `received.setAlgorithmConstraints(new AlgorithmConstraints(ConstraintType...` |
| java-jose4j | gt-0005 | `JweRsaOaep.java:27` | 24: `received.setContentEncryptionAlgorithmConstraints(new AlgorithmConstr...` |
| java-jose4j | gt-0014 | `JwsHs256.java:21` | 19: `jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.HMAC_SHA256);` |
| java-jose4j | gt-0015 | `JwsHs256.java:27` | 24: `received.setAlgorithmConstraints(new AlgorithmConstraints(Constrain...` |
| java-jose4j | gt-0017 | `JwsRs256.java:20` | 17: `jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.RSA_USING_SHA256);` |
| java-jose4j | gt-0018 | `JwsRs256.java:26` | 23: `received.setAlgorithmConstraints(new AlgorithmConstraints(Constrain...` |
| java-jose4j | gt-0021 | `JwtEs256.java:31` | 30: `jws.setAlgorithmHeaderValue(AlgorithmIdentifiers.ECDSA_USING_P256_C...` |
| java-jose4j | gt-0022 | `JwtEs256.java:42` | 40: `.setJwsAlgorithmConstraints(ConstraintType.PERMIT, AlgorithmIdentif...` |
| java-nimbus-jose-jwt | gt-0001 | `JweDirect.java:17` | 15: `KeyGenerator keyGen = KeyGenerator.getInstance("AES");` |
| java-nimbus-jose-jwt | gt-0003 | `JweDirect.java:23` | 20: `new JWEHeader(JWEAlgorithm.DIR, EncryptionMethod.A256GCM),` |
| java-nimbus-jose-jwt | gt-0008 | `JweRsaOaep.java:22` | 19: `JWEHeader header = new JWEHeader(JWEAlgorithm.RSA_OAEP_256, Encrypt...` |
| java-nimbus-jose-jwt | gt-0009 | `JweRsaOaep.java:22` | 19: `JWEHeader header = new JWEHeader(JWEAlgorithm.RSA_OAEP_256, Encrypt...` |
| java-nimbus-jose-jwt | gt-0021 | `JwsEd25519.java:26` | 24: `new JWSHeader.Builder(JWSAlgorithm.EdDSA).keyID(jwk.getKeyID()).bui...` |
| java-nimbus-jose-jwt | gt-0026 | `JwsEs256.java:26` | 24: `new JWSHeader.Builder(JWSAlgorithm.ES256).keyID(ecJWK.getKeyID()).b...` |
| java-nimbus-jose-jwt | gt-0029 | `JwsHs256.java:17` | 15: `SecureRandom random = new SecureRandom();` |
| java-nimbus-jose-jwt | gt-0031 | `JwsHs256.java:21` | 20: `JWSObject jwsObject = new JWSObject(new JWSHeader(JWSAlgorithm.HS25...` |
| java-nimbus-jose-jwt | gt-0036 | `JwtRs256.java:31` | 29: `new JWSHeader.Builder(JWSAlgorithm.RS256).keyID(rsaJWK.getKeyID())....` |
| java-spring-security-crypto | gt-0005 | `RandomKey.java:10` | 9: `BytesKeyGenerator generator = KeyGenerators.secureRandom(16);` |
| java-stdlib | gt-0001 | `AesCbc.java:14` | 12: `KeyGenerator keyGen = KeyGenerator.getInstance("AES");` |
| java-stdlib | gt-0006 | `AesGcm.java:14` | 12: `KeyGenerator keyGen = KeyGenerator.getInstance("AES");` |
| java-stdlib | gt-0007 | `AesGcm.java:18` | 16: `SecureRandom random = new SecureRandom();` |
| java-stdlib | gt-0014 | `ChaCha20Poly1305.java:14` | 12: `KeyGenerator keyGen = KeyGenerator.getInstance("ChaCha20");` |
| java-stdlib | gt-0019 | `EcdsaP256.java:13` | 11: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC");`; 12: `keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));` |
| java-stdlib | gt-0022 | `Ed25519Signature.java:11` | 10: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("Ed25519...` |
| java-stdlib | gt-0025 | `HmacSha256.java:15` | 14: `KeyGenerator keyGen = KeyGenerator.getInstance("HmacSHA256");` |
| java-stdlib | gt-0028 | `KeySerialization.java:16` | 14: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("EC");`; 15: `keyPairGen.initialize(new ECGenParameterSpec("secp256r1"));` |
| java-stdlib | gt-0038 | `RandomBytes.java:10` | 8: `SecureRandom random = new SecureRandom();` |
| java-stdlib | gt-0039 | `RandomBytes.java:14` | 12: `SecureRandom strong = SecureRandom.getInstanceStrong();` |
| java-stdlib | gt-0040 | `RandomBytes.java:16` | 8: `SecureRandom random = new SecureRandom();` |
| java-stdlib | gt-0041 | `RsaOaep.java:12` | 10: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA");` |
| java-stdlib | gt-0044 | `RsaSignature.java:12` | 10: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("RSA");` |
| java-stdlib | gt-0050 | `X25519Agreement.java:11` | 10: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("X25519");` |
| java-stdlib | gt-0051 | `X25519Agreement.java:12` | 10: `KeyPairGenerator keyPairGen = KeyPairGenerator.getInstance("X25519");` |
| java-xmlsec | gt-0001 | `XmlEncryption.java:27` | 25: `KeyGenerator keyGenerator = KeyGenerator.getInstance("AES");` |
| java-xmlsec | gt-0002 | `XmlEncryption.java:31` | 29: `KeyPairGenerator kpg = KeyPairGenerator.getInstance("RSA");` |
| java-xmlsec | gt-0007 | `XmlSignature.java:24` | 22: `KeyPairGenerator kpg = KeyPairGenerator.getInstance("RSA");` |
| javascript-node-forge | gt-0016 | `main.js:6` | 5: `const hmac = forge.hmac.create();` |
| rust-ring | gt-0009 | `pbkdf2.rs:8` | 6: `let rng = rand::SystemRandom::new();` |
| rust-ring | gt-0013 | `aead_aes256gcm.rs:7` | 5: `let rng = SystemRandom::new();` |
| rust-ring | gt-0015 | `aead_aes256gcm.rs:11` | 5: `let rng = SystemRandom::new();` |
| rust-ring | gt-0025 | `ed25519.rs:13` | 12: `let peer_public_key = signature::UnparsedPublicKey::new(&signature:...` |
| rust-ring | gt-0030 | `ecdsa_p256.rs:14` | 13: `let public_key = signature::UnparsedPublicKey::new(&ECDSA_P256_SHA2...` |
| rust-ring | gt-0033 | `rsa_pkcs1.rs:14` | 13: `let public_key = signature::UnparsedPublicKey::new(&signature::RSA_...` |
| rust-ring | gt-0040 | `random.rs:6` | 4: `let rng = SystemRandom::new();` |
| rust-ring | gt-0041 | `random.rs:7` | 4: `let rng = SystemRandom::new();` |

## Suite B, all five languages (2026-11-05)

The first annotation is the ground truth the fixture author wrote. The second came from one fresh agent that got only the fixture sources (the libraries and both apps of every scenario), the annotation guide and the `internal/algo` tables, and was told what an entry point is and what dependency distance counts. It never saw the ground truth. Each pair has the same assets in both variants, so the counts are per scenario: one row per language, over the seven scenarios. Reachability in both variants, distance and the call path were compared as fields too.

| Language | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|
| go | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |
| python | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |
| java | 10 | 8 | 8 | 0 | 2 | 0 | 0.800 |
| javascript | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |
| rust | 8 | 8 | 8 | 0 | 0 | 0 | 1.000 |

### Resolutions

- **Java `new SecretKeySpec(KEY, "HmacSHA256")`** (deadfn libb `Codec.java` line 16, depth3 libc `Codec.java` line 15). The second annotation left it out because it wraps a constant key. Kept, by the Java decision above that `new SecretKeySpec(...)` makes a key from bytes and is a key parse. The notes of both assets say so.

### Questions the guide leaves open

- **Distance in Go.** A Go module lists every module its build needs in its `go.mod`, the ones it does not import as `// indirect`. In the depth3 scenario liba's `go.mod` lists libc as indirect, so the requirement graph has an edge liba -> libc that the import graph does not. Both annotations count distance along the chain of imports (app -> liba -> libb -> libc, distance 3). The guide now says so ([Suite B](annotation-guide.md#suite-b)).

## Suite B, handler and callback (2026-11-05)

The 20 entries of the two scenarios added after the first suite B results: handler (net/http, gin, Flask, FastAPI, Spring Boot, Express, axum, actix-web) and callback (Go, Python). The first annotation is the ground truth the fixture author wrote. The second came from one fresh agent that got only the fixture sources of the 10 pairs, the guide as it stood before this pass, the schema (`internal/groundtruth`) and `internal/algo`, in a directory with no ground truth in it, and was told how each library is installed. It never saw the ground truth. Counts are per entry; reachability, distance and depth were compared as fields, and the call path by the functions it names.

| Language | Entries | First | Second | Agreed | Field | Only first | Only second | Agreement |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| go | 6 | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| python | 6 | 6 | 6 | 6 | 0 | 0 | 0 | 1.000 |
| java | 2 | 2 | 2 | 2 | 2 | 0 | 0 | 1.000 |
| javascript | 2 | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 |
| rust | 4 | 4 | 4 | 4 | 0 | 0 | 0 | 1.000 |

Both annotations give every entry one asset, SHA-256 in libb at distance 2, reachable in each reachable variant and unreachable in each unreachable variant.

### Resolutions

- **Java line** (`Digests.java`, both variants). The second annotation put the asset on line 23; the file has 14 lines and `MessageDigest.getInstance("SHA-256")` is line 11. Both meant the same call, so line 11 stands, as in the other Java suite B entries.
- **Path names.** The second annotation wrote short names (`main`, `fingerprint`); the first writes them qualified, as in the other suite B entries. They name the same functions in the same order. The path is not scored.

### Questions the guide leaves open

- **Functions handed to a framework or a library.** The guide said the entry point reaches what it runs on some path, and that an application function the entry point never calls is not an entry point, but not whether a handler it registers or a callback it passes is reached. Both annotations decided yes, because the framework or library function calls it while the program runs, and no when it is never handed over. The guide now says so ([Suite B](annotation-guide.md#suite-b)).

## Suite A, generated private keys (2026-11-08)

A key generation that returns a private key became also related crypto material of type `private-key` ([guide](annotation-guide.md#related-crypto-material)). Every suite A `algorithm` asset was read against the new rule, and one `related-crypto-material` asset with `material_type: private-key` was added on the lines of each asymmetric key generation: 150 assets (Python 33, JavaScript 41, Rust 15, Go 22, Java 39). Suite B has no key generation. The second annotator's role was taken by the rule itself: each added asset copies `file`, `lines`, `origin`, `reachable` and `depth` from the key generation it sits on, so there is nothing for two annotations to disagree on but which calls count.

### Resolutions

- **Secret keys** (`AESGCM.generate_key`, `ChaCha20Poly1305.generate_key`, `Fernet.generate_key`, JCA `KeyGenerator.generateKey`, `OctKey.generate_key`, ring `hmac::Key::generate`, jose `generateSecret`, WebCrypto AES `generateKey`). Not private keys, so no material.
- **A key pair kept in an object** (Node `createECDH(...).generateKeys()`, python-ecdsa `ECDH().generate_private_key()`). The call returns only the public key, so no material, by the rule on material a library keeps inside its own object.
- **Keys from a seed or from bytes** (`new EdDSAPrivateKeySpec(seed, curve)`). Not a generation, so no material.
- **Calls that generate a key pair among other work** (`selfsigned.generate`, which returns the private key's PEM with the certificate; `new JcaPGPKeyPair(..., gen.generateKeyPair(), ...)`; `add_server_key(paramiko.RSAKey.generate(...))`). Material on the lines of the call, as for the algorithm asset.
- **curve25519-dalek `Scalar::random`** (`ristretto_dh.rs` lines 8 and 9). The ground truth already counts it as the key generation of the ECDH secret, so it has the material too.

## Suite C, `files` scope (2026-11-13)

Two fresh agents annotated the 142 sampled files of the ten suite C repositories independently, each with only the checkout at the pinned commit, the sample list, the annotation guide, the schema and the `internal/algo` tables. Neither saw the other's annotation, and no tool was run on a suite C repository and no tool output was read, before or during adjudication. Matching is as in the [Python pilot](#suite-a-python-pilot-2026-11-04): same file, overlapping lines and same asset type, then compared on lines, anchors, primitive, algorithm, parameters and material type. Objects were compared as groupings of the matched assets, and every grouping agreed.

| Entry | First | Second | Agreed | Field | Only first | Only second | Agreement | Final |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| java-c-mall | 1 | 1 | 1 | 0 | 0 | 0 | 1.000 | 1 |
| java-c-jadx | 2 | 2 | 2 | 0 | 0 | 0 | 1.000 | 2 |
| python-c-yt-dlp | 27 | 45 | 27 | 0 | 0 | 18 | 0.600 | 45 |
| python-c-graphify | 16 | 16 | 16 | 0 | 0 | 0 | 1.000 | 16 |
| go-c-ollama | 31 | 31 | 30 | 1 | 0 | 0 | 0.968 | 31 |
| go-c-caveman | 29 | 30 | 29 | 0 | 0 | 1 | 0.967 | 30 |
| rust-c-clash-verge-rev | 9 | 8 | 8 | 0 | 1 | 0 | 0.889 | 9 |
| rust-c-cc-switch | 26 | 21 | 21 | 0 | 5 | 0 | 0.808 | 23 |
| javascript-c-freecodecamp | 21 | 21 | 20 | 0 | 1 | 1 | 0.909 | 21 |
| javascript-c-firecrawl | 36 | 36 | 32 | 4 | 0 | 0 | 0.889 | 36 |
| **Total** | **198** | **211** | **186** | **5** | **7** | **20** | **0.853** | **214** |

The final ground truth has 214 assets in 142 files. The line ranges of every annotated asset were checked for a call cut short (unbalanced parentheses or a chained call continuing on the next line); none was.

### Resolutions

Of the 32 disagreements, 24 were resolved by taking the asset or field value that one annotation had and 8 by leaving it out; 3 assets that neither annotation listed were added. Each resolution is in the asset's `notes`, prefixed `Adjudication:`.

- **python-c-yt-dlp, TLS sessions and loaders (18, only second).** Kept: 3 `wrap_socket`/`wrap_bio` calls as protocol assets, 6 `load_cert_chain` calls as a certificate and a key parse each (12 assets), and 3 `load_verify_locations` calls as certificate parses. The first annotation had read them as configuring the context.
- **go-c-caveman, `httptest.NewTLSServer` (1, only second).** Kept as a protocol asset.
- **rust-c-clash-verge-rev, reqwest `Client::builder().tls_backend_rustls()` (1, only first).** Removed: it selects a TLS backend for an HTTP client, which is configuration.
- **rust-c-cc-switch, `Uuid::new_v4()` (5, only first).** Removed: an identifier library is not a random draw, even where the UUID bytes feed a PKCE verifier or an OAuth state.
- **javascript-c-freecodecamp, `nanoid()` (1, only second).** Removed, for the same reason.
- **javascript-c-freecodecamp, second `jwt.sign` in `user.ts` (1, only first).** Kept; the second annotation missed it.
- **go-c-ollama, `ed25519.GenerateKey` primitive (1, field).** `signature`: the stored key is loaded and signs requests in a file outside the sample.
- **javascript-c-firecrawl, jose `generateKeyPair('RS256')` and `('RS512')` (3, field).** Padding `PKCS1v15` added, as in the suite A jose entry.
- **javascript-c-firecrawl, `SignJWT` chain with `alg: overrides.alg ?? "RS256"` (1, field).** No `hash`: one caller passes `RS512`, so only the padding is shared by every value that reaches the call.
- **Neither annotation (3 added).** tokio-rustls `connector.connect(name, stream)` creates a TLS session over a stream: one call in rust-c-clash-verge-rev and two in rust-c-cc-switch. Both annotations had read them as using an existing context.

### Questions the guide leaves open

Each is now a sentence in the [annotation guide](annotation-guide.md):

- **Project helpers.** A call to a project function that wraps crypto is not an asset at the call site; the library call inside it is, in its own file. A project's own implementation of an algorithm (yt-dlp's pure-Python AES) is annotated at its call sites.
- **Random draws and identifier libraries.** Which APIs count as a draw, and that `uuid4`, `Uuid::new_v4` and `nanoid` do not.
- **TLS sessions.** A session or listener created over a socket or stream from a context is a protocol asset; a connector without a stream and an HTTP client's TLS options are configuration.
- **Certificates and keys loaded into a context.** `load_cert_chain`, `load_verify_locations`, `load_native_certs` and `add_parsable_certificates` are parses; default trust stores are not.
- **External tools.** Crypto run through a command line is annotated at the call that names the algorithm.
- **Purpose.** Cache keys, ETags, test digests and `usedforsecurity=False` are assets; a digest output named `salt` or `cnonce` is not material.
- **Key generation primitive.** The use of the key is looked for anywhere in the repository.
- **Objects across files.** An imported secret given to calls in several files is one object.
- **Overridden algorithm literals.** Only the parameters every reaching value shares are set.
- **Reachability fields in suite C.** `reachable: unknown`, no `depth`.

## Suite C, `pool` scope (2026-11-13)

Two fresh agents judged every item of the ten `<id>-pool` entries independently. Each had the checkout at the pinned commit, the annotation guide, the schema, the `internal/algo` tables and the `files` ground truth of the entry, and saw each item only as a file, a line range and the source text, with no word on which pool source brought it. Neither saw the other's annotation, and no tool output was read before or during adjudication. Assets were matched as for the `files` scope: same file, overlapping lines and same asset type, then compared on lines, anchors, primitive, algorithm, parameters and material type; objects were compared as groupings of the matched assets. An item's verdict is `asset` when its lines overlap a ground-truth asset, which is how the score pairs a prediction with an asset.

| Entry | Items | Item verdicts differ | First | Second | Agreed | Field | Only first | Only second | Agreement | Final |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| java-c-mall | 52 | 0 | 5 | 5 | 5 | 0 | 0 | 0 | 1.000 | 4 |
| java-c-jadx | 62 | 0 | 7 | 7 | 7 | 0 | 0 | 0 | 1.000 | 7 |
| python-c-yt-dlp | 4,288 | 12 | 161 | 155 | 145 | 7 | 9 | 3 | 0.884 | 154 |
| python-c-graphify | 205 | 0 | 58 | 58 | 58 | 0 | 0 | 0 | 1.000 | 58 |
| go-c-ollama | 2,319 | 0 | 109 | 109 | 108 | 1 | 0 | 0 | 0.991 | 109 |
| go-c-caveman | 1,501 | 3 | 258 | 257 | 256 | 0 | 2 | 1 | 0.988 | 258 |
| rust-c-clash-verge-rev | 152 | 0 | 13 | 13 | 13 | 0 | 0 | 0 | 1.000 | 13 |
| rust-c-cc-switch | 975 | 0 | 43 | 43 | 43 | 0 | 0 | 0 | 1.000 | 43 |
| javascript-c-freecodecamp | 197 | 0 | 50 | 50 | 50 | 0 | 0 | 0 | 1.000 | 50 |
| javascript-c-firecrawl | 904 | 1 | 297 | 297 | 297 | 0 | 0 | 0 | 1.000 | 297 |
| **Total** | **10,655** | **16** | **1,001** | **994** | **982** | **8** | **11** | **4** | **0.977** | **993** |

The two annotations gave the same verdict on 10,639 of 10,655 items (0.998). Agreement on assets is agreed over agreed, field, only first and only second, as in the `files` table. The final ground truth has 993 assets, in `groundtruth/C/<id>-pool.yaml`, and the final verdict of every item is in `groundtruth/C/<id>-pool.verdicts.tsv`. Every pool asset in one of the 142 sampled files has the fields of its `files` ground-truth asset. The one object that differs is `ed25519.GenerateKey` in go-c-ollama's `cmd/cmd.go`: the pool covers `auth/auth.go`, which signs with the stored key, so the key generation and that signature share an object across files, where the `files` truth, with `auth/auth.go` outside its sample, has none.

### Resolutions

Of the 23 asset disagreements, 10 were resolved by taking the first annotation and 13 by taking the second. Of the 4 object-grouping disagreements, 2 went to the second annotation and 2 followed the asset they belong to (first). Two assets both annotations agreed on were changed by a ruling: one removed, one added. One range was corrected after a check of every range for a call cut short. Each resolution is in the asset's `notes`, prefixed `Adjudication:`; a removed asset's reason is in the verdict file.

- **python-c-yt-dlp, `aes_*_bytes` call sites (8 only first, 3 only second).** `aes_cbc_decrypt_bytes` and `aes_gcm_decrypt_and_verify_bytes` call pycryptodome's `AES.new`, with a pure-Python fallback, so they are helpers: their 8 call sites were removed, and their `AES.new` calls are the assets. `aes_cbc_encrypt_bytes` calls no library, only the project's own `aes_cbc_encrypt`, so it is part of the project's implementation and its call site in `tencent.py` was kept, with its range extended to line 42, where the call closes. The 3 calls the second annotation had inside `yt_dlp/aes.py` (lines 19, 23, 27) are calls between the implementation's building blocks and were removed.
- **python-c-yt-dlp, fields (7).** Padding `PKCS7` added to one `aes_cbc_encrypt` call that leaves the default padding and two `aes_ecb_encrypt` calls, which always pad with PKCS#7 (second). `CryptUnprotectData` keeps algorithm `DPAPI` with primitive `other`, not `ae` (first). The inline `pow(m, e, n)` RSA in `adn.py` takes no padding from the separate `pkcs1pad` helper (first). The inner `hashlib.md5` in `kankanews.py` ends on line 39, and the `hashlib.sha256` path step in `tiktok.py` is on line 241 alone (second).
- **python-c-yt-dlp, objects (2).** The literal keys of `test/test_aes.py` given to several of the project's AES calls are two objects, `key` and `self.key` (second).
- **python-c-yt-dlp, websockets (added).** `websockets.sync.client.connect(sock=sock, ..., ssl=ssl_ctx)` opens a TLS session over the program's socket from its context: a protocol asset, which both annotations had read as configuration.
- **go-c-ollama, `windows.CryptQueryObject` (1 field).** The call ends on line 287, and it shares a certificate object with the `x509.ParseCertificate` of the signer certificate taken from the store it returns (first).
- **go-c-caveman, `x509.Certificate.Verify` (2 only first).** Kept as `signature` assets: chain verification checks signatures. The test's CA key is the ECDSA key generated in the same test, so that verify is ECDSA and joins the key's object; the other one's CA comes from an operator file, so it has no algorithm.
- **go-c-caveman, `&rsa.PublicKey{N: n, E: e}` (1 only second).** Removed: a struct literal that assembles a key is not a call; the parse calls around it are annotated.
- **java-c-mall, `new DefaultAlipayClient(...)` (removed).** Both annotations had it as an RSA signer object. In the SDK version the repository pins, the constructor takes the sign type as an argument (RSA, RSA2 or SM2) and only stores the key string, so the library does not fix the algorithm, and the facade rule leaves it out.
- **javascript-c-firecrawl, one item.** Item 152, a line inside the `https.createServer({ key, cert }, ...)` call, overlaps that call's three assets, so its verdict is `asset`; both annotations had the same assets.

Library facades both annotations kept were checked in the published package source at the version the lockfile or build file pins: Hutool `JWTUtil.createToken` installs an HS256 signer when the token names no algorithm (hutool-all 5.8.40), `@fastify/cookie` `sign` defaults to HMAC with SHA-256 (11.0.2), Gatsby's `createContentDigest` hashes with MD5 (gatsby 5.16.1 through gatsby-core-utils 4.16.0), and tauri-plugin-updater's `Update::verify` runs minisign's Ed25519 verify (2.14.1 with minisign-verify 0.3.0).

### Known misses of the pool

These assets, by the guide, are in the repositories but no pool item covers them, so they are not in the pool ground truth: they are what pool recall cannot see. The annotators reported them, or they are `files` ground-truth assets with no pool item.

| Entry | Asset | Where |
|---|---|---|
| java-c-mall | `AlipaySignature.rsaCheckV1` one-shot verify | `mall-portal/.../service/impl/AlipayServiceImpl.java` 86 |
| python-c-yt-dlp | IV created by a literal (`files` truth) | `devscripts/generate_aes_testdata.py` 29 |
| python-c-yt-dlp | `load_verify_locations(cafile=...)` certificate parse (`files` truth) | `yt_dlp/networking/_helper.py` 25 |
| python-c-yt-dlp | IV created by a literal, passed to `aes_cbc_encrypt_bytes` | `yt_dlp/extractor/tencent.py` 41 |
| python-c-yt-dlp | zero IV created by `binascii.unhexlify('00...')`, passed to `AES.new` | `yt_dlp/extractor/tarangplus.py` 101 |
| python-c-graphify | `self._context.wrap_socket(...)` TLS session | `graphify/security.py` 229 |
| go-c-ollama | `httptest.NewTLSServer` ×5 | `server/download_test.go` 41, 144; `transfer/redirect_test.go` 317, 350, 552 |
| go-c-ollama | `windows.WinVerifyTrustEx` signature verify | `app/updater/updater_windows.go` 249 |
| go-c-ollama | crypto/rand draw in `auth.NewNonce` (`io.ReadFull` on `rand.Reader`) | `auth/auth.go` 46 |
| go-c-caveman | `httptest.NewTLSServer` ×7 (one in the `files` truth) | `proxy/internal/identity/oidc_test.go` 35, 296, 351; `shared/platform/postgresconfig/postgresconfig_test.go` 126, 207; `shared/platform/ssrf/ssrf_test.go` 785; `proxy/internal/standalone/standalone_test.go` 1770 |
| go-c-caveman | `httptest.Server.StartTLS` | `proxy/internal/gateway/reliability_test.go` 175 |
| go-c-caveman | `CertPool.AppendCertsFromPEM` certificate parse ×2 | `proxy/internal/identity/tls_test.go` 102, 251 |
| go-c-caveman | ECDSA `PublicKey.Bytes()` key serialization | `proxy/internal/identity/oidc_test.go` 56 |
| rust-c-cc-switch | `add_parsable_certificates` certificate parse (`files` truth) | `src-tauri/src/proxy/hyper_client.rs` 623 |

### Questions the guide leaves open

Each is now a sentence in the [annotation guide](annotation-guide.md):

- **A project implementation and its helpers.** Calls between the building blocks of a project's own algorithm are not assets, and a project function that calls a library's crypto is a helper even with a pure-code fallback.
- **Library facades.** A facade is an asset only when the pinned package source fixes its algorithm.
- **OS data protection.** `CryptUnprotectData` is `DPAPI` with primitive `other`.
- **Certificate chain verification.** A `signature` asset, with an algorithm only when the CA key's type is known.
- **Key literals.** A struct literal that assembles a key is not an asset.
- **TLS sessions from a client library.** `websockets` `connect(sock=..., ssl=ctx)` opens a session over a socket.
- **Salts of a plain hash.** A random value hashed as a salt is `salt` material.
- **Assets no item covers.** They stay out of a `pool` ground truth and are listed here.
