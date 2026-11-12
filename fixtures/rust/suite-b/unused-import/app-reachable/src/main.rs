#[allow(dead_code)]
fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

fn main() {
    println!("{} {}", cbench_unused_import_liba::VERSION, hex(&cbench_unused_import_liba::fingerprint(b"hello")));
}
