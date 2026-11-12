const crypto = require('node:crypto');
const sha256 = require('fast-sha256');

function main() {
  const password = new TextEncoder().encode('correct horse battery staple');
  const salt = crypto.randomBytes(16);
  const key = sha256.pbkdf2(password, salt, 600000, 32);
  console.log('ok', key.length);
}

main();
