const nacl = require('tweetnacl');

function main() {
  const digest = nacl.hash(new TextEncoder().encode('benchmark message'));
  console.log('ok', Buffer.from(digest).toString('hex'));
}

main();
