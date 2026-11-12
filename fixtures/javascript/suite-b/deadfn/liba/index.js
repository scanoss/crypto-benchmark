const codec = require('cbench-deadfn-libb/codec');

function sign(data) {
  return codec.sign(data);
}

function checksum(data) {
  return codec.checksum(data);
}

module.exports = { sign, checksum };
