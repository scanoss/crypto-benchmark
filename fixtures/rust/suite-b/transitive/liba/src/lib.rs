pub fn fingerprint(data: &[u8]) -> Vec<u8> {
    cbench_transitive_libb::digest(data)
}
