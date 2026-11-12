# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-hkdf.sbom.json fixtures/go/x-crypto-hkdf
evinse -i work/cdxgen/go-x-crypto-hkdf.sbom.json -o work/cdxgen/go-x-crypto-hkdf.cbom.json -l go fixtures/go/x-crypto-hkdf
