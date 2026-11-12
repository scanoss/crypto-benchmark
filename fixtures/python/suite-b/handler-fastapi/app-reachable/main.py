import sys

import uvicorn
from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
from fastapi.testclient import TestClient

import liba


def health():
    return PlainTextResponse("ok")


def fingerprint(data: str = ""):
    return PlainTextResponse(liba.fingerprint(data.encode()).hex())


def main():
    app = FastAPI()
    app.add_api_route("/health", health, methods=["GET"])
    app.add_api_route("/fingerprint", fingerprint, methods=["GET"])
    if sys.argv[1:] == ["serve"]:
        uvicorn.run(app, host="127.0.0.1", port=8080)
        return
    # Without "serve", answer one request in process and exit.
    response = TestClient(app).get("/fingerprint?data=hello")
    print(response.status_code, response.text)


if __name__ == "__main__":
    main()
