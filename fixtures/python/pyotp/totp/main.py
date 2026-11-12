import pyotp


def main():
    secret = pyotp.random_base32()
    totp = pyotp.TOTP(secret)
    code = totp.now()
    assert totp.verify(code)
    print("ok", totp.provisioning_uri(name="alice@example.com", issuer_name="Benchmark"))


if __name__ == "__main__":
    main()
