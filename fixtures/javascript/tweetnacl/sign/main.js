const nacl = require('tweetnacl');

function main() {
  const keyPair = nacl.sign.keyPair();
  const message = new TextEncoder().encode('benchmark message');

  const signedMessage = nacl.sign(message, keyPair.secretKey);
  const opened = nacl.sign.open(signedMessage, keyPair.publicKey);

  const signature = nacl.sign.detached(message, keyPair.secretKey);
  const valid = nacl.sign.detached.verify(message, signature, keyPair.publicKey);
  console.log('ok', opened !== null && valid);
}

main();
