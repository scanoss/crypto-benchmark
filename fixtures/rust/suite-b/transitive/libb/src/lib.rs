use sha2::{Digest, Sha512};

pub fn digest(data: &[u8]) -> Vec<u8> {
    Sha512::digest(data).to_vec()
}
