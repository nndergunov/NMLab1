package main

import (
	"fmt"
	"math"
)

func main() {
	phiDeriv := func(x float64) float64 {
		return 1 / 3 * math.Pow(6*math.Pow(x, 2)-5*x-12, -2/3) * (12*x - 5)
	} // derivative of phi function.

	var (
		a  float64 = -2
		b  float64 = 0
		x0 float64 = -1.5
		q  float64 = 1 / 3 * math.Pow(12, -2/3) * (-5)
	) // iteration method arguments.

	fmt.Println(satisfiesTheoremOne(x0, q, phiDeriv))
	fmt.Println(satisfiesTheoremTwo(a, b, x0, q, phiDeriv))
}
