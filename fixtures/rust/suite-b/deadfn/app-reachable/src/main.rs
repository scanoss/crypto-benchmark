#[allow(dead_code)]
fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

fn main() {
    println!("{}", hex(&cbench_deadfn_liba::sign(b"hello")));
}
