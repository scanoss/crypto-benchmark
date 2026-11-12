import { xchacha20poly1305 } from '@noble/ciphers/chacha.js';
import { randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const nonce = randomBytes(24);
  const plaintext = utf8ToBytes('secret message');

  const ciphertext = xchacha20poly1305(key, nonce).encrypt(plaintext);
  const recovered = xchacha20poly1305(key, nonce).decrypt(ciphertext);
  console.log('ok', new TextDecoder().decode(recovered) === 'secret message');
}

main();
