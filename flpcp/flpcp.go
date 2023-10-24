package flpcp

import (
	u "github.com/Weiqi97/Single-Server-Prio/utilities"
	"github.com/fentec-project/gofe/innerprod/simple"
	"math/big"
)

type FLPCP struct {
	N int
	P *big.Int
}

func InitFLPCP(vectorLength int, params *simple.DDHParams) (scheme FLPCP) {
	scheme.N = vectorLength
	scheme.P = params.P
	return
}

func (scheme FLPCP) Prover(x []u.Zp) (pi []u.Zp) {
	// Circuit uses a one; so we create the constant in Zp.
	one := u.NewZp(big.NewInt(1), scheme.P)

	// The prover output pi = (w, f_1(0), ..., f_L(0), c_p), where L = 3.
	var p = make([]u.Zp, scheme.N+1)

	// W is the witness, which we randomly sample from Zp.
	w := u.RandZp(scheme.P)
	pi = append(pi, w)

	// f_i(j) is the i-th input wired to j-th G gate, we have n + 1 G gate, where |x| = n.
	f10 := u.RandZp(scheme.P)
	f20 := u.RandZp(scheme.P)
	f30 := u.RandZp(scheme.P)
	pi = append(pi, f10, f20, f30)

	// Useful values for the circuit.
	r := u.RandZp(scheme.P)

	// p(j) is the value on the output wire from the j-th G-gate.
	for i := 0; i < scheme.N; i++ {
		p[i] = u.Mul(u.Sub(x[i], one), x[i])
	}

	// Temp holder for a bunch of summation.
	temp := u.NewZp(big.NewInt(0), scheme.P)

	// The final output G gate.
	for i := 0; i < scheme.N; i++ {
		temp = u.Add(temp, p[i])
	}
	p[scheme.N] = u.Mul(u.Sub(w, r), temp)

	// So the degree is N + 1.
	coeffs := u.LinSolver(p)

	pi = append(pi, coeffs...)

	return
}

//func (scheme FLPCP) verifier(pi []float64) (b float64) {
//	// Extract the variables.
//	w := pi[0]
//	f10 := pi[1]
//	f20 := pi[2]
//	f30 := pi[3]
//	coeffs := pi[4:]
//
//	// Execute the first check. There's j polynomials need to be created.
//
//	// Execute the second check.
//	b = 0.0
//	for i := 0; i < scheme.N; i++ {
//		b += math.Pow(float64(scheme.N), float64(i)) * coeffs[i]
//	}
//	// Add the constant term.
//	b += coeffs[len(coeffs)-1]
//
//	return
//}
