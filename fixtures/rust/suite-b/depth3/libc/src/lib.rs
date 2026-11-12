use hmac::{Hmac, Mac};
use sha2::Sha256;

const KEY: &[u8] = b"cbench-suite-b-key";

pub fn mac(data: &[u8]) -> Vec<u8> {
    let mut mac = Hmac::<Sha256>::new_from_slice(KEY).expect("HMAC takes any key size");
    mac.update(data);
    mac.finalize().into_bytes().to_vec()
}
