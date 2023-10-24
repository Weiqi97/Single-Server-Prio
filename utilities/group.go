package utilities

import "math/big"
import "crypto/rand"

type Zp struct {
	ele *big.Int
	mod *big.Int
}

func (x Zp) GetEle() *big.Int {
	return x.ele
}

func (x Zp) GetMod() *big.Int {
	return x.mod
}

func NewZp(ele *big.Int, mod *big.Int) (result Zp) {
	result.ele, result.mod = ele, mod
	return
}

func RandZp(mod *big.Int) (result Zp) {
	// Omit the error check for simple usage case.
	random, _ := rand.Int(rand.Reader, mod)
	result = NewZp(random, mod)
	return
}

func Add(x Zp, y Zp) (result Zp) {
	temp := new(big.Int)
	temp.Add(x.ele, y.ele)
	temp.Mod(temp, x.mod)
	result = NewZp(temp, x.mod)
	return
}

func Sub(x Zp, y Zp) (result Zp) {
	temp := new(big.Int)
	temp.Sub(x.ele, y.ele)
	temp.Mod(temp, x.mod)
	result = NewZp(temp, x.mod)
	return
}

func Mul(x Zp, y Zp) (result Zp) {
	temp := new(big.Int)
	temp.Mul(x.ele, y.ele)
	temp.Mod(temp, x.mod)
	result = NewZp(temp, x.mod)
	return
}

func Inv(x Zp) (result Zp) {
	temp := new(big.Int)
	temp.ModInverse(x.ele, x.mod)
	result = NewZp(temp, x.mod)
	return
}

func Div(x Zp, y Zp) (result Zp) {
	temp := new(big.Int)
	temp.ModInverse(y.ele, x.mod)
	temp.Mul(x.ele, temp)
	temp.Mod(temp, x.mod)
	result = NewZp(temp, x.mod)
	return
}
