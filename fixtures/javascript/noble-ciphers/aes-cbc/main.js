import { cbc } from '@noble/ciphers/aes.js';
import { randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const iv = randomBytes(16);
  const plaintext = utf8ToBytes('secret message');

  const ciphertext = cbc(key, iv).encrypt(plaintext);
  const recovered = cbc(key, iv).decrypt(ciphertext);
  console.log('ok', new TextDecoder().decode(recovered) === 'secret message');
}

main();
