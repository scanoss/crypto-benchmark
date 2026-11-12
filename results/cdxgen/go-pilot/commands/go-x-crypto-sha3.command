# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-sha3.sbom.json fixtures/go/x-crypto-sha3
evinse -i work/cdxgen/go-x-crypto-sha3.sbom.json -o work/cdxgen/go-x-crypto-sha3.cbom.json -l go fixtures/go/x-crypto-sha3
