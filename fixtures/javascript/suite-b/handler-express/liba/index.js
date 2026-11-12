const libb = require('cbench-handler-express-libb');

function fingerprint(data) {
  return libb.digest(data);
}

module.exports = { fingerprint };
