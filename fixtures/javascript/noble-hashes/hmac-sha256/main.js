import { hmac } from '@noble/hashes/hmac.js';
import { sha256 } from '@noble/hashes/sha2.js';
import { bytesToHex, randomBytes, utf8ToBytes } from '@noble/hashes/utils.js';

function main() {
  const key = randomBytes(32);
  const message = utf8ToBytes('benchmark message');
  const mac = hmac(sha256, key, message);
  const again = hmac.create(sha256, key).update(message).digest();
  console.log('ok', bytesToHex(mac) === bytesToHex(again));
}

main();
