package utilities

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestGetStandardBasis(t *testing.T) {
	// Generate the basis [0, 1, 0].
	e := GetStandardBasis(3, 1)

	// Check for correctness.
	assert.Assert(t, e[0].String() == "0")
	assert.Assert(t, e[1].String() == "1")
	assert.Assert(t, e[2].String() == "0")
}
