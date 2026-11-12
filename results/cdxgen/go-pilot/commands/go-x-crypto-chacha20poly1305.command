# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-chacha20poly1305.sbom.json fixtures/go/x-crypto-chacha20poly1305
evinse -i work/cdxgen/go-x-crypto-chacha20poly1305.sbom.json -o work/cdxgen/go-x-crypto-chacha20poly1305.cbom.json -l go fixtures/go/x-crypto-chacha20poly1305
