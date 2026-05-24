package main

import "fmt"

// ТЗ:
// 1) Напиши функцию doubleAll(nums *[]int), которая принимает указатель на срез.
// 2) Умножь каждый элемент среза на 2 через указатель.
// 3) В main создай срез из 5 произвольных целых чисел.
// 4) Выведи срез до и после вызова doubleAll(&numbers).
// 5) Ожидаемый итог: исходный срез изменен, все элементы увеличены в 2 раза.
func doubleAll(nums *[]int) {
	for i := 0; i < len((*nums)); i++ {
		(*nums)[i] = (*nums)[i] * 2
	}
}

func main() {
	numbers := []int{3, 7, 1, 9, 4}

	fmt.Println("До:   ", numbers)
	doubleAll(&numbers)
	fmt.Println("После:", numbers)
}

// Оценка: 5/5 — задание выполнено корректно.
