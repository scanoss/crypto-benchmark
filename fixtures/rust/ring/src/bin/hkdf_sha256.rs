use ring::hkdf;

fn main() -> Result<(), ring::error::Unspecified> {
    let salt = hkdf::Salt::new(hkdf::HKDF_SHA256, b"protocol salt v1");
    let prk = salt.extract(b"input keying material");
    let okm = prk.expand(&[b"client write key"], hkdf::HKDF_SHA256)?;
    let mut key = [0u8; 32];
    okm.fill(&mut key)?;
    println!("ok {:02x}", key[0]);
    Ok(())
}
