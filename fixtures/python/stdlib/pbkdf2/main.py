import hashlib
import hmac
import os


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)
    key = hashlib.pbkdf2_hmac("sha256", password, salt, 600000)
    again = hashlib.pbkdf2_hmac("sha256", password, salt, 600000)
    assert hmac.compare_digest(key, again)
    print("ok", len(key))


if __name__ == "__main__":
    main()
