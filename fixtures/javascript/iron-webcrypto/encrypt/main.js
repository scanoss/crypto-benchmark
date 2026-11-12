import * as Iron from 'iron-webcrypto';

async function main() {
  const password = process.env.IRON_PASSWORD ?? 'a_long_random_secret_please_change_me';
  const options = Iron.defaults.encryption;

  const { encrypted, key } = await Iron.encrypt(password, options, 'secret message');
  const decrypted = await Iron.decrypt(password, { ...options, salt: key.salt, iv: key.iv }, encrypted);
  console.log('ok', decrypted === 'secret message');
}

await main();
