#[allow(dead_code)]
fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

fn main() {
    let h = cbench_dispatch_liba::standard();
    println!("{}", hex(&h.sum(b"hello")));
}
