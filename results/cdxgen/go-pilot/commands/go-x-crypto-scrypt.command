# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-scrypt.sbom.json fixtures/go/x-crypto-scrypt
evinse -i work/cdxgen/go-x-crypto-scrypt.sbom.json -o work/cdxgen/go-x-crypto-scrypt.cbom.json -l go fixtures/go/x-crypto-scrypt
