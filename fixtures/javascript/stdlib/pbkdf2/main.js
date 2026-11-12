import { pbkdf2Sync, randomBytes } from 'node:crypto';

function main() {
  const salt = randomBytes(16);
  const key = pbkdf2Sync('correct horse battery staple', salt, 600000, 32, 'sha256');
  console.log('ok', key.length);
}

main();
