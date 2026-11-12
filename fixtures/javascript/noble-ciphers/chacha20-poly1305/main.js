import { chacha20poly1305 } from '@noble/ciphers/chacha.js';
import { randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const nonce = randomBytes(12);
  const plaintext = utf8ToBytes('secret message');

  const cipher = chacha20poly1305(key, nonce);
  const ciphertext = cipher.encrypt(plaintext);
  const recovered = cipher.decrypt(ciphertext);
  console.log('ok', new TextDecoder().decode(recovered) === 'secret message');
}

main();
