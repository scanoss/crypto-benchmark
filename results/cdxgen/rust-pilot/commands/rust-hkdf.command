# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/hkdf
evinse -i sbom.json -o work/cdxgen/rust-hkdf.cbom.json -l rust fixtures/rust/hkdf
