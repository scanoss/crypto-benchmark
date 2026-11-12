const forge = require('node-forge');

function main() {
  const keys = forge.pki.rsa.generateKeyPair({ bits: 2048, e: 0x10001 });
  const privatePem = forge.pki.privateKeyToPem(keys.privateKey);
  const publicPem = forge.pki.publicKeyToPem(keys.publicKey);
  const privateKey = forge.pki.privateKeyFromPem(privatePem);
  const publicKey = forge.pki.publicKeyFromPem(publicPem);
  console.log('ok', privateKey.n.equals(publicKey.n));
}

main();
