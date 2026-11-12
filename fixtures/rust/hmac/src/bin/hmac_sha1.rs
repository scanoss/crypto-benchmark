use hmac::{Hmac, KeyInit, Mac};
use sha1::Sha1;

fn main() {
    let mut mac = Hmac::<Sha1>::new_from_slice(b"webhook-secret").expect("any key size");
    mac.update(b"{\"event\":\"push\"}");
    let signature = mac.finalize().into_bytes();
    assert_eq!(signature.len(), 20);
    println!("ok {:02x}{:02x}", signature[0], signature[1]);
}
