package FLPCP

import (
	"crypto/rand"
	"github.com/Weiqi97/Single-Server-Prio/utilities"
	"github.com/fentec-project/gofe/data"
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

func (scheme FLPCP) prover(x data.Vector) (pi data.Vector) {
	// The prover output pi = (w, f_1(0), ..., f_L(0), c_p), where L = 3.
	var p = make([]*big.Int, scheme.N+1)

	// W is the witness, which we randomly sample from Zp.
	var w *big.Int
	w, _ = rand.Int(rand.Reader, scheme.P)
	pi = append(pi, w)

	// f_i(j) is the i-th input wired to j-th G gate, we have n + 1 G gate, where |x| = n.
	var f10, f20, f30 *big.Int
	f10, _ = rand.Int(rand.Reader, scheme.P)
	f20, _ = rand.Int(rand.Reader, scheme.P)
	f30, _ = rand.Int(rand.Reader, scheme.P)
	pi = append(pi, f10, f20, f30)

	// Useful values for the circuit.
	var one, r, temp, output *big.Int
	one.SetInt64(1)
	r, _ = rand.Int(rand.Reader, scheme.P)

	// p(j) is the value on the output wire from the j-th G-gate.
	for i := 0; i < scheme.N; i++ {
		temp.Sub(x[i], one)
		p[i].Mul(x[i], temp)
	}

	// The final output G gate.
	output = utilities.SumVector(scheme.N, p)
	temp.Sub(w, r)
	p[scheme.N+1].Mul(output, temp)

	// So the degree is N + 1.

	return
}

func (scheme FLPCP) verifier() {

}
