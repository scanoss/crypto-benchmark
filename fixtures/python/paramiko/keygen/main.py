import paramiko


def main():
    key = paramiko.RSAKey.generate(bits=3072)
    public = paramiko.RSAKey(data=key.asbytes())
    assert key.get_bits() == 3072 and public == key
    print("ok", key.fingerprint)


if __name__ == "__main__":
    main()
