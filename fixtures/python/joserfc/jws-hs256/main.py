from joserfc import jws
from joserfc.jwk import OctKey


def main():
    key = OctKey.generate_key(256)
    token = jws.serialize_compact({"alg": "HS256"}, b"benchmark payload", key)
    obj = jws.deserialize_compact(token, key)
    assert obj.payload == b"benchmark payload"
    print("ok", obj.headers()["alg"])


if __name__ == "__main__":
    main()
