package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("\nPorabola")
	porabola(5)
}

func porabola(x float64) {
	for index := 0.0; index <= 10; index++ {
		x = index
		y := math.Sqrt(index)

		fmt.Printf("\ny = %.2f AT x = %.2f", y, x)
	}
}
