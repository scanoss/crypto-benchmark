import { createCipheriv, createDecipheriv, randomBytes } from 'node:crypto';

function main() {
  const key = randomBytes(32);
  const iv = randomBytes(16);

  const cipher = createCipheriv('aes-256-cbc', key, iv);
  const ciphertext = Buffer.concat([cipher.update('secret message', 'utf8'), cipher.final()]);

  const decipher = createDecipheriv('aes-256-cbc', key, iv);
  const plaintext = Buffer.concat([decipher.update(ciphertext), decipher.final()]).toString('utf8');
  console.log('ok', plaintext === 'secret message');
}

main();
