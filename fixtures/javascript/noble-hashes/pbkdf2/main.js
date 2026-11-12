import { pbkdf2 } from '@noble/hashes/pbkdf2.js';
import { sha256 } from '@noble/hashes/sha2.js';
import { bytesToHex, randomBytes } from '@noble/hashes/utils.js';

function main() {
  const salt = randomBytes(16);
  const key = pbkdf2(sha256, 'correct horse battery staple', salt, { c: 600000, dkLen: 32 });
  console.log('ok', bytesToHex(key).length);
}

main();
