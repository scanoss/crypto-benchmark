const { createHash, hash } = require('blake3-wasm');

function main() {
  const digest = hash('benchmark message');

  const hasher = createHash();
  hasher.update('benchmark ');
  hasher.update('message');
  const streamed = hasher.digest();

  console.log('ok', digest.equals(streamed));
}

main();
