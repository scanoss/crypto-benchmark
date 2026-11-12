import { constants, generateKeyPairSync, sign, verify } from 'node:crypto';

function main() {
  const { privateKey, publicKey } = generateKeyPairSync('rsa', { modulusLength: 3072 });
  const data = Buffer.from('benchmark message');
  const signature = sign('sha256', data, {
    key: privateKey,
    padding: constants.RSA_PKCS1_PSS_PADDING,
    saltLength: 32,
  });
  const valid = verify(
    'sha256',
    data,
    { key: publicKey, padding: constants.RSA_PKCS1_PSS_PADDING, saltLength: 32 },
    signature,
  );
  console.log('ok', valid);
}

main();
