package main

import (
	"testing"

	"github.com/crate-crypto/go-ipa/bandersnatch/fp"
	"github.com/stretchr/testify/require"
	blst "github.com/supranational/blst/bindings/go"
)

func BenchmarkFpInverse(b *testing.B) {
	b.Run("go-ipa", func(b *testing.B) {
		var x fp.Element
		x.SetRandom()

		b.ResetTimer()
		var resIPA fp.Element
		for i := 0; i < b.N; i++ {
			resIPA.Inverse(&x)
		}
	})
	b.Run("go-blst", func(b *testing.B) {
		var xIPA fp.Element
		xIPA.SetRandom()
		var xBLST blst.Scalar
		beBytes := xIPA.Bytes()
		xBLST.FromBEndian(beBytes[:])

		var resBLST *blst.Scalar
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			resBLST = xBLST.Inverse()
		}
		_ = resBLST
	})
}

func TestVerifyGoIPAvsBLSTInverse(t *testing.T) {
	t.Parallel()

	for i := 0; i < 500_000; i++ {
		// Generate random field element
		var xIPA fp.Element
		xIPA.SetRandom()

		// Get the inverse with go-ipa
		var invIPA fp.Element
		invIPA.Inverse(&xIPA)

		// Transform the same x to blst, and calculate the multiplication.
		var xBLST blst.Scalar
		beBytes := xIPA.Bytes()
		xBLST.FromBEndian(beBytes[:])
		invBLST := xBLST.Inverse()

		// Both inverses should match.
		invIPABytes := invIPA.Bytes()
		require.Equal(t, invIPABytes[:], invBLST.Serialize())
	}
}

func BenchmarkFpMul(b *testing.B) {
	b.Run("go-ipa", func(b *testing.B) {
		var x fp.Element
		x.SetRandom()
		var y fp.Element
		y.SetRandom()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x.Mul(&x, &y)
		}
	})
	b.Run("go-blst", func(b *testing.B) {
		var xIPA fp.Element
		xIPA.SetRandom()
		var yIPA fp.Element
		yIPA.SetRandom()
		var xBLST blst.Scalar
		beXBytes := xIPA.Bytes()
		xBLST.FromBEndian(beXBytes[:])
		var yBLST blst.Scalar
		beYBytes := yIPA.Bytes()
		yBLST.FromBEndian(beYBytes[:])

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			xBLST.MulAssign(&yBLST)
		}
	})
}

func TestVerifyGoIPAvsBLSTMultiplication(t *testing.T) {
	t.Parallel()

	for i := 0; i < 500_000; i++ {
		// Generate random field element
		var xIPA fp.Element
		xIPA.SetRandom()
		var yIPA fp.Element
		yIPA.SetRandom()

		// Multiply x and y with go-ipa
		var mulIPA fp.Element
		mulIPA.Mul(&xIPA, &yIPA)

		// Transform x and y to blst, and calculate the multiplication.
		var xBLST blst.Scalar
		beXBytes := xIPA.Bytes()
		xBLST.FromBEndian(beXBytes[:])
		var yBLST blst.Scalar
		beYBytes := yIPA.Bytes()
		yBLST.FromBEndian(beYBytes[:])
		mulBLST, _ := xBLST.MulAssign(&yBLST)

		mulIPABytes := mulIPA.Bytes()
		require.Equal(t, mulIPABytes[:], mulBLST.Serialize())
	}
}
