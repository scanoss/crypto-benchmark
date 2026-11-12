# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-stdlib.sbom.json fixtures/go/stdlib
evinse -i work/cdxgen/go-stdlib.sbom.json -o work/cdxgen/go-stdlib.cbom.json -l go fixtures/go/stdlib
