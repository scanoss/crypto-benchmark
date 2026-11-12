use curve25519_dalek::montgomery::MontgomeryPoint;

fn main() {
    let mut alice_secret = [0u8; 32];
    getrandom::fill(&mut alice_secret).expect("system random");
    let mut bob_secret = [0u8; 32];
    getrandom::fill(&mut bob_secret).expect("system random");

    let alice_public = MontgomeryPoint::mul_base_clamped(alice_secret);
    let bob_public = MontgomeryPoint::mul_base_clamped(bob_secret);

    let alice_shared = bob_public.mul_clamped(alice_secret);
    let bob_shared = alice_public.mul_clamped(bob_secret);
    assert_eq!(alice_shared, bob_shared);
    println!("ok {:02x}", alice_shared.to_bytes()[0]);
}
