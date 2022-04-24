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
		return math.Pow(x, 3) - 4*math.Pow(x, 2) + x + 6
	} // goal function.

	var (
		a float64 = -2
		b float64 = -0.5
	) // bisection arguments.

	eps := getEpsilon()

	res, err := bisection(a, b, eps, f)
	if err != nil {
		log.Println(err)
	} // calculating...

	_, err = os.Stdout.WriteString(fmt.Sprintf("The result of calculations is: %f", res))
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
