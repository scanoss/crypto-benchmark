use openssl::hash::MessageDigest;
use openssl::pkcs5::pbkdf2_hmac;
use openssl::rand::rand_bytes;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let mut salt = [0u8; 16];
    rand_bytes(&mut salt)?;
    let mut key = [0u8; 32];
    pbkdf2_hmac(b"hunter2", &salt, 310_000, MessageDigest::sha256(), &mut key)?;
    println!("ok {:02x}", key[0]);
    Ok(())
}
