package main

import (
	"fmt"
)

func main() {
	arr := [5]int{51, 42, 33, 24, 15}

	fmt.Println("\nвывод всех элементов массива")
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[", i, "] =", arr[i])
	}

	fmt.Println("\nумножение четных элементов массива на 10")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 10
		}
		fmt.Println("arr[", i, "] =", arr[i])
	}

	fmt.Println("\nделение нечетных элементов массива на 10")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			arr[i] /= 10
		}
		fmt.Println("arr[", i, "] =", arr[i])
	}
}
