# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/openssl
evinse -i sbom.json -o work/cdxgen/rust-openssl.cbom.json -l rust fixtures/rust/openssl
