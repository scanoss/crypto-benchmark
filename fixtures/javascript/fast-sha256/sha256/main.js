const sha256 = require('fast-sha256');

function main() {
  const encoder = new TextEncoder();
  const digest = sha256(encoder.encode('benchmark message'));

  const hasher = new sha256.Hash();
  hasher.update(encoder.encode('benchmark '));
  hasher.update(encoder.encode('message'));
  const streamed = hasher.digest();

  console.log('ok', Buffer.from(digest).equals(Buffer.from(streamed)));
}

main();
