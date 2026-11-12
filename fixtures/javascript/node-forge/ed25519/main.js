const forge = require('node-forge');

function main() {
  const ed25519 = forge.pki.ed25519;
  const { publicKey, privateKey } = ed25519.generateKeyPair();
  const signature = ed25519.sign({ message: 'benchmark message', encoding: 'utf8', privateKey });
  const valid = ed25519.verify({ message: 'benchmark message', encoding: 'utf8', signature, publicKey });
  console.log('ok', valid);
}

main();
