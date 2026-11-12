use openssl::pkey::PKey;
use openssl::rsa::Rsa;
use openssl::symm::Cipher;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let rsa = Rsa::generate(2048)?;
    let pem = rsa.private_key_to_pem_passphrase(Cipher::aes_256_cbc(), b"passphrase")?;
    let public_pem = rsa.public_key_to_pem()?;

    let restored = PKey::private_key_from_pem_passphrase(&pem, b"passphrase")?;
    let public = PKey::public_key_from_pem(&public_pem)?;
    assert!(restored.public_eq(&public));
    println!("ok {}", restored.bits());
    Ok(())
}
