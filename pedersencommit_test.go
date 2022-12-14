package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gballet/go-verkle"
	"github.com/stretchr/testify/require"
)

func BenchmarkPedersenCommit(b *testing.B) {
	// Create a random stem.
	var stem [32]byte
	_, err := rand.Read(stem[:])
	require.NoError(b, err)

	// Create a 256-random 32 byte values
	var values [verkle.NodeWidth][]byte
	for i := range values {
		values[i] = make([]byte, 32)
		_, err = rand.Read(values[i][:])
		require.NoError(b, err)
	}

	// Create [1, 2, 4, ...] subtests.
	for numValues := 1; numValues <= 256; numValues *= 2 {
		// Create a vector of 256 zero elements.
		var leafValues [verkle.NodeWidth][]byte
		// Fill the first numValues with the ones we pre-generated.
		// leafValues is a 256-element vector with numValues non-zero elements.
		copy(leafValues[:], values[:numValues])

		// Warm-up the library so it computes the lagrange precomputed table out of the benchmark main loop.
		_ = verkle.NewLeafNode(stem[:31], leafValues[:])

		b.Run(fmt.Sprintf("non zero entries=%d", numValues), func(b *testing.B) {
			now := time.Now()

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Create a leaf node with leafValues
				ln := verkle.NewLeafNode(stem[:31], leafValues[:])
				// Trigger the polynomial commitment.
				ln.Commit()
				// Trigger the EC point -> Fr transformation.
				ln.Hash()
			}
			b.ReportMetric(float64(time.Since(now).Nanoseconds()/int64(numValues))/float64(b.N), "ns/value")
		})
	}
}

func BenchmarkPedersenUpdateCommitment(b *testing.B) {
	// Createa a random stem.
	var stem [32]byte
	_, err := rand.Read(stem[:])
	require.NoError(b, err)

	// Create a 256-random 32 byte values
	var values [verkle.NodeWidth][]byte
	for i := range values {
		values[i] = make([]byte, 32)
		_, err = rand.Read(values[i][:])
		require.NoError(b, err)
	}

	// Create a new random 32-byte value that will be inserted in an *existing* leaf.
	var newValue [32]byte
	_, err = rand.Read(newValue[:])
	require.NoError(b, err)

	// Create [1, 4, 16, ...] subtests.
	for numValues := 1; numValues <= 256; numValues *= 4 {
		// Create a vector of 256 zero elements.
		var leafValues [verkle.NodeWidth][]byte
		// Fill the first numValues with the ones we pre-generated.
		// leafValues is a 256-element vector with numValues non-zero elements.
		copy(leafValues[:], values[:numValues])

		// Create a leaf node, and compute the commitment and hash. We'll benchmark the cost of
		// updating a single entry of this vector.
		ln := verkle.NewLeafNode(stem[:31], leafValues[:])
		ln.Commit()
		ln.Hash()

		b.Run(fmt.Sprintf("vec_num_entries=%d", numValues), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Insert the new value into the existing leaf node. (Note we use the same stem as expected)
				ln.Insert(stem[:], newValue[:], nil)

				// Note that we *don't* ln.Commit(), since we don't need to do a fresh computation.

				// Trigger the EC point -> Fr transformation.
				ln.Hash()
			}
		})
	}
}
