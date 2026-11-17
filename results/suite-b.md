# Suite B: transitive reachability

## Summary

| | crypto-finder | sonar-cryptography | cdxgen |
|---|---|---|---|
| Tells crypto the app really uses apart from crypto it never calls, in its dependencies | **45 of 45 pairs** | no, it does not report reachability | no, it does not report reachability |
| Unused crypto wrongly flagged as used | **0 of 52** | n/a | n/a |
| Crypto found inside dependencies | **104 of 104 (100%)** | 0 of 52 (0%) | 16 of 84 (19%), JavaScript only |

sonar-cryptography and cdxgen ran on the first 35 pairs, before the handler and callback scenarios were added, so their counts have smaller totals. The sections below give the details.

## Scores in detail

crypto-finder v0.54.0 with remote ruleset v1.48.0 (2026-11-17) discriminates **all 45 pairs**: it detects all 104 suite B assets with no false positive, answers reachable with precision and recall 1.000, calls no unreachable asset reachable (0 / 52), and names the right dependency depth and distance on every asset. cdxgen and sonar-cryptography report no reachability for crypto assets, so they are shown as having none, not as scoring zero.

## What suite B measures

Whether a tool knows that an application can reach crypto that lives in its dependencies. Each pair is an application that depends on library A, which may depend on libraries B and C, where a library calls a real crypto library. The reachable and unreachable variants share the dependency graph and the crypto and differ in one change to the application ([fixtures/README-suite-b.md](../fixtures/README-suite-b.md)). The libraries are installed through each ecosystem's package manager, the way published packages are.

| Scenario | Languages | What the one change tests |
|---|---|---|
| direct, transitive, depth3, wrapper | all five | whether the app calls into the chain at all, at dependency depth 1, 2 and 3, and through a wrapper |
| deadfn | all five | a dead function inside a reachable file |
| unused-import | all five | a dependency that is imported but never called |
| dispatch | all five | a call through an interface, trait or base-class method |
| handler | all five, one pair per web framework (Go net/http and gin, Python Flask and FastAPI, Java Spring Boot, JavaScript Express, Rust axum and actix-web) | a handler the framework registers against one it never registers |
| callback | Go, Python | a function passed by value to the standard library against one never passed |

45 pairs, 90 entries, 104 ground-truth assets, annotated twice independently ([annotation log](../docs/annotation-log.md#suite-b-all-five-languages-2026-11-05)). Every ground-truth asset counts, detected or not, so a tool cannot raise its reachable recall by detecting less ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules), "Suite B reachability").

## crypto-finder, current run

v0.54.0, remote ruleset v1.48.0, `scan --scan-dependencies --export-callgraph` ([run record](crypto-finder/suite-b/README.md)). `score.md` next to the predictions is the full report.

| Group | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs correct |
|---|---|---|---|---|---|---|
| go | 1.000 | 1.000 | 0 / 11 | 22 / 22 | 22 / 22 | 10 / 10 |
| java | 1.000 | 1.000 | 0 / 11 | 22 / 22 | 22 / 22 | 8 / 8 |
| javascript | 1.000 | 1.000 | 0 / 9 | 18 / 18 | 18 / 18 | 8 / 8 |
| python | 1.000 | 1.000 | 0 / 11 | 22 / 22 | 22 / 22 | 10 / 10 |
| rust | 1.000 | 1.000 | 0 / 10 | 20 / 20 | 20 / 20 | 9 / 9 |
| all | 1.000 | 1.000 | 0 / 52 | 104 / 104 | 104 / 104 | 45 / 45 |

| Scenario | Reachable P | Reachable R | False reachable | Depth acc | Distance acc | Pairs correct |
|---|---|---|---|---|---|---|
| callback | 1.000 | 1.000 | 0 / 2 | 4 / 4 | 4 / 4 | 2 / 2 |
| deadfn | 1.000 | 1.000 | 0 / 11 | 22 / 22 | 22 / 22 | 5 / 5 |
| depth3 | 1.000 | 1.000 | 0 / 6 | 12 / 12 | 12 / 12 | 5 / 5 |
| direct | 1.000 | 1.000 | 0 / 5 | 10 / 10 | 10 / 10 | 5 / 5 |
| dispatch | 1.000 | 1.000 | 0 / 5 | 10 / 10 | 10 / 10 | 5 / 5 |
| handler | 1.000 | 1.000 | 0 / 8 | 16 / 16 | 16 / 16 | 8 / 8 |
| transitive | 1.000 | 1.000 | 0 / 5 | 10 / 10 | 10 / 10 | 5 / 5 |
| unused-import | 1.000 | 1.000 | 0 / 5 | 10 / 10 | 10 / 10 | 5 / 5 |
| wrapper | 1.000 | 1.000 | 0 / 5 | 10 / 10 | 10 / 10 | 5 / 5 |

## History of pinned runs

crypto-finder, all 45 pairs. Detection was complete from v0.51.0 on: TP 104, FP 0, FN 0.

| crypto-finder | Ruleset | Run | Pairs correct | Reachable P / R | False reachable | Detection TP / FP / FN |
|---|---|---|---|---|---|---|
| v0.50.0 | v1.42.0 | 2026-11-06 | 27 of 45 | 0.754 / 0.915 | 14 / 52 | 104 / 22 / 0 |
| v0.51.0 | v1.43.0 | 2026-11-09 | 40 of 45 | 0.977 / 0.915 | 1 / 52 | 104 / 0 / 0 |
| v0.51.0 | v1.44.0, v1.45.0, v1.46.0 | 2026-11-10, 2026-11-11 | 40 of 45 | 0.977 / 0.915 | 1 / 52 | 104 / 0 / 0 |
| v0.53.1 | v1.47.0 | 2026-11-12 | 45 of 45 | 1.000 / 1.000 | 0 / 52 | 104 / 0 / 0 |
| v0.53.1 | v1.48.0 | 2026-11-12 | 45 of 45 | 1.000 / 1.000 | 0 / 52 | 104 / 0 / 0 |
| v0.54.0 | v1.48.0 | 2026-11-17 | 45 of 45 | 1.000 / 1.000 | 0 / 52 | 104 / 0 / 0 |

What moved, in user terms:

- **v0.51.0:** a handler or callback that is never registered or passed is no longer treated as an entry point just because its name appears elsewhere (a route string, a comment, another module's qualified name, a Rust lint attribute), and the 22 digest and tag outputs that v0.50.0 reported as key material are gone. Every handler and callback pair became right; the five left were dispatch.
- **v0.53.1:** a method call on a local bound to a factory's result resolves to the class the factory constructs, so all five dispatch pairs became right.
- **Rulesets v1.44.0 to v1.48.0** changed no answer on a scored suite B asset.
- **v0.54.0:** only the ruleset download changed, to a leased bundle from `/v3/cryptography/rulesets`; every prediction is the same as v0.53.1's.

## Other tools

Run on 2026-11-05 on the first 35 pairs (the handler and callback scenarios came later), with the pinned suite A configurations. Neither reports reachability for crypto assets, so the only measurable question is whether each finds the dependency crypto at all.

| Language | Assets | cdxgen 13.6.0 | sonar-cryptography 1.7.0 |
|---|---|---|---|
| Go | 16 | 0 | 0 |
| Python | 16 | 0 | 0 |
| Java | 20 | 0 | 0 |
| JavaScript | 16 | 16 (precision 1.000), the same assets in both variants of every pair | not run (no JavaScript support) |
| Rust | 16 | 0 | not run (no Rust support) |

- **cdxgen** analyzes `node_modules` for JavaScript and so finds the dependency crypto there, with no way to tell reachable from unreachable. For Go, golem's default package pattern covers the application module only; with `--golem-patterns all` it finds 14 of the 16 Go assets, under versionless package URLs that do not match the dependency components. rusi analyzes the workspace crates only. For Python and Java it inventories libraries, not algorithm use. Details: [cdxgen/suite-b](cdxgen/suite-b/README.md).
- **sonar-cryptography** analyzes the scanned project's source only, and every suite B asset is in a dependency, so its recall here is 0 by design. Details: [sonar-cryptography/suite-b](sonar-cryptography/suite-b/README.md).

## Limitations

- The fixtures are small and uniform by intent: one crypto call per function, the standard crypto library where the language has one. They isolate reachability; they do not stand for the size or shape of real applications, which suite C covers.
- Each scenario fixes one construct. A tool can be right on a scenario here and wrong on a variant of the same construct elsewhere.
- Ecosystem mechanisms are realistic but local: Go `replace`, local wheels, a local Maven repository with sources jars, npm tarballs with `overrides`, Cargo path dependencies. None of the fixture libraries comes from a public registry, so a tool that only consults a registry would see none of them. crypto-finder ran with `--no-dependency-findings-api`.
- The ground truth's call `path` is recorded but not scored.
- The fixture author saw crypto-finder's output on early Go and Python fixtures. No fixture was changed to fit a result, and every miss of every run is published.
