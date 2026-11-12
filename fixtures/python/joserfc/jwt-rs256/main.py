from joserfc import jwt
from joserfc.jwk import RSAKey


def main():
    key = RSAKey.generate_key(2048)
    encoded = jwt.encode({"alg": "RS256"}, {"sub": "bench", "iss": "crypto-benchmark"}, key)
    token = jwt.decode(encoded, key)
    jwt.JWTClaimsRegistry(iss={"essential": True, "value": "crypto-benchmark"}).validate(token.claims)
    assert token.claims["sub"] == "bench"
    print("ok", token.header["alg"])


if __name__ == "__main__":
    main()
