const liba = require('cbench-direct-liba');

function run(data) {
  return liba.fingerprint(data);
}

function main() {
  console.log('hello');
}

main();
