from passlib.hash import sha512_crypt


def main():
    password = "correct horse battery staple"

    hashed = sha512_crypt.using(rounds=656000).hash(password)
    assert sha512_crypt.verify(password, hashed)
    assert not sha512_crypt.verify("wrong password", hashed)
    print("ok", hashed[:3])


if __name__ == "__main__":
    main()
