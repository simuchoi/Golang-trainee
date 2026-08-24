package main

import "fmt"

func main() {
	fmt.Println("\nEntry point to main\n")
	boo()
	foo()
}

func boo() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

func foo() {
	fmt.Println("11")
	fmt.Println("22")
	fmt.Println("33")
}
