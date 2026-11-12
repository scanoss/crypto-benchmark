from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import padding, rsa


def main():
    private_key = rsa.generate_private_key(public_exponent=65537, key_size=3072)
    message = b"benchmark payload"
    pss = padding.PSS(mgf=padding.MGF1(hashes.SHA256()), salt_length=padding.PSS.MAX_LENGTH)
    signature = private_key.sign(message, pss, hashes.SHA256())
    private_key.public_key().verify(signature, message, pss, hashes.SHA256())
    print("ok", len(signature))


if __name__ == "__main__":
    main()
