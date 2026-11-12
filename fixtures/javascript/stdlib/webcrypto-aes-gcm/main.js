const { subtle } = globalThis.crypto;

async function main() {
  const key = await subtle.generateKey({ name: 'AES-GCM', length: 256 }, true, ['encrypt', 'decrypt']);
  const iv = crypto.getRandomValues(new Uint8Array(12));
  const data = new TextEncoder().encode('secret message');

  const ciphertext = await subtle.encrypt({ name: 'AES-GCM', iv }, key, data);
  const plaintext = await subtle.decrypt({ name: 'AES-GCM', iv }, key, ciphertext);
  const rawKey = await subtle.exportKey('raw', key);
  console.log('ok', new TextDecoder().decode(plaintext) === 'secret message', rawKey.byteLength);
}

await main();
