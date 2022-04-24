package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f := func(x float64) float64 {
		return math.Pow(x, 3) - 7*x - 6
	} // goal function.

	firstDeriv := func(x float64) float64 {
		return 3*math.Pow(x, 2) - 7
	}

	secondDeriv := func(x float64) float64 {
		return 6 * x
	}

	var (
		a     float64 = 2.5
		b     float64 = 3.5
		x0    float64 = 3.5
		fDMin float64 = math.Abs(firstDeriv(a))
		sDMax float64 = math.Abs(secondDeriv(b))
	) // newton method arguments.

	eps := getEpsilon()

	res, approx, err := newton(a, b, x0, fDMin, sDMax, eps, f, firstDeriv, secondDeriv)
	if err != nil {
		log.Println(err)
	} // calculating...

	_, err = os.Stdout.WriteString(fmt.Sprintf("The result of calculations is: %f +- %f", res, approx))
	if err != nil {
		log.Println(err)
	}
}

func getEpsilon() float64 {
	eps := math.Pow(10, -3) // default epsilon.

	_, err := os.Stdout.WriteString(fmt.Sprintf("Default epsilon is '%s'. Do you want to change it [Y/(any)]?\n",
		strconv.FormatFloat(eps, 'f', -1, 64)))
	if err != nil {
		log.Println(err)

		return eps
	}

	var ans string // string variable to hold user answer.

	_, err = fmt.Scanln(&ans)
	if err != nil {
		log.Println(err)

		return eps
	}

	if ans == "Y" {
		_, err = os.Stdout.WriteString("Write your epsilon.\n")
		if err != nil {
			log.Println(err)

			return eps
		}

		var newEps float64 // variable to hold new epsilon variable (is needed if read error occurs)

		_, err = fmt.Scanln(&newEps)
		if err == nil {
			eps = newEps
		} else {
			_, err = os.Stdout.WriteString("Given epsilon is not float type.\n")
			if err != nil {
				log.Println(err)
			}
		}
	}

	_, err = os.Stdout.WriteString(fmt.Sprintf("Using epsilon '%s'.\n", strconv.FormatFloat(eps, 'f', -1, 64)))
	if err != nil {
		log.Println(err)

		return eps
	}

	return eps
}
