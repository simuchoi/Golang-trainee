package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

func main() {
	type User struct {
		name string
		age  int
	}

	userArray := []User{
		User{
			name: "Sam",
			age:  39,
		},
		User{
			name: "John",
			age:  37,
		},
	}

	fmt.Println("\nПользователи до append")

	for index, user := range userArray {
		pp.Println(user, index)
	}
	fmt.Println("len:", len(userArray))
	fmt.Println("cap:", cap(userArray))

	userArray = append(userArray, User{"Tom", 35})
	userArray = append(userArray, User{"will", 20})

	fmt.Println("\nПользователи после append")

	for index, user := range userArray {
		pp.Println(user, index)
	}
	fmt.Println("len:", len(userArray))
	fmt.Println("cap:", cap(userArray))
}
