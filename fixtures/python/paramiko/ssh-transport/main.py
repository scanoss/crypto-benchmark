import socket
import threading

import paramiko


class Server(paramiko.ServerInterface):
    def get_allowed_auths(self, username):
        return "password"

    def check_auth_password(self, username, password):
        if username == "bench" and password == "s3cret":
            return paramiko.AUTH_SUCCESSFUL
        return paramiko.AUTH_FAILED


def main():
    server_sock, client_sock = socket.socketpair()
    server = paramiko.Transport(server_sock)
    server.add_server_key(paramiko.RSAKey.generate(bits=3072))
    started = threading.Event()
    threading.Thread(target=server.start_server, kwargs={"event": started, "server": Server()}).start()
    client = paramiko.Transport(client_sock)
    client.connect(username="bench", password="s3cret")
    started.wait(10)
    assert client.is_authenticated()
    print("ok", client.remote_cipher)
    client.close()
    server.close()


if __name__ == "__main__":
    main()
