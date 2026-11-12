import nacl.encoding
import nacl.hash
import nacl.utils


def main():
    message = b"benchmark hashed message"
    key = nacl.utils.random(32)
    mac = nacl.hash.blake2b(message, key=key, encoder=nacl.encoding.HexEncoder)
    again = nacl.hash.blake2b(message, key=key, encoder=nacl.encoding.HexEncoder)
    digest = nacl.hash.sha256(message, encoder=nacl.encoding.HexEncoder)
    assert mac == again and len(digest) == 64
    print("ok", mac[:16].decode())


if __name__ == "__main__":
    main()
