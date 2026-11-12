use chacha20::XChaCha20;
use chacha20::cipher::{KeyIvInit, StreamCipher};

fn main() {
    let key = [0x42u8; 32];
    let nonce = [0x24u8; 24];
    let mut buffer = *b"extended nonce variant";

    let mut encryptor = XChaCha20::new(&key.into(), &nonce.into());
    encryptor.apply_keystream(&mut buffer);
    let mut decryptor = XChaCha20::new(&key.into(), &nonce.into());
    decryptor.apply_keystream(&mut buffer);
    assert_eq!(&buffer, b"extended nonce variant");
    println!("ok");
}
