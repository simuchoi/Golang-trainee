package main

import "fmt"

func main() {
	isEmployd := map[string]bool{
		"Tom":  false,
		"John": true,
	}

	c, ok := isEmployd["John"]

	if !ok { // ключ значение отсуствует
		fmt.Println("c:", c, "\nok:", ok)
		fmt.Println("\nPerson doesn't exist")
		return
	}

	if c { // ключ значение присутсвтует
		fmt.Println("\nPerson is employed")
	} else {
		fmt.Println("\nPerson is not employed")
	}

	fmt.Println("c:", c, "\nok:", ok)
}
