import hashlib
import hmac
import os


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)
    key = hashlib.scrypt(password, salt=salt, n=2**14, r=8, p=1, dklen=32)
    again = hashlib.scrypt(password, salt=salt, n=2**14, r=8, p=1, dklen=32)
    assert hmac.compare_digest(key, again)
    print("ok", len(key))


if __name__ == "__main__":
    main()
