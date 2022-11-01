# Verkle Trie vs Merkle Patricia Trie explorations

This repository contains multiple benchmarks that are relevant to compare a Merkle Patricia Trie with a Verkle Trie.

## Run

To run the benchmarks you need the Go, C, and Rust compilers installed, and run:

- `make bench-go`: to run the Go benchmarks.
- `make bench-blst`: to run the C benchmarks.
- `make bench-arkworks`: to run the Rust benchmarks.

## Benchmark results

### AMD Ryzen 7 3800XT 8-Core Processor

Go benchmarks (with _Intel ADX_ CPU instructions):
```
goos: linux
goarch: amd64
pkg: github.com/jsign/keccak-vs-pedersen
cpu: AMD Ryzen 7 3800XT 8-Core Processor            
BenchmarkFpInverse/go-ipa-16                      722480              1665 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-16                     831808              1406 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-16                        78093471                14.87 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-16                       10473631               113.0 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=32-16                       2707854               440.5 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=64-16                       2905962               412.9 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=128-16                      3171886               372.5 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=256-16                      1715360               685.3 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=512-16                       893442              1336 ns/op               0 B/op          0 allocs/op
BenchmarkKeccak/size=1024-16                      397701              2888 ns/op               0 B/op          0 allocs/op
BenchmarkPedersenCommit/non_zero_entries=1-16              29551             39907 ns/op             39907 ns/value        33929 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=2-16              25261             47237 ns/op             23619 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=4-16              19268             60985 ns/op             15247 ns/value        33920 B/op         18 allocs/op
BenchmarkPedersenCommit/non_zero_entries=8-16              13380             90209 ns/op             11276 ns/value        33929 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=16-16              8083            147664 ns/op              9229 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=32-16              4119            265643 ns/op              8301 ns/value        33930 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=64-16              2246            491729 ns/op              7683 ns/value        33929 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=128-16             1119            986195 ns/op              7705 ns/value        33932 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=256-16              598           1904488 ns/op              7440 ns/value        33943 B/op         20 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=1-16             82398             13985 ns/op             368 B/op          7 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=4-16             87585             13944 ns/op             368 B/op          7 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=16-16            83296             14103 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=64-16            86702             13843 ns/op             368 B/op          7 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=256-16           81349             13784 ns/op             368 B/op          7 allocs/op
BenchmarkHashAddress/keccak-16                                   2691130               445.9 ns/op             0 B/op          0 allocs/op
BenchmarkHashAddress/pedersen_hash-16                             160214              6766 ns/op             160 B/op          3 allocs/op
BenchmarkHashStorageSlot/keccak-16                               2706037               442.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashStorageSlot/pedersen_hash-16                         133964              8609 ns/op             416 B/op         11 allocs/op
```

Go (without _Intel ADX_ CPU instructions):
```
BenchmarkFpInverse/go-ipa-16              705946              1629 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-16             714726              1427 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-16                41319416                24.52 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-16               10759808               110.6 ns/op             0 B/op          0 allocs/op
```

C BLST benchmark:
```
Inv(x) takes 1314 ns/op
Mul(a,b) takes 73 ns/op
```

Rust `arkworks-rs` benchmarks:
```
inverse                 time:   [2.1817 µs 2.1942 µs 2.2123 µs]                     
mul                     time:   [16.340 ns 16.414 ns 16.504 ns]                 
```

### Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz

Go benchmarks (with _Intel ADX_ CPU instructions):

```
goos: darwin
goarch: amd64
pkg: github.com/jsign/verkle-vs-patricia
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkFpInverse/go-ipa-8               618098              1866 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-8              798846              1501 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-8                 72805527                16.61 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-8                 9406170               118.5 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=32-8                2694834               442.2 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=64-8                2566628               424.5 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=128-8               2991798               388.3 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=256-8               1601151               747.5 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=512-8                817632              1459 ns/op               0 B/op          0 allocs/op
BenchmarkKeccak/size=1024-8               414661              2900 ns/op               0 B/op          0 allocs/op
BenchmarkPedersenCommit/non_zero_entries=1-8               28960             41144 ns/op             41146 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=2-8               23982             47916 ns/op             23960 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=4-8               18398             63763 ns/op             15942 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=8-8               12680             93300 ns/op             11663 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=16-8               7726            157905 ns/op              9870 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=32-8               4231            283232 ns/op              8851 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=64-8               2155            533065 ns/op              8329 ns/value        33929 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=128-8              1032           1029765 ns/op              8045 ns/value        33939 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=256-8               538           2283879 ns/op              8922 ns/value        33948 B/op         21 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=1-8              67816             17282 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=4-8              70353             17237 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=16-8             69090             16906 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=64-8             67299             16929 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=256-8            71257             18147 ns/op             368 B/op          7 allocs/op
BenchmarkHashAddress/keccak-8                                    2593776               454.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashAddress/pedersen_hash-8                              164329              7201 ns/op             160 B/op          3 allocs/op
BenchmarkHashStorageSlot/keccak-8                                2643908               445.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashStorageSlot/pedersen_hash-8                          130675              9139 ns/op             416 B/op         11 allocs/op
```

Go (without _Intel ADX_ CPU instructions):
```
BenchmarkFpInverse/go-ipa-8               563982              2384 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-8              655209              2146 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-8                 39916938                34.68 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-8                 8970868               133.3 ns/op             0 B/op          0 allocs/op
```

C BLST benchmark:
```
Inv(x) takes 1418 ns/op
Mul(a,b) takes 142 ns/op
```

Rust `arkworks-rs` benchmarks:
```
inverse                 time:   [2.0268 µs 2.0381 µs 2.0498 µs]
mul                     time:   [23.864 ns 24.053 ns 24.313 ns]
```
