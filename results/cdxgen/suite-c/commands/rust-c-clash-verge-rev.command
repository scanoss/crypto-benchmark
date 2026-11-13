# CycloneDX Generator 13.6.0
cdxgen -t rust -o work/cdxgen/rust-c-clash-verge-rev.sbom.json work/cdx-src/rust-c-clash-verge-rev
evinse -i work/cdxgen/rust-c-clash-verge-rev.sbom.json -o work/cdxgen/rust-c-clash-verge-rev.cbom.json -l rust work/cdx-src/rust-c-clash-verge-rev
