# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-x-crypto-ssh-agent.sbom.json fixtures/go/x-crypto-ssh-agent
evinse -i work/cdxgen/go-x-crypto-ssh-agent.sbom.json -o work/cdxgen/go-x-crypto-ssh-agent.cbom.json -l go fixtures/go/x-crypto-ssh-agent
