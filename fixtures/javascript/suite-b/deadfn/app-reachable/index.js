const liba = require('cbench-deadfn-liba');

function main() {
  console.log(liba.sign(Buffer.from('hello')).toString('hex'));
}

main();
