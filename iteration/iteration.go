package main

import (
	"math"
)

// satisfiesTheoremOne checks if given variables satisfy second theorem.
func satisfiesTheoremOne(x0, q float64, phiDeriv func(float64) float64) bool {
	if math.Abs(phiDeriv(x0)) > q || q > 1 || q < 0 {
		return false
	}

	return true
}

// satisfiesTheoremTwo checks if given variables satisfy second theorem.
func satisfiesTheoremTwo(a, b, x0, q float64, phiDeriv func(float64) float64) bool {
	delta := math.Abs(math.Abs(x0) - math.Abs(b-a))

	rightCond := math.Abs(phiDeriv(x0) - x0)
	leftCond := (1 - q) * delta

	if rightCond > leftCond {
		return false
	}

	return true
}
