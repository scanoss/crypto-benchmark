import { EncryptJWT, generateKeyPair, jwtDecrypt } from 'jose';

async function main() {
  const { publicKey, privateKey } = await generateKeyPair('RSA-OAEP-256', { modulusLength: 2048 });
  const jwe = await new EncryptJWT({ sub: 'bench' })
    .setProtectedHeader({ alg: 'RSA-OAEP-256', enc: 'A256GCM' })
    .setIssuedAt()
    .setExpirationTime('2h')
    .encrypt(publicKey);

  const { payload } = await jwtDecrypt(jwe, privateKey, {
    keyManagementAlgorithms: ['RSA-OAEP-256'],
    contentEncryptionAlgorithms: ['A256GCM'],
  });
  console.log('ok', payload.sub === 'bench');
}

await main();
