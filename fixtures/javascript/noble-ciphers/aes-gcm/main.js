import { gcm } from '@noble/ciphers/aes.js';
import { randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const nonce = randomBytes(12);
  const aad = utf8ToBytes('header');
  const plaintext = utf8ToBytes('secret message');

  const ciphertext = gcm(key, nonce, aad).encrypt(plaintext);
  const recovered = gcm(key, nonce, aad).decrypt(ciphertext);
  console.log('ok', new TextDecoder().decode(recovered) === 'secret message');
}

main();
