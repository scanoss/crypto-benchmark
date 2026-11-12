import libb


class Hasher:
    def sum(self, data):
        raise NotImplementedError


class _Standard(Hasher):
    def sum(self, data):
        return libb.digest(data)


class _Plain(Hasher):
    def sum(self, data):
        return bytes(data)


def standard():
    return _Standard()


def plain():
    return _Plain()
