const hash = require('hash.js');

function main() {
  const key = process.env.HMAC_KEY || 'benchmark-hmac-key';
  const mac = hash.hmac(hash.sha256, key).update('benchmark message').digest('hex');
  console.log('ok', mac);
}

main();
