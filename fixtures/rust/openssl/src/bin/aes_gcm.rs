use openssl::rand::rand_bytes;
use openssl::symm::{Cipher, decrypt_aead, encrypt_aead};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let mut key = [0u8; 16];
    rand_bytes(&mut key)?;
    let mut nonce = [0u8; 12];
    rand_bytes(&mut nonce)?;
    let aad = b"header";

    let mut tag = [0u8; 16];
    let ciphertext = encrypt_aead(Cipher::aes_128_gcm(), &key, Some(&nonce), aad, b"payload", &mut tag)?;
    let plaintext = decrypt_aead(Cipher::aes_128_gcm(), &key, Some(&nonce), aad, &ciphertext, &tag)?;
    assert_eq!(plaintext, b"payload");
    println!("ok");
    Ok(())
}
