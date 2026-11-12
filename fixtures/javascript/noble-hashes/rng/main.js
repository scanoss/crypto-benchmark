import { bytesToHex, randomBytes } from '@noble/hashes/utils.js';

function main() {
  const token = randomBytes(32);
  console.log('ok', bytesToHex(token).length);
}

main();
