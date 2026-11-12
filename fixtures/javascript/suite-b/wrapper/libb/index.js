const crypto = require('node:crypto');

function encode(data) {
  return encodeInner(data);
}

function encodeInner(data) {
  return hashBlock(data);
}

function hashBlock(data) {
  return crypto.createHash('sha384').update(data).digest();
}

module.exports = { encode };
