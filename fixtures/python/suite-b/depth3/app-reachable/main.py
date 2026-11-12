import liba


def run(data):
    return liba.seal(data)


def main():
    print(run(b"hello").hex())


if __name__ == "__main__":
    main()
