# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/aes
evinse -i sbom.json -o work/cdxgen/rust-aes.cbom.json -l rust fixtures/rust/aes
