import { createHmac, randomBytes, timingSafeEqual } from 'node:crypto';

function main() {
  const key = randomBytes(32);
  const mac = createHmac('sha256', key).update('benchmark message').digest();
  const expected = createHmac('sha256', key).update('benchmark message').digest();
  console.log('ok', timingSafeEqual(mac, expected));
}

main();
