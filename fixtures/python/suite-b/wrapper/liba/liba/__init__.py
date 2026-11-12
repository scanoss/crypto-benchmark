import libb


def protect(data):
    return _protect(data)


def _protect(data):
    return libb.encode(data)
