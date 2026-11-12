use sha2::{Digest, Sha384};

fn main() {
    let hash = Sha384::new_with_prefix(b"hello world").finalize();
    assert_eq!(hash.len(), 48);
    println!("ok {:02x}{:02x}", hash[0], hash[1]);
}
