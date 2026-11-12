use curve25519_dalek::ristretto::RistrettoPoint;
use curve25519_dalek::scalar::Scalar;
use sha2::Sha512;

fn main() {
    let challenge = Scalar::hash_from_bytes::<Sha512>(b"transcript: commitment || message");
    let generator = RistrettoPoint::hash_from_bytes::<Sha512>(b"domain separated generator");
    let point = generator * challenge;
    println!("ok {:02x}", point.compress().to_bytes()[0]);
}
