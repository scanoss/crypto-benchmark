# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-curve25519.sbom.json fixtures/go/x-crypto-curve25519
evinse -i work/cdxgen/go-x-crypto-curve25519.sbom.json -o work/cdxgen/go-x-crypto-curve25519.cbom.json -l go fixtures/go/x-crypto-curve25519
