package main

import (
	"fmt"
	"task1_module/greeting"
)

func main() {
	greeting.SayHello()

	num := greeting.ReturnInt()
	fmt.Println("ReturnInt from main.go: ", num)

	greeting.SayBad()
}