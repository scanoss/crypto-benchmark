import jwt
from cryptography.hazmat.primitives.asymmetric import rsa


def main():
    private_key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
    payload = {"sub": "benchmark", "admin": False}
    token = jwt.encode(payload, private_key, algorithm="RS256")
    decoded = jwt.decode(token, private_key.public_key(), algorithms=["RS256"])
    assert decoded == payload
    print("ok", decoded == payload)


if __name__ == "__main__":
    main()
