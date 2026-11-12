use pki_types::pem::PemObject;
use pki_types::{CertificateDer, ServerName, UnixTime};
use webpki::{EndEntityCert, KeyUsage, anchor_from_trusted_cert};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let ca_der = CertificateDer::from_pem_slice(include_bytes!("../../certs/ca.pem"))?;
    let leaf_der = CertificateDer::from_pem_slice(include_bytes!("../../certs/localhost.pem"))?;

    let anchors = [anchor_from_trusted_cert(&ca_der)?];
    let leaf = EndEntityCert::try_from(&leaf_der)?;

    leaf.verify_for_usage(
        webpki::ALL_VERIFICATION_ALGS,
        &anchors,
        &[],
        UnixTime::now(),
        KeyUsage::server_auth(),
        None,
        None,
    )?;
    leaf.verify_is_valid_for_subject_name(&ServerName::try_from("localhost")?)?;
    println!("ok");
    Ok(())
}
