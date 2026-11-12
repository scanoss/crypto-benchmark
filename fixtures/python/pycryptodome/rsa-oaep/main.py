from Crypto.Cipher import PKCS1_OAEP
from Crypto.Hash import SHA256
from Crypto.PublicKey import RSA


def main():
    key = RSA.generate(3072)
    data = b"secret message"

    cipher = PKCS1_OAEP.new(key.public_key(), hashAlgo=SHA256)
    ciphertext = cipher.encrypt(data)

    cipher = PKCS1_OAEP.new(key, hashAlgo=SHA256)
    recovered = cipher.decrypt(ciphertext)
    assert recovered == data
    print("ok", recovered == data)


if __name__ == "__main__":
    main()
