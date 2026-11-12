from Crypto.Random import get_random_bytes


def main():
    token = get_random_bytes(32)
    other = get_random_bytes(32)
    assert len(token) == 32 and token != other
    print("ok", token.hex())


if __name__ == "__main__":
    main()
