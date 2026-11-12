import os

from cryptography.hazmat.primitives import padding
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes


def main():
    key = os.urandom(32)
    iv = os.urandom(16)
    plaintext = b"benchmark payload"
    cipher = Cipher(algorithms.AES(key), modes.CBC(iv))

    padder = padding.PKCS7(128).padder()
    padded = padder.update(plaintext) + padder.finalize()
    encryptor = cipher.encryptor()
    ciphertext = encryptor.update(padded) + encryptor.finalize()

    decryptor = cipher.decryptor()
    padded_out = decryptor.update(ciphertext) + decryptor.finalize()
    unpadder = padding.PKCS7(128).unpadder()
    recovered = unpadder.update(padded_out) + unpadder.finalize()
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
