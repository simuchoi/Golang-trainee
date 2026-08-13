package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

func main() {
	type User struct {
		Name    string
		Rating  float64
		Premium bool
	}

	userArray := [3]User{
		User{
			Name:    "Oleg",
			Rating:  10.0,
			Premium: true,
		},
		User{
			Name:    "Nikita",
			Rating:  8.5,
			Premium: true,
		},
		User{
			Name:    "Klim",
			Rating:  9.0,
			Premium: false,
		},
	}

	fmt.Println("\nПользователи c премиумом до:") // обычный цикл (зависит от длины массива)
	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}

	fmt.Println("\nПользователи c премиумом после:") // обычный цикл (зависит от длины массива)
	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium == true {
			userArray[i].Rating += 1
		}
		pp.Println(userArray[i])
	}

	fmt.Println("\nПробег по массиву с range (копия)") // работа с копией, только для вывода, индексы и элементы
	for index, user := range userArray {               // пробег по всем элементам (длина массива не важна)
		if user.Premium {
			user.Rating += 1
		}
		pp.Println(index, user)
	}

	fmt.Println("\nПробег по массиву с range (копия) без индексов") // работа с копией, только для вывода, только элементы
	for _, user := range userArray {                                // пробег по всем элементам  (длина массива не важна)
		pp.Println(user)
	}

	fmt.Println("\nПробег по массиву с range (оригинал)") // работа с оригиналом, для изменения элементов
	for index, user := range userArray {                  // пробег по всем элементам (длина массива не важна)
		if user.Premium {
			userArray[index].Rating += 1
		}
		pp.Println(index, user)
	}

	fmt.Println("\nПробег по массиву с range (оригинал)") // работа с оригиналом, для изменения элементов
	for index, user := range userArray {                  // пробег по всем элементам (длина массива не важна)
		pp.Println(index, user)
	}
}
