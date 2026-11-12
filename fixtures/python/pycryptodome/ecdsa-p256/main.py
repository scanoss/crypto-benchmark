from Crypto.Hash import SHA256
from Crypto.PublicKey import ECC
from Crypto.Signature import DSS


def main():
    key = ECC.generate(curve="P-256")
    message = b"message to sign"

    h = SHA256.new(message)
    signature = DSS.new(key, "fips-186-3").sign(h)

    verifier = DSS.new(key.public_key(), "fips-186-3")
    verifier.verify(SHA256.new(message), signature)
    print("ok", len(signature))


if __name__ == "__main__":
    main()
