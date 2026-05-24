package main

import "fmt"

// ТЗ:
// 1) Создай переменную value со значением 15.
// 2) Напиши функцию addFive(p *int), которая увеличивает число на 5 через указатель.
// 3) В main выведи value до и после вызова addFive(&value).
// 4) Для вывода используй fmt.Printf.
// 5) Ожидаемый итог: после вызова value = 20.

func main() {
	value := 15

	fmt.Printf("Value before: %d", value)

	addFive(&value)

	fmt.Printf("\nValue after:  %d", value)
}

func addFive(p *int) {
	*p += 5
}

// Оценка: 5/5 — задание выполнено корректно.
