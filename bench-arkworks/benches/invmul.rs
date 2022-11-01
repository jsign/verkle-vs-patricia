use ark_bls12_381::fr::Fr;
use ark_ff::Field;
use ark_std::{
    rand::{rngs::StdRng, SeedableRng},
    UniformRand,
};
use criterion::{criterion_group, criterion_main, Criterion};
use std::ops::Mul;

fn criterion_benchmark(c: &mut Criterion) {
    let mut rng = StdRng::seed_from_u64(42);
    let mut x = Fr::rand(&mut rng);
    let y = Fr::rand(&mut rng);

    c.bench_function("inverse", |b| {
        b.iter(|| {
            x = x.inverse().unwrap();
        })
    });
    c.bench_function("mul", |b| {
        b.iter(|| {
            x = x.mul(y);
        })
    });
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
