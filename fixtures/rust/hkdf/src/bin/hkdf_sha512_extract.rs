use hkdf::Hkdf;
use sha2::Sha512;

fn main() {
    let input_key_material = [0x42u8; 32];
    let salt = b"session-salt-v1";

    let (prk, hk) = Hkdf::<Sha512>::extract(Some(salt), &input_key_material);
    let mut enc_key = [0u8; 32];
    hk.expand(b"encryption key", &mut enc_key).expect("valid length");

    assert_eq!(prk.len(), 64);
    println!("ok {:02x}{:02x}", enc_key[0], enc_key[1]);
}
