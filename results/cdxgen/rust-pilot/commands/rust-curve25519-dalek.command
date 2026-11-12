# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/curve25519-dalek
evinse -i sbom.json -o work/cdxgen/rust-curve25519-dalek.cbom.json -l rust fixtures/rust/curve25519-dalek
