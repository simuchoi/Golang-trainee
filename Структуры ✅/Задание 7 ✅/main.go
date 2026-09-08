package main

import (
	"fmt"
	"structModule/car"
	"structModule/ice"
)

func main() {
	car1, err := car.NewCar(
		"nissan 350z",
		3.2,
		4,
		3,
		"black",
		true,
		20000.00,
	)

	car1.Price = -1

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("car1 struct before: ", car1)
	car1.ChangeEngine(10)
	fmt.Println("change engine")
	fmt.Println("car1 struct after: ", car1)

	iceCream1 := ice.NewIceCream("strawberry", 200.0)
	fmt.Println("iceCream1 struct: ", iceCream1)
}
