use openssl::rand::rand_bytes;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let mut csrf_token = [0u8; 32];
    rand_bytes(&mut csrf_token)?;
    println!("ok {:02x}", csrf_token[0]);
    Ok(())
}
