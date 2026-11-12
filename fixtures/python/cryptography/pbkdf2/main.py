import os

from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC


def kdf(salt):
    return PBKDF2HMAC(algorithm=hashes.SHA256(), length=32, salt=salt, iterations=600000)


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)
    key = kdf(salt).derive(password)
    kdf(salt).verify(password, key)
    print("ok", len(key))


if __name__ == "__main__":
    main()
