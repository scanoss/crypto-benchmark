const jwt = require('jsonwebtoken');

function main() {
  const secret = process.env.JWT_SECRET || 'benchmark-shared-secret-of-32-bytes';
  const token = jwt.sign({ sub: 'bench' }, secret, { algorithm: 'HS256', expiresIn: '1h' });
  const claims = jwt.verify(token, secret);
  console.log('ok', claims.sub === 'bench');
}

main();
