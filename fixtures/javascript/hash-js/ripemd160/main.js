const hash = require('hash.js');

function main() {
  const digest = hash.ripemd160().update('benchmark message').digest('hex');
  console.log('ok', digest);
}

main();
