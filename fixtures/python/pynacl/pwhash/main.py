import nacl.pwhash


def main():
    password = b"correct horse battery staple"
    hashed = nacl.pwhash.argon2id.str(password)
    assert nacl.pwhash.argon2id.verify(hashed, password)
    print("ok", hashed.split(b"$")[1].decode())


if __name__ == "__main__":
    main()
