import liba


def main():
    print(liba.VERSION, liba.fingerprint(b"hello").hex())


if __name__ == "__main__":
    main()
