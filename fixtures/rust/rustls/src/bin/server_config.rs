use std::sync::Arc;

use rustls::pki_types::pem::PemObject;
use rustls::pki_types::{CertificateDer, PrivateKeyDer};
use rustls::{ServerConfig, ServerConnection};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let certs = CertificateDer::pem_slice_iter(include_bytes!("../../certs/localhost.pem"))
        .collect::<Result<Vec<_>, _>>()?;
    let key = PrivateKeyDer::from_pem_slice(include_bytes!("../../certs/localhost.key"))?;

    let config = ServerConfig::builder()
        .with_no_client_auth()
        .with_single_cert(certs, key)?;

    let conn = ServerConnection::new(Arc::new(config))?;
    println!("ok {}", conn.is_handshaking());
    Ok(())
}
