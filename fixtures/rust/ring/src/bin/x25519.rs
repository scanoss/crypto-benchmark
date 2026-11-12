use ring::agreement::{self, EphemeralPrivateKey, UnparsedPublicKey, X25519};
use ring::rand::SystemRandom;

fn main() -> Result<(), ring::error::Unspecified> {
    let rng = SystemRandom::new();
    let my_private_key = EphemeralPrivateKey::generate(&X25519, &rng)?;
    let my_public_key = my_private_key.compute_public_key()?;

    let peer_private_key = EphemeralPrivateKey::generate(&X25519, &rng)?;
    let peer_public_key = UnparsedPublicKey::new(&X25519, peer_private_key.compute_public_key()?.as_ref().to_vec());

    let mine = agreement::agree_ephemeral(my_private_key, &peer_public_key, |key_material| key_material.to_vec())?;
    let theirs = agreement::agree_ephemeral(peer_private_key, &UnparsedPublicKey::new(&X25519, my_public_key.as_ref()), |key_material| key_material.to_vec())?;
    assert_eq!(mine, theirs);
    println!("ok");
    Ok(())
}
