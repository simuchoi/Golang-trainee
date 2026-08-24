package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\n2 roots")
	solveQuadratic(-2, 4, 6)

	fmt.Println("\n1 root")
	solveQuadratic(1, -6, 9)

	fmt.Println("\n0 roots")
	solveQuadratic(1, 2, 6)
}

func solveQuadratic(a float64, b float64, c float64) {
	D := math.Pow(b, 2) - 4*a*c
	sqrtD := math.Sqrt(D)

	if D < 0 {
		fmt.Println("---")
		return
	}

	if D == 0 {
		x := -b / (2 * a)

		fmt.Printf("x1: %.2f\n", x)

		return
	}

	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)

	fmt.Printf("\nx1: %.2f \nx2: %.2f\n", x1, x2)
}
