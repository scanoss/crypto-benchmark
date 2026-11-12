import { diffieHellman, generateKeyPairSync } from 'node:crypto';

function main() {
  const alice = generateKeyPairSync('x25519');
  const bob = generateKeyPairSync('x25519');
  const aliceShared = diffieHellman({ privateKey: alice.privateKey, publicKey: bob.publicKey });
  const bobShared = diffieHellman({ privateKey: bob.privateKey, publicKey: alice.publicKey });
  console.log('ok', aliceShared.equals(bobShared));
}

main();
