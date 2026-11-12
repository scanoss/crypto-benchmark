use siphasher::sip::SipHasher13;

fn main() {
    let hasher = SipHasher13::new_with_keys(0x0706050403020100, 0x0f0e0d0c0b0a0908);
    let digest = hasher.hash(b"request-id-42");
    println!("ok {digest:016x}");
}
