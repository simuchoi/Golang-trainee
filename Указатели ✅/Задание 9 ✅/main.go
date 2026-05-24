package main

import "fmt"

// ТЗ:
// 1) Объяви пустой указатель: var p *int.
// 2) Проверь условие p == nil.
// 3) Если указатель nil, выведи "указатель пустой".
// 4) Если не nil, выведи значение по адресу *p.
// 5) Сделай второй пример: привяжи p к переменной n и снова выведи результат.

func main() {
	var pointer *int
	number := 10

	if pointer == nil {
		fmt.Println("Указатель пустой")
	}

	pointer = &number

	if pointer != nil {
		fmt.Println("Указатель не пустой")
		fmt.Printf("*pointer: %d\n", *pointer)
	}
}

// Оценка: 5/5 — задание выполнено корректно.
