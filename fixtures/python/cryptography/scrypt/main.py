import os

from cryptography.hazmat.primitives.kdf.scrypt import Scrypt


def kdf(salt):
    return Scrypt(salt=salt, length=32, n=2**14, r=8, p=1)


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)
    key = kdf(salt).derive(password)
    kdf(salt).verify(password, key)
    print("ok", len(key))


if __name__ == "__main__":
    main()
