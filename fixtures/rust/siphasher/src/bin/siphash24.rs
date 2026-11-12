use siphasher::sip::SipHasher24;
use std::hash::Hasher;

fn main() {
    let key: [u8; 16] = *b"0123456789abcdef";
    let hasher = SipHasher24::new_with_key(&key);
    let tag = hasher.hash(b"bucket key");

    let mut streaming = SipHasher24::new_with_key(&key);
    streaming.write(b"bucket ");
    streaming.write(b"key");
    assert_eq!(streaming.finish(), tag);
    println!("ok {tag:016x}");
}
