use openssl::hash::{Hasher, MessageDigest, hash};
use openssl::sha::sha256;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let digest = hash(MessageDigest::sha256(), b"hello world")?;

    let mut hasher = Hasher::new(MessageDigest::sha256())?;
    hasher.update(b"hello ")?;
    hasher.update(b"world")?;
    let streamed = hasher.finish()?;

    let direct = sha256(b"hello world");
    assert_eq!(&*digest, &*streamed);
    assert_eq!(&*digest, &direct[..]);
    println!("ok {:02x}", direct[0]);
    Ok(())
}
