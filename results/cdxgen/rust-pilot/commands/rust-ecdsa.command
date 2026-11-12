# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/ecdsa
evinse -i sbom.json -o work/cdxgen/rust-ecdsa.cbom.json -l rust fixtures/rust/ecdsa
