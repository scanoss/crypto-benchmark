import { generateKeyPairSync, sign, verify } from 'node:crypto';

function main() {
  const { privateKey, publicKey } = generateKeyPairSync('ed25519');
  const data = Buffer.from('benchmark message');
  const signature = sign(null, data, privateKey);
  const valid = verify(null, data, publicKey, signature);
  console.log('ok', valid);
}

main();
