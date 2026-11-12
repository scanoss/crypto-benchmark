const liba = require('cbench-direct-liba');

function run(data) {
  return liba.fingerprint(data);
}

function main() {
  console.log(run(Buffer.from('hello')).toString('hex'));
}

main();
