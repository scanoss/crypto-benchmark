import { ed25519 } from '@noble/curves/ed25519.js';

function main() {
  const { secretKey, publicKey } = ed25519.keygen();
  const message = new TextEncoder().encode('benchmark message');
  const signature = ed25519.sign(message, secretKey);
  const valid = ed25519.verify(signature, message, publicKey);
  console.log('ok', valid);
}

main();
