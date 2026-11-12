from hashlib import sha256

from ecdsa import SECP256k1, SigningKey
from ecdsa.util import sigdecode_der, sigencode_der


def main():
    sk = SigningKey.generate(curve=SECP256k1, hashfunc=sha256)
    vk = sk.verifying_key
    message = b"message to sign"

    signature = sk.sign_deterministic(message, sigencode=sigencode_der)
    assert vk.verify(signature, message, sigdecode=sigdecode_der)
    print("ok", len(signature))


if __name__ == "__main__":
    main()
