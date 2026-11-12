use sha2::{Digest, Sha384};

pub fn encode(data: &[u8]) -> Vec<u8> {
    encode_inner(data)
}

fn encode_inner(data: &[u8]) -> Vec<u8> {
    hash_block(data)
}

fn hash_block(data: &[u8]) -> Vec<u8> {
    Sha384::digest(data).to_vec()
}
