use sha1::{Digest, Sha1};

fn main() {
    let mut hasher = Sha1::new();
    hasher.update(b"hello world");
    let hash = hasher.finalize();
    let hex: String = hash.iter().map(|b| format!("{b:02x}")).collect();
    assert_eq!(hex, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed");
    println!("ok {hex}");
}
