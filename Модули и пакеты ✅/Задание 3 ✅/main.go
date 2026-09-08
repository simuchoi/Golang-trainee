package main

import (
	apartment "task3/structs/apartments"
	car "task3/structs/cars"

	"github.com/k0kubun/pp/v3"
)

func main() {
	apartment1 := apartment.NewApartment(12, 20, 12)
	pp.Println(apartment1)

	car1 := car.NewCar("123", 122.0)
	pp.Println(car1)
}
