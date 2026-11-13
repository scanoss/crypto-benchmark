# CycloneDX Generator 13.6.0
cdxgen -t rust -o work/cdxgen/rust-c-cc-switch.sbom.json work/cdx-src/rust-c-cc-switch/src-tauri
evinse -i work/cdxgen/rust-c-cc-switch.sbom.json -o work/cdxgen/rust-c-cc-switch.cbom.json -l rust work/cdx-src/rust-c-cc-switch/src-tauri
