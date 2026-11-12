import { blake3 } from '@noble/hashes/blake3.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const digest = blake3(utf8ToBytes('benchmark message'));
  console.log('ok', bytesToHex(digest));
}

main();
