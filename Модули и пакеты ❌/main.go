package main

import (
	"fmt"
	"modules_packets/greeting"
	"modules_packets/user"
)

func main() {
	greeting.SayHello()

	num := greeting.ReturnInt()
	fmt.Println("ReturnInt from main.go: ", num)

	greeting.SayBad()

	u := user.User{}
	
	u.name = "sad"

	fmt.Println("u.name: ", u.name)
}