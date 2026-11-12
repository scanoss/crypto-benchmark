# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/siphasher
evinse -i sbom.json -o work/cdxgen/rust-siphasher.cbom.json -l rust fixtures/rust/siphasher
