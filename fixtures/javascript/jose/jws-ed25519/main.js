import { CompactSign, compactVerify, generateKeyPair } from 'jose';

async function main() {
  const { publicKey, privateKey } = await generateKeyPair('Ed25519');
  const payload = new TextEncoder().encode('benchmark message');
  const jws = await new CompactSign(payload).setProtectedHeader({ alg: 'Ed25519' }).sign(privateKey);
  const { payload: verified } = await compactVerify(jws, publicKey);
  console.log('ok', new TextDecoder().decode(verified) === 'benchmark message');
}

await main();
