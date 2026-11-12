const hash = require('hash.js');

function main() {
  const digest = hash.sha512().update('benchmark message').digest('hex');
  console.log('ok', digest);
}

main();
