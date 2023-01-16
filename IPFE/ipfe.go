package IPFE

import (
	"fmt"
	"github.com/Weiqi97/Single-Server-Prio/utilities"
	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/simple"
	"math/big"
)

// IPFE is a wrapper for the simple DDH inner product scheme instance from the GoFE library.
type IPFE struct {
	instance *simple.DDH
}

// InitDDHScheme returns a new IPFE scheme.
func InitDDHScheme(vectorLength int, modulusLength int, bound *big.Int) (scheme IPFE) {
	scheme.instance, _ = simple.NewDDHPrecomp(vectorLength, modulusLength, bound)
	return
}

// KeyGen generates the master secret and public keys.
func (scheme IPFE) KeyGen() (msk data.Vector, mpk data.Vector) {
	msk, mpk, _ = scheme.instance.GenerateMasterKeys()
	return
}

// KeyDer derives the function key on a vector y.
func (scheme IPFE) KeyDer(msk data.Vector, y data.Vector) (yKey *big.Int) {
	yKey, _ = scheme.instance.DeriveKey(msk, y)
	return
}

// Enc encrypts the input vector x.
func (scheme IPFE) Enc(mpk data.Vector, x data.Vector) (c data.Vector) {
	c, _ = scheme.instance.Encrypt(x, mpk)
	return
}

// Dec computes the desired inner product <x, y>.
func (scheme IPFE) Dec(c data.Vector, y data.Vector, yKey *big.Int) (xy *big.Int) {
	xy, _ = scheme.instance.Decrypt(c, yKey, y)
	return
}

// AddTwoCiphertexts homomorphically adds two ciphertext.
func (scheme IPFE) AddTwoCiphertexts(c1 data.Vector, c2 data.Vector) (c data.Vector) {
	// Initialize the result c to the desired length.
	c = make([]*big.Int, len(c1))

	// Multiply c1 and c2, such that the exponent is added.
	for i := range c1 {
		c[i] = new(big.Int).Mod(new(big.Int).Mul(c1[i], c2[i]), scheme.instance.Params.P)
	}

	return
}

// AddCiphertexts homomorphically adds multiple ciphertexts at once.
func (scheme IPFE) AddCiphertexts(ciphertexts []data.Vector) (c data.Vector, e error) {
	// Check whether the number of input ciphertexts is more than 2.
	if len(ciphertexts) < 2 {
		return nil, fmt.Errorf("not enough arguments to call this function")
	}

	// Add the first and second ciphertext.
	c = scheme.AddTwoCiphertexts(ciphertexts[0], ciphertexts[1])

	// Keep adding the rest ciphertexts.
	for i := 2; i < len(ciphertexts); i++ {
		c = scheme.AddTwoCiphertexts(c, ciphertexts[i])
	}

	return c, nil
}

// RecoverCiphertext decrypts the encrypted vector.
func (scheme IPFE) RecoverCiphertext(msk data.Vector, c data.Vector) (x []*big.Int) {
	// Initialize the result x to the desired length.
	x = make([]*big.Int, len(c)-1)

	// Use the standard basis to retrieve values one by one.
	for i := 0; i < len(c)-1; i++ {
		y := utilities.GetStandardBasis(len(c)-1, i)
		yKey := scheme.KeyDer(msk, y)
		x[i] = scheme.Dec(c, y, yKey)
	}

	return
}
