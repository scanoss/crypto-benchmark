import rsa


def main():
    public_key, private_key = rsa.newkeys(2048)
    plaintext = b"benchmark payload"
    ciphertext = rsa.encrypt(plaintext, public_key)
    recovered = rsa.decrypt(ciphertext, private_key)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
