use std::collections::HashMap;

use axum::body::Body;
use axum::extract::Query;
use axum::http::Request;
use axum::routing::get;
use axum::Router;
use tower::ServiceExt;

fn hex(bytes: &[u8]) -> String {
    bytes.iter().map(|b| format!("{b:02x}")).collect()
}

async fn health() -> &'static str {
    "ok"
}

#[allow(dead_code)]
async fn fingerprint(Query(params): Query<HashMap<String, String>>) -> String {
    let data = params.get("data").map(String::as_str).unwrap_or("");
    hex(&cbench_handler_axum_liba::fingerprint(data.as_bytes()))
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/health", get(health));
    if std::env::args().nth(1).as_deref() == Some("serve") {
        let listener = tokio::net::TcpListener::bind("127.0.0.1:8080").await.unwrap();
        axum::serve(listener, app).await.unwrap();
        return;
    }
    // Without "serve", answer one request in process and exit.
    let request = Request::get("/fingerprint?data=hello").body(Body::empty()).unwrap();
    let response = app.oneshot(request).await.unwrap();
    let status = response.status().as_u16();
    let body = axum::body::to_bytes(response.into_body(), usize::MAX).await.unwrap();
    println!("{status} {}", String::from_utf8_lossy(&body));
}
