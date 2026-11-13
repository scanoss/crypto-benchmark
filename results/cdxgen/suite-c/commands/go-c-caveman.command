# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-c-caveman.sbom.json work/cdx-src/go-c-caveman
evinse -i work/cdxgen/go-c-caveman.sbom.json -o work/cdxgen/go-c-caveman.cbom.json -l go work/cdx-src/go-c-caveman
