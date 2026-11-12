import ssl

import truststore


def main():
    ctx = truststore.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
    ctx.minimum_version = ssl.TLSVersion.TLSv1_2
    assert ctx.verify_mode == ssl.CERT_REQUIRED and ctx.check_hostname
    print("ok", ctx.minimum_version.name)


if __name__ == "__main__":
    main()
