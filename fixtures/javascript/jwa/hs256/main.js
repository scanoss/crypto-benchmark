const jwa = require('jwa');

function main() {
  const hmac = jwa('HS256');
  const input = 'eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJiZW5jaCJ9';
  const secret = process.env.JWT_SECRET || 'benchmark-shared-secret-of-32-bytes';
  const signature = hmac.sign(input, secret);
  console.log('ok', hmac.verify(input, signature, secret));
}

main();
