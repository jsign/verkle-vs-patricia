module github.com/jsign/verkle-vs-patricia

go 1.19

require (
	github.com/crate-crypto/go-ipa v0.0.0-20221107233728-e868de2bc50d
	github.com/ethereum/go-ethereum v1.10.25
	github.com/gballet/go-verkle v0.0.0-20220923150140-6c08cd337774
	github.com/holiman/uint256 v1.2.0
	github.com/jsign/go-ethereum v1.10.18-0.20221106222644-245b3f6459ad
	github.com/stretchr/testify v1.8.0
	github.com/supranational/blst v0.3.10
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a
	golang.org/x/exp v0.0.0-20220426173459-3bcf042a4bf5
)

require (
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/VictoriaMetrics/fastcache v1.6.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/tsdb v0.7.1 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	golang.org/x/sys v0.0.0-20221013171732-95e765b1cc43 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// References https://github.com/gballet/go-ethereum/tree/beverly-hills-head, to have GetTreeKey func.
replace github.com/ethereum/go-ethereum => github.com/gballet/go-ethereum v1.10.24-0.20221005144906-f0e4ab07b54a
