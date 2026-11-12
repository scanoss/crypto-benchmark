use std::sync::Arc;

use tokio::io::{AsyncReadExt, AsyncWriteExt, duplex};
use tokio_rustls::rustls::pki_types::pem::PemObject;
use tokio_rustls::rustls::pki_types::{CertificateDer, PrivateKeyDer, ServerName};
use tokio_rustls::rustls::{ClientConfig, RootCertStore, ServerConfig};
use tokio_rustls::{TlsAcceptor, TlsConnector};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let certs = CertificateDer::pem_slice_iter(include_bytes!("../../certs/localhost.pem"))
        .collect::<Result<Vec<_>, _>>()?;
    let key = PrivateKeyDer::from_pem_slice(include_bytes!("../../certs/localhost.key"))?;
    let server_config = ServerConfig::builder()
        .with_no_client_auth()
        .with_single_cert(certs, key)?;
    let acceptor = TlsAcceptor::from(Arc::new(server_config));

    let mut roots = RootCertStore::empty();
    roots.add(CertificateDer::from_pem_slice(include_bytes!("../../certs/ca.pem"))?)?;
    let client_config = ClientConfig::builder()
        .with_root_certificates(roots)
        .with_no_client_auth();
    let connector = TlsConnector::from(Arc::new(client_config));

    let (client_io, server_io) = duplex(64 * 1024);
    let server = tokio::spawn(async move {
        let mut stream = acceptor.accept(server_io).await?;
        let mut buf = [0u8; 5];
        stream.read_exact(&mut buf).await?;
        stream.write_all(&buf).await?;
        stream.shutdown().await?;
        Ok::<_, std::io::Error>(())
    });

    let domain = ServerName::try_from("localhost")?;
    let mut stream = connector.connect(domain, client_io).await?;
    stream.write_all(b"hello").await?;
    let mut echo = [0u8; 5];
    stream.read_exact(&mut echo).await?;
    server.await??;
    assert_eq!(&echo, b"hello");
    println!("ok {:?}", stream.get_ref().1.protocol_version());
    Ok(())
}
