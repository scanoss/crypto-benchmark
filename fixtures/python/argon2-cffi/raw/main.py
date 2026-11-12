import os

from argon2.low_level import Type, hash_secret_raw


def main():
    password = b"correct horse battery staple"
    salt = os.urandom(16)

    key = hash_secret_raw(
        password, salt, time_cost=3, memory_cost=2**16, parallelism=4, hash_len=32, type=Type.ID
    )
    again = hash_secret_raw(
        password, salt, time_cost=3, memory_cost=2**16, parallelism=4, hash_len=32, type=Type.ID
    )
    assert key == again
    print("ok", key.hex())


if __name__ == "__main__":
    main()
