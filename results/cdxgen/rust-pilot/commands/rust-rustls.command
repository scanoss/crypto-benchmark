# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/rustls
evinse -i sbom.json -o work/cdxgen/rust-rustls.cbom.json -l rust fixtures/rust/rustls
