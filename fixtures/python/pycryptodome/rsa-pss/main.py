from Crypto.Hash import SHA256
from Crypto.PublicKey import RSA
from Crypto.Signature import pss


def main():
    key = RSA.generate(3072)
    message = b"message to sign"

    h = SHA256.new(message)
    signature = pss.new(key).sign(h)

    verifier = pss.new(key.public_key())
    verifier.verify(SHA256.new(message), signature)
    print("ok", len(signature))


if __name__ == "__main__":
    main()
