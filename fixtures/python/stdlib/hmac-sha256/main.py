import hashlib
import hmac
import secrets


def main():
    key = secrets.token_bytes(32)
    message = b"benchmark payload"
    tag = hmac.new(key, message, hashlib.sha256).digest()
    expected = hmac.new(key, message, hashlib.sha256).digest()
    assert hmac.compare_digest(tag, expected)
    print("ok", tag.hex())


if __name__ == "__main__":
    main()
