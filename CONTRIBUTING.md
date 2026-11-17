# Contributing

The benchmark gets better with cases its authors did not think of. If you know code where a crypto inventory tool, any tool, finds too little, finds too much or names the wrong algorithm, we want it here. A case that crypto-finder misses is as welcome as any other.

Accepted contributions join the existing suites in the next version of the benchmark. Published versions never change. [METHODOLOGY.md](METHODOLOGY.md#versions-and-contributions) has the rules. This page is the how-to.

## Two ways to contribute

| You have | Do this | Effort |
|---|---|---|
| A snippet or a repository file where you expect a tool to find crypto | [Propose a case](https://github.com/scanoss/crypto-benchmark/issues/new?template=propose-case.yml): paste the code and say what crypto it uses. We write the fixture and the ground truth. | Minutes |
| A ground-truth asset you think is wrong or missing | [Report a ground-truth error](https://github.com/scanoss/crypto-benchmark/issues/new?template=ground-truth-error.yml): name the entry, the file, the lines and the guide rule. | Minutes |
| Time to write the case yourself | Open a pull request with the program and its ground truth, as below. | An hour or two |

You need no API key and no tool to contribute. Please do not attach tool output to a proposal: we finalize a case's ground truth before any tool runs on it.

## What we accept

- **Suite A, library coverage.** A small program that uses a crypto library or a standard library's crypto API, for Java, Python, Go, Rust or JavaScript/TypeScript. It can be a library on the [list](corpus/libraries.yaml) or one that is not.
- **Suite B, transitive reachability.** A new scenario, or an existing one in another language or framework, as a reachable and an unreachable pair that differ in one change ([fixtures/README-suite-b.md](fixtures/README-suite-b.md)).
- **Ground-truth corrections** in any suite, including the suite C real repositories.
- **Annotation guide gaps**, where the [guide](docs/annotation-guide.md) does not say whether something is an asset.

We do not take suite C repository suggestions. Those repositories come only from the [selection rule](METHODOLOGY.md#suite-c-selection), so no one, us included, can pick a repository after seeing how a tool does on it.

## Writing a suite A case

Read the [annotation guide](docs/annotation-guide.md) first. It decides what an asset is, which lines it takes and which fields it carries.

1. **The program.** Add a directory under `fixtures/<language>/<library>/<case>/`, one directory per primitive or use. Keep it small, runnable, and free of network access and real secrets. Pin every dependency to an exact version, the way the existing fixtures do: `requirements.txt` with `==`, `package.json` with an exact version and a committed `package-lock.json`, `Cargo.toml` with `=` and a committed `Cargo.lock`, `go.mod` with `go.sum`, `pom.xml` with fixed versions.
2. **The manifest.** For a library that already has an entry, nothing changes. For a new library, add an entry to `corpus/manifest.yaml`:

   ```yaml
   - id: python-example-lib
     language: python
     suite: A
     split: dev
     path: fixtures/python/example-lib
     notes: "Suite A library coverage: example-lib==1.2.3, one program per primitive. Contributed."
   ```

3. **The ground truth.** List every asset of every file of the entry in `groundtruth/A/<entry-id>.yaml`. For an existing entry, append your assets with new ids. One asset looks like this:

   ```yaml
   - id: gt-0042
     file: hmac-sha256/main.py
     lines: [7, 7]
     asset_type: algorithm
     primitive: mac
     algorithm: HMAC
     parameters:
       hash: SHA-256
     origin: project
     reachable: true
     depth: 0
     notes: "Why this call is an asset, when it is not obvious from the guide."
   ```

   The schema is the Go type in `internal/groundtruth`, and the existing files under `groundtruth/A/` show every field in use.

4. **Check it.**

   ```sh
   go run ./cmd/cbench validate   # the manifest and every ground-truth file
   go test ./...
   scripts/check-public.sh        # nothing private, no keys
   ```

5. **Open the pull request** and fill in its template.

## Writing a suite B pair

Follow [fixtures/README-suite-b.md](fixtures/README-suite-b.md): the libraries, `app-reachable/` and `app-unreachable/` under `fixtures/<language>/suite-b/<scenario>/`, two manifest entries with `pair`, `variant` and `scenario`, and two ground-truth files under `groundtruth/B/` with the [suite B fields](docs/annotation-guide.md#suite-b), including `reachable`, `distance` and `path`. `scripts/prepare-suite-b.sh` builds a pair the way the scans see it.

## What happens next

1. **Review.** A maintainer checks that the programs build and run, and that the case belongs in the suite it targets.
2. **Second annotation.** Your ground truth is the first of the [two annotations](METHODOLOGY.md#ground-truth-process). A second annotator labels the same files without seeing yours. Every disagreement is settled against the guide, and the resolution goes into the asset's `notes`. We may push those changes to your branch, or ask you about them.
3. **Merge.** The case lands on `main` with an entry in [docs/annotation-log.md](docs/annotation-log.md) that credits you.
4. **Release.** At the next version's release, every tool runs on the whole corpus, your case included, and the results are published under that version.

A case is never accepted, changed or rejected because of what a tool finds on it.

## License

By contributing, you agree that your code is released under the [Apache License 2.0](LICENSE) and your data (manifest entries, ground truth, documentation) under [CC BY 4.0](LICENSE-DATA), like the rest of the repository.
