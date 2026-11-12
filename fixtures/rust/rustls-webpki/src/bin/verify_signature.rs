use pki_types::CertificateDer;
use pki_types::pem::PemObject;
use ring::rand::SystemRandom;
use ring::signature::{ECDSA_P256_SHA256_ASN1_SIGNING, EcdsaKeyPair};
use webpki::EndEntityCert;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let leaf_der = CertificateDer::from_pem_slice(include_bytes!("../../certs/localhost.pem"))?;
    let leaf = EndEntityCert::try_from(&leaf_der)?;

    let pkcs8 = pki_types::PrivatePkcs8KeyDer::from_pem_slice(include_bytes!("../../certs/localhost.key"))?;
    let rng = SystemRandom::new();
    let key_pair = EcdsaKeyPair::from_pkcs8(&ECDSA_P256_SHA256_ASN1_SIGNING, pkcs8.secret_pkcs8_der(), &rng).map_err(|e| e.to_string())?;

    let message = b"server key exchange params";
    let signature = key_pair.sign(&rng, message).map_err(|_| "signing failed")?;
    leaf.verify_signature(webpki::ring::ECDSA_P256_SHA256, message, signature.as_ref())?;
    println!("ok");
    Ok(())
}
