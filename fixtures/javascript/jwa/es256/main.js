const crypto = require('node:crypto');
const jwa = require('jwa');

function main() {
  const { privateKey, publicKey } = crypto.generateKeyPairSync('ec', {
    namedCurve: 'P-256',
    publicKeyEncoding: { type: 'spki', format: 'pem' },
    privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
  });
  const ecdsa = jwa('ES256');
  const input = 'eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJiZW5jaCJ9';
  const signature = ecdsa.sign(input, privateKey);
  console.log('ok', ecdsa.verify(input, signature, publicKey));
}

main();
