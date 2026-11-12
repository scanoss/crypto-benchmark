import os

from cryptography.hazmat.primitives.ciphers.aead import ChaCha20Poly1305


def main():
    key = ChaCha20Poly1305.generate_key()
    chacha = ChaCha20Poly1305(key)
    nonce = os.urandom(12)
    plaintext = b"benchmark payload"
    aad = b"header"
    ciphertext = chacha.encrypt(nonce, plaintext, aad)
    recovered = chacha.decrypt(nonce, ciphertext, aad)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
