import { sha256 } from '@noble/hashes/sha2.js';
import { bytesToHex, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const message = utf8ToBytes('benchmark message');
  const digest = sha256(message);

  const hasher = sha256.create();
  hasher.update(utf8ToBytes('benchmark '));
  hasher.update(utf8ToBytes('message'));
  const streamed = hasher.digest();

  console.log('ok', bytesToHex(digest) === bytesToHex(streamed));
}

main();
