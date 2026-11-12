const liba = require('cbench-deadfn-liba');

function main() {
  console.log(liba.checksum(Buffer.from('hello')));
}

main();
