use openssl::hash::MessageDigest;
use openssl::pkey::PKey;
use openssl::rsa::Rsa;
use openssl::sign::{Signer, Verifier};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let keypair = PKey::from_rsa(Rsa::generate(3072)?)?;
    let data = b"hello, world!";

    let mut signer = Signer::new(MessageDigest::sha256(), &keypair)?;
    signer.update(data)?;
    let signature = signer.sign_to_vec()?;

    let mut verifier = Verifier::new(MessageDigest::sha256(), &keypair)?;
    verifier.update(data)?;
    assert!(verifier.verify(&signature)?);
    println!("ok {}", signature.len());
    Ok(())
}
