use sha2::{Digest, Sha256};

fn main() {
    let mut hasher = Sha256::new();
    hasher.update(b"hello ");
    hasher.update(b"world");
    let hash = hasher.finalize();
    let hex: String = hash.iter().map(|b| format!("{b:02x}")).collect();
    assert_eq!(hex.len(), 64);
    println!("ok {hex}");
}
