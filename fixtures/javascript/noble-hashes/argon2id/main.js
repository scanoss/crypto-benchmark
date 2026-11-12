import { argon2id } from '@noble/hashes/argon2.js';
import { bytesToHex, randomBytes } from '@noble/hashes/utils.js';

function main() {
  const salt = randomBytes(16);
  const key = argon2id('correct horse battery staple', salt, { t: 2, m: 19456, p: 1, dkLen: 32 });
  console.log('ok', bytesToHex(key).length);
}

main();
