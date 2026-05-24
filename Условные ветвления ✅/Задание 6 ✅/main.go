package main

import (
	"fmt"
	"math/rand"
)

func main() {
	age := 0
	dayOfWeek := ""
	price := 0
	ticketNumber := 0

	fmt.Println("Вы покупаете билет в кино")

	fmt.Print("Укажите ваш возраст (0-150): ")
	fmt.Scan(&age)

	fmt.Print("Укажите день недели (пн-вс): ")
	fmt.Scan(&dayOfWeek)

	if age >= 0 && age < 12 {
		price = 100
	} else if age >= 12 && age <= 17 {
		price = 200
	} else if age >= 18 && age <= 60 {
		price = 350
	} else if age > 60 && age <= 150 {
		price = 150
	} else {
		fmt.Println("Какой-то странный возраст")
		return
	}

	if dayOfWeek == "ср" {
		price = price - price*20/100
		fmt.Println("\nСегодня среда! Получите скидку 20%")
	}

	fmt.Println("\nИнформация о покупке:")

	ticketNumber = rand.Intn(100)
	fmt.Println("Номер билета:", ticketNumber)
	fmt.Println("Стоимость билета:", price)
	fmt.Println("Ждем вас в кинотеатре!")
}
