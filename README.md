# Verkle Trie vs Merkle Patricia Trie explorations

This repository contains multiple benchmarks that are relevant to compare a Merkle Patricia Trie with a Verkle Trie.

## Benchmark results

AMD Ryzen 7 3800XT 8-Core Processor:
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
