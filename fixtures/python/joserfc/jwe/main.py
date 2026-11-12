from joserfc import jwe
from joserfc.jwk import RSAKey


def main():
    key = RSAKey.generate_key(2048)
    plaintext = b"benchmark jwe message"
    algorithms = ["RSA-OAEP-256", "A256GCM"]
    token = jwe.encrypt_compact({"alg": "RSA-OAEP-256", "enc": "A256GCM"}, plaintext, key, algorithms=algorithms)
    obj = jwe.decrypt_compact(token, key, algorithms=algorithms)
    assert obj.plaintext == plaintext
    print("ok", obj.plaintext == plaintext)


if __name__ == "__main__":
    main()
