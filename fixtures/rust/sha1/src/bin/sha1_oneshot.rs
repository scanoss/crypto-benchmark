use sha1::{Digest, Sha1};

fn main() {
    let commit = b"blob 11\0hello world";
    let id = Sha1::digest(commit);
    assert_eq!(id.len(), 20);
    println!("ok {:02x}{:02x}", id[0], id[1]);
}
