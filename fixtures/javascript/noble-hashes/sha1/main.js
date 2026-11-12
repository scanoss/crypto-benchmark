import { sha1 } from '@noble/hashes/legacy.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const digest = sha1(utf8ToBytes('benchmark message'));
  console.log('ok', bytesToHex(digest));
}

main();
