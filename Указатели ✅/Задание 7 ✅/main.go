package main

import "fmt"

// ТЗ:
// 1) Создай переменную x со значением 99.
// 2) Напиши функцию reset(p *int), которая присваивает 0 через указатель.
// 3) В main вызови reset(&x).
// 4) Выведи значение x до и после вызова функции.
// 5) Ожидаемый итог: после reset x = 0.

func main() {
	x := 99

	fmt.Printf("x before: %d\n", x)

	reset(&x)

	fmt.Printf("x after: %d\n", x)
}

func reset(p *int) {
	*p = 0
}

// Оценка: 5/5 — задание выполнено корректно.
