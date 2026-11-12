import { createECDH } from 'node:crypto';

function main() {
  const alice = createECDH('prime256v1');
  const bob = createECDH('prime256v1');
  const alicePublic = alice.generateKeys();
  const bobPublic = bob.generateKeys();
  const aliceShared = alice.computeSecret(bobPublic);
  const bobShared = bob.computeSecret(alicePublic);
  console.log('ok', aliceShared.equals(bobShared));
}

main();
