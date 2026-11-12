from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric.x25519 import X25519PrivateKey
from cryptography.hazmat.primitives.kdf.hkdf import HKDF


def derive(shared_key):
    return HKDF(algorithm=hashes.SHA256(), length=32, salt=None, info=b"handshake data").derive(shared_key)


def main():
    alice = X25519PrivateKey.generate()
    bob = X25519PrivateKey.generate()
    alice_key = derive(alice.exchange(bob.public_key()))
    bob_key = derive(bob.exchange(alice.public_key()))
    assert alice_key == bob_key
    print("ok", alice_key == bob_key)


if __name__ == "__main__":
    main()
