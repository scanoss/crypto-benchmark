# sonar-cryptography, suite A JavaScript pilot

Not run: sonar-cryptography does not support JavaScript.

sonar-cryptography 1.7.0 ships detection rules for Java, Python and Go only, and its README lists no JavaScript or TypeScript support. `scripts/run-sonar-cryptography.sh` matches that: it accepts `SONAR_CRYPTO_LANGUAGE` of `java`, `python` or `go`, and an entry without `go.mod` or a Maven or Gradle build file is scanned as Python. Running it on `fixtures/javascript/*` would scan `.js` files with the Python frontend and report nothing about them, which would measure the script, not the tool. No predictions or score are published, and the comparison in [results/javascript-pilot.md](../../javascript-pilot.md) lists the tool as not supported.
