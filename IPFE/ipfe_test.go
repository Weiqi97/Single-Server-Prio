package IPFE

import (
	"github.com/fentec-project/gofe/data"
	"gotest.tools/v3/assert"
	"math/big"
	"reflect"
	"testing"
)

// TestInitDDHScheme is a dummy test to make sure the function runs.
func TestInitDDHScheme(t *testing.T) {
	_ = InitDDHScheme(3, 1024, big.NewInt(10))
}

// TestIPFE_KeyGen is a dummy test to make sure the function runs.
func TestIPFE_KeyGen(t *testing.T) {
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	_, _ = scheme.KeyGen()
}

// TestIPFE_KeyDer is a dummy test to make sure the correct type is returned.
func TestIPFE_KeyDer(t *testing.T) {
	// Set testing values.
	yValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	y := data.NewVector(yValues)

	// Initialize the scheme and generate keys.
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	msk, _ := scheme.KeyGen()

	// Run key derivation on y.
	yKey := scheme.KeyDer(msk, y)

	// Dummy test on type.
	assert.Assert(t, reflect.TypeOf(yKey) == reflect.TypeOf(big.NewInt(0)))
}

// TestIPFE_Enc is a dummy test to make sure the correct length of ciphertext is returned.
func TestIPFE_Enc(t *testing.T) {
	// Set testing values.
	xValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	x := data.NewVector(xValues)

	// Initialize the scheme and generate keys.
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	_, mpk := scheme.KeyGen()

	// Run encryption on x.
	c := scheme.Enc(mpk, x)

	// Dummy test on length.
	assert.Assert(t, len(c) == 4)
}

// TestIPFE_Dec checks whether the decryption result correctly returns <x, y>.
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
	assert.Assert(t, xy.Cmp(big.NewInt(28)) == 0)
}

// TestIPFE_AddCiphertext is a dummy test that checks whether ciphertext can be added.
func TestIPFE_AddCiphertext(t *testing.T) {
	// Set testing values.
	xValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	x := data.NewVector(xValues)

	// Initialize the scheme and generate keys.
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	_, mpk := scheme.KeyGen()

	// Run encryption on x.
	c := scheme.Enc(mpk, x)

	// Add x with itself.
	c = scheme.AddCiphertext(c, c)

	// Dummy test on length.
	assert.Assert(t, len(c) == 4)

}

// TestIPFE_RecoverCiphertext checks the correctness of adding ciphertext.
func TestIPFE_RecoverCiphertext(t *testing.T) {
	// Set testing values.
	xValues := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	x := data.NewVector(xValues)

	// Initialize the scheme and generate keys.
	scheme := InitDDHScheme(3, 1024, big.NewInt(10))
	msk, mpk := scheme.KeyGen()

	// Run encryption on x.
	c := scheme.Enc(mpk, x)

	// Add x with itself.
	c = scheme.AddCiphertext(c, c)

	// Recover the added result.
	xx := scheme.RecoverCiphertext(msk, c)

	// Check the length of the recovered string.
	assert.Assert(t, len(xx) == 3)

	// Check the correctness of the recovered string.
	assert.Assert(t, xx[0].String() == "2")
	assert.Assert(t, xx[1].String() == "4")
	assert.Assert(t, xx[2].String() == "6")
}
