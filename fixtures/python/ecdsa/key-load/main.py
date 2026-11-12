from hashlib import sha256

from ecdsa import NIST256p, SigningKey, VerifyingKey


def main():
    sk = SigningKey.generate(curve=NIST256p, hashfunc=sha256)
    private_pem = sk.to_pem(format="pkcs8")
    public_pem = sk.verifying_key.to_pem()

    loaded_sk = SigningKey.from_pem(private_pem, hashfunc=sha256)
    loaded_vk = VerifyingKey.from_pem(public_pem, hashfunc=sha256)
    message = b"message to sign"
    assert loaded_vk.verify(loaded_sk.sign(message), message)
    print("ok", loaded_sk.to_string() == sk.to_string())


if __name__ == "__main__":
    main()
