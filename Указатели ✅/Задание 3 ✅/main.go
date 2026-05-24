package main

import "fmt"

// ТЗ:
// 1) Реализуй три функции для счетчика: increment(n *int), decrement(n *int), reset(n *int).
// 2) increment увеличивает значение по указателю, decrement уменьшает, reset обнуляет.
// 3) В main объяви переменную counter.
// 4) Вызови increment 5 раз, decrement 2 раза, затем reset.
// 5) Ожидаемый итог: после операций counter = 3, после reset counter = 0.
func increment(n *int) {
	*n = *n + 1
}

func decrement(n *int) {
	*n = *n - 1
}

func reset(n *int) {
	*n = 0
}

func main() {
	counter := 0

	for i := 0; i < 5; i++ {
		increment(&counter)
	}
	for i := 0; i < 2; i++ {
		decrement(&counter)
	}
	fmt.Println("Счётчик после операций:", counter) // ожидается 3

	reset(&counter)
	fmt.Println("Счётчик после сброса:  ", counter) // ожидается 0
}

// Оценка: 5/5 — задание выполнено корректно.
