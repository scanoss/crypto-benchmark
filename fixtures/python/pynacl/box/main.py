from nacl.public import Box, PrivateKey


def main():
    alice = PrivateKey.generate()
    bob = PrivateKey.generate()
    plaintext = b"benchmark box message"
    encrypted = Box(alice, bob.public_key).encrypt(plaintext)
    recovered = Box(bob, alice.public_key).decrypt(encrypted)
    assert recovered == plaintext
    print("ok", plaintext == recovered)


if __name__ == "__main__":
    main()
