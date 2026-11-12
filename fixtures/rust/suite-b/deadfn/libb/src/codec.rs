use hmac::{Hmac, Mac};
use sha2::Sha256;

const KEY: &[u8] = b"cbench-suite-b-key";

pub fn sign(data: &[u8]) -> Vec<u8> {
    let mut mac = Hmac::<Sha256>::new_from_slice(KEY).expect("HMAC takes any key size");
    mac.update(data);
    mac.finalize().into_bytes().to_vec()
}

pub fn checksum(data: &[u8]) -> u8 {
    data.iter().fold(0, |c, b| c ^ b)
}

pub fn legacy_digest(data: &[u8]) -> Vec<u8> {
    use sha1::Digest;
    sha1::Sha1::digest(data).to_vec()
}
