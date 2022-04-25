package main

import (
	"errors"
	"fmt"
	"math"
)

var errNotNegative = errors.New("multiplication of given functions with given args a and b is not negative")

func bisection(a, b, epsilon float64, f func(float64) float64) (float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, fmt.Errorf("%w. a: %f; f(a): %f; b: %f; f(b): %f", errNotNegative, a, f(a), b, f(b))
	}

	iterations := int(math.Floor(math.Log2((b-a)/epsilon)) + 1)

	var c float64

	for i := 0; i < iterations; i++ {
		c = (a + b) / 2

		fmt.Println("xn: ", c)

		if f(c) == 0.0 {
			break
		}

		fmt.Println("f(xn): ", f(c))

		if f(c)*f(a) > 0 {
			a = c
		}

		if f(c)*f(b) > 0 {
			b = c
		}
	}

	return c, nil
}
