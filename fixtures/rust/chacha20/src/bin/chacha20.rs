use chacha20::ChaCha20;
use chacha20::cipher::{KeyIvInit, StreamCipher, StreamCipherSeek};

fn main() {
    let mut key = [0u8; 32];
    getrandom::fill(&mut key).expect("system random");
    let mut nonce = [0u8; 12];
    getrandom::fill(&mut nonce).expect("system random");

    let mut cipher = ChaCha20::new(&key.into(), &nonce.into());
    let mut buffer = *b"attack at dawn!!";
    cipher.apply_keystream(&mut buffer);
    cipher.seek(0u32);
    cipher.apply_keystream(&mut buffer);
    assert_eq!(&buffer, b"attack at dawn!!");
    println!("ok");
}
