import * as Iron from 'iron-webcrypto';

async function main() {
  const password = process.env.IRON_PASSWORD ?? 'a_long_random_secret_please_change_me';
  const payload = { userId: 123, scope: ['user'] };

  const sealed = await Iron.seal(payload, password, Iron.defaults);
  const unsealed = await Iron.unseal(sealed, password, Iron.defaults);
  console.log('ok', unsealed.userId === 123);
}

await main();
