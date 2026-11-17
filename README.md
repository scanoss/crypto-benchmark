# crypto-benchmark

A public, reproducible accuracy benchmark for crypto inventory tools. It measures how well a tool finds the cryptographic assets in source code, and whether it can tell which of them an application actually reaches. It compares [SCANOSS crypto-finder](https://github.com/scanoss/crypto-finder), cdxgen and sonar-cryptography on Java, Python, Go, Rust and JavaScript/TypeScript, and publishes every prediction, score and miss.

## Results at a glance

| Suite | Question | crypto-finder | sonar-cryptography 1.7.0 | cdxgen 13.6.0 |
|---|---|---|---|---|
| **A**, library coverage | Precision / recall, operation level | **0.994 / 0.998** | 0.777 / 0.177 (3 languages) | 0.593 / 0.079 |
| **B**, transitive reachability | Pairs told apart correctly | **45 of 45** | no reachability reported | no reachability reported |
| **C**, real code, `files` (headline) | Precision / recall, first score | **0.931 / 0.505** | 1.000 / 0.224 (3 languages) | 0.353 / 0.084 |
| **C**, real code, `pool` | Precision / recall **relative to the pool**, first score | **0.949 / 0.447** | 0.983 / 0.286 (3 languages) | 0.099 / 0.077 |

- crypto-finder ran with remote ruleset v1.48.0: v0.54.0 for suites A and B, scored on 2026-11-17, and v0.53.1 for the suite C first score, scored on 2026-11-13. v0.54.0 changes only how the ruleset is downloaded, and gave the same predictions as v0.53.1 on every suite, C included.
- Suite C shows the **first score**, frozen before any tool ran on it. Most of crypto-finder's suite C misses are test files that its default configuration left out. A later run with `--include-tests` scores 0.929 / 0.860 on `files` and 0.951 / 0.886 on the pool. It is published [beside the first score](#suite-c-real-repositories), never in its place.
- Pool recall counts only the assets some tool or a grep found. It overstates true recall and is not comparable to suite A or `files` recall.

## Where to read more

| You want | Read |
|---|---|
| Every suite A table: targets, per language, per asset type, known gaps, run history | [results/README.md](results/README.md) |
| Suite B by language and scenario, and why the other tools score none | [results/suite-b.md](results/suite-b.md) |
| Suite C, both scopes, every miss by cause, and the later run | [results/suite-c.md](results/suite-c.md) |
| How ground truth is made and how a finding is matched | [METHODOLOGY.md](METHODOLOGY.md) |
| What counts as an asset | [docs/annotation-guide.md](docs/annotation-guide.md) |
| Agreement counts and resolutions per annotation round | [docs/annotation-log.md](docs/annotation-log.md) |
| How to contribute a case or a correction | [CONTRIBUTING.md](CONTRIBUTING.md) |
| One tool run: commands, predictions, `score.md`, deviations | `results/<tool>/<run>/` |

## What it measures, and why

A crypto inventory (a CBOM) is only useful if it is complete and correct. A missed algorithm hides a migration risk, and a false one costs review time. Much of an application's crypto also sits in its dependencies, so an inventory must say which of that crypto the application can reach.

The benchmark asks three questions, one per suite:

| Suite | Question | Corpus | Scores |
|---|---|---|---|
| **A**, library coverage | Does the tool find every crypto operation and name it correctly? | 79 entries in 5 languages: the top 15 crypto packages of each ecosystem plus its standard library, one small program set each. 1,416 assets in 1,015 objects, under `fixtures/` | Detection, attribution, targets |
| **B**, transitive reachability | Can the tool tell crypto the application reaches through its dependencies from crypto it does not? | 45 pairs (90 entries, 104 assets) over the five languages: app -> lib A -> lib B -> crypto. The reachable and unreachable variant of a pair differ in one change | Reachability, depth, distance, pairs |
| **C**, real repositories | Do the scores hold on real code that no tool was tuned on? | 10 open-source applications, 2 per language, chosen by a pre-registered rule and fetched at a pinned commit, never vendored | Detection in two scopes |

Suite C has two scopes:

- **`files`, the headline.** 142 files sampled from the ten repositories (15 per repository, 7 where fewer exist), every asset annotated: 214 assets. Precision and recall are both exact within the sample.
- **`pool`, whole repositories.** Every finding of every tool and every line of a grep for crypto terms, adjudicated: 10,655 items, 993 assets. Precision is exact. Recall is relative to the pool.

Suite C repositories: Go caveman and ollama, Java jadx and mall, JavaScript firecrawl and freeCodeCamp, Python graphify and yt-dlp, Rust cc-switch and clash-verge-rev. [corpus/suite-c.yaml](corpus/suite-c.yaml) records every candidate examined and why it was taken or rejected.

## Tools compared

| Tool | Version and configuration | Languages run |
|---|---|---|
| crypto-finder | v0.54.0 released binary (v0.53.1 for the frozen suite C first score), remote ruleset v1.48.0 pinned with `--ruleset-version` (checksum recorded in every `score.md`) | all five |
| sonar-cryptography | 1.7.0, headless in the pinned CBOMkit-action v2.3.0 container, no SonarQube server | Go, Java, Python. It has no JavaScript or Rust support |
| cdxgen | 13.6.0: `evinse` (golem) for Go, `evinse` (rusi) for Rust, `cbom` for the rest | all five. For Python and Java it inventories libraries, not algorithm use in source, so it scores 0 there |

Each tool is configured for its best fair run. [METHODOLOGY.md](METHODOLOGY.md#competitor-configuration) gives every command line and why it was chosen.

## Results

All counts are TP / FP / FN at operation level unless a row says otherwise. Each number comes from a committed `score.md`.

### Suite A: library coverage

| Tool | Languages | TP / FP / FN | P / R, operation | P / R, object | Target cells passed, object level |
|---|---|---|---|---|---|
| crypto-finder | 5 | 1,413 / 9 / 3 | 0.994 / 0.998 | 0.991 / 0.998 | 28 of 30 |
| sonar-cryptography | 3 | 150 / 43 / 696 | 0.777 / 0.177 | 0.728 / 0.197 | 0 of 20 |
| cdxgen | 5 | 112 / 77 / 1,304 | 0.593 / 0.079 | 0.560 / 0.097 | 0 of 30 |

- crypto-finder fails two cells, Python protocol and Rust protocol. Both fail on precision: it reports a TLS configuration setter as a protocol of its own.
- The main difference is coverage. crypto-finder finds 954 of 955 algorithm assets. sonar-cryptography finds 24% in its three languages, cdxgen 6%.

Details: [results/README.md](results/README.md), with [known gaps](results/README.md#known-gaps) and the [history of pinned runs](results/README.md#history-of-pinned-runs).

### Suite B: transitive reachability

| Tool | Pairs correct | Reachable P / R | Unreachable assets called reachable | Dependency crypto detected |
|---|---|---|---|---|
| crypto-finder | 45 of 45 | 1.000 / 1.000 | 0 of 52 | 104 / 0 / 0 |
| sonar-cryptography | no reachability reported | n/a | n/a | 0 of 52 assets (Go, Python, Java) |
| cdxgen | no reachability reported | n/a | n/a | 16 of 84 assets, JavaScript only, the same in both variants |

- crypto-finder also names the right dependency depth and distance on all 104 assets.
- sonar-cryptography and cdxgen report no reachability for crypto assets, so they are shown as having none, not as scoring zero. Both ran on the first 35 pairs, before the handler and callback scenarios were added.

Details: [results/suite-b.md](results/suite-b.md).

### Suite C: real repositories

**`files` scope (headline): 142 files, 214 assets**

| Tool | TP / FP / FN | P / R, operation | P / R, object |
|---|---|---|---|
| crypto-finder, first score | 108 / 8 / 106 | 0.931 / 0.505 | 0.924 / 0.485 |
| crypto-finder `--include-tests`, later run | 184 / 14 / 30 | 0.929 / 0.860 | 0.925 / 0.860 |
| sonar-cryptography, 3 languages (125 assets) | 28 / 0 / 97 | 1.000 / 0.224 | 1.000 / 0.228 |
| cdxgen | 18 / 33 / 196 | 0.353 / 0.084 | 0.353 / 0.090 |

**`pool` scope: whole repositories, 993 adjudicated assets, recall relative to the pool**

| Tool | TP / FP / FN | P / R, operation | P / R, object | Assets only this tool found |
|---|---|---|---|---|
| crypto-finder, first score | 444 / 24 / 549 | 0.949 / 0.447 | 0.944 / 0.435 | 284 |
| crypto-finder `--include-tests`, later run | 880 / 45 / 113 | 0.951 / 0.886 | 0.949 / 0.896 | 649, with the later run in place of the first score |
| sonar-cryptography, 3 languages (590 assets) | 169 / 3 / 421 | 0.983 / 0.286 | 0.983 / 0.300 | 71, all in Go test files; 0 against the later run |
| cdxgen | 76 / 689 / 917 | 0.099 / 0.077 | 0.098 / 0.081 | 9 |

**The first score and the later run.** The first score is the result. It used the tool versions frozen before any tool ran on suite C, and it is published as is ([first score freeze](METHODOLOGY.md#first-score-freeze)). crypto-finder's recorded configuration kept its default, which leaves test sources out, while the other tools were configured to scan tests. Test files hold 90 of its 106 `files` misses and 436 of its 549 pool misses. The `--include-tests` run changed only that flag, after the first score was seen, so it is a separate, dated run. Outside test paths its findings are identical to the first score.

Details: [results/suite-c.md](results/suite-c.md), with every miss grouped by cause and the [assets only one tool found](results/suite-c-pool-unique/README.md).

## How scoring works

| Rule | In short |
|---|---|
| Ground truth | Two annotators label each entry independently from the code and the [annotation guide](docs/annotation-guide.md). Every disagreement is resolved by reading the code against the guide, and the resolution is recorded. |
| Strict matching | A finding matches an asset with the same origin, component, file and asset type, on overlapping lines (`LINE_TOLERANCE = 0`), one to one. In suite A, an algorithm asset also matches on its anchors: the call that names its algorithm for the same object. |
| Operation and object view | The operation level counts every annotated operation and is the primary view. The object level counts each crypto object (for example, a key and everything done with it) once, so a tool that reports an object once, where it is created, is not penalized for its granularity. |
| Attribution | A match is attributed when it also names the right algorithm and parameters. |
| Targets, suite A only | Pre-registered on 2026-11-05: precision >= 0.90 and recall >= 0.85 in every language × asset type cell, per language and per asset type. Suite C has no targets. |
| Change log | Rules were fixed before results were published. Every later change is dated, applied to every tool alike, and followed by a re-score of every tool from its stored predictions ([change log](METHODOLOGY.md#change-log)). |

Full rules: [METHODOLOGY.md](METHODOLOGY.md).

## Known limitations

- **Who publishes it.** The benchmark is published by SCANOSS, the maker of crypto-finder. The pre-registered rules, independent annotation, first-score freeze and published misses exist so that a reader can check every number without trusting the publisher. Outside cases are [welcome](#contributing) and are accepted blind, before any tool runs on them.
- **Annotators are agents.** Each entry is annotated by two independent agents and adjudicated. There is no human annotator.
- **Suite A library list.** The Java candidates include the libraries crypto-finder's knowledge base names, because Maven Central search missed several. The list can favor libraries the tool already covers ([METHODOLOGY.md](METHODOLOGY.md#suite-a-library-list)).
- **Suite B fixtures are small and uniform** by intent, and the fixture author saw crypto-finder's output on early Go and Python fixtures. No fixture was changed to fit a result.
- **The suite C sample is small.** One repository can dominate a row, and the two Java repositories hold only 3 `files` assets between them.
- **Pool recall overstates true recall.** It cannot see assets no source found, and it favors the tools that built the pool. The annotators list 25 known assets that no pool item covers ([annotation log](docs/annotation-log.md#known-misses-of-the-pool)).
- **A project's own crypto code is out of crypto-finder's reach.** It detects calls into crypto libraries and standard modules, not an algorithm written in the project itself, such as yt-dlp's pure-Python AES.
- **Reachability on real repositories is not scored yet.** Suite C reachability is phase 2. Suite B scores reachability on controlled pairs.
- **Competitors ran once.** Their scores come from their stored predictions, re-scored against the current ground truth.
- **Hold-out.** This repository holds only the `dev` split. A hold-out set with the same layout lives outside it and is passed with `--overlay DIR`.

## Reproduce

Requires Go (see `go.mod`) and `git`.

### Re-score the published predictions

No tool or key is needed to check a published score.

```sh
go run ./cmd/cbench validate          # check the manifest and every ground-truth file
go run ./cmd/cbench score --tool crypto-finder \
  --predictions results/crypto-finder/go-pilot/predictions    # markdown report on stdout
go run ./scripts/scoreboard           # every suite A table of results/README.md
```

Every `results/<tool>/<run>/score.md` is the output of `cbench score` on the predictions committed next to it.

### Run a tool yourself

1. **Fetch** the suite C repositories at their pinned commits, if you need them: `go run ./cmd/cbench fetch` clones each into `work/corpus/<id>`.
2. **Run** the tool on an entry.
3. **Adapt** its output into a prediction file with `cbench adapt`.
4. **Score** the predictions with `cbench score`.

**crypto-finder.** Its detection rules are served remotely, so a scan needs a SCANOSS API key in the `SCANOSS_API_KEY` environment variable. Do not commit a key to this repository. Pin the ruleset with `--ruleset-version`. From v0.54.0, crypto-finder downloads the ruleset from `/v3/cryptography/rulesets` as a bundle with a signed lease, so the key needs access to that endpoint. The interim report still records the ruleset checksum, the same for every key, so a run can be checked against the published one.

```sh
crypto-finder scan --ruleset-version v1.48.0 --include-tests --export-callgraph-project-reachability \
  --export-callgraph work/ID.callgraph.json --output work/ID.interim.json fixtures/go/stdlib
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/ID.interim.json \
  --callgraph work/ID.callgraph.json --out predictions/crypto-finder/ID.json
go run ./cmd/cbench score --tool crypto-finder --predictions predictions
```

- `--include-tests` is the current configuration. The suite C first score ran without it, and suites A and B have no test files.
- Reachability comes only from the call graph export. Leave out `--callgraph` only for a run without reachability.
- Suite B entries are real dependency graphs, built with each package manager first: `scripts/run-crypto-finder-suite-b.sh CF_BIN PREDICTIONS_DIR LANG...` builds, scans and adapts every pair. It needs the Go, Python, Node.js and Cargo toolchains, a JDK 21 with Maven, and network access for the first build ([fixtures/README-suite-b.md](fixtures/README-suite-b.md)).
- Field mappings: [internal/adapter/cryptofinder/README.md](internal/adapter/cryptofinder/README.md).

**cdxgen and sonar-cryptography.** Any CycloneDX 1.6 or 1.7 CBOM with `evidence.occurrences` goes through the generic adapter ([internal/adapter/cbom/README.md](internal/adapter/cbom/README.md)):

```sh
go run ./cmd/cbench adapt cbom --tool NAME --tool-version V --entry ID --input bom.json \
  --out predictions/NAME/ID.json [--source-root DIR] [--cdxgen]
```

| Tool | Runner | Needs |
|---|---|---|
| cdxgen 13.6.0 | `scripts/run-cdxgen.sh ENTRY_DIR OUT`, then adapt with `--cdxgen --source-root "$(realpath ENTRY_DIR)"` | Node.js (`npm exec`) or `CDXGEN_BIN` |
| sonar-cryptography 1.7.0 | `scripts/run-sonar-cryptography.sh ENTRY_DIR OUT` | Docker only. Build a Java entry first (`mvn package dependency:copy-dependencies`) |

Both scripts refuse any other tool version and write the commands they ran to `OUT.command`.

### cbench commands

| Command | Does |
|---|---|
| `cbench validate [--overlay DIR]` | Checks the manifest and every ground-truth file. Exits non-zero on any error. |
| `cbench fetch [--only ID] [--dest DIR]` | Clones every URL entry at its pinned commit into `work/corpus/<id>`. Safe to rerun. |
| `cbench frame (--entry ID \| --dir DIR --language L --commit SHA) [--sample \| --grep]` | Prints a suite C checkout's crypto signal files, the seeded `files` sample, or the `pool` grep candidates. |
| `cbench adapt crypto-finder\|cbom ...` | Converts one tool run into a prediction file. |
| `cbench score --tool NAME [--predictions DIR] [--format json] [--out FILE]` | Validates the benchmark, then scores the predictions. |
| `cbench unique --run LABEL=DIR/TOOL --run ... [--out FILE]` | Counts, per run, the `pool` assets it matched and those no other compared run matched. |

## Repository layout

```
corpus/manifest.yaml              every corpus entry: id, language, suite, split, url+sha or path
corpus/libraries.yaml             the frozen suite A library list
corpus/suite-c.yaml               every suite C candidate examined and its outcome
groundtruth/<suite>/<entry>.yaml  annotated crypto assets for one entry
fixtures/<language>/...           in-repo programs for suites A and B
cmd/cbench                        the CLI
internal/score                    the matching rules and metrics
internal/adapter/<tool>           converters from each tool's output to predictions
results/<tool>/<run>/             published predictions, score reports and run records
docs/                             annotation guide, annotation log, strict-matching record
METHODOLOGY.md                    how ground truth is made and how tools are scored
scripts/                          tool runners, the scoreboard and check-public.sh
```

## Contributing

Found code where a crypto inventory tool, any tool, finds too little, finds too much or names the wrong algorithm? [Propose a case](https://github.com/scanoss/crypto-benchmark/issues/new?template=propose-case.yml) with a snippet and the crypto you expect, or [report a ground-truth error](https://github.com/scanoss/crypto-benchmark/issues/new?template=ground-truth-error.yml). You need no tool and no API key. Accepted cases join the existing suites in the next version of the benchmark, and published versions never change. No case is accepted or rejected because of what a tool finds on it. [CONTRIBUTING.md](CONTRIBUTING.md) has the how-to, and [METHODOLOGY.md](METHODOLOGY.md#versions-and-contributions) the rules.

## License

- Code (`cmd/`, `internal/`, `fixtures/`, scripts and workflows): [Apache License 2.0](LICENSE).
- Data (the corpus manifest, ground truth, methodology and published results): [CC BY 4.0](LICENSE-DATA).
- Repositories fetched by `cbench fetch` keep their own licenses and are never redistributed here.
