use ring::{hmac, rand};

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = rand::SystemRandom::new();
    let key = hmac::Key::generate(hmac::HMAC_SHA256, &rng)?;

    let msg = "hello, world";
    let tag = hmac::sign(&key, msg.as_bytes());
    hmac::verify(&key, msg.as_bytes(), tag.as_ref())?;
    println!("ok {}", tag.as_ref().len());
    Ok(())
}
