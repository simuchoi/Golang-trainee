package main

import "fmt"

func main() {
	testFunc()
	sumCounter(8, 4)
	divideCounter(8, 4)

	fmt.Println("main 1") // 7

	defer func() {
		fmt.Println("defer main") // 9
	}()

	fmt.Println("main 2") //8
}

func sumCounter(a int, b int) {
	result := a + b

	defer func() {
		fmt.Println("defer sumCounter") // 4
	}()

	fmt.Println("result sum:", result) // 3
}

func divideCounter(a int, b int) {
	result := a / b

	fmt.Println("result divide:", result) // 5

	defer func() {
		fmt.Println("defer divideCounter()") // 6
	}()
}

func testFunc() {
	fmt.Println("i'm test function") // 1
	defer func() {
		fmt.Println("defer testFunc()") // 2
	}()
}
