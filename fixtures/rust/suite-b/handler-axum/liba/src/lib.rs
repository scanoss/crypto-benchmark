pub fn fingerprint(data: &[u8]) -> Vec<u8> {
    cbench_handler_axum_libb::digest(data)
}
