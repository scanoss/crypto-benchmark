import jwt
from cryptography.hazmat.primitives.asymmetric import ec


def main():
    private_key = ec.generate_private_key(ec.SECP256R1())
    payload = {"sub": "benchmark", "admin": False}
    token = jwt.encode(payload, private_key, algorithm="ES256")
    decoded = jwt.decode(token, private_key.public_key(), algorithms=["ES256"])
    assert decoded == payload
    print("ok", decoded == payload)


if __name__ == "__main__":
    main()
