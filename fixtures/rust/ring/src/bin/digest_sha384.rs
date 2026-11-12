use ring::digest;

fn main() {
    let fingerprint = digest::digest(&digest::SHA384, b"-----BEGIN CERTIFICATE-----");
    assert_eq!(fingerprint.as_ref().len(), 48);
    println!("ok {:02x}", fingerprint.as_ref()[0]);
}
