# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-golang-jwt-v5.sbom.json fixtures/go/golang-jwt-v5
evinse -i work/cdxgen/go-golang-jwt-v5.sbom.json -o work/cdxgen/go-golang-jwt-v5.cbom.json -l go fixtures/go/golang-jwt-v5
