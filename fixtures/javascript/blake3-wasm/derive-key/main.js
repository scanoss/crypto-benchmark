const crypto = require('node:crypto');
const { deriveKey } = require('blake3-wasm');

function main() {
  const material = crypto.randomBytes(32);
  const sessionKey = deriveKey('crypto-benchmark 2026-11-04 session key', material);
  console.log('ok', sessionKey.length);
}

main();
