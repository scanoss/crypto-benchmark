from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import padding, rsa


def main():
    private_key = rsa.generate_private_key(public_exponent=65537, key_size=3072)
    oaep = padding.OAEP(mgf=padding.MGF1(algorithm=hashes.SHA256()), algorithm=hashes.SHA256(), label=None)
    plaintext = b"benchmark payload"
    ciphertext = private_key.public_key().encrypt(plaintext, oaep)
    recovered = private_key.decrypt(ciphertext, oaep)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
