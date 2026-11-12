import { hkdfSync, randomBytes } from 'node:crypto';

function main() {
  const inputKey = randomBytes(32);
  const salt = randomBytes(16);
  const derived = hkdfSync('sha256', inputKey, salt, 'application-key', 32);
  console.log('ok', derived.byteLength);
}

main();
