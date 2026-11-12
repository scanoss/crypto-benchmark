import datetime

from cryptography import x509
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.x509.oid import NameOID
from OpenSSL import SSL


def self_signed(key):
    name = x509.Name([x509.NameAttribute(NameOID.COMMON_NAME, "localhost")])
    now = datetime.datetime.now(datetime.timezone.utc)
    return (
        x509.CertificateBuilder()
        .subject_name(name)
        .issuer_name(name)
        .public_key(key.public_key())
        .serial_number(x509.random_serial_number())
        .not_valid_before(now)
        .not_valid_after(now + datetime.timedelta(days=365))
        .sign(key, hashes.SHA256())
    )


def main():
    key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
    cert = self_signed(key)
    ctx = SSL.Context(SSL.TLS_SERVER_METHOD)
    ctx.set_min_proto_version(SSL.TLS1_2_VERSION)
    ctx.use_certificate(cert)
    ctx.use_privatekey(key)
    ctx.check_privatekey()
    assert cert.public_key().key_size == 2048
    print("ok", cert.subject.rfc4514_string())


if __name__ == "__main__":
    main()
