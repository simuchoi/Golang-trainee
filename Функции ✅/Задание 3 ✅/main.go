package main

import (
	"fmt"
)

// Задание 3 — Рекурсия: числа Фибоначчи
//
// Напиши две версии функции вычисления n-го числа Фибоначчи:
// - fibRecursive(n int) int — через рекурсию
// - fibIterative(n int) int — через цикл
//
// Затем напиши функцию compareFib(n int) которая вызывает обе, сравнивает результаты и выводит:
// compareFib(10) // → Recursive: 55 | Iterative: 55 | Results match: true
//
// Подсказка: задание помогает понять разницу между рекурсивным и итеративным подходами

func main() {
	compareFib(4)
}

func compareFib(n int) {
	iter, err := fibIterative(n)

	rec := fibRecursive(n)

	if err != nil {
		fmt.Printf("Recursive: %d | Iterative: error: %v\n", rec, err)
		return
	}

	match := rec == iter
	fmt.Printf("Recursive: %d | Iterative: %d | Match: %v\n", rec, iter, match)
}

func fibIterative(n int) (int, error) {
	sum := 0
	num1 := 0
	num2 := 1

	for i := 0; i <= n; i++ {
		sum = num1 + num2

		fmt.Printf("fibIterative #%d: %d", i, sum)

		num1 = num2
		num2 = sum
	}
	return sum, nil
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}
