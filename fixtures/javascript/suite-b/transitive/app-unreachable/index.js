const liba = require('cbench-transitive-liba');

function run(data) {
  return liba.fingerprint(data);
}

function main() {
  console.log('hello');
}

main();
