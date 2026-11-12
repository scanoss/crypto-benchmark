use ring::aead::{AES_256_GCM, Aad, LessSafeKey, NONCE_LEN, Nonce, UnboundKey};
use ring::rand::{SecureRandom, SystemRandom};

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = SystemRandom::new();
    let mut key_bytes = [0u8; 32];
    rng.fill(&mut key_bytes)?;
    let key = LessSafeKey::new(UnboundKey::new(&AES_256_GCM, &key_bytes)?);

    let mut nonce_bytes = [0u8; NONCE_LEN];
    rng.fill(&mut nonce_bytes)?;
    let mut in_out = b"secret message".to_vec();
    key.seal_in_place_append_tag(Nonce::assume_unique_for_key(nonce_bytes), Aad::empty(), &mut in_out)?;

    let plaintext = key.open_in_place(Nonce::assume_unique_for_key(nonce_bytes), Aad::empty(), &mut in_out)?;
    assert_eq!(plaintext, b"secret message");
    println!("ok");
    Ok(())
}
