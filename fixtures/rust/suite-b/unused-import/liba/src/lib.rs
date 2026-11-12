pub const VERSION: &str = "1.0.0";

pub fn fingerprint(data: &[u8]) -> Vec<u8> {
    cbench_unused_import_libb::digest(data)
}
