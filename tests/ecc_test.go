package tests

import (
	"math/big"
	"testing"

	"github.com/renatosaksanni/mathx/numerical/cryptography/ecc"
	"github.com/stretchr/testify/assert"
)

func TestIsOnCurve(t *testing.T) {
	a, _ := new(big.Int).SetString("2", 10)
	b, _ := new(big.Int).SetString("3", 10)
	p, _ := new(big.Int).SetString("97", 10)
	curve := ecc.NewEllipticCurve(a, b, p)

	point := &ecc.Point{
		X: new(big.Int).SetInt64(3),
		Y: new(big.Int).SetInt64(6),
	}

	assert.True(t, curve.IsOnCurve(point), "Point should be on the curve")
}

func TestScalarMult(t *testing.T) {
	a, _ := new(big.Int).SetString("2", 10)
	b, _ := new(big.Int).SetString("3", 10)
	p, _ := new(big.Int).SetString("97", 10)
	curve := ecc.NewEllipticCurve(a, b, p)

	point := &ecc.Point{
		X: new(big.Int).SetInt64(3),
		Y: new(big.Int).SetInt64(6),
	}
	scalar := new(big.Int).SetInt64(2)

	result, err := curve.ScalarMult(point, scalar)
	assert.NoError(t, err)
	assert.True(t, curve.IsOnCurve(result), "Resulting point should be on the curve")
}
