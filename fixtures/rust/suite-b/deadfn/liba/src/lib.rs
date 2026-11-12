pub fn sign(data: &[u8]) -> Vec<u8> {
    cbench_deadfn_libb::codec::sign(data)
}

pub fn checksum(data: &[u8]) -> u8 {
    cbench_deadfn_libb::codec::checksum(data)
}
