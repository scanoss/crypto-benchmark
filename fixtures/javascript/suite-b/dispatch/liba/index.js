const libb = require('cbench-dispatch-libb');

class Hasher {
  sum(data) {
    throw new Error('not implemented');
  }
}

class Standard extends Hasher {
  sum(data) {
    return libb.digest(data);
  }
}

class Plain extends Hasher {
  sum(data) {
    return Buffer.from(data);
  }
}

function standard() {
  return new Standard();
}

function plain() {
  return new Plain();
}

module.exports = { Hasher, standard, plain };
