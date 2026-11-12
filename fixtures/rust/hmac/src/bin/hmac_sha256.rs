use hmac::{Hmac, KeyInit, Mac};
use sha2::Sha256;

type HmacSha256 = Hmac<Sha256>;

fn main() {
    let key = b"my secret and secure key";
    let message = b"input message";

    let mut mac = HmacSha256::new_from_slice(key).expect("HMAC can take key of any size");
    mac.update(message);
    let code = mac.finalize().into_bytes();

    let mut verifier = HmacSha256::new_from_slice(key).expect("HMAC can take key of any size");
    verifier.update(message);
    verifier.verify_slice(&code).expect("tag matches");
    println!("ok {:02x}{:02x}", code[0], code[1]);
}
