# JavaScript suite A fixtures

One directory per row of the frozen JavaScript list in [corpus/libraries.yaml](../../corpus/libraries.yaml), plus `stdlib` for node:crypto and WebCrypto. Each `<primitive>/main.js` is a canonical program written the way the library's documentation shows it, and each directory pins exact versions in `package.json` and `package-lock.json`.

Run one directory on Node.js 24 from a scratch copy:

```sh
cp -r fixtures/javascript/jose /tmp/jose && cd /tmp/jose && npm ci && for m in */main.js; do node "$m"; done
```

Every program prints `ok` and a check value. A library's module system is the one it ships: ESM for jose, the noble packages, crypto-random-string, @peculiar/x509, iron-webcrypto and stdlib, CommonJS for the rest. No program needs TypeScript.

Notes on the rows that fit programs poorly:

- **blake3-wasm** is pinned at 2.1.5. The frozen row names 3.0.0, but 3.0.0 cannot be installed: it depends on `@c4312/blake3-internal@2.1.7`, which is not on the npm registry.
- **crypto-random-string** only draws random strings, so its entry is one program with three draws.
- **@peculiar/x509** needs a Reflect polyfill and pins `reflect-metadata` next to it.
- **jwa**, **jsonwebtoken** and **@peculiar/x509** generate their keys with node:crypto or WebCrypto, and **selfsigned** parses its certificate with node:crypto, as their documentation does, so those entries also annotate node:crypto and WebCrypto calls.
