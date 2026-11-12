from Crypto.PublicKey import RSA


def main():
    key = RSA.generate(3072)
    passphrase = "correct horse battery staple"

    encrypted = key.export_key(
        passphrase=passphrase,
        pkcs=8,
        protection="scryptAndAES256-GCM",
        prot_params={"iteration_count": 2**14},
    )
    public_pem = key.public_key().export_key()

    loaded = RSA.import_key(encrypted, passphrase=passphrase)
    loaded_public = RSA.import_key(public_pem)
    assert loaded == key and loaded_public == key.public_key()
    print("ok", loaded.size_in_bits())


if __name__ == "__main__":
    main()
