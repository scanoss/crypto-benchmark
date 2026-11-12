import nacl.secret
import nacl.utils


def main():
    key = nacl.utils.random(nacl.secret.SecretBox.KEY_SIZE)
    box = nacl.secret.SecretBox(key)
    plaintext = b"benchmark secretbox message"
    encrypted = box.encrypt(plaintext)
    recovered = box.decrypt(encrypted)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
