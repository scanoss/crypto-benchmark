import { createHash, hash } from 'node:crypto';

function main() {
  const digest = createHash('sha256').update('benchmark message').digest('hex');
  const oneShot = hash('sha256', 'benchmark message');
  console.log('ok', digest === oneShot);
}

main();
