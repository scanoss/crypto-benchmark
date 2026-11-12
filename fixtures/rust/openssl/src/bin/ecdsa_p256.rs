use openssl::ec::{EcGroup, EcKey};
use openssl::ecdsa::EcdsaSig;
use openssl::nid::Nid;
use openssl::sha::sha256;

fn main() -> Result<(), openssl::error::ErrorStack> {
    let group = EcGroup::from_curve_name(Nid::X9_62_PRIME256V1)?;
    let key = EcKey::generate(&group)?;

    let digest = sha256(b"firmware image");
    let signature = EcdsaSig::sign(&digest, &key)?;
    assert!(signature.verify(&digest, &key)?);
    println!("ok {}", signature.to_der()?.len());
    Ok(())
}
