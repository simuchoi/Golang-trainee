package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() *PayPal {
	return &PayPal{}
}

func (c PayPal) Pay(usd float64) int {
	fmt.Println("\nОплата через PayPal")
	fmt.Println("Итого:", usd, "USD")

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("\nОтмена операции PayPal \nID:", id)
}
