import { randomBytes, randomInt, randomUUID } from 'node:crypto';

function main() {
  const token = randomBytes(32).toString('hex');
  const requestId = randomUUID();
  const pin = randomInt(0, 1000000);
  const challenge = crypto.getRandomValues(new Uint8Array(16));
  console.log('ok', token.length, requestId.length, pin >= 0, challenge.length);
}

main();
