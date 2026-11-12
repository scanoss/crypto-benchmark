const nacl = require('tweetnacl');

function main() {
  const key = nacl.randomBytes(nacl.secretbox.keyLength);
  const nonce = nacl.randomBytes(nacl.secretbox.nonceLength);
  const message = new TextEncoder().encode('secret message');

  const box = nacl.secretbox(message, nonce, key);
  const opened = nacl.secretbox.open(box, nonce, key);
  console.log('ok', new TextDecoder().decode(opened) === 'secret message');
}

main();
