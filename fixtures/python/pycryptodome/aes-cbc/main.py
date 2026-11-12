from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes
from Crypto.Util.Padding import pad, unpad


def main():
    key = get_random_bytes(32)
    data = b"secret message"

    cipher = AES.new(key, AES.MODE_CBC)
    ciphertext = cipher.encrypt(pad(data, AES.block_size))
    iv = cipher.iv

    cipher = AES.new(key, AES.MODE_CBC, iv=iv)
    recovered = unpad(cipher.decrypt(ciphertext), AES.block_size)
    assert recovered == data
    print("ok", recovered == data)


if __name__ == "__main__":
    main()
