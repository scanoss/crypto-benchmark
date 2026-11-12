import { SignJWT, generateKeyPair, jwtVerify } from 'jose';

async function main() {
  const { publicKey, privateKey } = await generateKeyPair('RS256', { modulusLength: 2048 });
  const jwt = await new SignJWT({ sub: 'bench' })
    .setProtectedHeader({ alg: 'RS256' })
    .setIssuedAt()
    .setExpirationTime('2h')
    .sign(privateKey);

  const { payload } = await jwtVerify(jwt, publicKey, { algorithms: ['RS256'] });
  console.log('ok', payload.sub === 'bench');
}

await main();
