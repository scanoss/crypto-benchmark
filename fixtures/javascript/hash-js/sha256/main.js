const hash = require('hash.js');

function main() {
  const digest = hash.sha256().update('benchmark message').digest('hex');

  const hasher = hash.sha256();
  hasher.update('benchmark ');
  hasher.update('message');
  const streamed = hasher.digest('hex');

  console.log('ok', digest === streamed);
}

main();
