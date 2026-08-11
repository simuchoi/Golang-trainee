package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 5) // размер массива, предвыделение памяти

	fmt.Println("\nДо:")
	fmt.Println("slice", intSlice)
	fmt.Println("len", len(intSlice))
	fmt.Println("cap", cap(intSlice))

	intSlice = append(intSlice, 1, 2, 3, 4, 5, 6)

	fmt.Println("\nПосле:")
	fmt.Println("slice", intSlice)
	fmt.Println("len", len(intSlice))
	fmt.Println("cap", cap(intSlice))
}
