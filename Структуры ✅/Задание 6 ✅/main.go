package main

import "fmt"

type Car struct {
	model  string
	engine float64
	wheels int
	doors  int
	color  string
	isNew  bool
	price  float64
}

func main() {

	car1 := Car{
		model:  "nissan 350z",
		engine: 3.2,
		wheels: 4,
		doors:  3,
		color:  "silver",
		isNew:  true,
		price:  20000.00,
	}
	car2 := Car{
		model:  "porcshe 911",
		engine: 4.2,
		wheels: 4,
		doors:  3,
		color:  "black",
		isNew:  true,
		price:  55000.00,
	}

	fmt.Println("\nmodel: ", car1.model)
	fmt.Println("engine: ", car1.engine)
	fmt.Println("wheels: ", car1.wheels)
	fmt.Println("doors: ", car1.doors)
	fmt.Println("color: ", car1.color)
	fmt.Println("isNew: ", car1.isNew)
	fmt.Println("price: ", car1.price)

	fmt.Println("\nA crash has happened...")

	car1.isNew = false
	car1.price = 3000.00

	fmt.Println("\nmodel: ", car1.model)
	fmt.Println("engine: ", car1.engine)
	fmt.Println("wheels: ", car1.wheels)
	fmt.Println("doors: ", car1.doors)
	fmt.Println("color: ", car1.color)
	fmt.Println("isNew: ", car1.isNew)
	fmt.Println("price: ", car1.price)

	fmt.Println("\nBuying another car...")

	fmt.Println("\nmodel: ", car2.model)
	fmt.Println("engine: ", car2.engine)
	fmt.Println("wheels: ", car2.wheels)
	fmt.Println("doors: ", car2.doors)
	fmt.Println("color: ", car2.color)
	fmt.Println("isNew: ", car2.isNew)
	fmt.Println("price: ", car2.price)

}
