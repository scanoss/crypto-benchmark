const { subtle } = globalThis.crypto;

async function main() {
  const keyPair = await subtle.generateKey({ name: 'ECDSA', namedCurve: 'P-384' }, true, ['sign', 'verify']);
  const data = new TextEncoder().encode('benchmark message');
  const signature = await subtle.sign({ name: 'ECDSA', hash: 'SHA-384' }, keyPair.privateKey, data);

  const spki = await subtle.exportKey('spki', keyPair.publicKey);
  const publicKey = await subtle.importKey('spki', spki, { name: 'ECDSA', namedCurve: 'P-384' }, true, ['verify']);
  const valid = await subtle.verify({ name: 'ECDSA', hash: 'SHA-384' }, publicKey, signature, data);
  console.log('ok', valid);
}

await main();
