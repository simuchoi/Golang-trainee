package main

import (
	"fmt"
	"modules_packets/greeting"
	"modules_packets/user"

	"github.com/k0kubun/pp"
)

func main() {
	greeting.SayHello()

	num := greeting.ReturnInt()
	fmt.Println("ReturnInt from main.go: ", num)

	greeting.SayBad()

	user1 := user.NewUser("Oleg", 23)
	user1.SetNewName("Helge")

	user2 := user.User{}
	user2.SetNewName("John")
	user2.SetNewAge(39)

	fmt.Println("user1: ", user1.GetName())
	pp.Println(user1)

	fmt.Println("user2: ", user2.GetName())
	pp.Println(user2)
}