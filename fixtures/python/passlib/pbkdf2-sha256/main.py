from passlib.hash import pbkdf2_sha256


def main():
    password = "correct horse battery staple"

    hashed = pbkdf2_sha256.using(rounds=600000).hash(password)
    assert pbkdf2_sha256.verify(password, hashed)
    assert not pbkdf2_sha256.verify("wrong password", hashed)
    print("ok", hashed.split("$")[2])


if __name__ == "__main__":
    main()
