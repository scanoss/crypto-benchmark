use aes::Aes192;
use aes::cipher::{Array, BlockCipherDecrypt, BlockCipherEncrypt, KeyInit};

fn main() {
    let key = Array::from([7u8; 24]);
    let mut block = Array::from(*b"sixteen byte msg");
    let cipher = Aes192::new(&key);
    cipher.encrypt_block(&mut block);
    cipher.decrypt_block(&mut block);
    assert_eq!(&block[..], b"sixteen byte msg");
    println!("ok");
}
