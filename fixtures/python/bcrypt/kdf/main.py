import os

import bcrypt


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)
    key = bcrypt.kdf(password=password, salt=salt, desired_key_bytes=32, rounds=100)
    again = bcrypt.kdf(password=password, salt=salt, desired_key_bytes=32, rounds=100)
    assert key == again and len(key) == 32
    print("ok", len(key))


if __name__ == "__main__":
    main()
