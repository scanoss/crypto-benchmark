import libc


def seal(data):
    return libc.mac(data)
