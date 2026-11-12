import nacl.utils


def main():
    token = nacl.utils.random(32)
    other = nacl.utils.random(32)
    assert len(token) == 32 and token != other
    print("ok", token.hex()[:16])


if __name__ == "__main__":
    main()
