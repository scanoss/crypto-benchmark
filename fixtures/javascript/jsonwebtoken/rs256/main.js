const crypto = require('node:crypto');
const jwt = require('jsonwebtoken');

function main() {
  const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', { modulusLength: 2048 });
  const token = jwt.sign({ sub: 'bench' }, privateKey, {
    algorithm: 'RS256',
    issuer: 'crypto-benchmark',
    expiresIn: '1h',
  });
  const claims = jwt.verify(token, publicKey, { algorithms: ['RS256'], issuer: 'crypto-benchmark' });
  console.log('ok', claims.sub === 'bench');
}

main();
