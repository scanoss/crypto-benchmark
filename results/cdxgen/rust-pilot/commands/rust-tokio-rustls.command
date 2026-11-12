# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/tokio-rustls
evinse -i sbom.json -o work/cdxgen/rust-tokio-rustls.cbom.json -l rust fixtures/rust/tokio-rustls
