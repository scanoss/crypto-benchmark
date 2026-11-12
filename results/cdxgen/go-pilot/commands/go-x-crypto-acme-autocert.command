# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-acme-autocert.sbom.json fixtures/go/x-crypto-acme-autocert
evinse -i work/cdxgen/go-x-crypto-acme-autocert.sbom.json -o work/cdxgen/go-x-crypto-acme-autocert.cbom.json -l go fixtures/go/x-crypto-acme-autocert
