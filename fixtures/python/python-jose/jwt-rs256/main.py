from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from jose import jwt


def main():
    private_key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
    private_pem = private_key.private_bytes(
        serialization.Encoding.PEM,
        serialization.PrivateFormat.PKCS8,
        serialization.NoEncryption(),
    )
    public_pem = private_key.public_key().public_bytes(
        serialization.Encoding.PEM, serialization.PublicFormat.SubjectPublicKeyInfo
    )
    claims = {"sub": "alice", "iss": "benchmark"}

    token = jwt.encode(claims, private_pem, algorithm="RS256")
    decoded = jwt.decode(token, public_pem, algorithms=["RS256"], issuer="benchmark")
    assert decoded == claims
    print("ok", decoded["sub"])


if __name__ == "__main__":
    main()
