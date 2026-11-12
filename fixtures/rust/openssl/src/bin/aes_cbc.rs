use openssl::rand::rand_bytes;
use openssl::symm::{Cipher, decrypt, encrypt};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let cipher = Cipher::aes_256_cbc();
    let mut key = [0u8; 32];
    rand_bytes(&mut key)?;
    let mut iv = [0u8; 16];
    rand_bytes(&mut iv)?;

    let ciphertext = encrypt(cipher, &key, Some(&iv), b"Some Crypto Text")?;
    let plaintext = decrypt(cipher, &key, Some(&iv), &ciphertext)?;
    assert_eq!(plaintext, b"Some Crypto Text");
    println!("ok {}", ciphertext.len());
    Ok(())
}
