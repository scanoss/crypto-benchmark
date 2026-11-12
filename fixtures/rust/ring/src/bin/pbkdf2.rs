use ring::rand::SecureRandom;
use ring::{digest, pbkdf2, rand};
use std::num::NonZeroU32;

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = rand::SystemRandom::new();
    let mut salt = [0u8; 16];
    rng.fill(&mut salt)?;
    let iterations = NonZeroU32::new(600_000).unwrap();
    let password = b"correct horse battery staple";

    let mut hash = [0u8; digest::SHA256_OUTPUT_LEN];
    pbkdf2::derive(pbkdf2::PBKDF2_HMAC_SHA256, iterations, &salt, password, &mut hash);
    pbkdf2::verify(pbkdf2::PBKDF2_HMAC_SHA256, iterations, &salt, password, &hash)?;
    println!("ok {:02x}", hash[0]);
    Ok(())
}
