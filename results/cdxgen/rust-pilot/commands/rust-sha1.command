# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/sha1
evinse -i sbom.json -o work/cdxgen/rust-sha1.cbom.json -l rust fixtures/rust/sha1
