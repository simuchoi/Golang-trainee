package main

import "fmt"

func main() {

	fmt.Println("main 1")

	defer func() {
		fmt.Println("defer 1")
	}()

	fmt.Println("main 2")

	defer func() {
		fmt.Println("defer 2")
	}()

	fmt.Println("main 3")

	defer func() {
		fmt.Println("defer 3")
	}()
}
