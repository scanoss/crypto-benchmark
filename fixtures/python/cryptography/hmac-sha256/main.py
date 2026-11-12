import os

from cryptography.hazmat.primitives import hashes, hmac


def main():
    key = os.urandom(32)
    message = b"benchmark payload"
    h = hmac.HMAC(key, hashes.SHA256())
    h.update(message)
    tag = h.finalize()

    verifier = hmac.HMAC(key, hashes.SHA256())
    verifier.update(message)
    verifier.verify(tag)
    print("ok", len(tag))


if __name__ == "__main__":
    main()
