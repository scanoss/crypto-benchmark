const libb = require('cbench-wrapper-libb');

function protect(data) {
  return protectInner(data);
}

function protectInner(data) {
  return libb.encode(data);
}

module.exports = { protect };
