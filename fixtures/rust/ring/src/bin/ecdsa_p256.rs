use ring::rand::SystemRandom;
use ring::signature::{self, ECDSA_P256_SHA256_FIXED, ECDSA_P256_SHA256_FIXED_SIGNING, EcdsaKeyPair, KeyPair};

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = SystemRandom::new();
    let pkcs8 = EcdsaKeyPair::generate_pkcs8(&ECDSA_P256_SHA256_FIXED_SIGNING, &rng)?;
    let key_pair = EcdsaKeyPair::from_pkcs8(&ECDSA_P256_SHA256_FIXED_SIGNING, pkcs8.as_ref(), &rng)
        .map_err(|_| ring::error::Unspecified)?;

    let message = b"order #1234";
    let sig = key_pair.sign(&rng, message)?;

    let public_key = signature::UnparsedPublicKey::new(&ECDSA_P256_SHA256_FIXED, key_pair.public_key().as_ref());
    public_key.verify(message, sig.as_ref())?;
    println!("ok");
    Ok(())
}
