from passlib.hash import argon2


def main():
    password = "correct horse battery staple"

    hashed = argon2.using(type="ID").hash(password)
    assert argon2.verify(password, hashed)
    assert not argon2.verify("wrong password", hashed)
    print("ok", hashed.split("$")[1])


if __name__ == "__main__":
    main()
