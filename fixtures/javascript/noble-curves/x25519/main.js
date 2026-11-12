import { x25519 } from '@noble/curves/ed25519.js';
import { bytesToHex } from '@noble/curves/utils.js';

function main() {
  const alice = x25519.keygen();
  const bob = x25519.keygen();
  const aliceShared = x25519.getSharedSecret(alice.secretKey, bob.publicKey);
  const bobShared = x25519.getSharedSecret(bob.secretKey, alice.publicKey);
  console.log('ok', bytesToHex(aliceShared) === bytesToHex(bobShared));
}

main();
