package main

import "fmt"

// ТЗ: считать положительное целое число и посчитать сумму его цифр
// с помощью цикла и операций %10 и /10.
func main() {
	num := 0
	sum := 0

	fmt.Println("Сумма цифр числа")
	fmt.Print("Введите положительное число: ")
	fmt.Scan(&num)

	for num > 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	fmt.Println("Сумма цифр:", sum)
}
