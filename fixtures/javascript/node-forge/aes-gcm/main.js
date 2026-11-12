const forge = require('node-forge');

function main() {
  const key = forge.random.getBytesSync(32);
  const iv = forge.random.getBytesSync(12);
  const additionalData = 'header';

  const cipher = forge.cipher.createCipher('AES-GCM', key);
  cipher.start({ iv, additionalData, tagLength: 128 });
  cipher.update(forge.util.createBuffer('secret message', 'utf8'));
  cipher.finish();
  const ciphertext = cipher.output.getBytes();
  const tag = cipher.mode.tag;

  const decipher = forge.cipher.createDecipher('AES-GCM', key);
  decipher.start({ iv, additionalData, tagLength: 128, tag });
  decipher.update(forge.util.createBuffer(ciphertext));
  const authentic = decipher.finish();
  console.log('ok', authentic && decipher.output.toString() === 'secret message');
}

main();
