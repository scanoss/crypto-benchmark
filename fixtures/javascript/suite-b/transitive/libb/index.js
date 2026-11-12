const crypto = require('node:crypto');

function digest(data) {
  return crypto.createHash('sha512').update(data).digest();
}

module.exports = { digest };
