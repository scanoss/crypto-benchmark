# Methodology

How the benchmark builds its ground truth and scores a tool against it. The rules here are pre-registered: they are written before results are published, and every later change is dated in the [change log](#change-log), applied to every tool alike, and followed by a re-score of every tool from its stored predictions.

| Topic | In short |
|---|---|
| [Ground truth](#ground-truth-process) | Two independent annotations per entry, every disagreement resolved against the [annotation guide](docs/annotation-guide.md) |
| [Suite A library list](#suite-a-library-list) | Top 15 crypto-purpose packages per ecosystem plus the standard library, frozen on 2026-11-04 |
| [Suite C](#suite-c) | 2 applications per language by stars and six fixed filters; `files` sample (headline) and `pool` scopes; first score frozen at pinned versions |
| [Matching](#matching-rules) | Same origin, component, file and asset type, overlapping lines (`LINE_TOLERANCE = 0`), one-to-one; anchors for algorithm assets |
| [Object level](#object-level-view) | Each crypto object counts once, next to the operation-level primary view |
| [Metrics](#metrics) | Precision, recall, F1, attribution, reachability, suite B pairs |
| [Targets](#targets) | Precision >= 0.90 and recall >= 0.85 per language × asset type cell |
| [Competitors](#competitor-configuration) | The exact configuration of cdxgen and sonar-cryptography, chosen for their best fair run |

## Ground truth process

Annotation is done by agents only; there is no human annotator. Two annotators label each entry independently, from the program files and [docs/annotation-guide.md](docs/annotation-guide.md). One may be the author of the programs. The other is a fresh agent that sees only the program files, the guide, the schema and the `internal/algo` tables. The two annotations are diffed, and every disagreement is resolved by reading the code against the guide. The resolution goes into the asset's `notes`. When a resolution decides a question the guide left open, the guide gains a sentence that settles it. [docs/annotation-log.md](docs/annotation-log.md) records the agreement counts for each entry.

`scope.kind: exhaustive` commits that every asset in every file of the entry is listed. `files` commits the same for every file in `scope.files` ([files scope](#files-scope)), and `pool` commits that every item of a pool of tool findings and grep candidates was adjudicated ([pool scope](#pool-scope)).

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

## Suite C

Suite C asks whether the suite A scores hold on real code that no tool was tuned on. It holds ten open-source applications, two per language, fetched at a pinned commit by `cbench fetch` and never vendored. The rule, the signal, the scopes and the freeze were written on 2026-11-13 (commit `361edb8`), before any candidate was retrieved and before any tool ran on any of them; the list and three clarifications came after, as the [change log](#change-log) records.

| Part | In short |
|---|---|
| [Selection](#suite-c-selection) | GitHub search by stars per language, six filters in a fixed order, the top 2 that pass; every candidate examined is recorded with its outcome |
| [Signal](#crypto-signal) | A file carries crypto when it imports a crypto module by a fixed, tool-independent pattern (`internal/frame`, `cbench frame`) |
| [`files` scope](#files-scope) | 15 signal files per repository, drawn with the pinned commit as seed, annotated exhaustively; the headline |
| [`pool` scope](#pool-scope) | Every finding of the three tools plus grep candidates, each adjudicated; exact precision, recall relative to the pool |
| [Freeze](#first-score-freeze) | The first score is taken at tool versions pinned now and published as is; later runs never replace it |
| [Decisions](#suite-c-decisions) | The choices the brief left open, with their reasons |

Reachability on real repositories is phase 2 and not scored here. It needs an entry point model per repository, and suite B already scores it on controlled pairs.

### Suite C selection

[corpus/suite-c.yaml](corpus/suite-c.yaml) records every candidate examined, in rank order, with its outcome: `selected`, or `rejected` with the filter that failed and why. The ten selected rows are the suite C entries of `corpus/manifest.yaml`, at the same URL and commit, and `cbench validate` checks they agree. GitHub search results change over time, so the record keeps the snapshot that decided each outcome: rank, stars, license, last commit and counts.

**The list**, applied on 2026-11-13. 49 candidates were examined: 10 selected, 13 rejected on license, 15 on application, 1 as a suite A library's source, 9 on size and 1 on crypto use. Full commits are in `corpus/manifest.yaml`.

| Entry | Repository | Stars | Rank | Commit | Source files | Signal files | Sampled |
|---|---|---|---|---|---|---|---|
| `java-c-mall` | [macrozheng/mall](https://github.com/macrozheng/mall) | 85,044 | 3 | `dcaa93b31503` | 524 | 28 | 15 |
| `java-c-jadx` | [skylot/jadx](https://github.com/skylot/jadx) | 51,792 | 12 | `fca5c2c89d08` | 2,002 | 26 | 15 |
| `python-c-yt-dlp` | [yt-dlp/yt-dlp](https://github.com/yt-dlp/yt-dlp) | 207,300 | 8 | `412e7f0a4b05` | 1,122 | 65 | 15 |
| `python-c-graphify` | [Graphify-Labs/graphify](https://github.com/Graphify-Labs/graphify) | 146,744 | 19 | `52cb28218eea` | 687 | 49 | 15 |
| `go-c-ollama` | [ollama/ollama](https://github.com/ollama/ollama) | 186,812 | 2 | `5632619189ab` | 1,097 | 39 | 15 |
| `go-c-caveman` | [JuliusBrussee/caveman](https://github.com/JuliusBrussee/caveman) | 119,423 | 5 | `f5523a8e7bd0` | 666 | 91 | 15 |
| `rust-c-clash-verge-rev` | [clash-verge-rev/clash-verge-rev](https://github.com/clash-verge-rev/clash-verge-rev) | 158,022 | 2 | `119ee86082dd` | 147 | 7 | 7 |
| `rust-c-cc-switch` | [farion1231/cc-switch](https://github.com/farion1231/cc-switch) | 153,939 | 3 | `3c76d3bda670` | 414 | 32 | 15 |
| `javascript-c-freecodecamp` | [freeCodeCamp/freeCodeCamp](https://github.com/freeCodeCamp/freeCodeCamp) | 459,052 | 1 | `7fce48a0f6b7` | 1,196 | 19 | 15 |
| `javascript-c-firecrawl` | [firecrawl/firecrawl](https://github.com/firecrawl/firecrawl) | 201,283 | 10 | `ab3acea889c7` | 1,176 | 92 | 15 |

The JavaScript rank is in the merged JavaScript and TypeScript list. Rejections in short, by filter:

| Filter | Rejected |
|---|---|
| 1 License | java-design-patterns, advanced-java, elasticsearch, termux-app; free-programming-books, system-design-primer, awesome-python, anthropics/skills, HelloGitHub, AutoGPT, dify; developer-roadmap, n8n |
| 3 Application | spring-boot, TheAlgorithms/Java, spring-framework, guava; public-apis, project-based-learning, TheAlgorithms/Python, transformers, spec-kit, langchain, ui-ux-pro-max-skill; awesome-go; claw-code; ECC, react |
| 4 Crypto library | golang/go, the source of the suite A Go standard library entry |
| 5 Size | ghidra, dbeaver; hermes-agent, langflow; kubernetes; rea, openclaw, deepseek-harness, opencode |
| 6 Crypto use | markitdown (3 signal files) |

A candidate rejected on size was counted from its git tree at the pinned commit (GitHub trees API) with the same path rules as `cbench frame`, without a clone. Every candidate counted both ways got the same number.

**Candidates.** For each language, the GitHub repository search ranked by stars, retrieved on 2026-11-13:

```
https://api.github.com/search/repositories?q=language:<L>+stars:>=1000+fork:false+archived:false+pushed:>=2026-08-15&sort=stars&order=desc&per_page=100&page=<n>
```

`<L>` is `Java`, `Python`, `Go` and `Rust`; the JavaScript entries come from `JavaScript` and `TypeScript`, merged by stars. GitHub's `language` is the repository's primary language. Candidates are taken in rank order, and the search stops at the second one that passes.

**Filters**, applied in this order; the first that fails is the recorded reason:

| # | Filter | Passes when | Source |
|---|---|---|---|
| 1 | License | The repository's license is OSI approved | GitHub `license.spdx_id`, checked against `isOsiApproved` in the [SPDX license list](https://github.com/spdx/license-list-data/blob/main/json/licenses.json) |
| 2 | Maintained | The default branch's last commit is on or after 2026-08-15, 90 days before retrieval | GitHub commits API; that commit is the pinned `sha` |
| 3 | Application | The description or the README's opening presents a program to run (a server, service, CLI tool, desktop, mobile or web app, bot or platform you deploy), not code to import (library, framework, SDK, engine, toolkit, plugin, component set) or material to read (tutorials, guides, books, courses, interview prep, lists, examples, algorithm collections) | GitHub description and README |
| 4 | Not a crypto library | The repository is not a crypto library, and is not the source of a [suite A library](#suite-a-library-list) | `corpus/libraries.yaml` |
| 5 | Tractable | 50 to 3,000 source files of the language outside excluded paths at the pinned commit | `cbench frame --dir DIR --language L` |
| 6 | Uses crypto | 5 to 150 of those files carry the [crypto signal](#crypto-signal) | same |

A crypto-purpose application (a password manager, a VPN) passes filter 4; only libraries fail it. Filters 5 and 6 read the counts that `cbench frame` prints, never which files carry the signal, and never any tool output.

**Blind selection.** No tool, crypto-finder, cdxgen, sonar-cryptography or any other, is run on a candidate before the [freeze](#first-score-freeze). Selection uses only GitHub metadata, the README and the two counts above.

**Excluded paths**, for selection and for every scope: any directory named `vendor`, `third_party`, `third-party`, `thirdparty`, `node_modules` or `bower_components`, and `*.min.js`, `*.d.ts`, `*.d.mts` and `*.d.cts` files. Git submodules are not fetched. The source files of a language are `.java`; `.py`; `.go`; `.rs`; and for JavaScript `.js`, `.jsx`, `.mjs`, `.cjs`, `.ts`, `.tsx`, `.mts` and `.cts`. Test files are in scope, as everywhere in the benchmark.

### Crypto signal

A source file carries the crypto signal when it imports a module whose path names crypto, or uses one of the few crypto APIs that need no import. `internal/frame` is the normative implementation and `cbench frame` prints it; changing either is a change to these rules.

- **Imports.** Python `import` and `from ... import`, Java `import`, Go imports, Rust `use`, `extern crate` and the first segment of any `a::b` path, and JavaScript `import ... from`, bare `import`, `require(...)` and `import(...)`.
- **Crypto module.** The module path is split at `.`, `/`, `:` and `@`, lowercased, with `-` read as `_`. It names crypto when a segment contains `crypt`, `jwt`, `jws`, `jwe`, `jwk`, `x509`, `pkcs`, `openssl` or `rustls`, or equals one of a fixed list of 80 names: standard library modules (`hashlib`, `hmac`, `secrets`, `ssl`, `tls`, `security`), algorithm names (`sha256`, `aes`, `rsa`, `ed25519`, `hkdf`, ...), and generic crypto words (`cipher`, `digest`, `signature`, `pki`). Library names in the list (`ring`, `nacl`, `paramiko`, `bouncycastle`, `tink`, ...) cover crates and packages whose name says nothing about crypto; the list is broader than the suite A libraries, so the frame does not favor code the tools were tuned on.
- **Without an import.** Python `os.urandom` and `random.SystemRandom`, Java fully qualified `java.security.`, `javax.crypto.` and `javax.net.ssl.`, and JavaScript `crypto.subtle` and `crypto.getRandomValues`.
- **Own module.** In Go, an import of the repository's own module (any `go.mod` in it) is judged on the path after the module, so a repository named `crypto-something` does not flag every file.

The signal is broad on purpose: a signal file with no crypto asset is still annotated, and tests precision. Crypto in a file without the signal, such as a project helper that wraps a library imported elsewhere, is outside the `files` frame; the `pool` scope covers it.

### `files` scope

`scope.kind: files` commits that every asset in every file of `scope.files` is listed, by the same rules and two-annotator process as `exhaustive`. The score ignores predictions in other files and reports them as a count, so precision and recall are both exact within the sample.

- **Sample.** 15 signal files per repository, or every signal file when there are 15 or fewer. Each signal file is ranked by the hex SHA-256 of `<pinned commit>:<path>`, and the 15 lowest are taken (`internal/frame.Sample`). The pinned commit is the seed: it was fixed by selection, so no one chooses a seed, and anyone can redraw the sample with `cbench frame --entry ID --sample`.
- **What it estimates.** Precision and recall over files that import crypto. The ten repositories give 142 files: 15 each, and all 7 signal files of `rust-c-clash-verge-rev`.
- **Headline.** The `files` scope is the suite C headline: it is the only scope whose recall does not depend on what the tools found.

### `pool` scope

`scope.kind: pool` commits that every item of the pool was adjudicated, over the whole repository (source files of the entry's language outside excluded paths).

- **Pool.** The union of every finding of the three tools at their [frozen versions](#first-score-freeze), on the whole repository, every finding of the later crypto-finder run with `--include-tests` ([results](results/suite-c.md#later-run-crypto-finder-with---include-tests)), and every grep candidate line that `cbench frame --entry ID --grep` prints: a line, outside import lines, that names a crypto operation, algorithm or API (`internal/frame`, `grepTerms`). Only findings in source files of the entry's language outside excluded paths join it. There is no depth cutoff. A tool with no support for a language contributes no findings there.
- **Why the later run is pooled.** The pool is a union, so adding a source only enlarges it: every item the other sources bring is still adjudicated, and an asset only that run found joins the denominator of every tool's pool recall instead of going uncounted. The run was added on 2026-11-13, before any pool item was adjudicated and before any pool result existed. It changes no tool's first score: crypto-finder's first pool score is still that of its frozen configuration, and the `--include-tests` run is scored beside it as the later run it is.
- **Entry.** An entry has one ground-truth file with one scope, so the pool of a repository is a second manifest entry, `<id>-pool`, with the same language, URL and commit; `cbench validate` checks they agree. The ten `<id>-pool` entries were added on 2026-11-13 with no ground truth, which `cbench validate` reports as a warning until it lands. Before matching, the score drops a pool entry's predictions in files of other languages or in excluded paths (`frame.InScope`, the same rule as `cbench frame`) and reports their number (`ignored_outside_pool_frame` in the JSON report).
- **Adjudication.** Two annotators judge each item independently against the [annotation guide](docs/annotation-guide.md), as for any ground truth. An item that is an asset becomes a ground-truth asset with its full fields; one that is not is recorded as rejected, so the adjudication can be audited. Items on the same asset are one asset.
- **Size.** On the pinned commits the grep gives 9,899 candidate lines over the ten repositories, from 52 (`java-c-mall`) to 4,296 (`python-c-yt-dlp`). Every one is adjudicated; the pool is not sampled.
- **Precision** is exact over the whole repository, because every finding of every pooled tool was adjudicated.
- **Recall is relative to the pool.** The denominator is the assets found by some tool or the grep, not every asset in the repository. As in TREC-style pooling, it overstates true recall by the assets nobody found, and it favors the tools that built the pool. Every row that includes a pool entry is labeled "recall relative to pool" (the score already does this, [Scope](#matching-rules)), and pool recall is never shown next to suite A or `files` recall without that label. The report gives, per tool, the number of adjudicated assets that only that tool found.
- **A tool that joins later.** Its findings outside the pool are adjudicated and added, and every tool is scored again against the larger pool. That score is a new dated run; it never replaces the [first score](#first-score-freeze).

### Targets

None are pre-registered for suite C. The suite A targets are not applied to it. Suite C is published as a measurement with its counts, so a reader can see how far real code moves each number.

### First score freeze

- **Pinned now.** The first suite C score is taken with the versions of the published suite A and B runs: crypto-finder v0.53.1 with remote ruleset v1.48.0, cdxgen 13.6.0 and sonar-cryptography 1.7.0, each run as recorded in [Competitor configuration](#competitor-configuration). They are recorded here on 2026-11-13, before annotation starts and before any tool runs on a suite C repository.
- **Published as is.** The first score is published whatever it shows, with every miss.
- **Later runs are separate.** Any change to a tool, a rule, a knowledge base or a configuration made after its suite C results were seen is reported as a separate, later, dated run. It never replaces or edits the first score, and the results page shows the first score first.
- **Rules of scoring.** A change to these rules or the matching rules after the first score is applied to every tool, dated in the [change log](#change-log), and published next to the first score, never in its place.

### Suite C decisions

The brief left these open; each is decided here so the work can go ahead without another round.

| Question | Decision | Why |
|---|---|---|
| Sample size per repository | 15 files | 142 files in all (one repository has 7 signal files), about the annotation volume of suite A, and enough files per repository that one unusual file does not decide its score |
| Seed | The pinned commit, through SHA-256 of `<commit>:<path>` | The commit was fixed by selection, so the seed cannot be chosen after looking at a draw |
| Headline scope | `files` | Its recall does not depend on the tools; pool recall does |
| Targets | None | There is no prior on real code; a target set now would be a guess, and one set later would be fitted |
| Tool versions of the first score | crypto-finder v0.53.1, ruleset v1.48.0, cdxgen 13.6.0, sonar-cryptography 1.7.0 | The same versions as the published suites A and B, so suite C measures the same tools on other code; fixed before any tool sees a suite C repository |
| Files of other languages in a repository | Out of scope; for `pool` entries, the scorer drops predictions outside the entry's language files and excluded paths before matching, and reports their number. Made on 2026-11-13, before any pool ground truth or pool result | An entry has one language, and suite A scores one language per entry |
| Test files | In scope | The benchmark scores every file, and the competitors are configured not to skip tests |
| Ambiguous "application" calls in filter 3 | Decided by the selector, with the reason recorded in `corpus/suite-c.yaml`; a challenge is a dated change made before the first score only | The record makes each call auditable; reselecting after the first score would be selection on results |
| Split | `dev`, public | Suite C tests generalization by being unseen by the tool authors at tuning time, which the freeze enforces; a private hold-out adds nothing to that |
| Size counts of rejected candidates | Read from the git tree at the pinned commit | Same path rules as `cbench frame`, and equal to it wherever both were run, without cloning repositories of up to 55,000 files |
| A large grep pool (4,296 lines in yt-dlp) | Adjudicated in full, no sampling | A sampled pool would make pool precision an estimate; most lines are rejected at a glance |
| Reachability | Phase 2 | Real entry points need a model per repository; suite B scores reachability on controlled pairs |

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

**Scope.** For `scope.kind: files`, predictions in files outside `scope.files` are ignored and reported only as a count. For `scope.kind: pool`, predictions outside the entry's language source files or in excluded paths ([Suite C](#pool-scope)) are dropped before matching and reported only as a count, and recall is labeled "relative to pool", in every row that includes such an entry.

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
| 2026-11-13 | Suite C: the [selection rule](#suite-c-selection), the [crypto signal](#crypto-signal), what the [`files`](#files-scope) and [`pool`](#pool-scope) scopes commit to, and the [first score freeze](#first-score-freeze), written before any candidate was retrieved (commit `361edb8`). Then the ten repositories were selected by it and added to the manifest with no ground truth. Three clarifications were written after retrieval, and none changes an outcome: counting size-rejected candidates from the git tree, adjudicating the grep pool in full, and the `<id>-pool` entry. `cbench validate` checks `corpus/suite-c.yaml` against the manifest. | Suite C tests whether the suite A scores hold on code no tool was tuned on, so the selection, the scopes and the freeze had to be fixed before anyone sees a suite C result. No scoring rule for suites A and B changed. |
| 2026-11-13 | Suite C `files` ground truth: the 142 sampled files of the ten entries, 214 assets, from two independent annotations adjudicated against the [annotation guide](docs/annotation-guide.md) ([annotation log](docs/annotation-log.md#suite-c-files-scope-2026-11-13)). The guide gained sentences on project helpers and a project's own algorithm implementations, random draws versus identifier libraries, TLS sessions, certificates and keys loaded into a TLS context, crypto run through a command line, non-security uses, where a key generation's primitive is looked for, objects across files, overridden algorithm literals, and suite C reachability fields. | Made before any tool output on suite C was seen: no tool had run on a suite C repository, so neither the annotations nor the guide clarifications could follow a result. No scoring rule changed, and no suite A or B ground truth changed. |
| 2026-11-13 | Suite C `pool` preparation. The scorer drops a `pool` entry's predictions outside the entry's language source files or in excluded paths before matching, by the same rule as `cbench frame` (`frame.InScope`), and reports their number, as the [decisions](#suite-c-decisions) had pre-registered. The ten `<id>-pool` entries join the manifest with no ground truth yet. The [pool](#pool-scope) also takes in the findings of the later crypto-finder run with `--include-tests`. | Made before any pool item was adjudicated and before any pool result existed. The pool is a union, so adding a source only enlarges it, and no first score changes. No scoring rule for suites A and B, and no `files` score, changed: the drop applies only to `pool` entries. |
