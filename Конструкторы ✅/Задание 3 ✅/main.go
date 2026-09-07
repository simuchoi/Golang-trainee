package main

import (
	"errors"
	"fmt"
)

func main() {
	user, err := newUser(
		"Oleg",
		23,
		"+79154078632",
		10.0,
		false,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.String())
}

func (u User) String() string {
	return fmt.Sprintf(
		"{Name: %s, Age: %d, Phone: %s, Rating: %.1f, isClose: %v}",
		u.Name, u.Age, u.Phone, u.Rating, u.isClose,
	)
}

type User struct {
	Name    string
	Age     int
	Phone   string
	Rating  float64
	isClose bool
}

func newUser(name string, age int, phone string, rating float64, isClose bool) (User, error) {
	if name == "" {
		return User{}, errors.New("Empty name!")
	}

	if age < 0 || age > 150 {
		return User{}, errors.New("Age out of range!")
	}

	if phone == "" {
		return User{}, errors.New("Empty phone number!")
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}, errors.New("Rating out of range!")
	}

	return User{
		Name:    name,
		Age:     age,
		Phone:   phone,
		Rating:  rating,
		isClose: isClose,
	}, nil
}
