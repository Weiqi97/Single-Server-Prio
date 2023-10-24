package test

import (
	"github.com/Weiqi97/Single-Server-Prio/utilities"
	"gotest.tools/v3/assert"
	"math/big"
	"testing"
)

// Set the modular for global usage in this testing file.
var mod = big.NewInt(11)

func TestNewZp(t *testing.T) {
	// Create a new testing Zp.
	ele := big.NewInt(6)
	x := utilities.NewZp(ele, mod)
	assert.Assert(t, x.GetEle().Cmp(big.NewInt(6)) == 0)
}

func TestRand(t *testing.T) {
	// Create a new testing Zp.
	x := utilities.RandZp(mod)
	assert.Assert(t, x.GetEle().Cmp(mod) == -1)
}

func TestAdd(t *testing.T) {
	ele1 := big.NewInt(6)
	ele2 := big.NewInt(7)

	x1 := utilities.NewZp(ele1, mod)
	x2 := utilities.NewZp(ele2, mod)

	result := utilities.Add(x1, x2)

	assert.Assert(t, result.GetEle().Cmp(big.NewInt(2)) == 0)
}

func TestSub(t *testing.T) {
	ele1 := big.NewInt(6)
	ele2 := big.NewInt(7)

	x1 := utilities.NewZp(ele1, mod)
	x2 := utilities.NewZp(ele2, mod)

	result := utilities.Sub(x1, x2)

	assert.Assert(t, result.GetEle().Cmp(big.NewInt(10)) == 0)
}

func TestMul(t *testing.T) {
	ele1 := big.NewInt(6)
	ele2 := big.NewInt(7)

	x1 := utilities.NewZp(ele1, mod)
	x2 := utilities.NewZp(ele2, mod)

	result := utilities.Mul(x1, x2)

	assert.Assert(t, result.GetEle().Cmp(big.NewInt(9)) == 0)
}

func TestInv(t *testing.T) {
	ele := big.NewInt(6)
	x := utilities.NewZp(ele, mod)

	result := utilities.Inv(x)

	assert.Assert(t, result.GetEle().Cmp(big.NewInt(2)) == 0)
}

func TestDiv(t *testing.T) {
	ele1 := big.NewInt(6)
	ele2 := big.NewInt(7)

	x1 := utilities.NewZp(ele1, mod)
	x2 := utilities.NewZp(ele2, mod)

	result := utilities.Div(x1, x2)

	assert.Assert(t, result.GetEle().Cmp(big.NewInt(4)) == 0)
}
