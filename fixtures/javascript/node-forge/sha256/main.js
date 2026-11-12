const forge = require('node-forge');

function main() {
  const md = forge.md.sha256.create();
  md.update('benchmark message', 'utf8');
  console.log('ok', md.digest().toHex());
}

main();
