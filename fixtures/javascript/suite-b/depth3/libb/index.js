const libc = require('cbench-depth3-libc');

function seal(data) {
  return libc.mac(data);
}

module.exports = { seal };
