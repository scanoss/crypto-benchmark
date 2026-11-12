import liba


def run(data):
    return liba.fingerprint(data)


def main():
    print(run(b"hello").hex())


if __name__ == "__main__":
    main()
