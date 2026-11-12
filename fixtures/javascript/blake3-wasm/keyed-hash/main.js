const crypto = require('node:crypto');
const { keyedHash } = require('blake3-wasm');

function main() {
  const key = crypto.randomBytes(32);
  const mac = keyedHash(key, 'benchmark message');
  console.log('ok', mac.toString('hex'));
}

main();
