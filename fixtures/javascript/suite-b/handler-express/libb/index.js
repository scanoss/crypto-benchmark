const crypto = require('node:crypto');

function digest(data) {
  return crypto.createHash('sha256').update(data).digest();
}

module.exports = { digest };
