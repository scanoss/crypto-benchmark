const forge = require('node-forge');

function main() {
  const keys = forge.pki.rsa.generateKeyPair({ bits: 2048, e: 0x10001 });
  const md = forge.md.sha256.create();
  md.update('benchmark message', 'utf8');
  const pss = forge.pss.create({
    md: forge.md.sha256.create(),
    mgf: forge.mgf.mgf1.create(forge.md.sha256.create()),
    saltLength: 32,
  });
  const signature = keys.privateKey.sign(md, pss);
  const valid = keys.publicKey.verify(md.digest().bytes(), signature, pss);
  console.log('ok', valid);
}

main();
