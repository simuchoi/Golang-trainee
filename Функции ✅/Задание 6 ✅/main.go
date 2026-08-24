package main

import "fmt"

func main() {
	greeting("Oleg", 1)
}

func greeting(name string, number int) {
	fmt.Println("Good morning", name)
	fmt.Println("Provide you to room number", number)
}
