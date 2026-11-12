const nacl = require('tweetnacl');

function main() {
  const alice = nacl.box.keyPair();
  const bob = nacl.box.keyPair();
  const nonce = nacl.randomBytes(nacl.box.nonceLength);
  const message = new TextEncoder().encode('secret message');

  const box = nacl.box(message, nonce, bob.publicKey, alice.secretKey);
  const opened = nacl.box.open(box, nonce, alice.publicKey, bob.secretKey);
  console.log('ok', new TextDecoder().decode(opened) === 'secret message');
}

main();
