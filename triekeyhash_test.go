package main

import (
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	verkleutils "github.com/ethereum/go-ethereum/trie/utils"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func BenchmarkHashAddress(b *testing.B) {
	addrBytes := make([]byte, common.AddressLength)
	_, err := rand.Read(addrBytes[:])
	require.NoError(b, err)

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
	addrBytes := make([]byte, common.AddressLength)
	_, err := rand.Read(addrBytes[:])
	require.NoError(b, err)

	storageSlot := uint256.NewInt(rand.Uint64())
	storageSlotAddr := crypto.Keccak256Hash(storageSlot.Bytes()).Bytes()

	// Warmup go-ipa generating precomputes before main bench.
	sink = verkleutils.GetTreeKeyStorageSlot(addrBytes, storageSlot)

	b.Run("keccak", benchKeccak(storageSlotAddr))
	b.Run("pedersen hash", func(b *testing.B) {
		var res []byte
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			res = verkleutils.GetTreeKeyStorageSlot(addrBytes, storageSlot)
		}
		sink = res
	})
}
