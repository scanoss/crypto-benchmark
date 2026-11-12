# Methodology

How the benchmark builds its ground truth and scores a tool against it. The rules here are pre-registered: they are written before results are published, and every later change is dated in the [change log](#change-log), applied to every tool alike, and followed by a re-score of every tool from its stored predictions.

| Topic | In short |
|---|---|
| [Ground truth](#ground-truth-process) | Two independent annotations per entry, every disagreement resolved against the [annotation guide](docs/annotation-guide.md) |
| [Suite A library list](#suite-a-library-list) | Top 15 crypto-purpose packages per ecosystem plus the standard library, frozen on 2026-11-04 |
| [Matching](#matching-rules) | Same origin, component, file and asset type, overlapping lines (`LINE_TOLERANCE = 0`), one-to-one; anchors for algorithm assets |
| [Object level](#object-level-view) | Each crypto object counts once, next to the operation-level primary view |
| [Metrics](#metrics) | Precision, recall, F1, attribution, reachability, suite B pairs |
| [Targets](#targets) | Precision >= 0.90 and recall >= 0.85 per language × asset type cell |
| [Competitors](#competitor-configuration) | The exact configuration of cdxgen and sonar-cryptography, chosen for their best fair run |

## Ground truth process

Annotation is done by agents only; there is no human annotator. Two annotators label each entry independently, from the program files and [docs/annotation-guide.md](docs/annotation-guide.md). One may be the author of the programs. The other is a fresh agent that sees only the program files, the guide, the schema and the `internal/algo` tables. The two annotations are diffed, and every disagreement is resolved by reading the code against the guide. The resolution goes into the asset's `notes`. When a resolution decides a question the guide left open, the guide gains a sentence that settles it. [docs/annotation-log.md](docs/annotation-log.md) records the agreement counts for each entry.

`scope.kind: exhaustive` commits that every asset in every file of the entry is listed. What `files` and `pool` commit to is to be written with the first suite C entries.

## Suite A library list

[corpus/libraries.yaml](corpus/libraries.yaml) is the frozen list of libraries that suite A covers, and `cbench validate` checks it (`internal/libraries`). The values were retrieved on 2026-11-04, and the list does not change after results are published. Each row records its source URL and retrieval date.

**Rule.** The list takes the top 15 crypto-purpose packages per ecosystem, ranked by one primary metric. A package is crypto-purpose when a project calls it directly for a cryptographic operation: a hash, cipher, MAC, signature, KDF or password hash, key agreement, random generation, JWT/JOSE, TLS, SSH, X.509/PKI or OpenPGP. Helpers are excluded, even when they rank inside the top 15, and are listed in `excluded` with a reason. Helpers are:

- trait-only and type-only packages;
- encoding (ASN.1, PEM, DER);
- constant-time comparison and zeroizing;
- buffer plumbing (block-buffer);
- low-level bindings that projects reach through another listed package;
- polyfills (browserify);
- packages that only carry the keyword, such as non-cryptographic hashes and OAuth frameworks.

When one library ships under two package names, it gets one row, under the name with the higher value. For example, pycryptodomex is pycryptodome under another namespace, and Bouncy Castle ships one artifact per JDK. The standard library's crypto is always included, as a `stdlib` row: Python's hashlib, hmac, ssl and secrets; node:crypto and WebCrypto; Go's crypto/*; Java's JCA/JCE.

**Candidates come from the registry, not from a hand-picked list:**

| Ecosystem | Candidate route | Primary metric |
|---|---|---|
| Python | The [top 15,000 PyPI packages](https://hugovk.github.io/top-pypi-packages/top-pypi-packages.min.json) (2026-11-01 snapshot), each checked through the PyPI JSON API for the classifier `Topic :: Security :: Cryptography`. Because many crypto packages publish no classifier, a second route matches a fixed list of crypto terms against the same packages' names and summaries. | pypistats `last_month` |
| JavaScript | The npm registry search for `keywords:crypto`, plus the same search for 19 more crypto keywords, because `keywords:crypto` misses packages such as jsonwebtoken that tag themselves only `jwt`. The search index returns neither crypto-js (discontinued, 26,105,106 weekly downloads) nor elliptic, so neither is a candidate. | npm downloads, last week |
| Rust | The crates.io category `cryptography`, first 200 by recent downloads. | crates.io `recent_downloads` |
| Go | Every `golang.org/x/crypto/*` package, plus pkg.go.dev package search for 14 crypto terms. | pkg.go.dev "Imported by" |
| Java | Maven Central search for 17 crypto terms, plus the Maven coordinates that crypto-finder's knowledge base names. | deps.dev direct dependents |

The PyPI search page cannot be reached without a browser, and libraries.io needs an API key, so the classifier is read package by package through the PyPI JSON API. The PyPI term list is `crypt`, `cipher`, `hash`, `signature`, `signing`, `jwt`, `jose`, `json web`, `ssl`, `tls`, `ssh`, `nacl`, `sodium`, `ecdsa`, `ed25519`, `rsa`, `aes`, `certificate`, `x509`, `pgp`, `gpg`, `kdf`, `argon2`, `scrypt`, `password`, `hmac`, `keccak`, `blake`, `otp`, `pkcs` and `elliptic`. Without the term route, bcrypt, PyJWT, PyNaCl, ecdsa, passlib, python-jose, pyotp and truststore would all be missing: none of them publishes the classifier. Go ranks import paths, not modules, and folds the major versions of one module into one row.

**Secondary metric.** The secondary metric is the deps.dev `dependentCount` of the package's default version, read from `https://api.deps.dev/v3alpha/systems/<system>/packages/<name>/versions/<version>:dependents`. A row leaves it out where deps.dev has no data, which is every Go row. It counts dependents of one version only, so a package whose default version was released recently shows few dependents. It is recorded for context and does not rank anything. Java's primary metric has the same per-version limitation: an artifact whose last release is old (bcprov-jdk15on 1.70) ranks above its maintained successor.

**Known bias.** The Java candidate route includes crypto-finder's own knowledge base. Maven Central search matches whole words only, and so missed jjwt, jsch, xmlsec and jbcrypt. No neutral category source for Maven could be reached (mvnrepository.com refuses scripted access). Eight of the 15 Java rows came only from that route. The route can therefore favor libraries the tool already covers, and a neutral Maven source should replace it.

**Python fixtures.** The frozen Python top 15 replaced oscrypto, which ranks 16th with 17,110,247 downloads, with pyotp, which ranks 13th. `fixtures/python/oscrypto` and its ground truth were removed, and `fixtures/python/pyotp` was added.

## Matching rules

How a tool finding is matched to a ground-truth asset. These rules are pre-registered here before any results are published, and they do not change after results are seen. `cbench score` implements them in `internal/score`; the constants below are the ones in the code.

The rules changed on four dates, before the pinned crypto-finder runs; each change is in the [change log](#change-log) with its reason.

Every tool's output is first normalized by an adapter into `predictions/<tool>/<entry-id>.json` (`internal/prediction`), which uses the ground-truth vocabulary: `asset_type`, `primitive`, `material_type`, `origin`, `reachable` (`true`, `false` or `unknown`), `depth` and, when the tool states it, the exact dependency `distance`. `depth` is the distance in the dependency graph from the project to the component that holds the asset: 0 for project code, 1 for a direct dependency, 2 for a transitive dependency at any distance. It is not the length of the call chain that reaches the asset. A file with `reachability_reported: false` comes from a tool that never reports reachability; all its assets answer `unknown`.

**Candidate match.** A prediction is a candidate for a ground-truth asset of the same entry when all of these hold:

- the same `origin`;
- for `origin: dependency`, the same `component`, compared as exact strings after dropping the package URL's qualifiers (`?...`) and subpath (`#...`);
- the same file path, after cleaning both paths and making them slash-separated. Paths are relative to the entry root, or, for dependency assets, to the component's source root (for example the root of a Maven source jar or of a Go module);
- the same `asset_type`;
- overlapping line ranges, after widening each ground-truth range by `LINE_TOLERANCE = 0` lines on each side, so the ranges themselves must share a line. A prediction on line 12 is a candidate for a ground-truth range `[10, 12]`; one on line 13 is not. The ground-truth range is the whole expression ([annotation guide](docs/annotation-guide.md#lines)), so a tool that reports any line of a multi-line call matches it. An `algorithm` asset has its `lines`, the call that performs the operation, and may have `anchors`, the call that names its algorithm for the same object in the same function ([annotation guide](docs/annotation-guide.md#anchors)); a prediction that overlaps any of these ranges is a candidate. A key generation with `lines` on `generateKeyPair()` (line 14) and an anchor on `KeyPairGenerator.getInstance("EC")` (line 11) matches a prediction on line 11 or line 14, and not one on line 12.

**Assignment.** One-to-one and greedy. All candidate pairs are sorted by line distance (0 when the prediction overlaps one of the asset's unwidened ranges, otherwise the number of lines to the nearest one; with `LINE_TOLERANCE = 0` every candidate is at distance 0), then attributed pairs (by the rules under **Attributed** below) before unattributed ones, then by ground-truth `id`, then by the prediction's position after sorting predictions by file, start line, end line, asset type and algorithm. Pairs are taken in that order whenever both sides are still free, so a prediction takes at most one asset and an asset at most one prediction, whatever anchors they share: a tool that reports both the `getInstance` and the `generateKeyPair` of one key generation scores one TP and one duplicate, and a prediction on a `new SecureRandom()` that anchors two draws matches one of them. The attributed tie-break decides which of several candidates on the same lines is matched, so a tool that reports a generic and a specific name for one call (`AES` and `AES-GCM`) is matched on the name that attributes, and two findings on a call with two ground-truth assets pair the way that attributes both. It never makes a pair that is not a candidate, and it does not change TP, FP, FN or duplicates, only which pairs are attributed. A prediction left over that was a candidate for some ground-truth asset is a **duplicate**: every asset it could match was already taken. Duplicates are reported separately and count as neither TP nor FP.

**Scope.** For `scope.kind: files`, predictions in files outside `scope.files` are ignored and reported only as a count. For `scope.kind: pool`, recall is labeled "relative to pool", in every row that includes such an entry.

**Detected.** TP is the number of matched pairs. FP is the number of unmatched predictions in scope that are not duplicates. FN is the number of unmatched ground-truth assets.

**Attributed.** Ground-truth `algorithm` is the specific algorithm without the parameters that have their own fields: `SHA-256`, `SHA-512`, `MD5`, `AES`, `ChaCha20-Poly1305`, `RSA`, `RSA-OAEP`, `ECDSA`, `Ed25519`, `HMAC`, `PBKDF2`, `ML-KEM-768`. It is never a family such as `SHA-2`. The digest a signature or MAC computes is the `hash` parameter, not part of the name.

Before comparing, both the ground truth and the prediction are canonicalized by the same function, `algo.Canonicalize`. The tables in `internal/algo/algo.go` (algorithm names and aliases, digests, block ciphers and their key sizes, EC schemes, RSA schemes, the schemes that a digest splits off, JCA signature names, JWS and JWA names, modes, paddings, curves) are the normative lists; adding to them is a change to these rules.

1. A name's key is the name uppercased with `-`, `_` and spaces removed. A name whose key is in the algorithm table becomes the table's spelling, so `SHA256`, `SHA-256` and `sha_256` are `SHA-256`, and `3DES`, `DESede` and `TripleDES` are `3DES`. `Curve25519` and `X25519` are different algorithms.
2. A block cipher (AES, ARIA, Camellia, SM4, DES, 3DES, Blowfish, Twofish, SEED) followed by a key size, a mode, or a key size and then a mode, splits into the cipher, `key_size` and `mode`. A mode may be followed by a padding from the padding table (rule 9), which splits into `padding`. Because the key drops separators, the key size may be glued to the cipher, as in the names sonar-cryptography and CBOMkit write. The key size must be one the cipher has: 128, 192 or 256 for AES, ARIA, Camellia and Twofish; 128 for SM4 and SEED; 56 or 64 for DES; 112, 128, 168 or 192 for 3DES (effective sizes and sizes with parity bits); a multiple of 8 from 32 to 448 for Blowfish. Any other number keeps the name whole. The modes are ECB, CBC, CFB, OFB, CTR, GCM, CCM, XTS, SIV, GCM-SIV, OCB, KW and KWP. A CFB or OFB followed by a feedback size of 1, 8, 16, 32, 64 or 128 bits (`CFB8`, `OFB128`) is the base mode, and the size is dropped. `AES-256-GCM` and `AES256-GCM` are `AES` with `key_size` 256 and `mode` GCM; `AES-GCM` is `AES` with `mode` GCM; `AES-CFB8` is `AES` with `mode` CFB; `AES128-CBC-PKCS5` and `AES-128-CBC-PKCS5` are `AES` with `key_size` 128, `mode` CBC and `padding` PKCS7. `AES-512` stays whole, and so does `AES128-PKCS5`, a padding with no mode.
3. A Java (JCA) transformation `ALG/MODE/PADDING`, a name with at least two `/`, splits at the first two into the algorithm, `mode` and `padding`. The algorithm part is canonicalized by these rules, so `AES_256/GCM/NoPadding` is `AES` with `key_size` 256, `mode` GCM and `padding` none. The mode part goes through the mode table, and a mode the table does not list (`NONE`) is reduced to its key. The padding part may hold a slash itself (`OAEPWithSHA-512/256AndMGF1Padding`). A name in the algorithm table is never split, so `SHA-512/256` stays whole, and a name with one `/` is not a transformation.
4. An EC scheme (EC, ECDSA, ECDH, ECIES) followed by a curve from the curve table splits into the scheme and `curve`. Curves are canonicalized wherever they appear: `secp256r1` and `prime256v1` are `P-256`, `secp384r1` is `P-384`, `secp521r1` is `P-521`; `secp256k1` and the brainpool curves keep their names.
5. An RSA scheme name is `RSA` with the padding it names: `RSA-OAEP` and `RSAES-OAEP` are `RSA` with `padding` OAEP; `RSA-PSS` and `RSASSA-PSS` are `RSA` with `padding` PSS; `RSA-PKCS1`, `RSA-PKCS1v15`, `RSASSA-PKCS1-v1_5` and `RSAES-PKCS1-v1_5` are `RSA` with `padding` PKCS1v15. `RSA` or any of these followed by a positive integer also splits off `key_size`: `RSA-2048` is `RSA` with `key_size` 2048, and `RSA-OAEP-3072` is `RSA` with `padding` OAEP and `key_size` 3072. A digest after a signature scheme or after RSA-OAEP splits by rule 6 (`RSA-PSS-SHA256`, `RSA-OAEP-SHA256`). The JWA names `RSA-OAEP-256`, `RSA-OAEP-384` and `RSA-OAEP-512` name a digest, not a key size (rule 11). Anything else after the scheme keeps the name whole (`RSA-KEM`).
6. A signature, MAC, RSA-OAEP or PBKDF2 scheme followed by a digest splits into the scheme's algorithm, the padding the scheme names and `hash`. The digests are the hashes in the algorithm table: MD4, MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256, the four SHA3 sizes, RIPEMD-160, BLAKE2b, BLAKE2s, BLAKE3 and SM3. The schemes are `HMAC`, `ECDSA`, `DSA` and `RSA` with no padding, `RSA-PSS` and `RSASSA-PSS` with padding PSS, `RSA-PKCS1`, `RSA-PKCS1v15`, `RSASSA-PKCS1` and `RSASSA-PKCS1-v1_5` with padding PKCS1v15, `RSA-OAEP` and `RSAES-OAEP` with padding OAEP, and `PBKDF2`, `PBKDF2-HMAC` and the JCA `PBKDF2WithHmac` with no padding. `HMAC-SHA-256` and `HmacSHA256` are `HMAC` with `hash` SHA-256; `ECDSA-SHA256` is `ECDSA` with `hash` SHA-256; `RSASSA-PKCS1-SHA-256` is `RSA` with `padding` PKCS1v15 and `hash` SHA-256; `RSA-OAEP-SHA-256` is `RSA` with `padding` OAEP and `hash` SHA-256; `PBKDF2-SHA256`, `PBKDF2-HMAC-SHA256` and `PBKDF2WithHmacSHA256` are `PBKDF2` with `hash` SHA-256. `PBKDF2-HMAC` and `PBKDF2WithHmac` with no digest are aliases of `PBKDF2`, since HMAC is the PRF PBKDF2 uses. A Java (JCA) signature name `<digest>with<signature>` splits the same way: `SHA256withRSA` is `RSA` with `padding` PKCS1v15, which is what JCA's plain RSA signature is, and `hash` SHA-256; `SHA256withRSA/PSS` and `SHA256withRSAandMGF1` have padding PSS; `SHA384withECDSA` and `SHA1withDSA` have none. Anything else keeps the name whole: `HMACSHA2` (no digest size), `Ed25519-SHA512` and `SHA256withEd25519` (Ed25519 is not a scheme here) and `HKDF-SHA256` (HKDF is not a scheme here). A bare `HMAC` has no hash, so it does not attribute against a ground truth that sets one. The ground truth sets no `hash` on RSA-OAEP or PBKDF2 ([annotation guide](docs/annotation-guide.md#fields)), so a prediction's split-off digest there is a parameter the ground truth does not set.
7. Any other name is reduced to its key and kept whole: `SHA-256` does not split. In sonar-cryptography's names, `HMACSHA2` (no digest size), `MGF1`, `RAW` and `ConcatenationKDF` stay whole this way.
8. A name containing `$` (`$ALG`, `SHA-$bits`) is an unresolved placeholder and counts as no algorithm.
9. Paddings are canonicalized by the padding table, in the name and in the `padding` field: `NoPadding` is `none`; `PKCS5Padding`, `PKCS7Padding` and the CycloneDX `pkcs5` and `pkcs7` are `PKCS7`; `PKCS1Padding` and the CycloneDX `pkcs1v15` are `PKCS1v15`; `OAEPPadding`, `OAEPWith<digest>AndMGF1Padding` for each JCA OAEP digest, and the CycloneDX `oaep` are `OAEP`, which drops the digest; `PSSPadding` and the CycloneDX `pss` are `PSS`. Any other padding is reduced to its key.
10. Explicit `mode` and `curve` fields are canonicalized by the mode and curve tables, an explicit `hash` by the algorithm table (`SHA256` is `SHA-256`), and `key_size` is a number. A parameter split from the name fills its field only when the field is empty. When both are set and disagree, the explicit field wins and the score report lists the disagreement (`name_conflicts` in the JSON report).
11. A JWS or JWA algorithm name is what its table entry computes: `HS256`, `HS384` and `HS512` are `HMAC` with that `hash`; `RS256` to `RS512` are `RSA` with `padding` PKCS1v15 and that `hash`; `PS256` to `PS512` are `RSA` with `padding` PSS and that `hash`; `ES256`, `ES384` and `ES512` are `ECDSA` with `curve` P-256, P-384 and P-521 and `hash` SHA-256, SHA-384 and SHA-512; `ES256K` is `ECDSA` with `curve` secp256k1 and `hash` SHA-256; `RSA-OAEP-256`, `RSA-OAEP-384` and `RSA-OAEP-512` are `RSA` with `padding` OAEP and that `hash`. The name may follow the algorithm it computes, so `RSA-RS256` (crypto-finder's name for python-jose's RS256) is `RS256` and `ECDSA-ES256` is `ES256`; with any other prefix the name stays whole (`HMAC-RS256`). `EdDSA` names no curve, so it is an algorithm of its own in the algorithm table and equals neither `Ed25519` nor `Ed448`.

A matched pair is attributed when, after canonicalization, the ground truth's algorithm, if it has one, equals the prediction's, every parameter the ground truth has (`mode`, `padding`, `key_size`, `curve`, `hash`, including those split from its name) is present in the prediction and equal, and, when the ground truth sets `material_type`, the prediction's `material_type` is the same. `material_type` is a CycloneDX 1.7 `relatedCryptoMaterialProperties.type` and is compared as is. A prediction may carry parameters the ground truth does not set.

**Reachability.** Over matched pairs whose ground-truth `reachable` is not `unknown`, in entries whose predictions report reachability, a confusion matrix of ground-truth reachable (rows) by predicted reachable (columns). A predicted `unknown` is the "no answer" column. The matrix is reported overall and by ground-truth depth (0, 1, 2+). When no scored entry's predictions report reachability, reachability is n/a, never zero.

**Suite B pair discrimination.** For each pair, the tool discriminates correctly when every ground-truth asset of the reachable variant annotated `reachable: true` is matched with predicted `true`, and no ground-truth asset of the unreachable variant annotated `reachable: false` is matched with predicted `true`. Assets annotated otherwise in a variant, such as application crypto that is reachable in both, do not decide the pair. A pair is skipped, with the reason, when either variant has no ground truth or no prediction file, when either variant's predictions do not report reachability, or when a variant has no asset of the kind above. The score is correct pairs over scored pairs.

**Suite B reachability.** A view of its own over the suite B pair entries (`suite_b` in the JSON report), broken down by language, by scenario and by both. Unlike the confusion matrix above, it counts every ground-truth asset of those entries, matched or not: the outcome of an asset is the matched prediction's answer (`reachable`, `unreachable` or no answer), or `not detected` when no prediction matched it.

- Reachable precision = reachable assets answered reachable / all assets answered reachable.
- Reachable recall = reachable assets answered reachable / all reachable assets. A not-detected asset and no answer both count as misses, so a tool cannot raise its recall by detecting less.
- False reachable = `reachable: false` assets of unreachable variants answered reachable / those assets detected. It is the rate at which a tool calls crypto reachable when the application has no call path to it.
- Depth accuracy = matched assets whose answered depth (0, 1, 2+) equals the ground truth's / matched assets answered `reachable` or `unreachable` with a depth. An answer without a depth is counted apart (`depth_missing`). Distance accuracy compares the exact dependency distance (`distance`, 3 for app -> A -> B -> C) over the matched assets whose prediction states one.
- Pairs: the pair discrimination above, grouped by the reachable variant's language and scenario.

A tool whose suite B predictions report no reachability is shown as having none. Its outcome table only shows what it detected, and no reachability metric is printed for it, never a zero.

`scope.components`, when the ground truth sets it, lists the dependency components the annotation covers. A prediction in any other dependency component is ignored like one outside `scope.files`. Project code stays in scope.

**Entries.** An entry with ground truth but no prediction file is reported as "not run", not as all FN. An entry with no ground truth is skipped with a note.

### Object-level view

The operation-level rules above are the primary view. The object-level view is computed from the same match, with the same constants, and is reported after it (`object_level` in the JSON report).

- **Object.** A ground-truth object is the set of an entry's assets that share an `object` id, or a single asset with no `object`. Object ids are compared within one entry only. All the assets of an object have the same `asset_type`, which `cbench validate` checks.
- **Detected.** An object is detected when any of its assets is in a matched pair under the candidate-match and assignment rules above. TP is the number of detected objects, and FN the number of objects none of whose assets is matched.
- **Predictions.** For each detected object, one of the predictions matched to its assets counts toward TP, and every other prediction matched to it is a duplicate. An operation-level duplicate is a duplicate here too, because every asset it was a candidate for is matched, so its object is detected. FP is the operation-level FP: the unmatched in-scope predictions that are not duplicates. Duplicates count as neither TP nor FP.
- **Attributed.** An object is attributed when at least one of its matched assets is attributed under the rules above.
- **Breakdowns.** Precision, recall, F1 and attribution are reported overall, by language, by asset type and by entry (one entry per library in suite A), with the formulas in [Metrics](#metrics). An object is broken down by its asset type, an FP or duplicate by the prediction's. Rows that include a `pool` entry are labeled as in the operation-level view.

An example: an RSA key generated on one line and used to encrypt and decrypt on two later lines is one object with three assets. A tool that reports only the key generation scores TP 1 and FN 2 at operation level, and TP 1 and FN 0 at object level. A tool that reports all three scores TP 3 at operation level and, at object level, TP 1 with two duplicates.

## Metrics

Detection, for each row of the breakdowns (overall, by suite, by language, by asset type, language × asset type, and by entry):

- precision = TP / (TP + FP);
- recall = TP / (TP + FN);
- F1 = 2·TP / (2·TP + FP + FN), which equals the harmonic mean of precision and recall wherever both are defined;
- TP, FP, FN and duplicates as counts.

A predicted asset is broken down by the entry's suite and language and by its own `asset_type`; a matched pair and an FN by the ground truth's. Any ratio whose denominator is 0 is printed `n/a`.

Attribution: attributed TP / TP, attributed precision = attributed TP / (TP + FP), and attributed recall = attributed TP / (TP + FN). A detected but unattributed pair counts against both.

Reachability: the confusion matrices above, with reachable precision (pairs answered `true` that are reachable, over pairs answered `true`) and reachable recall (reachable pairs answered `true`, over reachable pairs; no answer counts as a miss).

Suite B: correct pairs over scored pairs.

Object level: the detection and attribution formulas above, over objects as defined in [Object-level view](#object-level-view), overall, by language, by asset type and by entry.

## Targets

Pre-registered on 2026-11-05, before any tool is tuned to them: **precision >= 0.90 and recall >= 0.85**, for each tool, in every language × asset type cell of suite A (`algorithm`, `certificate`, `protocol`, `related-crypto-material`), in each language over all asset types, and in each asset type over all languages. The object-level view is the headline; the operation-level view is reported next to it against the same targets.

- A cell passes when both precision and recall meet their target. A ratio whose denominator is 0 counts as met (a cell with no predictions cannot miss precision, one with no ground truth cannot miss recall), and a cell with neither ground truth nor predictions is `n/a`, not a pass.
- The comparison is exact on the counts: precision meets its target when 10·TP >= 9·(TP + FP), recall when 100·TP >= 85·(TP + FN). Printed ratios are rounded to three decimals and do not decide a cell.
- `go run ./scripts/scoreboard` prints the target tables, one per tool and view, and [results/README.md](results/README.md#targets) holds them.

## Hold-out policy

The public repository holds only the `dev` split. The `holdout` split lives in a private overlay with the same layout, passed with `--overlay`, and no tool configuration is tuned against it. To be written in full.

## Competitor configuration

The exact version, command line and configuration of every competing tool, chosen to give each tool its best fair run. Each run records its exact commands next to its output.

**cdxgen 13.6.0** (`scripts/run-cdxgen.sh ENTRY_DIR OUT`). A Go entry runs `cdxgen -t go -o sbom.json DIR` and then `evinse -i sbom.json -o OUT -l go DIR`, which runs cdxgen's golem analyzer. A Rust entry runs the same with `-t rust` and `-l rust` (rusi). Any other entry runs `cbom -o OUT DIR`, whose source analysis covers JavaScript and TypeScript; for Java and Python, cdxgen 13.6.0 inventories crypto libraries, keystores and certificates but not algorithm use in source. The output goes through `cbench adapt cbom --cdxgen --source-root DIR`. Golem runs on its default package pattern, `./...`. `--golem-patterns all` also analyzes dependency modules, but golem then keeps one location per algorithm across the whole build, so the project's own uses disappear when a dependency uses the same algorithm, and on the Go pair it scored lower (F1 0.286 against 0.667).

**sonar-cryptography 1.7.0** (`scripts/run-sonar-cryptography.sh ENTRY_DIR OUT`). The plugin runs headless in the CBOMkit-action v2.3.0 container (`ghcr.io/cbomkit/cbomkit-action`, pinned by digest), which embeds sonar-cryptography-plugin 1.7.0 through cbomkit-lib 1.3.0 and drives its detection rules through the SonarQube Java (8.22), Python (5.29) and Go frontends, with no SonarQube server, account or token. The 1.7.0 release publishes no plugin jar, so a SonarQube server would need the plugin built from source; the container is both lighter and the published build. The script checks the three embedded versions and refuses any other. It scans a copy of the entry with no network, one language per entry (go.mod is Go, a Maven or Gradle build file is Java, anything else Python), and three settings that matter:

1. The JVM runs inside the workspace, since occurrence locations are written relative to its working directory; elsewhere they carry the mount path.
2. A Python entry without `pyproject.toml`, `setup.cfg` or `setup.py` gets an empty `setup.cfg` in the copy. Without a build file at the root, cbomkit-lib scans only the first leaf directory it finds and silently drops the rest.
3. `CBOMKIT_EXCLUDE` matches nothing, because the default excludes test files and the benchmark scores every file.

A Java entry must be built before the scan, with its dependency jars inside the entry (`mvn package dependency:copy-dependencies`); the scanner resolves types from the `.class` files and jars under the entry, and the script refuses a Java entry without `.class` files. The CBOM is CycloneDX 1.6 with `evidence.occurrences`, which `cbench adapt cbom` reads without flags. The plugin analyzes the scanned source only: it has no dependency analysis and no reachability.

**CodeQL** quantum queries are not adapted yet. They emit SARIF results, and DGML graphs, not CBOM, so an adapter would:

1. pin the CodeQL CLI and query pack versions and build one database per language per entry;
2. read SARIF `results[].locations[].physicalLocation` (`artifactLocation.uri` relative to the source root, `region.startLine` and `endLine`) as file and lines;
3. map each query's `ruleId` and result message to `asset_type`, primitive and algorithm through a table kept in the adapter, since SARIF has no crypto vocabulary;
4. take no reachability and no dependency origin, since the queries analyze project source only, unless a dataflow query from entry points is added and documented as such;
5. ignore the DGML graphs, which describe data flow for display rather than assets.

## Reproducibility

Every result records the versions of every tool that produced it. crypto-finder results also record the ruleset version and checksum, and every crypto-finder run pins the ruleset with `--ruleset-version`. Corpus repositories are fetched at the exact commit in `corpus/manifest.yaml`.

## Change log

Every change applied to every tool, and every tool was scored again from its stored predictions. [results/README.md](results/README.md#how-the-scoring-changed) has the effect on the scores.

| Date | Change | Why |
|---|---|---|
| 2026-11-04 | A digest splits off signature and MAC names into a `hash` parameter (rule 6). Related crypto material that CycloneDX names (IVs, nonces, salts, tags, shared secrets, seeds) becomes annotated assets with a `material_type` ([annotation guide](docs/annotation-guide.md#related-crypto-material)). | Made after the Python pilot results were published; the pilot pages showed the scores before and after. |
| 2026-11-04 | The [object-level view](#object-level-view) is added and ground-truth assets gain an optional `object` field. Rule 6 also splits a digest off RSA-OAEP and PBKDF2 names, `PBKDF2-HMAC` becomes an alias of `PBKDF2`, and rule 11 adds JWS and JWA names. | Some tools, such as sonar-cryptography, report a crypto object once, where it is created; without the second view they lose recall for their granularity rather than for crypto they miss. The Python pilot showed names the rules did not recognize (`RSA-OAEP-SHA-256`, `PBKDF2-HMAC`, `PBKDF2-SHA256`, `RSA-RS256`). The operation level stays the primary view. |
| 2026-11-04 | Strict matching: `LINE_TOLERANCE` goes from 2 to 0, so ranges must overlap, and among candidates at the same distance an attributed pair is taken first. | The old tolerance paired 111 predictions with an asset 1 or 2 lines away, and none was the same operation on a neighbouring line: 57 were the call that names the algorithm before the call that performs it, 9 an IV, nonce or salt reported where it is passed on rather than drawn, and 45 another call nearby. No ground-truth range cut a multi-line expression short. [docs/strict-matching.md](docs/strict-matching.md) lists all 111. |
| 2026-11-05 | An `algorithm` asset may carry `anchors`, the call that names its algorithm for the same object in the same function; a prediction on the range or any anchor is a candidate. The allowed anchors are listed in the [annotation guide](docs/annotation-guide.md#anchors). Fernet's `decrypt` becomes its own operation. | Strict matching broke 57 pairs whose finding sat on the naming call (`KeyPairGenerator.getInstance` before `generateKeyPair`, `jwt.NewWithClaims` before `SignedString`). A tool that reports an object where its algorithm is named is not reporting the wrong line, and the anchors accept that call without accepting unrelated neighbours. 75 assets in 14 entries gained anchors; 49 of the 57 pairs match again. |
| 2026-11-05 | Suite B lands: the suite B reachability view, the `distance` field and `scope.components`. Later the same day, the handler and callback scenarios add 10 pairs and 20 assets; no existing fixture or ground truth changed. | The fixture author had seen crypto-finder's output on early Go and Python fixtures; no fixture was changed to fit a result after that, and every miss is published. |
| 2026-11-08 | A key generation that returns a private key, alone or in a key pair, also creates related crypto material of type `private-key` on the same lines. The value a digest, MAC or signature computes is restated as not material. | CycloneDX 1.7 models a private key as related crypto material, and an inventory is expected to list the private keys a program generates. 150 suite A assets added; suite B has no key generation. No matching rule changed. |
