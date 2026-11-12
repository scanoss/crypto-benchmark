import pyotp


def main():
    secret = pyotp.random_base32()
    hotp = pyotp.HOTP(secret)
    code = hotp.at(0)
    assert hotp.verify(code, 0)
    print("ok", code)


if __name__ == "__main__":
    main()
