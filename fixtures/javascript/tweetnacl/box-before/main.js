const nacl = require('tweetnacl');

function main() {
  const alice = nacl.box.keyPair();
  const bob = nacl.box.keyPair();
  const aliceShared = nacl.box.before(bob.publicKey, alice.secretKey);
  const bobShared = nacl.box.before(alice.publicKey, bob.secretKey);
  const nonce = nacl.randomBytes(nacl.box.nonceLength);
  const message = new TextEncoder().encode('secret message');

  const box = nacl.box.after(message, nonce, aliceShared);
  const opened = nacl.box.open.after(box, nonce, bobShared);
  console.log('ok', new TextDecoder().decode(opened) === 'secret message');
}

main();
