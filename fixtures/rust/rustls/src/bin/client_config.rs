use std::sync::Arc;

use rustls::{ClientConfig, ClientConnection, RootCertStore};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let root_store = RootCertStore {
        roots: webpki_roots::TLS_SERVER_ROOTS.into(),
    };
    let config = ClientConfig::builder()
        .with_root_certificates(root_store)
        .with_no_client_auth();

    let server_name = "example.com".try_into()?;
    let conn = ClientConnection::new(Arc::new(config), server_name)?;
    assert!(conn.is_handshaking());
    println!("ok {}", conn.wants_write());
    Ok(())
}
