use ecdsa::signature::{Signer, Verifier};
use ecdsa::{SigningKey, VerifyingKey, der};
use p256::NistP256;
use p256::elliptic_curve::Generate;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let original: SigningKey<NistP256> = SigningKey::generate();
    let secret_bytes = original.to_bytes();

    let signing_key = SigningKey::<NistP256>::from_slice(&secret_bytes)?;
    let signature: der::Signature<NistP256> = signing_key.sign(b"config file v2");

    let verifying_key: &VerifyingKey<NistP256> = signing_key.verifying_key();
    verifying_key.verify(b"config file v2", &signature)?;
    println!("ok {}", signature.as_bytes().len());
    Ok(())
}
