const liba = require('cbench-depth3-liba');

function run(data) {
  return liba.seal(data);
}

function main() {
  console.log(run(Buffer.from('hello')).toString('hex'));
}

main();
