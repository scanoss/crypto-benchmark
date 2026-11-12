import rsa


def main():
    _, private_key = rsa.newkeys(2048)
    pem = private_key.save_pkcs1(format="PEM")
    loaded = rsa.PrivateKey.load_pkcs1(pem, format="PEM")
    assert loaded == private_key
    print("ok", loaded.n.bit_length())


if __name__ == "__main__":
    main()
