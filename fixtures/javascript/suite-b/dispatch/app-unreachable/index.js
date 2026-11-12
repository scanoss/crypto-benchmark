const liba = require('cbench-dispatch-liba');

function main() {
  const h = liba.plain();
  console.log(h.sum(Buffer.from('hello')).toString('hex'));
}

main();
