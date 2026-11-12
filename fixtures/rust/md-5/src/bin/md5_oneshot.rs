use md5::{Digest, Md5};

fn main() {
    let etag = Md5::digest(b"file contents");
    assert_eq!(etag.len(), 16);
    println!("ok {:02x}{:02x}", etag[0], etag[1]);
}
