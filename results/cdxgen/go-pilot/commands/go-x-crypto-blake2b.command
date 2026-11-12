# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-blake2b.sbom.json fixtures/go/x-crypto-blake2b
evinse -i work/cdxgen/go-x-crypto-blake2b.sbom.json -o work/cdxgen/go-x-crypto-blake2b.cbom.json -l go fixtures/go/x-crypto-blake2b
