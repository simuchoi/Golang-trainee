// Задание 2 — Работа со срезами Напиши 3 функции для работы с числовым срезом []int:
// - findMin(nums []int) int — находит минимальный элемент
// - findMax(nums []int) int — находит максимальный элемент
// - average(nums []int) float64 — считает среднее арифметическое
// - analyze(nums []int) — которая вызывает все три и красиво выводит результаты в консоль. analyze([]int{3, 7, 1, 9, 4, 6}) // → Min: 1 | Max: 9 | Average: 5.00

package main

import "fmt"

func main() {
	arr := []int{-5, 20, -1}
	analyze(arr)
}

func analyze(nums []int) float64 {
	fmt.Print("Min: ", findMinValue(nums), " | ")
	fmt.Print("Max: ", findMaxValue(nums), " | ")
	fmt.Print("Avg: ", findAverageValue(nums))

	return 0
}

// принимает массив типа int, возвращает число типа int — находит минимальный элемент
func findMinValue(arr []int) int {
	minValue := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	return minValue
}

// принимает массив типа int, возвращает число типа int — находит максимальный элемент
func findMaxValue(arr []int) int {
	maxValue := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	return maxValue
}

// принимает массив типа int, возвращает число типа float64 — считает среднее арифметическое
func findAverageValue(arr []int) float64 {
	var sum float64 = 0

	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}
	sum /= float64(len(arr))
	return sum
}
