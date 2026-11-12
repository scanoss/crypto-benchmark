from Crypto.Cipher import ChaCha20_Poly1305
from Crypto.Random import get_random_bytes


def main():
    key = get_random_bytes(32)
    header = b"header"
    data = b"secret message"

    cipher = ChaCha20_Poly1305.new(key=key)
    cipher.update(header)
    ciphertext, tag = cipher.encrypt_and_digest(data)
    nonce = cipher.nonce

    cipher = ChaCha20_Poly1305.new(key=key, nonce=nonce)
    cipher.update(header)
    recovered = cipher.decrypt_and_verify(ciphertext, tag)
    assert recovered == data
    print("ok", recovered == data)


if __name__ == "__main__":
    main()
