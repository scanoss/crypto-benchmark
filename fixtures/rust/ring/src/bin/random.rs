use ring::rand::{SecureRandom, SystemRandom};

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = SystemRandom::new();
    let mut session_id = [0u8; 32];
    rng.fill(&mut session_id)?;
    let token: [u8; 16] = ring::rand::generate(&rng)?.expose();
    println!("ok {:02x}{:02x}", session_id[0], token[0]);
    Ok(())
}
