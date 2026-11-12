import { createCipheriv, createDecipheriv, randomBytes } from 'node:crypto';

function main() {
  const key = randomBytes(32);
  const iv = randomBytes(12);
  const aad = Buffer.from('header');

  const cipher = createCipheriv('aes-256-gcm', key, iv);
  cipher.setAAD(aad);
  const ciphertext = Buffer.concat([cipher.update('secret message', 'utf8'), cipher.final()]);
  const tag = cipher.getAuthTag();

  const decipher = createDecipheriv('aes-256-gcm', key, iv);
  decipher.setAAD(aad);
  decipher.setAuthTag(tag);
  const plaintext = Buffer.concat([decipher.update(ciphertext), decipher.final()]).toString('utf8');
  console.log('ok', plaintext === 'secret message');
}

main();
