package ecc

import (
	"errors"
	"math/big"
)

type Point struct {
	X, Y *big.Int
}

func (curve *EllipticCurve) AddPoints(p1, p2 *Point) (*Point, error) {
	if !curve.IsOnCurve(p1) {
		return nil, errors.New("p1 is not on the curve")
	}
	if !curve.IsOnCurve(p2) {
		return nil, errors.New("p2 is not on the curve")
	}

	if p1.X == nil {
		return p2, nil
	}

	if p2.X == nil {
		return p1, nil
	}

	if p1.X.Cmp(p2.X) == 0 && p1.Y.Cmp(new(big.Int).Neg(p2.Y)) == 0 {
		return &Point{X: nil, Y: nil}, nil // Point at infinity
	}

	lambda := new(big.Int)
	if p1.X.Cmp(p2.X) == 0 && p1.Y.Cmp(p2.Y) == 0 {
		// Point doubling
		num := new(big.Int).Mul(big.NewInt(3), new(big.Int).Exp(p1.X, big.NewInt(2), curve.P))
		num.Add(num, curve.A)
		denom := new(big.Int).Mul(big.NewInt(2), p1.Y)
		denom.ModInverse(denom, curve.P)
		lambda.Mul(num, denom)
	} else {
		// Point addition
		num := new(big.Int).Sub(p2.Y, p1.Y)
		denom := new(big.Int).Sub(p2.X, p1.X)
		denom.ModInverse(denom, curve.P)
		lambda.Mul(num, denom)
	}

	lambda.Mod(lambda, curve.P)

	x3 := new(big.Int).Exp(lambda, big.NewInt(2), curve.P)
	x3.Sub(x3, p1.X)
	x3.Sub(x3, p2.X)
	x3.Mod(x3, curve.P)

	y3 := new(big.Int).Sub(p1.X, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p1.Y)
	y3.Mod(y3, curve.P)

	return &Point{X: x3, Y: y3}, nil
}
