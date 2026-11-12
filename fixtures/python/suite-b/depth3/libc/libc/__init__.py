import hashlib
import hmac

KEY = b"cbench-suite-b-key"


def mac(data):
    return hmac.new(KEY, data, hashlib.sha256).digest()
