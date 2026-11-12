const crypto = require('node:crypto');

const KEY = Buffer.from('cbench-suite-b-key');

function sign(data) {
  return crypto.createHmac('sha256', KEY).update(data).digest();
}

function checksum(data) {
  let c = 0;
  for (const b of data) {
    c ^= b;
  }
  return c;
}

function legacyDigest(data) {
  return crypto.createHash('sha1').update(data).digest();
}

module.exports = { sign, checksum, legacyDigest };
