package utilities

import "math/big"

func GetStandardBasis(length int, index int) (e []*big.Int) {
	e = make([]*big.Int, length)
	for i := 0; i < length; i++ {
		e[i] = big.NewInt(0)
	}
	e[index] = big.NewInt(1)
	return
}
