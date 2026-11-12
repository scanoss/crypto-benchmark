const { subtle } = globalThis.crypto;

async function main() {
  const keyBytes = crypto.getRandomValues(new Uint8Array(32));
  const key = await subtle.importKey('raw', keyBytes, { name: 'HMAC', hash: 'SHA-256' }, false, ['sign', 'verify']);
  const data = new TextEncoder().encode('benchmark message');
  const mac = await subtle.sign('HMAC', key, data);
  const valid = await subtle.verify('HMAC', key, mac, data);
  console.log('ok', valid);
}

await main();
