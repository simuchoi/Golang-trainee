package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello")

	num := ReturnInt()
	
	fmt.Println("ReturnInt from greeting.go: ", num)
}