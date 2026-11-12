use k256::ecdsa::signature::{Signer, Verifier};
use k256::ecdsa::{Signature, SigningKey};
use k256::elliptic_curve::Generate;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let signing_key = SigningKey::generate();
    let message = b"transaction payload";
    let signature: Signature = signing_key.sign(message);

    let verifying_key = signing_key.verifying_key();
    verifying_key.verify(message, &signature)?;
    println!("ok {}", signature.to_bytes().len());
    Ok(())
}
