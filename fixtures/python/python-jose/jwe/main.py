import secrets

from jose import jwe


def main():
    key = secrets.token_bytes(32)
    plaintext = b"secret message"

    token = jwe.encrypt(plaintext, key, algorithm="dir", encryption="A256GCM")
    recovered = jwe.decrypt(token, key)
    assert recovered == plaintext
    print("ok", recovered == plaintext)


if __name__ == "__main__":
    main()
