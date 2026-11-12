from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import ec


def main():
    password = b"correct horse battery staple"
    key = ec.generate_private_key(ec.SECP256R1())
    pem = key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.BestAvailableEncryption(password),
    )
    loaded = serialization.load_pem_private_key(pem, password=password)
    assert loaded.private_numbers() == key.private_numbers()
    print("ok", loaded.curve.name)


if __name__ == "__main__":
    main()
