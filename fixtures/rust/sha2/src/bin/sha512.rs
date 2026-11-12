use sha2::{Digest, Sha512};

fn main() {
    let hash = Sha512::digest(b"hello world");
    let hex: String = hash.iter().map(|b| format!("{b:02x}")).collect();
    assert_eq!(hex.len(), 128);
    println!("ok {}", &hex[..16]);
}
