package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	verkleutils "github.com/ethereum/go-ethereum/trie/utils"
	"github.com/holiman/uint256"
)

func BenchmarkHashAddress(b *testing.B) {
	addrBytes := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Bytes()

	b.Run("keccak", benchKeccak(addrBytes))

	b.Run("pedersen hash", func(b *testing.B) {
		var res []byte
		// Warmup go-ipa generating precomputes before main bench.
		res = verkleutils.GetTreeKeyBalance(addrBytes)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res = verkleutils.GetTreeKeyBalance(addrBytes)
		}
		sink = res
	})
}

func BenchmarkHashStorageSlot(b *testing.B) {
	addrBytes := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Bytes()
	storageSlot := uint256.NewInt(42)
	storageSlotAddr := crypto.Keccak256Hash(storageSlot.Bytes()).Bytes()

	b.Run("keccak", benchKeccak(storageSlotAddr))
	b.Run("pedersen hash", func(b *testing.B) {
		var res []byte
		// Warmup go-ipa generating precomputes before main bench.
		res = verkleutils.GetTreeKeyStorageSlot(addrBytes, storageSlot)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res = verkleutils.GetTreeKeyStorageSlot(addrBytes, storageSlot)
		}
		sink = res
	})
}
