from Crypto.Protocol.KDF import scrypt
from Crypto.Random import get_random_bytes


def main():
    password = b"correct horse battery staple"
    salt = get_random_bytes(16)

    key = scrypt(password, salt, key_len=32, N=2**14, r=8, p=1)
    again = scrypt(password, salt, key_len=32, N=2**14, r=8, p=1)
    assert key == again
    print("ok", key.hex())


if __name__ == "__main__":
    main()
