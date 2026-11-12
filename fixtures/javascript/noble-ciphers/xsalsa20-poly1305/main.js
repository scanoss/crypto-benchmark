import { secretbox } from '@noble/ciphers/salsa.js';
import { randomBytes, utf8ToBytes } from '@noble/ciphers/utils.js';

function main() {
  const key = randomBytes(32);
  const nonce = randomBytes(24);
  const box = secretbox(key, nonce);
  const sealed = box.seal(utf8ToBytes('secret message'));
  const opened = box.open(sealed);
  console.log('ok', new TextDecoder().decode(opened) === 'secret message');
}

main();
