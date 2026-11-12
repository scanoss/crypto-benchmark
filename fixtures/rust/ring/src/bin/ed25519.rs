use ring::rand::SystemRandom;
use ring::signature::{self, Ed25519KeyPair, KeyPair};

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = SystemRandom::new();
    let pkcs8_bytes = Ed25519KeyPair::generate_pkcs8(&rng)?;
    let key_pair = Ed25519KeyPair::from_pkcs8(pkcs8_bytes.as_ref()).map_err(|_| ring::error::Unspecified)?;

    const MESSAGE: &[u8] = b"hello, world";
    let sig = key_pair.sign(MESSAGE);

    let peer_public_key = signature::UnparsedPublicKey::new(&signature::ED25519, key_pair.public_key().as_ref());
    peer_public_key.verify(MESSAGE, sig.as_ref())?;
    println!("ok");
    Ok(())
}
