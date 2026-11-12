from OpenSSL import SSL


def main():
    ctx = SSL.Context(SSL.TLS_METHOD)
    ctx.set_min_proto_version(SSL.TLS1_2_VERSION)
    ctx.set_verify(SSL.VERIFY_PEER | SSL.VERIFY_FAIL_IF_NO_PEER_CERT)
    ctx.set_cipher_list(b"ECDHE+AESGCM:ECDHE+CHACHA20:!aNULL:!MD5")
    ctx.set_default_verify_paths()
    assert ctx.get_verify_mode() & SSL.VERIFY_PEER
    print("ok", ctx.get_verify_mode())


if __name__ == "__main__":
    main()
