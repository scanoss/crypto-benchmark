import { SignJWT, jwtVerify } from 'jose';

async function main() {
  const secret = new TextEncoder().encode(process.env.JWT_SECRET ?? 'cc7e0d44fd473002f1c42167459001140ec6389b7353f8088f4d9a95f2f596f2');
  const jwt = await new SignJWT({ 'urn:example:claim': true })
    .setProtectedHeader({ alg: 'HS256' })
    .setIssuedAt()
    .setIssuer('urn:example:issuer')
    .setAudience('urn:example:audience')
    .setExpirationTime('2h')
    .sign(secret);

  const { payload } = await jwtVerify(jwt, secret, {
    issuer: 'urn:example:issuer',
    audience: 'urn:example:audience',
  });
  console.log('ok', payload['urn:example:claim']);
}

await main();
