#[allow(dead_code)]
fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

#[allow(dead_code)]
fn run(data: &[u8]) -> Vec<u8> {
    cbench_direct_liba::fingerprint(data)
}

fn main() {
    println!("hello");
}
