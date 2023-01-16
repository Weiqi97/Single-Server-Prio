package utilities

import "math/big"

// GetStandardBasis returns a vector of desired length where it is zero everywhere and one at index.
func GetStandardBasis(length int, index int) (e []*big.Int) {
	// Initialize the result to the desired length.
	e = make([]*big.Int, length)

	// Put zero everywhere.
	for i := 0; i < length; i++ {
		e[i] = big.NewInt(0)
	}

	// Flip the bit to one at the desired place.
	e[index] = big.NewInt(1)

	return
}
