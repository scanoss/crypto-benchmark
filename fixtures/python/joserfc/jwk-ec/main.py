from joserfc.jwk import ECKey


def main():
    key = ECKey.generate_key("P-256")
    private_jwk = key.as_dict(private=True)
    public_jwk = key.as_dict(private=False)
    restored = ECKey.import_key(private_jwk)
    public = ECKey.import_key(public_jwk)
    assert restored.is_private and not public.is_private
    assert "d" in private_jwk and "d" not in public_jwk
    print("ok", public_jwk["crv"])


if __name__ == "__main__":
    main()
