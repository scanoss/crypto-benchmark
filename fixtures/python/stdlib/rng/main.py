import secrets


def main():
    key = secrets.token_bytes(32)
    token = secrets.token_urlsafe(32)
    pick = secrets.randbelow(100)
    assert len(key) == 32 and 0 <= pick < 100
    print("ok", len(key), len(token))


if __name__ == "__main__":
    main()
