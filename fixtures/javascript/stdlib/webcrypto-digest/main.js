const { subtle } = globalThis.crypto;

async function main() {
  const data = new TextEncoder().encode('benchmark message');
  const sha256 = await subtle.digest('SHA-256', data);
  const sha512 = await subtle.digest('SHA-512', data);
  console.log('ok', sha256.byteLength, sha512.byteLength);
}

await main();
