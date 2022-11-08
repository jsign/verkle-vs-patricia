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
BenchmarkFpInverse/go-ipa-16                     1048089              1116 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-16                     834487              1409 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-16                        79013084                14.83 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-16                        9497001               108.3 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=32-16                       2811439               406.0 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=64-16                       3040315               395.3 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=128-16                      3415230               349.8 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=256-16                      1752966               661.4 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=512-16                       904863              1325 ns/op               0 B/op          0 allocs/op
BenchmarkKeccak/size=1024-16                      383635              2620 ns/op               0 B/op          0 allocs/op
BenchmarkPedersenCommit/non_zero_entries=1-16              30967             38261 ns/op             38261 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=2-16              25597             45282 ns/op             22641 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=4-16              19479             59399 ns/op             14850 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=8-16              13435             87227 ns/op             10903 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=16-16              8191            144080 ns/op              9005 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=32-16              4604            264342 ns/op              8261 ns/value   33937 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=64-16              2194            486668 ns/op              7604 ns/value   33930 B/op          19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=128-16             1051            982158 ns/op              7673 ns/value   33940 B/op          20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=256-16              602           1935915 ns/op              7562 ns/value   33951 B/op          21 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=1-16            132586              8819 ns/op             376 B/op              8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=4-16            128914              8798 ns/op             376 B/op              8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=16-16           130920              8826 ns/op             376 B/op              8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=64-16           124934              8804 ns/op             376 B/op              8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=256-16          135568              8607 ns/op             368 B/op              7 allocs/op
BenchmarkHashAddress/keccak-16                                   2738773               427.8 ns/op             0 B/op              0 allocs/op
BenchmarkHashAddress/pedersen_hash-16                             176554              6162 ns/op             160 B/op              3 allocs/op
BenchmarkHashStorageSlot/keccak-16                               2821916               429.1 ns/op             0 B/op              0 allocs/op
BenchmarkHashStorageSlot/pedersen_hash-16                         142660              7916 ns/op             416 B/op             11 allocs/op
```

Go (without _Intel ADX_ CPU instructions):
```
BenchmarkFpInverse/go-ipa-16             1030958              1168 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-16             870020              1391 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-16                40603137                24.69 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-16               10218465               113.1 ns/op             0 B/op          0 allocs/op
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
BenchmarkFpInverse/go-ipa-8               807909              1406 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-8              692365              1557 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-8                 65012094                17.57 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-8                 9384998               126.2 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=32-8                2491617               470.7 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=64-8                2692922               442.9 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=128-8               2892056               411.3 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=256-8               1511790               790.1 ns/op             0 B/op          0 allocs/op
BenchmarkKeccak/size=512-8                699243              1562 ns/op               0 B/op          0 allocs/op
BenchmarkKeccak/size=1024-8               361258              3070 ns/op               0 B/op          0 allocs/op
BenchmarkPedersenCommit/non_zero_entries=1-8               28230             42195 ns/op             42197 ns/value        33937 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=2-8               24967             46683 ns/op             23343 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=4-8               19010             62857 ns/op             15715 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=8-8               12634             92911 ns/op             11614 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=16-8               7234            155541 ns/op              9722 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=32-8               3674            281022 ns/op              8782 ns/value        33936 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=64-8               1977            532915 ns/op              8327 ns/value        33929 B/op         19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=128-8              1077           1036918 ns/op              8102 ns/value        33938 B/op         20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=256-8               565           2033419 ns/op              7943 ns/value        33948 B/op         21 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=1-8             109887             11084 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=4-8              98496             11511 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=16-8            110120             11647 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=64-8             94174             11570 ns/op             376 B/op          8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=256-8           106506             10452 ns/op             368 B/op          7 allocs/op
BenchmarkHashAddress/keccak-8                                    2528505               469.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashAddress/pedersen_hash-8                              165478              7131 ns/op             160 B/op          3 allocs/op
BenchmarkHashStorageSlot/keccak-8                                2419062               512.0 ns/op             0 B/op          0 allocs/op
BenchmarkHashStorageSlot/pedersen_hash-8                          121276              9981 ns/op             416 B/op         11 allocs/op
```

Go (without _Intel ADX_ CPU instructions):
```
BenchmarkFpInverse/go-ipa-8               825962              1424 ns/op               0 B/op          0 allocs/op
BenchmarkFpInverse/go-blst-8              648957              1658 ns/op              32 B/op          1 allocs/op
BenchmarkFpMul/go-ipa-8                 39319251                33.39 ns/op            0 B/op          0 allocs/op
BenchmarkFpMul/go-blst-8                 7675392               134.8 ns/op             0 B/op          0 allocs/op
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

### Nvidia Jetson Orin 12-Core Tegra @ 2.2GHz

Go benchmarks:
```
goarch: arm64
pkg: github.com/jsign/verkle-vs-patricia
BenchmarkFpInverse/go-ipa-12            	  442096	      2819 ns/op	       0 B/op	       0 allocs/op
BenchmarkFpInverse/go-blst-12           	  475665	      2471 ns/op	      32 B/op	       1 allocs/op
BenchmarkFpMul/go-ipa-12                	30143986	        38.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkFpMul/go-blst-12               	 4883612	       216.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=32-12              	 2010416	       600.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=64-12              	 1966846	       582.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=128-12             	 2104431	       538.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=256-12             	 1000000	      1006 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=512-12             	  622038	      1924 ns/op	       0 B/op	       0 allocs/op
BenchmarkKeccak/size=1024-12            	  314476	      3772 ns/op	       0 B/op	       0 allocs/op
BenchmarkPedersenCommit/non_zero_entries=1-12         	   13968	     81729 ns/op	     81729 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=2-12         	   10000	    102697 ns/op	     51349 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=4-12         	    8658	    140834 ns/op	     35209 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=8-12         	    5784	    210640 ns/op	     26330 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=16-12        	    3428	    379230 ns/op	     23702 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=32-12        	    1759	    653771 ns/op	     20431 ns/value	   33937 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=64-12        	     936	   1239902 ns/op	     19374 ns/value	   33931 B/op	      19 allocs/op
BenchmarkPedersenCommit/non_zero_entries=128-12       	     511	   2340999 ns/op	     18289 ns/value	   33943 B/op	      20 allocs/op
BenchmarkPedersenCommit/non_zero_entries=256-12       	     252	   4689499 ns/op	     18318 ns/value	   33958 B/op	      21 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=1-12         	   50266	     23384 ns/op	     376 B/op	       8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=4-12         	   51420	     23655 ns/op	     376 B/op	       8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=16-12        	   51490	     24926 ns/op	     376 B/op	       8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=64-12        	   50168	     23342 ns/op	     376 B/op	       8 allocs/op
BenchmarkPedersenUpdateCommitment/vec_num_entries=256-12       	   51333	     23607 ns/op	     368 B/op	       7 allocs/op
BenchmarkHashAddress/keccak-12                                 	 1864634	       603.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashAddress/pedersen_hash-12                          	   87103	     13799 ns/op	     160 B/op	       3 allocs/op
BenchmarkHashStorageSlot/keccak-12                             	 1897682	       597.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashStorageSlot/pedersen_hash-12                      	   68104	     17827 ns/op	     416 B/op	      11 allocs/op
PASS
ok  	github.com/jsign/verkle-vs-patricia	53.355s
```

C BLST benchmark:
```
Inv(x) takes 2360 ns/op
Mul(a,b) takes 126 ns/op
```

Rust `arkworks-rs` benchmarks:
```
inverse                 time:   [2.9362 µs 2.9450 µs 2.9545 µs]
mul                     time:   [35.116 ns 35.119 ns 35.123 ns]
```
