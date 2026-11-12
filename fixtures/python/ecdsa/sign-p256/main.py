from hashlib import sha256

from ecdsa import NIST256p, SigningKey


def main():
    sk = SigningKey.generate(curve=NIST256p, hashfunc=sha256)
    vk = sk.verifying_key
    message = b"message to sign"

    signature = sk.sign(message)
    assert vk.verify(signature, message)
    print("ok", len(signature))


if __name__ == "__main__":
    main()
