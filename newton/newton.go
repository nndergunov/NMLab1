package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	errNotNegative    = errors.New("multiplication of given functions with given args is not of expected result")
	errTheoremFailure = errors.New("given conditions do not satisfy seecond theorem")
)

func newton(a, b, x0, fDMin, sDMax, epsilon float64, f, firstDeriv, secondDeriv func(float64) float64) (float64, float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, 0, fmt.Errorf("%w. a: %f; f(a): %f; b: %f; f(b): %f", errNotNegative, a, f(a), b, f(b))
	}

	if f(x0)*secondDeriv(b) <= 0 {
		return 0, 0, fmt.Errorf("%w. x0: %f; f(x0): %f; f\"(x0): %f", errNotNegative, x0, f(x0), secondDeriv(x0))
	}

	maxDist := math.Max(math.Abs(x0-a), math.Abs(b-x0))

	q := (sDMax * maxDist) / (2 * fDMin)

	if q >= 1 {
		return 0, 0, fmt.Errorf("%w", errTheoremFailure)
	}

	iterations := int(math.Floor(math.Log2((math.Log(maxDist/epsilon))/(math.Log(1/q)))+1) + 1)

	var res float64

	for i := 0; i < iterations; i++ {
		res = res - f(res)/firstDeriv(x0)
	}

	approx := math.Pow(q, math.Pow(2, float64(iterations)-1)) * maxDist

	return res, approx, nil
}
