import os
import tempfile

import paramiko


def main():
    key = paramiko.RSAKey.generate(bits=3072)
    with tempfile.TemporaryDirectory() as tmp:
        path = os.path.join(tmp, "id_rsa")
        key.write_private_key_file(path, password="s3cret-passphrase")
        loaded = paramiko.RSAKey.from_private_key_file(path, password="s3cret-passphrase")
    assert loaded == key
    print("ok", loaded.get_bits())


if __name__ == "__main__":
    main()
