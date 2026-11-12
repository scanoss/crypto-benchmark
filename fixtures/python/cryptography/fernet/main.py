from cryptography.fernet import Fernet


def main():
    key = Fernet.generate_key()
    f = Fernet(key)
    plaintext = b"benchmark payload"
    token = f.encrypt(plaintext)
    recovered = f.decrypt(token)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
