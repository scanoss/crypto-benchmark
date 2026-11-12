# cdxgen, suite B (transitive reachability)

Run on 2026-11-05 over the 70 suite B entries, each built first with `scripts/prepare-suite-b.sh` into `work/cdx-prep/<entry>`, the same build crypto-finder scanned. `score.md` is the `cbench score` report for the pinned configuration; `golem-all/` is a second Go run that also analyzes dependency modules.

**cdxgen reports no reachability for crypto assets.** Its CBOM carries no reachable flag on a cryptographic asset in any mode we found, so every prediction answers `unknown` and the score shows cdxgen as having no reachability, not as scoring zero. What it can be measured on is whether it finds the crypto in the dependencies at all.

## Run record

- **cdxgen 13.6.0** through `scripts/run-cdxgen.sh` with `CDXGEN_BIN` pointing at an install of the same version. Each entry's commands are in `work/cdx-prep/<entry>/cdxgen.cbom.json.command`: `cdxgen -t go` then `evinse -l go` for Go, `cdxgen -t rust` then `evinse -l rust` for Rust, `cbom` for Python, Java and JavaScript, all on the built application directory (`work/cdx-prep/<entry>/app`), which holds `node_modules` for JavaScript. The configuration is the one [METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration) records.
- **Adapter:** `cbench adapt cbom --tool cdxgen --tool-version 13.6.0 --cdxgen --source-root <app>`.
- **golem-all:** for Go only, `cdxgen -t go -o sbom.json <app>` then `evinse -i sbom.json -o OUT -l go --golem-patterns all <app>`, adapted the same way.

## Results

| Language | Assets | Detected | Reachability |
|---|---|---|---|
| JavaScript | 16 | 16 (P 1.000) | none reported |
| Go | 16 | 0 | none reported |
| Rust | 16 | 0 | none reported |
| Python | 16 | 0 | none reported |
| Java | 20 | 0 | none reported |

- **JavaScript.** `cbom` analyzes `node_modules`, and the adapter attributes each location to its package (`node_modules/cbench-deadfn-libb/codec.js` is `codec.js` in `pkg:npm/cbench-deadfn-libb@1.0.0`). It finds all 16 assets, and the same assets in both variants of every pair: the dead SHA-1 and the crypto behind uncalled functions are reported like the rest, with no way to tell them apart.
- **Go.** golem's default package pattern, `./...`, covers the application module only, which holds no crypto. With `--golem-patterns all` (`golem-all/`) golem finds the crypto in 14 of the 16 dependency assets (all but the wrapper scenario's SHA-384), with locations under each library and the property `cdx:golem:crypto:usageScope: runtime` in both variants. It names a replaced module without its version (`pkg:golang/example.com/cbench/transitive/libb`), and the score compares components as exact strings after dropping qualifiers and subpath ([METHODOLOGY.md](../../../METHODOLOGY.md#matching-rules)), so those 22 predictions are outside `scope.components` and none matches. The 14 is context, not a score.
- **Rust.** rusi analyzes the workspace crates, here the application only, so it reports nothing from the path dependencies.
- **Python and Java.** cdxgen 13.6.0's `cbom` inventories crypto libraries, keystores and certificates for these languages but not algorithm use in source ([METHODOLOGY.md](../../../METHODOLOGY.md#competitor-configuration)), so it scores 0 as in suite A.

## Reachability modes checked

cdxgen's evinse has an evidence "research" profile and `--with-reachables` ("auto-tagged reachable slicing", which requires an SBOM made with `--deep`; `evinse --help` of 13.6.0). On `javascript-b-transitive-unreachable`, `cdxgen -t js --deep --profile research` then `evinse -l js --with-reachables --profile research` produced usage evidence (`occurrences`) on the two library components and no cryptographic-asset components; atom produced no callsites slice. The reachables slice attaches call-stack evidence to packages, not to crypto assets, so there is still nothing per asset to score. golem's `usageScope` separates runtime code from test code and reads `runtime` in both variants of every Go pair.
