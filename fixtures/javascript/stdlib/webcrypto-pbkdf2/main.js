const { subtle } = globalThis.crypto;

async function main() {
  const password = new TextEncoder().encode('correct horse battery staple');
  const baseKey = await subtle.importKey('raw', password, 'PBKDF2', false, ['deriveBits']);
  const salt = crypto.getRandomValues(new Uint8Array(16));
  const bits = await subtle.deriveBits({ name: 'PBKDF2', salt, iterations: 600000, hash: 'SHA-256' }, baseKey, 256);
  console.log('ok', bits.byteLength);
}

await main();
