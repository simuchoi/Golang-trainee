package main

import (
	"errors"
	"fmt"
)

type apartments struct {
	number int
	rooms  int
	floor  int
	area   float64
	price  float64
}

func newApartments(
	number int,
	rooms int,
	floor int,
	area float64,
	price float64,
) (apartments, error) {
	if number < 0 || number > 100 {
		return apartments{}, errors.New("wrong number")
	}

	if rooms < 0 || rooms > 5 {
		return apartments{}, errors.New("wrong rooms amount")
	}

	if floor < 1 || floor > 10 {
		return apartments{}, errors.New("wrong floor")
	}

	if area < 20 || area > 80 {
		return apartments{}, errors.New("wrong area")
	}

	if price < 0 {
		return apartments{}, errors.New("wrong price")
	}

	return apartments{
		number: number,
		rooms:  rooms,
		area:   area,
		price:  price,
	}, nil
}

func (a *apartments) changePrice(price float64) {
	a.price = price
}

func main() {
	apartment1, err := newApartments(49, 1, 6, 50, 6450000.00)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("apartment1 before: ", apartment1)
	apartment1.changePrice(7500000.00)
	fmt.Println("apartment1 after: ", apartment1)
}
