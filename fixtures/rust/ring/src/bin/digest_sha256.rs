use ring::digest;

fn main() {
    let one_shot = digest::digest(&digest::SHA256, b"hello, world");

    let mut ctx = digest::Context::new(&digest::SHA256);
    ctx.update(b"hello");
    ctx.update(b", world");
    let multi_part = ctx.finish();

    assert_eq!(one_shot.as_ref(), multi_part.as_ref());
    println!("ok {:02x}", one_shot.as_ref()[0]);
}
