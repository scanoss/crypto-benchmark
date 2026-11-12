const crypto = require('node:crypto');
const jwt = require('jsonwebtoken');

function main() {
  const { privateKey, publicKey } = crypto.generateKeyPairSync('ec', { namedCurve: 'P-256' });
  const token = jwt.sign({ sub: 'bench' }, privateKey, { algorithm: 'ES256', expiresIn: '1h' });
  jwt.verify(token, publicKey, { algorithms: ['ES256'] }, (err, claims) => {
    if (err) throw err;
    console.log('ok', claims.sub === 'bench');
  });
}

main();
