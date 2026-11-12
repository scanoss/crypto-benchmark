# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-bcrypt.sbom.json fixtures/go/x-crypto-bcrypt
evinse -i work/cdxgen/go-x-crypto-bcrypt.sbom.json -o work/cdxgen/go-x-crypto-bcrypt.cbom.json -l go fixtures/go/x-crypto-bcrypt
