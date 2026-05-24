package main

import "fmt"

func main() {
	fmt.Println("Калькулятор")

	num1 := 0
	num2 := 0
	opr := ""

	fmt.Print("Число 1: ")
	fmt.Scan(&num1)

	fmt.Print("Число 2: ")
	fmt.Scan(&num2)

	fmt.Print("Введите оператор: ")
	fmt.Scan(&opr)

	if opr == "+" {
		res := num1 + num2
		fmt.Println("Результат:", res)
	} else if opr == "-" {
		res := num1 - num2
		fmt.Println("Результат:", res)
	} else if opr == "*" {
		res := num1 * num2
		fmt.Println("Результат:", res)
	} else if opr == "/" {
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			res := num1 / num2
			fmt.Println("Результат:", res)
		}
	} else {
		fmt.Println("Такого оператора не существует")
	}
}
