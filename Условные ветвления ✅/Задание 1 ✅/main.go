package main

import "fmt"

func main() {
	grade := 0

	fmt.Print("Введите балл (0-100): ")
	fmt.Scan(&grade)

	if grade >= 0 && grade <= 39 {
		fmt.Println("Твоя оценка: F")
	} else if grade >= 40 && grade <= 59 {
		fmt.Println("Твоя оценка: D")
	} else if grade >= 60 && grade <= 74 {
		fmt.Println("Твоя оценка: C")
	} else if grade >= 75 && grade <= 89 {
		fmt.Println("Твоя оценка: B")
	} else if grade >= 90 && grade <= 100 {
		fmt.Println("Твоя оценка: A")
	} else {
		fmt.Println("Странные какие-то баллы")
	}
}
