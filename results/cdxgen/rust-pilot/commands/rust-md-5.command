# CycloneDX Generator 13.6.0
cdxgen -t rust -o sbom.json fixtures/rust/md-5
evinse -i sbom.json -o work/cdxgen/rust-md-5.cbom.json -l rust fixtures/rust/md-5
