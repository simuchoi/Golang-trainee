package main

import "fmt"

// ТЗ:
// 1) Создай две переменные a = 10 и b = 20.
// 2) Напиши функцию swap(x, y *int), которая меняет значения местами через указатели.
// 3) В main вызови swap(&a, &b).
// 4) Выведи значения до и после вызова функции.
// 5) Ожидаемый итог: после swap a = 20, b = 10.

func main() {
	a := 10
	b := 20

	fmt.Println("До swap():")
	fmt.Printf("a = %d\nb = %d\n", a, b)

	swap(&a, &b)

	fmt.Println("После swap():")
	fmt.Printf("a = %d\nb = %d\n", a, b)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// Оценка: 5/5 — задание выполнено корректно.
