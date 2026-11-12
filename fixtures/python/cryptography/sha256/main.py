from cryptography.hazmat.primitives import hashes


def main():
    digest = hashes.Hash(hashes.SHA256())
    digest.update(b"benchmark ")
    digest.update(b"payload")
    value = digest.finalize()
    assert len(value) == 32
    print("ok", value.hex())


if __name__ == "__main__":
    main()
