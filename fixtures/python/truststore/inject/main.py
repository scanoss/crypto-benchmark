import ssl

import truststore


def main():
    truststore.inject_into_ssl()
    try:
        ctx = ssl.create_default_context()
        ctx.minimum_version = ssl.TLSVersion.TLSv1_2
        assert isinstance(ctx, truststore.SSLContext)
    finally:
        truststore.extract_from_ssl()
    assert not isinstance(ssl.create_default_context(), truststore.SSLContext)
    print("ok", ctx.minimum_version.name)


if __name__ == "__main__":
    main()
