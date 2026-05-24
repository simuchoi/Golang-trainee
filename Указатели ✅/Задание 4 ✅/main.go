package main

import "fmt"

// ТЗ:
// 1) Создай переменную number со значением 25.
// 2) Создай указатель ptr, который хранит адрес number.
// 3) Через указатель измени значение number на 50 (не напрямую).
// 4) Выведи в консоль значение number, адрес ptr и значение по адресу *ptr.
// 5) Ожидаемый итог: number и *ptr равны 50.

func main() {
	number := 25
	ptr := &number

	fmt.Printf("До изменения:\n")
	fmt.Printf("number = %d\n", number)
	fmt.Printf("ptr    = %p\n", ptr)

	*ptr = 50

	fmt.Printf("После изменения:\n")
	fmt.Printf("number = %d\n", number)
	fmt.Printf("*ptr   = %d\n", *ptr)
}

// Оценка: 5/5 — задание выполнено корректно.
