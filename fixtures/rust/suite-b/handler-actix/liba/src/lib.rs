pub fn fingerprint(data: &[u8]) -> Vec<u8> {
    cbench_handler_actix_libb::digest(data)
}
