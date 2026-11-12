from ecdsa import ECDH, NIST256p


def main():
    alice = ECDH(curve=NIST256p)
    bob = ECDH(curve=NIST256p)
    alice_public = alice.generate_private_key()
    bob_public = bob.generate_private_key()

    alice.load_received_public_key(bob_public)
    bob.load_received_public_key(alice_public)
    alice_secret = alice.generate_sharedsecret_bytes()
    bob_secret = bob.generate_sharedsecret_bytes()
    assert alice_secret == bob_secret
    print("ok", alice_secret == bob_secret)


if __name__ == "__main__":
    main()
