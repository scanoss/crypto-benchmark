import hashlib

import pyotp


def main():
    secret = pyotp.random_base32(length=64)
    totp = pyotp.TOTP(secret, digits=8, digest=hashlib.sha256)
    code = totp.now()
    assert totp.verify(code, valid_window=1)
    print("ok", len(code))


if __name__ == "__main__":
    main()
