package main

import "fmt"

func main() {
	fmt.Println("Калькулятор (switch):")
	opr := ""
	num1 := 0
	num2 := 0

	fmt.Print("Число 1: ")
	fmt.Scan(&num1)

	fmt.Print("Число 2: ")
	fmt.Scan(&num2)

	fmt.Print("Оператор: ")
	fmt.Scan(&opr)

	switch opr {
	case "+":
		fmt.Println("Результат:", num1+num2)
	case "-":
		fmt.Println("Результат:", num1-num2)
	case "*":
		fmt.Println("Результат:", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Println("Результат:", num1/num2)
		}
	default:
		fmt.Println("Ошибка: неизвестный оператор")
	}
}
