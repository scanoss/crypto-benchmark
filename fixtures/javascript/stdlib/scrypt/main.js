import { randomBytes, scryptSync } from 'node:crypto';

function main() {
  const salt = randomBytes(16);
  const key = scryptSync('correct horse battery staple', salt, 32, { N: 2 ** 17, r: 8, p: 1, maxmem: 256 * 1024 * 1024 });
  console.log('ok', key.length);
}

main();
