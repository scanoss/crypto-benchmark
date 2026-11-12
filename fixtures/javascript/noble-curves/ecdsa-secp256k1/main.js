import { secp256k1 } from '@noble/curves/secp256k1.js';

function main() {
  const { secretKey, publicKey } = secp256k1.keygen();
  const message = new TextEncoder().encode('benchmark message');
  const signature = secp256k1.sign(message, secretKey);
  const valid = secp256k1.verify(signature, message, publicKey);
  console.log('ok', valid);
}

main();
