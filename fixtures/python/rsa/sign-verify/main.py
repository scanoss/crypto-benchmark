import rsa


def main():
    public_key, private_key = rsa.newkeys(2048)
    message = b"benchmark payload"
    signature = rsa.sign(message, private_key, "SHA-256")
    method = rsa.verify(message, signature, public_key)
    assert method == "SHA-256"
    print("ok", method)


if __name__ == "__main__":
    main()
