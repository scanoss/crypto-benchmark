use siphasher::sip128::{Hasher128, SipHasher24};
use std::hash::Hasher;

fn main() {
    let mut hasher = SipHasher24::new_with_keys(0x0706050403020100, 0x0f0e0d0c0b0a0908);
    hasher.write(b"content address");
    let digest = hasher.finish128();
    println!("ok {:032x}", digest.as_u128());
}
