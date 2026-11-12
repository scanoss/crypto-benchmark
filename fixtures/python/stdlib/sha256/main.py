import hashlib


def main():
    data = b"benchmark payload"
    h = hashlib.sha256()
    h.update(data)
    digest = h.hexdigest()
    assert digest == hashlib.sha256(data).hexdigest()
    print("ok", digest)


if __name__ == "__main__":
    main()
