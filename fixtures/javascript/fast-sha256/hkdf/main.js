const crypto = require('node:crypto');
const sha256 = require('fast-sha256');

function main() {
  const inputKey = crypto.randomBytes(32);
  const salt = crypto.randomBytes(16);
  const info = new TextEncoder().encode('application-key');
  const derived = sha256.hkdf(inputKey, salt, info, 32);
  console.log('ok', derived.length);
}

main();
