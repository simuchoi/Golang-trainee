package main

import "fmt"

func main() {
	myMap := map[string]int{
		"isComplete": 1,
		"isNew":      0,
	}

	for i, v := range myMap {
		fmt.Println(i, v)
	}
}
