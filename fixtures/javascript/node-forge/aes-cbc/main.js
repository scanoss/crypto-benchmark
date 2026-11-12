const forge = require('node-forge');

function main() {
  const key = forge.random.getBytesSync(32);
  const iv = forge.random.getBytesSync(16);

  const cipher = forge.cipher.createCipher('AES-CBC', key);
  cipher.start({ iv });
  cipher.update(forge.util.createBuffer('secret message', 'utf8'));
  cipher.finish();
  const ciphertext = cipher.output.getBytes();

  const decipher = forge.cipher.createDecipher('AES-CBC', key);
  decipher.start({ iv });
  decipher.update(forge.util.createBuffer(ciphertext));
  decipher.finish();
  console.log('ok', decipher.output.toString() === 'secret message');
}

main();
