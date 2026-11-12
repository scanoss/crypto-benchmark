pub fn protect(data: &[u8]) -> Vec<u8> {
    protect_inner(data)
}

fn protect_inner(data: &[u8]) -> Vec<u8> {
    cbench_wrapper_libb::encode(data)
}
