import { generateKeyPairSync, sign, verify } from 'node:crypto';

function main() {
  const { privateKey, publicKey } = generateKeyPairSync('rsa', { modulusLength: 2048 });
  const data = Buffer.from('benchmark message');
  const signature = sign('sha256', data, privateKey);
  const valid = verify('sha256', data, publicKey, signature);
  console.log('ok', valid);
}

main();
