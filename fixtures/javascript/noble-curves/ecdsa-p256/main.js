import { p256 } from '@noble/curves/nist.js';

function main() {
  const { secretKey, publicKey } = p256.keygen();
  const message = new TextEncoder().encode('benchmark message');
  const signature = p256.sign(message, secretKey);
  const valid = p256.verify(signature, message, publicKey);
  console.log('ok', valid);
}

main();
