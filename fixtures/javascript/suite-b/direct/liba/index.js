const crypto = require('node:crypto');

function fingerprint(data) {
  return crypto.createHash('sha256').update(data).digest();
}

module.exports = { fingerprint };
