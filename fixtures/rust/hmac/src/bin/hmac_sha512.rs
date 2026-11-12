use hmac::{KeyInit, Mac, SimpleHmac};
use sha2::Sha512;

fn main() {
    let mut mac = <SimpleHmac<Sha512> as KeyInit>::new_from_slice(b"api signing key").expect("any key size");
    mac.update(b"GET /v1/items");
    let tag = mac.finalize().into_bytes();
    assert_eq!(tag.len(), 64);
    println!("ok {:02x}{:02x}", tag[0], tag[1]);
}
