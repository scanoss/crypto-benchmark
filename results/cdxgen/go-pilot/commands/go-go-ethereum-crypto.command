# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-go-ethereum-crypto.sbom.json fixtures/go/go-ethereum-crypto
evinse -i work/cdxgen/go-go-ethereum-crypto.sbom.json -o work/cdxgen/go-go-ethereum-crypto.cbom.json -l go fixtures/go/go-ethereum-crypto
