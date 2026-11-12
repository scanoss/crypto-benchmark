# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-ripemd160.sbom.json fixtures/go/x-crypto-ripemd160
evinse -i work/cdxgen/go-x-crypto-ripemd160.sbom.json -o work/cdxgen/go-x-crypto-ripemd160.cbom.json -l go fixtures/go/x-crypto-ripemd160
