# CycloneDX Generator 13.6.0
cdxgen -t go -o work/cdxgen/go-dgrijalva-jwt-go.sbom.json fixtures/go/dgrijalva-jwt-go
evinse -i work/cdxgen/go-dgrijalva-jwt-go.sbom.json -o work/cdxgen/go-dgrijalva-jwt-go.cbom.json -l go fixtures/go/dgrijalva-jwt-go
