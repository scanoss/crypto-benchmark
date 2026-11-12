from cryptography.hazmat.primitives.asymmetric.ed25519 import Ed25519PrivateKey


def main():
    private_key = Ed25519PrivateKey.generate()
    message = b"benchmark payload"
    signature = private_key.sign(message)
    private_key.public_key().verify(signature, message)
    print("ok", len(signature))


if __name__ == "__main__":
    main()
