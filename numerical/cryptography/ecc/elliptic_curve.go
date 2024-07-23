package ecc

import (
	"math/big"
)

type EllipticCurve struct {
	A, B, P *big.Int
}

func NewEllipticCurve(a, b, p *big.Int) *EllipticCurve {
	return &EllipticCurve{A: a, B: b, P: p}
}

func (curve *EllipticCurve) IsOnCurve(point *Point) bool {
	if point.X == nil || point.Y == nil {
		return point.X == nil && point.Y == nil
	}

	y2 := new(big.Int).Exp(point.Y, big.NewInt(2), curve.P)
	y2.Mod(y2, curve.P)

	x3 := new(big.Int).Exp(point.X, big.NewInt(3), curve.P)
	ax := new(big.Int).Mul(curve.A, point.X)
	rhs := new(big.Int).Add(x3, ax)
	rhs.Add(rhs, curve.B)
	rhs.Mod(rhs, curve.P)

	return y2.Cmp(rhs) == 0
}
