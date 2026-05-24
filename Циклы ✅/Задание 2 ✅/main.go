package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ТЗ: загадать случайное число от 1 до 50 и принимать попытки пользователя,
// после каждой попытки подсказывать "Больше" или "Меньше", пока число не угадано.
func main() {
	rand.Seed(time.Now().UnixNano())

	num := 0
	randNum := rand.Intn(50) + 1

	fmt.Println("Я загадал число от 1 до 50, попробуй угадать!")

	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&num)

		if num == randNum {
			fmt.Println("Поздравляю! Ты угадал!")
			break
		} else if num > randNum {
			fmt.Println("Меньше")
		} else {
			fmt.Println("Больше")
		}
	}
}
