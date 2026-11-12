# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/rustls-webpki
evinse -i sbom.json -o work/cdxgen/rust-rustls-webpki.cbom.json -l rust fixtures/rust/rustls-webpki
