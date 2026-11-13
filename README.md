# crypto-benchmark

A public, reproducible accuracy benchmark for crypto inventory tools: how well a tool finds the cryptographic assets in source code, and whether it can tell which of them an application actually reaches.
It scores [SCANOSS crypto-finder](https://github.com/scanoss/crypto-finder) next to two open-source tools, cdxgen and sonar-cryptography, in Java, Python, Go, Rust and JavaScript/TypeScript.

## Results

crypto-finder v0.53.1 with remote ruleset v1.48.0, scored on 2026-11-12 (suites A and B) and 2026-11-13 (suite C). Full tables: [results/README.md](results/README.md) (suite A), [results/suite-b.md](results/suite-b.md) (suite B) and [results/suite-c.md](results/suite-c.md) (suite C).

| Suite | What is scored | crypto-finder | sonar-cryptography 1.7.0 | cdxgen 13.6.0 |
|---|---|---|---|---|
| A | Operation level, TP / FP / FN | 1,413 / 9 / 3 | 150 / 43 / 696 (3 languages) | 112 / 77 / 1,304 |
| A | Operation level, precision / recall | 0.994 / 0.998 | 0.777 / 0.177 | 0.593 / 0.079 |
| A | Object level, precision / recall | 0.991 / 0.998 | 0.728 / 0.197 | 0.560 / 0.097 |
| A | Target cells passed, object level | 28 of 30 | 0 of 20 | 0 of 30 |
| B | Pairs discriminated correctly | 45 of 45 | no reachability reported | no reachability reported |
| B | Reachable precision / recall | 1.000 / 1.000 | n/a | n/a |
| B | Dependency crypto detected | 104 / 0 / 0 (TP / FP / FN) | 0 of 52 assets (Go, Python, Java) | 16 of 84 assets (JavaScript only) |
| C | `files` scope, operation level, TP / FP / FN, first score | 108 / 8 / 106 | 28 / 0 / 97 (3 languages) | 18 / 33 / 196 |
| C | `files` scope, precision / recall, first score | 0.931 / 0.505 | 1.000 / 0.224 | 0.353 / 0.084 |

The two target cells crypto-finder misses are Python protocol and Rust protocol. Both fail on precision, because crypto-finder reports a TLS configuration setter as a protocol of its own. [Known gaps](results/README.md#known-gaps) lists every remaining error.

In suite C, 90 of crypto-finder's 106 misses are in test files: its recorded configuration used the default that leaves test sources out, while the competitors scan them. A later, separate run with `--include-tests` scores 184 / 14 / 30 (0.929 / 0.860); the first score stands ([suite C](results/suite-c.md)).

sonar-cryptography has no JavaScript or Rust support, so it ran on three languages. cdxgen 13.6.0 does not analyze algorithm use in Python or Java source, so it scores 0 there. Neither tool reports reachability for crypto assets, so suite B shows them as having none, not as scoring zero. Both were run on the first 35 suite B pairs only.

## What the suites measure

| Suite | Question | Corpus | Status |
|---|---|---|---|
| **A**, library coverage | Does the tool find every crypto operation, and name it correctly? | One small program set per library: the top 15 crypto packages of each ecosystem plus the standard library, 79 entries, 1,416 assets in 1,015 objects, under `fixtures/` | Scored |
| **B**, transitive reachability | Can the tool tell crypto the app reaches through its dependencies from crypto it does not? | 45 pairs over five languages (app -> lib A -> lib B -> crypto). Each pair has a reachable and an unreachable variant that differ in one change. Under `fixtures/<language>/suite-b/` | Scored |
| **C**, real repositories | Does it hold on real code? | Ten open-source applications, two per language, chosen by a pre-registered rule ([METHODOLOGY.md](METHODOLOGY.md#suite-c)) and fetched at a pinned commit, never vendored | `files` scope scored: first score 2026-11-13, 214 assets in 142 files ([results](results/suite-c.md)); `pool` scope not yet scored |

**Targets**, pre-registered on 2026-11-05: precision >= 0.90 and recall >= 0.85 in every language × asset type cell of suite A, in each language over all asset types, and in each asset type over all languages. The object level is the headline view ([METHODOLOGY.md](METHODOLOGY.md#targets)).

**Two views of detection.** The operation level counts every annotated operation. The object level counts each crypto object (a key with everything done with it) once, so a tool that reports an object once, where it is created, is not penalized for its granularity ([METHODOLOGY.md](METHODOLOGY.md#object-level-view)).

## Reproduce a score

Requires Go (see `go.mod`) and `git`.

```sh
go run ./cmd/cbench validate        # check the manifest and every ground-truth file
go run ./cmd/cbench score --tool crypto-finder \
  --predictions results/crypto-finder/go-pilot/predictions   # markdown report on stdout
go run ./scripts/scoreboard         # the tables of results/README.md
```

Every `results/<tool>/<run>/score.md` is the output of `cbench score` on the predictions committed next to it.

## Run a tool yourself

### crypto-finder

crypto-finder's detection rules are remote, so a scan needs a SCANOSS API key in `SCANOSS_API_KEY`. Never commit keys to this repository. Pin the ruleset with `--ruleset-version`; every result records the ruleset version and checksum.

```sh
crypto-finder scan --ruleset-version v1.48.0 --export-callgraph-project-reachability \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json fixtures/go/stdlib
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json --out predictions/crypto-finder/ID.json
```

Reachability comes only from the call graph export, so leave out `--callgraph` only for a run without reachability. [internal/adapter/cryptofinder/README.md](internal/adapter/cryptofinder/README.md) documents every field mapping.

Suite B entries are real dependency graphs, so each one is built with its package manager first. `scripts/run-crypto-finder-suite-b.sh CF_BIN PREDICTIONS_DIR LANG...` builds, scans and adapts every pair of the languages given:

```sh
scripts/run-crypto-finder-suite-b.sh ./crypto-finder work/preds go python java javascript rust
go run ./cmd/cbench score --tool crypto-finder --predictions work/preds
```

Go, Python, Node.js and Cargo need their toolchains; Java needs a full JDK 21 in `JAVA_HOME` and Maven. The web frameworks of the handler scenario come from the public registries, so the first build needs network access ([fixtures/README-suite-b.md](fixtures/README-suite-b.md)).

### cdxgen and sonar-cryptography

Any tool that writes a CycloneDX 1.6 or 1.7 CBOM with `evidence.occurrences` is scored through the generic CBOM adapter ([internal/adapter/cbom/README.md](internal/adapter/cbom/README.md)):

```sh
go run ./cmd/cbench adapt cbom --tool NAME --tool-version V --entry ID --input bom.json \
  --out predictions/NAME/ID.json [--source-root DIR] [--cdxgen]
```

| Tool | Runner | Needs | Notes |
|---|---|---|---|
| cdxgen 13.6.0 | `scripts/run-cdxgen.sh ENTRY_DIR OUT` | Node.js (`npm exec`) or `CDXGEN_BIN` | Go through `evinse -l go`, Rust through `evinse -l rust`, anything else through `cbom`. Adapt with `--cdxgen --source-root "$(realpath ENTRY_DIR)"`. |
| sonar-cryptography 1.7.0 | `scripts/run-sonar-cryptography.sh ENTRY_DIR OUT` | Docker only, no SonarQube server or account | Headless in the pinned CBOMkit-action v2.3.0 container. Build a Java entry first (`mvn package dependency:copy-dependencies`). |

Both scripts refuse any other tool version and write the commands they ran to `OUT.command`. [METHODOLOGY.md](METHODOLOGY.md#competitor-configuration) records each configuration and why it was chosen.

## Details

### cbench commands

| Command | Does |
|---|---|
| `cbench validate [--overlay DIR]` | Checks the manifest and every ground-truth file. Prints `severity: file: yaml.path: message` and exits non-zero on any error. |
| `cbench fetch [--only ID] [--dest DIR]` | Clones every URL entry at its pinned commit into `work/corpus/<id>`. Safe to rerun: it skips a checkout at the right commit and refuses, without deleting anything, when the directory holds anything else. |
| `cbench frame (--entry ID \| --dir DIR --language L --commit SHA) [--sample \| --grep]` | Prints a suite C checkout's signal files, the seeded `files` sample, or the `pool` grep candidates ([METHODOLOGY.md](METHODOLOGY.md#crypto-signal)). |
| `cbench adapt crypto-finder\|cbom ...` | Converts one tool run into a prediction file. |
| `cbench score --tool NAME [--predictions DIR] [--format json] [--out FILE]` | Validates the benchmark, then scores the predictions by the rules in [METHODOLOGY.md](METHODOLOGY.md#matching-rules). The report names the tool configuration and the entries scored, not run and skipped. |

### Layout

```
corpus/manifest.yaml              every corpus entry: id, language, suite, split, and url+sha or path
corpus/libraries.yaml             the frozen suite A library list
corpus/suite-c.yaml               every suite C candidate examined and its outcome
groundtruth/<suite>/<entry>.yaml  annotated crypto assets for one entry
fixtures/<language>/...           in-repo programs for suites A and B
cmd/cbench                        the CLI
internal/corpus, groundtruth, prediction   the Go types that define the schemas
internal/score                    the matching rules and metrics
internal/adapter/<tool>           converters from each tool's output to predictions
results/<tool>/<run>/             published predictions, score reports and run records
docs/                             annotation guide, annotation log, strict-matching record
METHODOLOGY.md                    how ground truth is made and how tools are scored
scripts/check-public.sh           fails when the tree holds anything that must not be public
```

`corpus/manifest.yaml` and `groundtruth/B/go-b-transitive-reachable.yaml` are working examples of each schema.

### Hold-out policy

This repository holds only the `dev` split. A hold-out set with the same layout lives outside it and is passed with `--overlay DIR`, so results on it cannot be tuned against.

## License

Code (`cmd/`, `internal/`, `fixtures/`, scripts and workflows) is licensed under the [Apache License 2.0](LICENSE).
Data, meaning the corpus manifest, ground truth, methodology and published results, is licensed under [CC BY 4.0](LICENSE-DATA).
Corpus repositories fetched by `cbench fetch` keep their own licenses and are never redistributed here.
