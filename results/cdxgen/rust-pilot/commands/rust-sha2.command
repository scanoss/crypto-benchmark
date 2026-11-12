# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/sha2
evinse -i sbom.json -o work/cdxgen/rust-sha2.cbom.json -l rust fixtures/rust/sha2
