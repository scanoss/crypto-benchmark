import liba


def main():
    h = liba.standard()
    print(h.sum(b"hello").hex())


if __name__ == "__main__":
    main()
