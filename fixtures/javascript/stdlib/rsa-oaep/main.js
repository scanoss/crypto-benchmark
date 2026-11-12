import { constants, generateKeyPairSync, privateDecrypt, publicEncrypt } from 'node:crypto';

function main() {
  const { privateKey, publicKey } = generateKeyPairSync('rsa', { modulusLength: 2048 });
  const encrypted = publicEncrypt(
    { key: publicKey, padding: constants.RSA_PKCS1_OAEP_PADDING, oaepHash: 'sha256' },
    Buffer.from('secret message'),
  );
  const decrypted = privateDecrypt(
    { key: privateKey, padding: constants.RSA_PKCS1_OAEP_PADDING, oaepHash: 'sha256' },
    encrypted,
  );
  console.log('ok', decrypted.toString() === 'secret message');
}

main();
