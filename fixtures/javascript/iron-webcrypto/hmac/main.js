import * as Iron from 'iron-webcrypto';

async function main() {
  const password = process.env.IRON_PASSWORD ?? 'a_long_random_secret_please_change_me';
  const options = Iron.defaults.integrity;

  const { digest, salt } = await Iron.hmacWithPassword(password, options, 'benchmark message');
  const check = await Iron.hmacWithPassword(password, { ...options, salt }, 'benchmark message');
  console.log('ok', digest === check.digest);
}

await main();
