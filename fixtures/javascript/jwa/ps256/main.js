const crypto = require('node:crypto');
const jwa = require('jwa');

function main() {
  const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
    modulusLength: 3072,
    publicKeyEncoding: { type: 'spki', format: 'pem' },
    privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
  });
  const pss = jwa('PS256');
  const input = 'eyJhbGciOiJQUzI1NiJ9.eyJzdWIiOiJiZW5jaCJ9';
  const signature = pss.sign(input, privateKey);
  console.log('ok', pss.verify(input, signature, publicKey));
}

main();
