from argon2 import PasswordHasher


def main():
    ph = PasswordHasher()
    password = "correct horse battery staple"

    hashed = ph.hash(password)
    assert ph.verify(hashed, password)
    assert not ph.check_needs_rehash(hashed)
    print("ok", hashed.split("$")[1])


if __name__ == "__main__":
    main()
