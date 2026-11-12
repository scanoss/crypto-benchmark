const liba = require('cbench-wrapper-liba');

function run(data) {
  return liba.protect(data);
}

function main() {
  console.log('hello');
}

main();
