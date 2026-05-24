package main

import "fmt"

// ТЗ:
// 1) Создай переменную n со значением 30 и указатель p на нее.
// 2) Напиши функцию setValue(p *int, newValue int), которая меняет значение по указателю.
// 3) В main вызови setValue(p, 77).
// 4) Выведи n до и после вызова функции.
// 5) Ожидаемый итог: после вызова n = 77.

func main() {
	n := 30
	p := &n

	fmt.Printf("n before: %d", n)

	setValue(p, 77)

	fmt.Printf("\nn after: %d", n)
}

func setValue(p *int, newValue int) {
	*p = newValue
}

// Оценка: 5/5 — задание выполнено корректно.
