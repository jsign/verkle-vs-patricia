package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

// sink defined globally only to avoid Go compiler to consider the benchmark code 'dead code'.
var sink []byte

func BenchmarkKeccak(b *testing.B) {
	// Generate a random byte blob to slice in sub-benchmarks.
	var baseRandomValue [1024]byte
	_, err := rand.Read(baseRandomValue[:])
	require.NoError(b, err)

	for valueSize := 32; valueSize <= 1024; valueSize *= 2 {
		subTestName := fmt.Sprintf("size=%d", valueSize)
		b.Run(subTestName, benchKeccak(baseRandomValue[:valueSize]))
	}
}

func benchKeccak(value []byte) func(b *testing.B) {
	return func(b *testing.B) {
		hasher := crypto.NewKeccakState()
		res := make([]byte, 32)

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Write the received blob.
			hasher.Write(value)
			hasher.Read(res)
			// Reset the hasher, so we can reuse it in the next loop. Again, to avoid noise from creating objects.
			hasher.Reset()
		}
		sink = res // Dummy assignation to avoid the Go compiler consider deadcode.
	}
}

func TestKeccakGethVsStd(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1_000_000; i++ {
		var value [512]byte
		_, err := rand.Read(value[:])
		require.NoError(t, err)

		// Use go-ethereum sha3 gethHasher, which exploits internal Read(...) method.
		gethHasher := crypto.NewKeccakState()
		gethHasher.Write(value[:])
		gethRes := make([]byte, 32)
		gethHasher.Read(gethRes)

		// Use base x/crypto/sha3.
		stdHasher := sha3.NewLegacyKeccak256()
		stdHasher.Write(value[:])
		stdRes := stdHasher.Sum(nil)

		// They should match.
		require.Equal(t, gethRes, stdRes)
	}
}
