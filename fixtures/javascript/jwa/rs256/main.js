const crypto = require('node:crypto');
const jwa = require('jwa');

function main() {
  const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
    modulusLength: 2048,
    publicKeyEncoding: { type: 'spki', format: 'pem' },
    privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
  });
  const input = 'eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJiZW5jaCJ9';
  const signature = jwa('RS256').sign(input, privateKey);
  const valid = jwa('RS256').verify(input, signature, publicKey);
  console.log('ok', valid);
}

main();
