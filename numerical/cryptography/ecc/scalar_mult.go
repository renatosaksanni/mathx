package ecc

import (
	"math/big"
)

func (curve *EllipticCurve) ScalarMult(point *Point, k *big.Int) (*Point, error) {
	result := &Point{X: big.NewInt(0), Y: big.NewInt(1), Z: big.NewInt(0)}
	temp := point

	for i := k.BitLen() - 1; i >= 0; i-- {
		result, _ = curve.AddPoints(result, result)
		if k.Bit(i) == 1 {
			result, _ = curve.AddPoints(result, temp)
		}
	}

	return result, nil
}
