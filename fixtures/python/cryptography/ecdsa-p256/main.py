from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import ec


def main():
    private_key = ec.generate_private_key(ec.SECP256R1())
    message = b"benchmark payload"
    signature = private_key.sign(message, ec.ECDSA(hashes.SHA256()))
    private_key.public_key().verify(signature, message, ec.ECDSA(hashes.SHA256()))
    print("ok", len(signature) > 0)


if __name__ == "__main__":
    main()
