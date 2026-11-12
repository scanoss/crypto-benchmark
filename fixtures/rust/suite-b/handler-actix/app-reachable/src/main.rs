use std::collections::HashMap;
use std::io::{Read, Write};
use std::net::{SocketAddr, TcpStream};

use actix_web::{web, App, HttpServer};

fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

async fn health() -> &'static str {
    "ok"
}

#[allow(dead_code)]
async fn fingerprint(query: web::Query<HashMap<String, String>>) -> String {
    let data = query.get("data").map(String::as_str).unwrap_or("");
    hex(&cbench_handler_actix_liba::fingerprint(data.as_bytes()))
}

// get sends one HTTP/1.1 request and returns the status code and body.
fn get(addr: SocketAddr, path: &str) -> std::io::Result<String> {
    let mut stream = TcpStream::connect(addr)?;
    write!(stream, "GET {path} HTTP/1.1\r\nHost: {addr}\r\nConnection: close\r\n\r\n")?;
    let mut response = String::new();
    stream.read_to_string(&mut response)?;
    let status = response.split(' ').nth(1).unwrap_or("").to_string();
    let body = response.split("\r\n\r\n").nth(1).unwrap_or("").to_string();
    Ok(format!("{status} {body}"))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let server = HttpServer::new(|| {
        App::new()
            .route("/health", web::get().to(health))
            .route("/fingerprint", web::get().to(fingerprint))
    })
    .workers(1);
    if std::env::args().nth(1).as_deref() == Some("serve") {
        return server.bind(("127.0.0.1", 8080))?.run().await;
    }
    // Without "serve", answer one request on an ephemeral port and exit.
    let server = server.bind(("127.0.0.1", 0))?;
    let addr = server.addrs()[0];
    let running = server.run();
    let handle = running.handle();
    let task = actix_web::rt::spawn(running);
    let response = actix_web::rt::task::spawn_blocking(move || get(addr, "/fingerprint?data=hello"))
        .await
        .expect("request task")?;
    println!("{response}");
    handle.stop(true).await;
    task.await.expect("server task")
}
