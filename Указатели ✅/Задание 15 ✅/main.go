package main

import "fmt"

// ТЗ:
// 1) Создай срез nums := []int{1, 2, 3, 4}.
// 2) Напиши функцию multiplyByThree(data *[]int), которая умножает каждый элемент на 3.
// 3) В main выведи nums до и после вызова multiplyByThree(&nums).
// 4) Измени значения только через указатель на срез.
// 5) Ожидаемый итог: nums = [3 6 9 12].

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println("\nnums before:")

	for i := 0; i < len(nums); i++ {
		fmt.Println("num: ", nums[i])
	}

	multiplyByThree(&nums)

	fmt.Println("\nnums after:")

	for i := 0; i < len(nums); i++ {
		fmt.Println("num: ", nums[i])
	}
}

func multiplyByThree(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 3
	}
}

// Оценка: 5/5 — задание выполнено корректно.
