use p256::ecdsa::signature::{Signer, Verifier};
use p256::ecdsa::{Signature, SigningKey, VerifyingKey};
use p256::elliptic_curve::Generate;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let signing_key = SigningKey::generate();
    let verifying_key_bytes = signing_key.verifying_key().to_sec1_point(false);

    let message = b"ECDSA proves knowledge of a secret number in the context of a single message";
    let signature: Signature = signing_key.sign(message);

    let verifying_key = VerifyingKey::from_sec1_bytes(verifying_key_bytes.as_ref())?;
    verifying_key.verify(message, &signature)?;
    println!("ok {}", signature.to_bytes().len());
    Ok(())
}
