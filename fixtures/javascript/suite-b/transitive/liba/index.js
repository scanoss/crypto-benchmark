const libb = require('cbench-transitive-libb');

function fingerprint(data) {
  return libb.digest(data);
}

module.exports = { fingerprint };
