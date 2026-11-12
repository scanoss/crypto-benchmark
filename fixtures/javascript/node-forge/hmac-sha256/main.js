const forge = require('node-forge');

function main() {
  const key = forge.random.getBytesSync(32);
  const hmac = forge.hmac.create();
  hmac.start('sha256', key);
  hmac.update('benchmark message', 'utf8');
  console.log('ok', hmac.digest().toHex());
}

main();
