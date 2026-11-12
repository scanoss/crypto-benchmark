import secrets

import jwt


def main():
    key = secrets.token_bytes(32)
    payload = {"sub": "benchmark", "admin": False}
    token = jwt.encode(payload, key, algorithm="HS256")
    decoded = jwt.decode(token, key, algorithms=["HS256"])
    assert decoded == payload
    print("ok", decoded == payload)


if __name__ == "__main__":
    main()
