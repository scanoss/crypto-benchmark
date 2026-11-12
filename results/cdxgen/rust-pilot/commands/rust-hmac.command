# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/hmac
evinse -i sbom.json -o work/cdxgen/rust-hmac.cbom.json -l rust fixtures/rust/hmac
