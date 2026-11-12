use hex_literal::hex;
use hkdf::Hkdf;
use sha2::Sha256;

fn main() {
    let ikm = hex!("0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b");
    let salt = hex!("000102030405060708090a0b0c");
    let info = hex!("f0f1f2f3f4f5f6f7f8f9");

    let hk = Hkdf::<Sha256>::new(Some(&salt[..]), &ikm);
    let mut okm = [0u8; 42];
    hk.expand(&info, &mut okm).expect("42 is a valid length for Sha256 to output");

    assert_eq!(okm[..4], hex!("3cb25f25"));
    println!("ok {:02x}{:02x}", okm[0], okm[1]);
}
