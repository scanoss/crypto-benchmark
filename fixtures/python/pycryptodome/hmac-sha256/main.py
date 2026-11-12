from Crypto.Hash import HMAC, SHA256
from Crypto.Random import get_random_bytes


def main():
    key = get_random_bytes(32)
    msg = b"message to authenticate"

    h = HMAC.new(key, digestmod=SHA256)
    h.update(msg)
    mac = h.hexdigest()

    verifier = HMAC.new(key, digestmod=SHA256)
    verifier.update(msg)
    verifier.hexverify(mac)
    print("ok", mac)


if __name__ == "__main__":
    main()
