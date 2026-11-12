const crypto = require('node:crypto');
const sha256 = require('fast-sha256');

function main() {
  const key = crypto.randomBytes(32);
  const mac = sha256.hmac(key, new TextEncoder().encode('benchmark message'));
  console.log('ok', Buffer.from(mac).toString('hex'));
}

main();
