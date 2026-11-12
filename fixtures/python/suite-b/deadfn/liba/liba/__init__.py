from libb import codec


def sign(data):
    return codec.sign(data)


def checksum(data):
    return codec.checksum(data)
