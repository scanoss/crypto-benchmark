use md5::{Digest, Md5};

fn main() {
    let mut hasher = Md5::new();
    hasher.update(b"hello world");
    let hash = hasher.finalize();
    let hex: String = hash.iter().map(|b| format!("{b:02x}")).collect();
    assert_eq!(hex, "5eb63bbbe01eeed093cb22bb8f5acdc3");
    println!("ok {hex}");
}
