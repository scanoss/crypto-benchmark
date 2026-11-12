import { p256 } from '@noble/curves/nist.js';
import { bytesToHex } from '@noble/curves/utils.js';

function main() {
  const alice = p256.keygen();
  const bob = p256.keygen();
  const aliceShared = p256.getSharedSecret(alice.secretKey, bob.publicKey);
  const bobShared = p256.getSharedSecret(bob.secretKey, alice.publicKey);
  console.log('ok', bytesToHex(aliceShared) === bytesToHex(bobShared));
}

main();
