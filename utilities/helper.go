package utilities

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
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

// SumVector sums all elements in a vector.
func SumVector(length int, x []*big.Int) (sum *big.Int) {
	sum = x[0]
	for i := 1; i < length; i++ {
		sum.Add(sum, x[i])
	}
	return
}

// LinSolver finds the solution of Ax = b to x.
func LinSolver(numGate int, y []float64) (x []float64) {
	matrixA := mat.NewDense(numGate, numGate+1, nil)

	// Fill in the matrix.
	for i := 0; i < numGate; i++ {
		for j := 0; j < numGate; j++ {
			matrixA.Set(i, j, math.Pow(float64(i), float64(j)))
		}
	}

	// Add 1 for the constant.
	for i := 0; i < numGate; i++ {
		matrixA.Set(i, numGate, 1.0)
	}

	// Convert y values to vec.
	yVec := mat.NewVecDense(numGate, y)

	// Create the coefficients.
	coeffs := mat.NewVecDense(numGate+1, nil)

	// Solve the linear system.
	err := coeffs.SolveVec(matrixA, yVec)

	// Error handling.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Assign values and return.
	x = coeffs.RawVector().Data
	return
}
