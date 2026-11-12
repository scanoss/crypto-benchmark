import { hkdf } from '@noble/hashes/hkdf.js';
import { sha256 } from '@noble/hashes/sha2.js';
import { bytesToHex, randomBytes, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const inputKey = randomBytes(32);
  const salt = randomBytes(32);
  const info = utf8ToBytes('application-key');
  const derived = hkdf(sha256, inputKey, salt, info, 32);
  console.log('ok', bytesToHex(derived).length);
}

main();
