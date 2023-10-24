package test

import (
	"github.com/Weiqi97/Single-Server-Prio/utilities"
	"gotest.tools/v3/assert"
	"math/big"
	"testing"
)

func TestGetStandardBasis(t *testing.T) {
	// Generate the basis [0, 1, 0].
	e := utilities.GetStandardBasis(3, 1)

	// Check for correctness.
	assert.Assert(t, e[0].String() == "0")
	assert.Assert(t, e[1].String() == "1")
	assert.Assert(t, e[2].String() == "0")
}

func TestIntPow(t *testing.T) {
	x := utilities.IntPow(2, 10)
	assert.Assert(t, x == 1024)
}

func TestLinSolver(t *testing.T) {
	// Set the testing values.
	bigMod := big.NewInt(4001)
	ele1 := big.NewInt(50)
	ele2 := big.NewInt(352)
	ele3 := big.NewInt(1374)
	ele4 := big.NewInt(3848)
	x1 := utilities.NewZp(ele1, bigMod)
	x2 := utilities.NewZp(ele2, bigMod)
	x3 := utilities.NewZp(ele3, bigMod)
	x4 := utilities.NewZp(ele4, bigMod)

	// Run the solver.
	x := utilities.LinSolver([]utilities.Zp{x1, x2, x3, x4})

	// Check for correctness.
	assert.Assert(t, x[0].GetEle().String() == "11")
	assert.Assert(t, x[1].GetEle().String() == "12")
	assert.Assert(t, x[2].GetEle().String() == "13")
	assert.Assert(t, x[3].GetEle().String() == "14")
}
