from Crypto.Hash import SHA256


def main():
    h = SHA256.new()
    h.update(b"hello ")
    h.update(b"world")
    digest = h.hexdigest()
    assert digest == SHA256.new(b"hello world").hexdigest()
    print("ok", digest)


if __name__ == "__main__":
    main()
