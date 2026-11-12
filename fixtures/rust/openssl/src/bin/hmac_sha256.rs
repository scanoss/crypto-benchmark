use openssl::hash::MessageDigest;
use openssl::memcmp;
use openssl::pkey::PKey;
use openssl::sign::Signer;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let key = PKey::hmac(b"my secret")?;
    let data = b"hello, world!";

    let mut signer = Signer::new(MessageDigest::sha256(), &key)?;
    signer.update(data)?;
    let hmac = signer.sign_to_vec()?;

    let mut checker = Signer::new(MessageDigest::sha256(), &key)?;
    checker.update(data)?;
    let expected = checker.sign_to_vec()?;
    assert!(memcmp::eq(&hmac, &expected));
    println!("ok {}", hmac.len());
    Ok(())
}
