# crypto-finder adapter

`cbench adapt crypto-finder` turns one crypto-finder scan into `predictions/crypto-finder/<entry>.json`.

```sh
crypto-finder scan --ruleset-version vX.Y.Z --scan-dependencies \
  --output interim.json --export-callgraph callgraph.json <entry checkout>
cbench adapt crypto-finder --entry <entry> --input interim.json --callgraph callgraph.json \
  --out predictions/crypto-finder/<entry>.json
```

## Inputs

- **Interim report** (`scan --output`), versions 1.5 to 1.9. The 1.8 `pqc` block and the 1.9 `rules.lease` are accepted and not scored. It is decoded strictly against the published schema (`pkg/schema/interim.go`, `schemas/interim-report-schema.json`). An undeclared property fails the adapter, because it means the contract moved.
- **Call graph export** (`scan --export-callgraph`), schema 6.17 or 6.18. Only `schema_version` and `finding_graphs[]` are read. The interim report carries no reachability, so without this file the prediction has `reachability_reported: false` and every asset answers `unknown`.

An interim asset is joined to its finding graph on `(finding_id, occurrence_key)`. `finding_id` alone is not unique: it hashes file, start line and rule.

## Mapping

| Prediction field | Source | Notes |
|---|---|---|
| `tool.version` | `tool.version` | |
| `tool.config` | `rules.source`, `rules.version`, `rules.checksum_sha256` | As `ruleset_source`, `ruleset_version`, `ruleset_checksum`. Empty values are left out; a local ruleset has no version. `rules.name` is not recorded: version and checksum identify the ruleset. |
| `file` | finding `file_path` | Relative to the scan target for project findings and to the dependency root for dependency findings, which is the benchmark's convention. Since crypto-finder 0.50.0 a PyPI dependency's path is relative to site-packages, where the wheel's `RECORD` names it (`libb/codec.py`), so the adapter takes it as written. Earlier versions wrote it relative to the top-level import package (`codec.py`); their output no longer matches PyPI dependency assets. |
| `lines` | `start_line`, `end_line` | |
| `asset_type` | `metadata.assetType` | Lowercased. A value that is not a CycloneDX asset type drops the asset, with a note on stderr. |
| `primitive` | `metadata.algorithmPrimitive` | Lowercased. Algorithm assets only. Left out when it is not a CycloneDX primitive. |
| `material_type` | `metadata.materialType` | Lowercased. Related crypto material only. Left out when it is not a CycloneDX 1.7 `relatedCryptoMaterialProperties.type`. |
| `algorithm` | `metadata.algorithmName`, else `metadata.algorithmFamily` | The specific name (`SHA-256`, `AES-256-GCM`, `ECDSA-secp256k1`) verbatim. The family (`SHA-2`) only when the rule sets no name; it rarely attributes, because the ground truth names the specific algorithm. An unresolved name such as `SHA-$bits` is kept as is and the scorer counts it as no algorithm. Key size, mode and curve inside the name are split out by the scorer (`internal/algo`), not here. |
| `parameters.mode` | `metadata.algorithmMode` | Verbatim. |
| `parameters.hash` | `metadata.algorithmHashFunction` | Verbatim. Algorithm assets only. The scorer canonicalizes it (`SHA256` is `SHA-256`), and also splits a digest out of the name (`ECDSA-SHA-256`). |
| `parameters.padding` | `metadata.algorithmPadding` | Verbatim, so a Java name such as `PKCS5Padding` is not turned into the CycloneDX `pkcs5`. |
| `parameters.key_size` | `metadata.algorithmParameterSetIdentifier` | Only when it is a positive integer. For related crypto material, `metadata.materialSize`. |
| `parameters.curve` | `metadata.algorithmParameterSetIdentifier`, `metadata.materialCurve` | For algorithm assets, the parameter set when it is a curve in the `internal/algo` curve table, canonicalized (`prime256v1` is `P-256`). Any other parameter set (`ML-KEM-768`) is left out. For related crypto material, `materialCurve` verbatim. |
| `origin` | `source` | `direct` is `project`. An empty source is `project` too, which is what crypto-finder itself defaults it to. `dependency` is `dependency`. Any other value drops the asset. |
| `component` | `dependency_info.purl` | A dependency finding with no package URL (an unknown ecosystem) is dropped with a note, since it cannot be matched to a component. |
| `reachable` | finding graph `reachability` | `reachable` is `true`, `unreachable` is `false`. `unknown`, `not_applicable` and a missing finding graph are `unknown`. |
| `depth` | finding graph `dependency` | Only with `reachable` `true` or `false`. A project finding is 0. A dependency finding is the length of `dependency.path` (direct dependency first, the finding's dependency last), capped at 2; without a path, `relationship` `direct` is 1 and `transitive` is 2; with neither, depth is left out. This is the distance in the resolved dependency graph, not the length of a call chain. |
| `distance` | finding graph `dependency.path` | The uncapped length of `dependency.path`, so app -> A -> B -> C is 3; 0 for a project finding. Left out when the depth came from `relationship` alone or is absent. |

`rules`, `match`, `oid`, columns, `parameter_conditions` and the rest of the metadata are not used. Per-value assets (`conditioned_value`) that crypto-finder reports at one call become separate predictions at the same lines; beyond the first match they score as duplicates.
