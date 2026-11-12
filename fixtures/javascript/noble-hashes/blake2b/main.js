import { blake2b } from '@noble/hashes/blake2.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const digest = blake2b(utf8ToBytes('benchmark message'), { dkLen: 32 });
  console.log('ok', bytesToHex(digest));
}

main();
