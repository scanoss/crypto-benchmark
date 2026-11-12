import paramiko


def main():
    key = paramiko.RSAKey.generate(bits=3072)
    data = b"benchmark ssh signed data"
    signature = key.sign_ssh_data(data, algorithm="rsa-sha2-256")
    signature.rewind()
    assert key.verify_ssh_sig(data, signature)
    print("ok", signature.asbytes()[4:16].decode())


if __name__ == "__main__":
    main()
