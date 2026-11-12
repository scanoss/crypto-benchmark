const forge = require('node-forge');

function main() {
  const keys = forge.pki.rsa.generateKeyPair({ bits: 2048, e: 0x10001 });
  const encrypted = keys.publicKey.encrypt('secret message', 'RSA-OAEP', {
    md: forge.md.sha256.create(),
    mgf1: { md: forge.md.sha256.create() },
  });
  const decrypted = keys.privateKey.decrypt(encrypted, 'RSA-OAEP', {
    md: forge.md.sha256.create(),
    mgf1: { md: forge.md.sha256.create() },
  });
  console.log('ok', decrypted === 'secret message');
}

main();
