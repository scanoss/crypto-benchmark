import { sha3_256 } from '@noble/hashes/sha3.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const digest = sha3_256(utf8ToBytes('benchmark message'));
  console.log('ok', bytesToHex(digest));
}

main();
