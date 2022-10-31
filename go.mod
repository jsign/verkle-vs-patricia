module github.com/jsign/keccak-vs-pedersen

go 1.19

require (
	github.com/crate-crypto/go-ipa v0.0.0-20220916134416-c5abbdbdf644
	github.com/ethereum/go-ethereum v1.10.25
	github.com/gballet/go-verkle v0.0.0-20220923150140-6c08cd337774
	github.com/holiman/uint256 v1.2.0
	github.com/stretchr/testify v1.8.0
	github.com/supranational/blst v0.3.10
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// References https://github.com/gballet/go-ethereum/tree/beverly-hills-head, to have GetTreeKey func.
replace github.com/ethereum/go-ethereum => github.com/gballet/go-ethereum v1.10.24-0.20221005144906-f0e4ab07b54a
