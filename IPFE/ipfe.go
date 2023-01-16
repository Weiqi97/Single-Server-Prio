package IPFE

import (
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

func (scheme IPFE) AddCiphertext(c1 data.Vector, c2 data.Vector) (c data.Vector) {
	for i := range c1 {
		c[i] = new(big.Int).Mod(new(big.Int).Mul(c1[i], c2[i]), scheme.instance.Params.P)
	}
	return
}

func (scheme IPFE) RecoverCiphertext(msk data.Vector, c data.Vector) (x []*big.Int) {
	for i := range c {
		y := utilities.GetStandardBasis(len(c), i)
		yKey := scheme.KeyDer(msk, y)
		x[i] = scheme.Dec(c, y, yKey)
	}
	return
}
