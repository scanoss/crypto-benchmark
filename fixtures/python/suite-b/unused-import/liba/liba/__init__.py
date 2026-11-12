import libb

VERSION = "1.0.0"


def fingerprint(data):
    return libb.digest(data)
