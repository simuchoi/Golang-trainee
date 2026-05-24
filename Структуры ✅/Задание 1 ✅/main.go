package main

import "fmt"

func main() {
	userOleg := User{
		Name:    "Oleg",
		Age:     23,
		Phone:   "+7 (915) 407-86-32",
		isClose: false,
		Rating:  10.0,
	}

	fmt.Println("Name: ", userOleg.Name)
	fmt.Println("Age: ", userOleg.Age)
	fmt.Println("Phone: ", userOleg.Phone)
	fmt.Println("isClose: ", userOleg.isClose)
	fmt.Printf("Rating: %.2f\n", userOleg.Rating)

	userOleg.Name = "Olegan"

	fmt.Println("\nnew Name: ", userOleg.Name)
}

type User struct {
	Name    string
	Age     int
	Phone   string
	isClose bool
	Rating  float64
}
