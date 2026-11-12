from Crypto.Hash import SHA256
from Crypto.Protocol.KDF import PBKDF2
from Crypto.Random import get_random_bytes


def main():
    password = b"correct horse battery staple"
    salt = get_random_bytes(16)

    key = PBKDF2(password, salt, 32, count=600000, hmac_hash_module=SHA256)
    again = PBKDF2(password, salt, 32, count=600000, hmac_hash_module=SHA256)
    assert key == again
    print("ok", key.hex())


if __name__ == "__main__":
    main()
