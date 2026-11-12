# image ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b
# sonar-cryptography-plugin 1.7.0, cbomkit-lib 1.3.0, CBOMkit-action v2.3.0
# language python
cp -a ENTRY_DIR WS
touch WS/setup.cfg
docker run --rm --network none --user "$(id -u):$(id -g)" -e HOME=/tmp -w /scan/ws -v WS:/scan/ws:ro -v OUTDIR:/scan/out -e GITHUB_WORKSPACE=/scan/ws -e CBOMKIT_OUTPUT_DIR=/scan/out -e CBOMKIT_LANGUAGES=python -e CBOMKIT_GENERATE_MODULE_CBOMS=false -e CBOMKIT_EXCLUDE=\(\?\!\) ghcr.io/cbomkit/cbomkit-action@sha256:47c2505267dc5f055152616ffa0d0076b7d0824fe9e0de670b572d1ebad2fc4b
cp OUTDIR/cbom.json OUT
