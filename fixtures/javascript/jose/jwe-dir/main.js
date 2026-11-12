import { CompactEncrypt, compactDecrypt, generateSecret } from 'jose';

async function main() {
  const secret = await generateSecret('A256GCM');
  const plaintext = new TextEncoder().encode('secret message');
  const jwe = await new CompactEncrypt(plaintext).setProtectedHeader({ alg: 'dir', enc: 'A256GCM' }).encrypt(secret);
  const { plaintext: decrypted } = await compactDecrypt(jwe, secret);
  console.log('ok', new TextDecoder().decode(decrypted) === 'secret message');
}

await main();
