use sha2::{Digest, Sha512_256};

fn main() {
    let mut hasher = Sha512_256::new();
    hasher.update(b"hello world");
    let hash = hasher.finalize();
    assert_eq!(hash.len(), 32);
    println!("ok {:02x}{:02x}", hash[0], hash[1]);
}
