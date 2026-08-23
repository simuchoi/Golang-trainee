package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (c Crypto) Pay(usd float64) int {
	fmt.Println("\nОплата через криптовалюту")
	fmt.Println("Итого:", usd, "USDT")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("\nОтмена операции \nID:", id)
}
