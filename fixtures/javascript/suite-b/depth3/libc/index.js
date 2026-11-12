const crypto = require('node:crypto');

const KEY = Buffer.from('cbench-suite-b-key');

function mac(data) {
  return crypto.createHmac('sha256', KEY).update(data).digest();
}

module.exports = { mac };
