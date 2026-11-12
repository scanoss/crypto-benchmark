use aes::Aes128;
use aes::cipher::{Array, BlockCipherDecrypt, BlockCipherEncrypt, KeyInit};

fn main() {
    let mut key = [0u8; 16];
    getrandom::fill(&mut key).expect("system random");
    let cipher = Aes128::new(&Array::from(key));

    let mut block = Array::from([42u8; 16]);
    let original = block;
    cipher.encrypt_block(&mut block);
    assert_ne!(block, original);
    cipher.decrypt_block(&mut block);
    assert_eq!(block, original);
    println!("ok");
}
