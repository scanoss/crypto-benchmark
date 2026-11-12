use ring::aead::{Aad, CHACHA20_POLY1305, LessSafeKey, Nonce, UnboundKey};

fn main() -> Result<(), ring::error::Unspecified> {
    let key_bytes = [0x11u8; 32];
    let sealing_key = LessSafeKey::new(UnboundKey::new(&CHACHA20_POLY1305, &key_bytes)?);
    let opening_key = LessSafeKey::new(UnboundKey::new(&CHACHA20_POLY1305, &key_bytes)?);

    let nonce = [0u8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1];
    let mut in_out = b"record payload".to_vec();
    let tag = sealing_key.seal_in_place_separate_tag(Nonce::assume_unique_for_key(nonce), Aad::from(b"header"), &mut in_out)?;
    in_out.extend_from_slice(tag.as_ref());

    let plaintext = opening_key.open_in_place(Nonce::assume_unique_for_key(nonce), Aad::from(b"header"), &mut in_out)?;
    assert_eq!(plaintext, b"record payload");
    println!("ok");
    Ok(())
}
