import bcrypt


def main():
    password = b"correct horse battery staple"
    hashed = bcrypt.hashpw(password, bcrypt.gensalt(rounds=12))
    assert bcrypt.checkpw(password, hashed)
    assert not bcrypt.checkpw(b"wrong password", hashed)
    print("ok", hashed[:7].decode())


if __name__ == "__main__":
    main()
