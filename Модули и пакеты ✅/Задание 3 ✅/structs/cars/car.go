package car

import "fmt"

type Car struct {
	brand    string
	maxSpeed float64
}

func NewCar(
	brand string,
	maxSpeed float64,
) Car {
	if brand == "" {
		fmt.Println("empty brand")
		return Car{}
	}

	if maxSpeed < 0 {
		fmt.Println("wrong speed")
		return Car{}
	}

	return Car{
		brand:    brand,
		maxSpeed: maxSpeed,
	}
}
