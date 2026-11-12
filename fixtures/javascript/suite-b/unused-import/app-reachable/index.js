const liba = require('cbench-unused-import-liba');

function main() {
  console.log(liba.VERSION, liba.fingerprint(Buffer.from('hello')).toString('hex'));
}

main();
