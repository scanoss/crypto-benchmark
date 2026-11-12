pub trait Hasher {
    fn sum(&self, data: &[u8]) -> Vec<u8>;
}

struct Standard;

impl Hasher for Standard {
    fn sum(&self, data: &[u8]) -> Vec<u8> {
        cbench_dispatch_libb::digest(data)
    }
}

struct Plain;

impl Hasher for Plain {
    fn sum(&self, data: &[u8]) -> Vec<u8> {
        data.to_vec()
    }
}

pub fn standard() -> Box<dyn Hasher> {
    Box::new(Standard)
}

pub fn plain() -> Box<dyn Hasher> {
    Box::new(Plain)
}
