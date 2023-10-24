package test

import (
	"fmt"
	"github.com/Weiqi97/Single-Server-Prio/flpcp"
	"github.com/Weiqi97/Single-Server-Prio/ipfe"
	u "github.com/Weiqi97/Single-Server-Prio/utilities"
	"github.com/fentec-project/gofe/data"
	"log"
	"math/big"
	"testing"
	"time"
)

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func BenchmarkProofGen(b *testing.B) {
	for _, n := range []int{9} {
		// Create a large modulo to start with.
		scheme := ipfe.InitDDHScheme(3, 1024, big.NewInt(100))
		mod := scheme.GetParam().P

		// Create the testing vector.
		x := make([]u.Zp, intPow(2, n))
		for i := 0; i < intPow(2, n); i++ {
			x[i] = u.NewZp(big.NewInt(1), mod)
		}

		proof := flpcp.InitFLPCP(intPow(2, n), scheme.GetParam())

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			proof.Prover(x)
		})
	}
}

func BenchmarkProofVal(b *testing.B) {
	// Set testing values.

	yValues := make([]*big.Int, 512)
	xValues := make([]*big.Int, 512)

	for i := 0; i < 512; i++ {
		xValues[i] = big.NewInt(10000)
		yValues[i] = big.NewInt(50000)
	}

	y := data.NewVector(yValues)
	x := data.NewVector(xValues)

	// Initialize the scheme and generate keys.
	scheme := ipfe.InitDDHScheme(512, 1024, big.NewInt(1009800097800))
	msk, mpk := scheme.KeyGen()

	// Run key derivation on y.

	// Run encryption on x.
	c := scheme.Enc(mpk, x)

	n := 512
	// Run decryption.
	b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
		start := time.Now()

		yKey := scheme.KeyDer(msk, y)
		_ = scheme.Dec(c, y, yKey)
		elapsed := time.Since(start)
		log.Printf("Execution time: %s", elapsed)
	})
}
