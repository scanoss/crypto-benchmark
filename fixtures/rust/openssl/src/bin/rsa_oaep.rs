use openssl::rsa::{Padding, Rsa};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let rsa = Rsa::generate(2048)?;
    let data = b"session key bytes";

    let mut encrypted = vec![0; rsa.size() as usize];
    let len = rsa.public_encrypt(data, &mut encrypted, Padding::PKCS1_OAEP)?;

    let mut decrypted = vec![0; rsa.size() as usize];
    let n = rsa.private_decrypt(&encrypted[..len], &mut decrypted, Padding::PKCS1_OAEP)?;
    assert_eq!(&decrypted[..n], data);
    println!("ok {len}");
    Ok(())
}
