import { sha512 } from '@noble/hashes/sha2.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const digest = sha512(utf8ToBytes('benchmark message'));
  console.log('ok', bytesToHex(digest));
}

main();
