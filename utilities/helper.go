package utilities

import (
	"math/big"
)

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

// IntPow raises an integer a to power b and return the result in int64 type.
func IntPow(a, b int) int64 {
	if b == 0 {
		return 1
	}
	result := a
	for i := 2; i <= b; i++ {
		result *= a
	}
	return int64(result)
}

// LinSolver Given a square matrix A and result b, finds the solution of Ax = b to x.
func LinSolver(b []Zp) (x []Zp) {
	// Get the modulo.
	mod := b[0].GetMod()

	// We know that A is a square matrix whose side length is same as b.
	size := len(b) // Size here in most cases is the same as number of G gate.

	// Create the composed matrix Ab.
	matrix := make([][]Zp, size)

	// Fill in the matrix, suppose the degree goes from high to low.
	for i := 0; i < size; i++ {
		matrix[i] = make([]Zp, size+1)
		for j := 0; j < size; j++ {
			temp := IntPow(i+1, size-j)
			matrix[i][j] = NewZp(big.NewInt(temp), mod)
		}
		matrix[i][size] = b[i]
	}

	// Bottom left half to all zeros.
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if i == j && matrix[i][j].GetEle().Cmp(big.NewInt(1)) != 0 {
				multiplier := Inv(matrix[i][i])
				for k := i; k <= size; k++ {
					matrix[j][k] = Mul(matrix[j][k], multiplier)
				}
			}

			if i != j {
				multiplier := matrix[j][i]
				for k := i; k <= size; k++ {
					matrix[j][k] = Sub(matrix[j][k], Mul(matrix[i][k], multiplier))
				}
			}
		}
	}

	// Top right half to all zeros.
	for i := size - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			multiplier := matrix[j][i]
			for k := i; k <= size; k++ {
				matrix[j][k] = Sub(matrix[j][k], Mul(matrix[i][k], multiplier))
			}
		}
	}

	// Get the intended result.
	for i := 0; i < size; i++ {
		x = append(x, matrix[i][size])
	}

	return
}
