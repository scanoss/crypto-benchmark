const forge = require('node-forge');

function main() {
  const salt = forge.random.getBytesSync(16);
  const derivedKey = forge.pkcs5.pbkdf2('correct horse battery staple', salt, 600000, 32, 'sha256');
  console.log('ok', forge.util.bytesToHex(derivedKey).length);
}

main();
