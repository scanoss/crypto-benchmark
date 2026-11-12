const libb = require('cbench-depth3-libb');

function seal(data) {
  return libb.seal(data);
}

module.exports = { seal };
