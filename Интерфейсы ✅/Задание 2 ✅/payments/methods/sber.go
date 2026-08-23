package methods

import (
	"fmt"
	"math/rand"
)

type Sber struct{}

func NewSber() *Sber {
	return &Sber{}
}

func (c Sber) Pay(usd float64) int {
	fmt.Println("\nОплата через СБЕР")
	fmt.Println("Итого:", usd, "долларов")

	return rand.Int()
}

func (c Sber) Cancel(id int) {
	fmt.Println("\nОтмена операции \nID:", id)
}
