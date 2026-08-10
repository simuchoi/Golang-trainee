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

	userSIMUCHOI := User{
		Name:    "Oleg",
		Rating:  10.0,
		Premium: true,
	}

	userPERSEID := User{
		Name:    "Nikita",
		Rating:  8.5,
		Premium: true,
	}

	userCOFFEEK := User{
		Name:    "Klim",
		Rating:  9.0,
		Premium: false,
	}

	fmt.Println("\nПользователи c премиумом до:")
	pp.Println(userSIMUCHOI)
	pp.Println(userPERSEID)
	pp.Println(userCOFFEEK)

	if userSIMUCHOI.Premium == true {
		userSIMUCHOI.Rating += 1
	}

	if userPERSEID.Premium == true {
		userPERSEID.Rating += 1
	}

	fmt.Println("\nПользователи c премиумом после:")
	pp.Println(userSIMUCHOI)
	pp.Println(userPERSEID)
	pp.Println(userCOFFEEK)

	// userArray := [3]User{userSIMUCHOI, userPERSEID, userCOFFEEK} // копии пользователей!
}
