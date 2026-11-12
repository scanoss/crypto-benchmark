import { createSign, createVerify, generateKeyPairSync } from 'node:crypto';

function main() {
  const { privateKey, publicKey } = generateKeyPairSync('ec', { namedCurve: 'P-256' });
  const signature = createSign('SHA256').update('benchmark message').sign(privateKey);
  const valid = createVerify('SHA256').update('benchmark message').verify(publicKey, signature);
  console.log('ok', valid);
}

main();
