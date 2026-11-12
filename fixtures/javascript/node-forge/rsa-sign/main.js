const forge = require('node-forge');

function main() {
  const keys = forge.pki.rsa.generateKeyPair({ bits: 2048, e: 0x10001 });
  const md = forge.md.sha256.create();
  md.update('benchmark message', 'utf8');
  const signature = keys.privateKey.sign(md);
  const valid = keys.publicKey.verify(md.digest().bytes(), signature);
  console.log('ok', valid);
}

main();
