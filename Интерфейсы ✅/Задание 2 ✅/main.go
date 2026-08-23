package main

import (
	"fmt"
	"interfaces/payments"
	"interfaces/payments/methods"

	"github.com/k0kubun/pp/v3"
)

func main() {
	method := methods.NewPayPal() //

	paymentModule := payments.NewPaymentModule(method)

	idFood := paymentModule.Pay("Бургер", 5.0)
	foodInfo := paymentModule.GetInfo(idFood)

	idPhone := paymentModule.Pay("iPhone 17 pro max", 1200.0)
	phoneInfo := paymentModule.GetInfo(idPhone)

	idGame := paymentModule.Pay("GTA VI", 100.0)
	gameInfo := paymentModule.GetInfo(idGame)
	paymentModule.Cancel(idGame)

	allInfo := paymentModule.GetAllInfo()

	fmt.Println("\nВсе транзакции:")
	pp.Println(allInfo)

	fmt.Println("\nТранзакции по-отдельности:")
	pp.Println("Игра:", gameInfo)
	pp.Println("Телефон:", phoneInfo)
	pp.Println("Еда:", foodInfo)

}
