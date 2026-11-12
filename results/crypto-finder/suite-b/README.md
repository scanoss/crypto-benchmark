# crypto-finder, suite B (transitive reachability)

crypto-finder v0.53.1 with remote ruleset v1.48.0, run on 2026-11-12 over the 90 suite B entries (45 pairs): **45 of 45 pairs** right, detection TP 104, FP 0, FN 0, reachable precision and recall 1.000, false reachable 0 / 52, depth and distance right on all 104 assets. `score.md` is the full `cbench score` report; [../../suite-b.md](../../suite-b.md) has the tables by language and scenario and the history of earlier runs.

## Run record

- **crypto-finder:** v0.53.1, the released `crypto-finder_0.53.1_linux_amd64.tar.gz` (sha256 `3a7db29d52a911fe8d02dec1bfc36aece2f64e66ab0d70f63a58ea3fd93bd8e8`, checked against the release's checksums file). The binary reports version `0.53.1`.
- **Ruleset:** remote, v1.48.0, checksum `sha256:a142fbe607446599dac42ec979b92d6610f96c9880cc5ca813f22ab449d8e3e6`, as the interim reports record it. The prediction files carry the same values in `tool.config`.
- **Toolchains:** Go 1.26.0, Python 3.12.3 (venv, pip 24.0), Maven 3.8.7 on Temurin JDK 21.0.12 (in `JAVA_HOME`), Node.js 24.19.0 with npm, cargo 1.98.0 (registry crates fetched with `cargo fetch --locked`, then built offline). The handler frameworks come from the public registries at the versions the fixtures pin.
- **Commands:** `scripts/run-crypto-finder-suite-b.sh CF PREDICTIONS go java javascript python rust`, with `CF` the v0.53.1 binary. For each entry it builds the fixture with `scripts/prepare-suite-b.sh` into `work/suite-b/<entry>` and runs:

```sh
env $(cat work/suite-b/ID/scan.env) crypto-finder scan --ruleset-version v1.48.0 \
  --scan-dependencies --no-dependency-findings-api --findings-cache none \
  --export-callgraph work/suite-b/ID/callgraph.json --output work/suite-b/ID/interim.json work/suite-b/ID/app
go run ./cmd/cbench adapt crypto-finder --entry ID --input work/suite-b/ID/interim.json \
  --callgraph work/suite-b/ID/callgraph.json --out PREDICTIONS/crypto-finder/ID.json
```

| Setting | Why |
|---|---|
| `scan.env` | Sets `VIRTUAL_ENV` to the entry's virtual environment for Python, and `MAVEN_OPTS=-Dmaven.repo.local=...` plus `JAVA_HOME` for Java. |
| `--no-dependency-findings-api` | Keeps the SCANOSS API from serving published findings for a dependency instead of scanning it. The fixture libraries are not published, so the flag makes that deterministic. |
| `--findings-cache none` | Keeps one entry's findings from serving another. |

Rust handler entries take 15 to 30 minutes each to scan, so the Rust entries were run with the runner's exact per-entry commands, six at a time.

## Out of scope

Predictions in third-party components that the ground truth does not annotate (for example pip's vendored packages, web framework internals, golang.org/x/net) are left out by `scope.components` and reported only as a count in `score.md`.
