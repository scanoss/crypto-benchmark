import hashlib


def encode(data):
    return _encode(data)


def _encode(data):
    return _hash_block(data)


def _hash_block(data):
    return hashlib.sha384(data).digest()
