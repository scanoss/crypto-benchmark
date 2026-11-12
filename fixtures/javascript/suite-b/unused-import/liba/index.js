const libb = require('cbench-unused-import-libb');

const VERSION = '1.0.0';

function fingerprint(data) {
  return libb.digest(data);
}

module.exports = { VERSION, fingerprint };
