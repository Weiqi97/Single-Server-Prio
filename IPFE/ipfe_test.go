package IPFE

import (
	"github.com/fentec-project/gofe/data"
	"gotest.tools/v3/assert"
	"math/big"
	"testing"
)

// TestInitDDHScheme is a dummy test to make sure the function runs.
func TestInitDDHScheme(t *testing.T) {
	_ = InitDDHScheme(3, 1024, big.NewInt(10))
}

func TestIPFE_KeyGen(t *testing.T) {
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	_, _ = scheme.KeyGen()
}

func TestIPFE_KeyDer(t *testing.T) {
	yValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	y := data.NewVector(yValues)
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	msk, _ := scheme.KeyGen()
	_ = scheme.KeyDer(msk, y)

}

func TestIPFE_Enc(t *testing.T) {
	xValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	x := data.NewVector(xValues)
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	_, mpk := scheme.KeyGen()
	_ = scheme.Enc(mpk, x)
}

func TestIPFE_Dec(t *testing.T) {
	// Set testing values.
	yValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	y := data.NewVector(yValues)
	xValues := []*big.Int{big.NewInt(2), big.NewInt(4), big.NewInt(6)}
	x := data.NewVector(xValues)

	// Initialize the scheme and generate keys.
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	msk, mpk := scheme.KeyGen()

	// Run key derivation on y.
	yKey := scheme.KeyDer(msk, y)

	// Run encryption on x.
	c := scheme.Enc(mpk, x)

	// Run decryption.
	xy := scheme.Dec(c, y, yKey)

	// Check equality.
	assert.Assert(t, xy.String() == "28")
}

func TestIPFE_AddCiphertext(t *testing.T) {

}

func TestIPFE_RecoverCiphertext(t *testing.T) {

}
