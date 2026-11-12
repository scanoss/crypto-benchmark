import { xchacha20poly1305 } from '@noble/ciphers/chacha.js';
import { managedNonce, randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const cipher = managedNonce(xchacha20poly1305)(key);
  const ciphertext = cipher.encrypt(utf8ToBytes('secret message'));
  const recovered = cipher.decrypt(ciphertext);
  console.log('ok', new TextDecoder().decode(recovered) === 'secret message');
}

main();
