import { SignJWT, generateKeyPair, jwtVerify } from 'jose';

async function main() {
  const { publicKey, privateKey } = await generateKeyPair('ES256');
  const jwt = await new SignJWT({ sub: 'bench' }).setProtectedHeader({ alg: 'ES256' }).sign(privateKey);
  const { payload, protectedHeader } = await jwtVerify(jwt, publicKey);
  console.log('ok', payload.sub === 'bench', protectedHeader.alg);
}

await main();
