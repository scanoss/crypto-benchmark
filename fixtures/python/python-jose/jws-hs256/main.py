import secrets

from jose import jws


def main():
    key = secrets.token_bytes(32)
    payload = {"sub": "alice", "admin": False}

    token = jws.sign(payload, key, algorithm="HS256")
    verified = jws.verify(token, key, algorithms=["HS256"])
    assert verified == b'{"sub":"alice","admin":false}'
    print("ok", token.count("."))


if __name__ == "__main__":
    main()
