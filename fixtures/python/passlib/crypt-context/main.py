from passlib.context import CryptContext


def main():
    context = CryptContext(schemes=["bcrypt"], bcrypt__rounds=12)
    password = "correct horse battery staple"

    hashed = context.hash(password)
    assert context.verify(password, hashed)
    assert not context.verify("wrong password", hashed)
    assert not context.needs_update(hashed)
    print("ok", hashed[:7])


if __name__ == "__main__":
    main()
