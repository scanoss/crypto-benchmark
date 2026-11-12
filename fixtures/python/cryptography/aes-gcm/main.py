import os

from cryptography.hazmat.primitives.ciphers.aead import AESGCM


def main():
    key = AESGCM.generate_key(bit_length=256)
    aesgcm = AESGCM(key)
    nonce = os.urandom(12)
    plaintext = b"benchmark payload"
    aad = b"header"
    ciphertext = aesgcm.encrypt(nonce, plaintext, aad)
    recovered = aesgcm.decrypt(nonce, ciphertext, aad)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
