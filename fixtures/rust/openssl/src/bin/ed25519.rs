use openssl::pkey::PKey;
use openssl::sign::{Signer, Verifier};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let key = PKey::generate_ed25519()?;
    let message = b"signed release note";

    let mut signer = Signer::new_without_digest(&key)?;
    let signature = signer.sign_oneshot_to_vec(message)?;

    let mut verifier = Verifier::new_without_digest(&key)?;
    assert!(verifier.verify_oneshot(&signature, message)?);
    println!("ok {}", signature.len());
    Ok(())
}
