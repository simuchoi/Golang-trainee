package main

import "fmt"

func main() {
	fmt.Println("\nchange wheel radius")
	radius := 16.5
	radiusPtr := &radius

	fmt.Println("radius before: ", radius)
	changeRadius(radiusPtr, 17.1)
	fmt.Println("radius after: ", radius)

}

func changeRadius(ptr *float64, num float64) {
	*ptr = num
}
