package ecc

import (
	"math/big"
)

func (curve *EllipticCurve) ScalarMult(point *Point, k *big.Int) (*Point, error) {
	result := &Point{X: nil, Y: nil}
	temp := point

	for i := k.BitLen() - 1; i >= 0; i-- {
		result, _ = curve.AddPoints(result, result)
		if k.Bit(i) == 1 {
			result, _ = curve.AddPoints(result, temp)
		}
	}

	return result, nil
}
