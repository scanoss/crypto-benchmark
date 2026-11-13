# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-c-ollama.sbom.json work/cdx-src/go-c-ollama
evinse -i work/cdxgen/go-c-ollama.sbom.json -o work/cdxgen/go-c-ollama.cbom.json -l go work/cdx-src/go-c-ollama
