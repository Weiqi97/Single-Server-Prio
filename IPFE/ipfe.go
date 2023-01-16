package IPFE

import (
	"fmt"
	"github.com/Weiqi97/Single-Server-Prio/utilities"
	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/simple"
	"math/big"
)

type IPFE struct {
	instance *simple.DDH
}

// InitDDHScheme This function
func InitDDHScheme(vectorLength int, modulusLength int, bound *big.Int) (scheme IPFE) {

	scheme.instance, _ = simple.NewDDHPrecomp(vectorLength, modulusLength, bound)

	return
}
func (scheme IPFE) KeyGen() (msk data.Vector, mpk data.Vector) {
	msk, mpk, _ = scheme.instance.GenerateMasterKeys()
	return
}

func (scheme IPFE) KeyDer(msk data.Vector, y data.Vector) (yKey *big.Int) {
	yKey, _ = scheme.instance.DeriveKey(msk, y)
	return
}

func (scheme IPFE) Enc(mpk data.Vector, x data.Vector) (c data.Vector) {
	c, _ = scheme.instance.Encrypt(x, mpk)
	return
}

func (scheme IPFE) Dec(c data.Vector, y data.Vector, yKey *big.Int) (xy *big.Int) {
	xy, _ = scheme.instance.Decrypt(c, yKey, y)
	return
}

func (scheme IPFE) AddTwoCiphertexts(c1 data.Vector, c2 data.Vector) (c data.Vector) {
	c = make([]*big.Int, len(c1))
	for i := range c1 {
		c[i] = new(big.Int).Mod(new(big.Int).Mul(c1[i], c2[i]), scheme.instance.Params.P)
	}
	return
}

func (scheme IPFE) AddCiphertexts(ciphertexts []data.Vector) (c data.Vector, e error) {
	if len(ciphertexts) < 2 {
		return nil, fmt.Errorf("not enough arguments to call this function")
	}
	c = scheme.AddTwoCiphertexts(ciphertexts[0], ciphertexts[1])
	for i := 2; i < len(ciphertexts); i++ {
		c = scheme.AddTwoCiphertexts(c, ciphertexts[i])
	}
	return c, nil
}

func (scheme IPFE) RecoverCiphertext(msk data.Vector, c data.Vector) (x []*big.Int) {
	x = make([]*big.Int, len(c)-1)

	for i := 0; i < len(c)-1; i++ {
		y := utilities.GetStandardBasis(len(c)-1, i)
		yKey := scheme.KeyDer(msk, y)
		x[i] = scheme.Dec(c, y, yKey)
	}
	return
}
