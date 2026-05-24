package main

import (
	"fmt"
)

func main() {
	user := User{
		Name:    "Oleg",
		Phone:   "+79154078632",
		Age:     23,
		Rating:  5.5,
		isClose: false,
	}

	fmt.Println("User Rating: ", user.Rating)

	user.addRating(5.2)

	fmt.Println("User Rating: ", user.Rating)
}

type User struct {
	Name    string
	Phone   string
	Age     int
	Rating  float64
	isClose bool
}

func (user *User) addRating(balls float64) {
	user.Rating += balls
}

func newUser(
	name string,
	phone string,
	age int,
	rating float64,
	isClose bool,
) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	if phone == "" {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
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
