use ring::rand::SystemRandom;
use ring::signature::{self, KeyPair, RsaKeyPair};

fn main() -> Result<(), ring::error::Unspecified> {
    let private_key_der = include_bytes!("../../keys/rsa-2048.pk8");
    let key_pair = RsaKeyPair::from_pkcs8(private_key_der).map_err(|_| ring::error::Unspecified)?;

    let rng = SystemRandom::new();
    let message = b"release-manifest.json";
    let mut sig = vec![0; key_pair.public().modulus_len()];
    key_pair.sign(&signature::RSA_PKCS1_SHA256, &rng, message, &mut sig)?;

    let public_key = signature::UnparsedPublicKey::new(&signature::RSA_PKCS1_2048_8192_SHA256, key_pair.public_key().as_ref());
    public_key.verify(message, &sig)?;
    println!("ok {}", sig.len());
    Ok(())
}
