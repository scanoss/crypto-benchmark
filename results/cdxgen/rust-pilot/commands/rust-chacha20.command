# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/chacha20
evinse -i sbom.json -o work/cdxgen/rust-chacha20.cbom.json -l rust fixtures/rust/chacha20
