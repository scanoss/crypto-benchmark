from nacl.signing import SigningKey


def main():
    signing_key = SigningKey.generate()
    verify_key = signing_key.verify_key
    signed = signing_key.sign(b"benchmark signed message")
    recovered = verify_key.verify(signed)
    assert recovered == b"benchmark signed message"
    print("ok", len(signed.signature))


if __name__ == "__main__":
    main()
