bench-go:
	go test . -run=none -bench=. -benchmem
.PHONY: bench-go

bench-go-noadx:
	go test -tags noadx . -run=none -bench=BenchmarkFp -benchmem

bench-blst:
	cd blst-bench && ([ -f libblst.a ] || (git clone https://github.com/supranational/blst.git && ./blst/build.sh ; rm -rf blst))
	cd blst-bench && cc -I . main.c libblst.a -o main && ./main && rm main
.PHONY: bench-blst

bench-arkworks:
	cd bench-arkworks && cargo install -q cargo-criterion && cargo criterion
.PHONY: bench-arkworks