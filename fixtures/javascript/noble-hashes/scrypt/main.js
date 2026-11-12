import { scrypt } from '@noble/hashes/scrypt.js';
import { bytesToHex, randomBytes } from '@noble/hashes/utils.js';

function main() {
  const salt = randomBytes(16);
  const key = scrypt('correct horse battery staple', salt, { N: 2 ** 17, r: 8, p: 1, dkLen: 32 });
  console.log('ok', bytesToHex(key).length);
}

main();
