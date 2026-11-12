use std::io::{Read, Write};
use std::sync::Arc;

use rustls::pki_types::pem::PemObject;
use rustls::pki_types::{CertificateDer, PrivateKeyDer};
use rustls::{ClientConfig, ClientConnection, Connection, RootCertStore, ServerConfig, ServerConnection};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let ca = CertificateDer::from_pem_slice(include_bytes!("../../certs/ca.pem"))?;
    let mut roots = RootCertStore::empty();
    roots.add(ca)?;
    let client_config = ClientConfig::builder_with_protocol_versions(&[&rustls::version::TLS13])
        .with_root_certificates(roots)
        .with_no_client_auth();

    let chain = CertificateDer::pem_slice_iter(include_bytes!("../../certs/localhost.pem"))
        .collect::<Result<Vec<_>, _>>()?;
    let key = PrivateKeyDer::from_pem_slice(include_bytes!("../../certs/localhost.key"))?;
    let server_config = ServerConfig::builder()
        .with_no_client_auth()
        .with_single_cert(chain, key)?;

    let mut client: Connection = ClientConnection::new(Arc::new(client_config), "localhost".try_into()?)?.into();
    let mut server: Connection = ServerConnection::new(Arc::new(server_config))?.into();

    while client.is_handshaking() || server.is_handshaking() {
        transfer(&mut client, &mut server)?;
        transfer(&mut server, &mut client)?;
    }

    client.writer().write_all(b"hello over TLS 1.3")?;
    transfer(&mut client, &mut server)?;
    let mut received = Vec::new();
    let _ = server.reader().read_to_end(&mut received);
    assert_eq!(received, b"hello over TLS 1.3");
    println!("ok {:?}", client.negotiated_cipher_suite().map(|s| s.suite()));
    Ok(())
}

fn transfer(from: &mut Connection, to: &mut Connection) -> Result<(), Box<dyn std::error::Error>> {
    let mut buf = Vec::new();
    while from.wants_write() {
        from.write_tls(&mut buf)?;
    }
    let mut slice = &buf[..];
    while !slice.is_empty() {
        to.read_tls(&mut slice)?;
    }
    to.process_new_packets()?;
    Ok(())
}
