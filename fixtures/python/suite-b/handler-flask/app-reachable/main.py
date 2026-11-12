import sys

from flask import Flask, request

import liba


def health():
    return "ok"


def fingerprint():
    return liba.fingerprint(request.args.get("data", "").encode()).hex()


def main():
    app = Flask(__name__)
    app.add_url_rule("/health", view_func=health)
    app.add_url_rule("/fingerprint", view_func=fingerprint)
    if sys.argv[1:] == ["serve"]:
        app.run(host="127.0.0.1", port=8080)
        return
    # Without "serve", answer one request in process and exit.
    response = app.test_client().get("/fingerprint?data=hello")
    print(response.status_code, response.get_data(as_text=True))


if __name__ == "__main__":
    main()
