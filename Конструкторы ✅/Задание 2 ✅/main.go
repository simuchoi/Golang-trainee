package main

import (
	"fmt"
)

func main() {
	user := newUser(
		"Oleg",
		"+79154078632",
		23,
		9.9,
		false,
	)

	fmt.Print(user)
}

type User struct {
	Name    string
	Phone   string
	Age     int
	Rating  float64
	isClose bool
}

func newUser(
	name string,
	phone string,
	age int,
	rating float64,
	isClose bool,
) User {
	fmt.Println("\nВалидация имени")
	if name == "" {
		fmt.Println("Не прошло валидацию")
		return User{}
	}

	fmt.Println("\nВалидация возраста")
	if age <= 0 || age >= 150 {
		fmt.Println("Не прошло валидацию")
		return User{}
	}

	fmt.Println("\nВалидация номера телефона")
	if phone == "" {
		fmt.Println("Не прошло валидацию")
		return User{}
	}

	fmt.Println("\nВалидация рейтинга")
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Не прошло валидацию")
		return User{}
	}

	return User{
		Name:    name,
		Phone:   phone,
		Age:     age,
		Rating:  rating,
		isClose: isClose,
	}
}
