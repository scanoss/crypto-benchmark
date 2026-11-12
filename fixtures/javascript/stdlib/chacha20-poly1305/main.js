import { createCipheriv, createDecipheriv, randomBytes } from 'node:crypto';

function main() {
  const key = randomBytes(32);
  const nonce = randomBytes(12);

  const cipher = createCipheriv('chacha20-poly1305', key, nonce, { authTagLength: 16 });
  const ciphertext = Buffer.concat([cipher.update('secret message', 'utf8'), cipher.final()]);
  const tag = cipher.getAuthTag();

  const decipher = createDecipheriv('chacha20-poly1305', key, nonce, { authTagLength: 16 });
  decipher.setAuthTag(tag);
  const plaintext = Buffer.concat([decipher.update(ciphertext), decipher.final()]).toString('utf8');
  console.log('ok', plaintext === 'secret message');
}

main();
