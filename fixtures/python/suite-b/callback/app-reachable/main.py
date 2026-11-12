import threading

import liba


def report():
    print(liba.fingerprint(b"hello").hex())


def greet():
    print("hello")


def main():
    worker = threading.Thread(target=report)
    worker.start()
    worker.join()


if __name__ == "__main__":
    main()
