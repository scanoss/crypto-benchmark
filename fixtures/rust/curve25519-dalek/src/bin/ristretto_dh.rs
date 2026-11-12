use curve25519_dalek::ristretto::{CompressedRistretto, RistrettoPoint};
use curve25519_dalek::scalar::Scalar;
use getrandom::SysRng;
use getrandom::rand_core::UnwrapErr;

fn main() {
    let mut rng = UnwrapErr(SysRng);
    let alice_secret = Scalar::random(&mut rng);
    let bob_secret = Scalar::random(&mut rng);

    let alice_public = RistrettoPoint::mul_base(&alice_secret).compress();
    let bob_public = RistrettoPoint::mul_base(&bob_secret).compress();

    let alice_shared = alice_secret * decode(&bob_public);
    let bob_shared = bob_secret * decode(&alice_public);
    assert_eq!(alice_shared.compress(), bob_shared.compress());
    println!("ok {:02x}", alice_shared.compress().to_bytes()[0]);
}

fn decode(point: &CompressedRistretto) -> RistrettoPoint {
    point.decompress().expect("valid ristretto encoding")
}
