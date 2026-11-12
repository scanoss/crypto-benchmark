# CycloneDX CBOM adapter

`cbench adapt cbom` turns one CycloneDX CBOM, written by any tool, into `predictions/<tool>/<entry>.json`.

```sh
cbench adapt cbom --tool <name> --tool-version <version> --entry <entry> --input bom.json \
  --out predictions/<name>/<entry>.json [--source-root <absolute scanned dir>] [--cdxgen]
```

## Input

A CycloneDX JSON document with `bomFormat: CycloneDX` and `specVersion` 1.6 or 1.7, the versions that define `cryptographic-asset` components. Any other version, or another format, fails the adapter. The document is not decoded strictly: a CBOM carries many fields this adapter does not need. Every `cryptographic-asset` component is read, at any nesting depth. Other components are read only to place crypto assets in a dependency.

`--tool` and `--tool-version` name the tool, since CBOM writers record themselves inconsistently under `metadata.tools`.

## Mapping

| Prediction field | Source | Notes |
|---|---|---|
| `tool.name`, `tool.version` | `--tool`, `--tool-version` | |
| `tool.config` | `--cdxgen` | `cdxgen_locations: "true"` when the flag is set, since it changes what is read. |
| `file`, `lines` | `evidence.occurrences[].location`, `.line` | One prediction per occurrence, at `[line, line]`. A relative location is taken as relative to the scanned directory. An absolute one is made relative to `--source-root`. `offset` is a column and `additionalContext` is free text (cdxgen writes the primitive, sonar-cryptography the matched API), so neither is read. |
| `asset_type` | `cryptoProperties.assetType` | Lowercased. A value that is not a CycloneDX asset type drops the component, with a note. |
| `primitive` | `cryptoProperties.algorithmProperties.primitive` | Lowercased. Algorithm assets only. Left out when it is not a CycloneDX primitive. |
| `algorithm` | `name` | Verbatim, for every asset type. `internal/algo` canonicalizes it when scoring, so `AES128-GCM`, `sha-256` and `AES/GCM/NoPadding` all work. A name such as `key@<uuid>` on related crypto material is kept too and only matters when the ground truth names an algorithm there. |
| `material_type` | `relatedCryptoMaterialProperties.type` | Lowercased. Related crypto material only. Left out when it is not a CycloneDX 1.7 material type. |
| `parameters.mode`, `parameters.padding` | `algorithmProperties.mode`, `.padding` | Verbatim (`gcm`, `pkcs5`). |
| `parameters.curve` | `algorithmProperties.ellipticCurve`, `.curve`, `.parameterSetIdentifier` | 1.7's `ellipticCurve` without its namespace (`nist/P-256` is `P-256`), else 1.6's `curve` verbatim, else the parameter set when it is a curve in the `internal/algo` curve table, canonicalized. |
| `parameters.key_size` | `algorithmProperties.parameterSetIdentifier`; `relatedCryptoMaterialProperties.size` | The parameter set only when it is a positive integer. sonar-cryptography also puts a hash's output size here (`SHA256` with 256); the extra parameter does not hurt attribution. |
| `parameters.hash` | none | CycloneDX 1.6 and 1.7 have no digest field in `algorithmProperties`. A digest in the name (`SHA256withECDSA`, `HMAC-SHA256`) is split out by the scorer. |
| `origin`, `component` | see below | |
| `reachable`, `depth` | none | `reachability_reported` is always `false` and every asset answers `unknown`. See below. |
| `metadata.components_without_location` | | Components with no usable location. |
| `metadata.locations_dropped` | | Locations that could not be matched: no line, or a path outside the scanned directory that no dependency claims. |

`algorithmProperties.classicalSecurityLevel` and `nistQuantumSecurityLevel` are strength ratings, not parameters the ground truth records, so they are not read. Neither are `oid`, `cryptoFunctions`, `certificateProperties` or `protocolProperties`.

### Components that cannot be matched

A component with no occurrence, or none with a location, cannot be matched to the ground truth. It produces no prediction, is counted in `metadata.components_without_location`, and is named on stderr. A location without a line, or one outside the scanned directory, is dropped the same way and counted in `metadata.locations_dropped`. The adapter never invents a location.

### Origin and component

Each location is placed by the first rule that decides it:

1. **Path inside a dependency.** The location lies under the source directory of exactly one library component. The file is then made relative to that directory, which is the benchmark's convention for dependency assets. Source directories come from `node_modules/<name>/` for npm package URLs (the last `node_modules` in the path), and, with `--cdxgen`, from the `internal:SrcGoMod` directory of a Go module that cdxgen marks as replaced by a local directory (`cdx:go:local_dir`).
2. **Nested.** The component sits in the `components` of a library component (`type` library or framework, with a package URL). The nearest such library is the component.
3. **Linked.** The `dependencies` section links the component, through `provides` or `dependsOn`, to exactly one library component and not to the project (`metadata.component`). Links between crypto assets are ignored.
4. Otherwise the origin is `project`.

The package URL is used verbatim. When two libraries claim the same path, or a crypto asset is linked to two libraries or to a library and the project, the origin is `project`. A dependency location that stays outside the scanned directory after these rules is dropped.

### Reachability

CycloneDX 1.6 and 1.7 have no reachability field for cryptographic assets, so the adapter reports none. It looked for these signals, and none qualifies for the languages this benchmark covers:

- `evidence.callstack` on a crypto component: a call path, not an answer to whether the asset is reachable from an entry point. cdxgen puts call stacks on library components, never on crypto assets.
- cdxgen `cdx:golem:crypto:usageScope` (`runtime`, `test`): the scope of the source file, not reachability.
- cdxgen `cdx:dosai:crypto:reachableFromEntryPoint` and `cdx:kosi:reachableFromRoots`: real reachability, but only from cdxgen's .NET (dosai) and Kotlin (kosi) analyzers, which this benchmark does not run.
- `vulnerabilities[].analysis` (VEX): applies to vulnerabilities, not to crypto assets.

## cdxgen (`--cdxgen`)

cdxgen writes source locations in two ways. Its JavaScript and TypeScript analysis (`cbom`, `cdxgen --include-crypto`) fills `evidence.occurrences` and needs no flag. Its per-language analyzers run by `evinse` (golem for Go, rusi for Rust) write no occurrences. They record each location in a property named `cdx:crypto:sourceLocation` or `cdx:<analyzer>:crypto:sourceLocation`, as `FILE:LINE:COLUMN`, `FILE:LINE` or `FILE#LINE`, and golem writes absolute paths. With `--cdxgen`, a component with no occurrences takes its locations from these properties, and Go modules replaced by a local directory get a source directory (rule 1 above). Pass `--source-root` with the directory cdxgen scanned so that absolute paths can be matched. `scripts/run-cdxgen.sh` writes the CBOM with the pinned cdxgen version and command; the top-level README shows the adapter call that follows.
