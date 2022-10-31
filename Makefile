bench-go:
	go test . -run=none -bench=.

bench-blst:
	cd blst-bench && ([ -f libblst.a ] || (git clone https://github.com/supranational/blst.git && ./blst/build.sh ; rm -rf blst))
	cd blst-bench && cc -I . main.c libblst.a -o main && ./main && rm main
    