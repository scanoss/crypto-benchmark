# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-pbkdf2.sbom.json fixtures/go/x-crypto-pbkdf2
evinse -i work/cdxgen/go-x-crypto-pbkdf2.sbom.json -o work/cdxgen/go-x-crypto-pbkdf2.cbom.json -l go fixtures/go/x-crypto-pbkdf2
