import hashlib
import hmac

KEY = b"cbench-suite-b-key"


def sign(data):
    return hmac.new(KEY, data, hashlib.sha256).digest()


def checksum(data):
    c = 0
    for b in data:
        c ^= b
    return c


def legacy_digest(data):
    return hashlib.sha1(data).digest()
