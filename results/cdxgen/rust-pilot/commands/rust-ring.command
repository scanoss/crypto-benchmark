# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/ring
evinse -i sbom.json -o work/cdxgen/rust-ring.cbom.json -l rust fixtures/rust/ring
