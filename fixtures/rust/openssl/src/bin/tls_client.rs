use openssl::ssl::{SslConnector, SslMethod, SslVersion};

fn main() -> Result<(), openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;
    builder.set_min_proto_version(Some(SslVersion::TLS1_2))?;
    builder.set_cipher_list("ECDHE+AESGCM:ECDHE+CHACHA20")?;
    let connector = builder.build();

    let config = connector.configure()?;
    let ssl = config.into_ssl("example.com")?;
    println!("ok {}", ssl.version_str());
    Ok(())
}
