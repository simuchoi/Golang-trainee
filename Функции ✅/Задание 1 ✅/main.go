// Задание 1 — Калькулятор Напиши программу-калькулятор, где каждая операция — это отдельная функция. Требования: - Функции: add, subtract, multiply, divide - Функция divide должна возвращать два значения — результат и ошибку (error), если делитель равен нулю - Написать функцию calculate(a, b float64, operation string) которая вызывает нужную функцию в зависимости от переданной операции

package main

import (
	"errors"
	"fmt"
)

func main() {
	operator := ""
	fmt.Scan(&operator)
	calculate(3, 4, operator)
}

// Задание 1
func calculate(a, b float64, opr string) {
	switch opr {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(subtract(a, b))
	case "*":
		fmt.Println(multiply(a, b))
	case "/":
		fmt.Println(divide(a, b))
	}
}

func add(a, b float64) float64 {
	s := a + b
	return s
}

func subtract(a, b float64) float64 {
	s := a - b
	return s
}

func multiply(a, b float64) float64 {
	s := a * b
	return s
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}
