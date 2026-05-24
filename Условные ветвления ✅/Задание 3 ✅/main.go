package main

import "fmt"

func main() {
	sideA := 0
	sideB := 0
	sideC := 0

	fmt.Print("Введите сторону A: ")
	fmt.Scan(&sideA)

	fmt.Print("Введите сторону B: ")
	fmt.Scan(&sideB)

	fmt.Print("Введите сторону C: ")
	fmt.Scan(&sideC)

	if (sideA+sideB > sideC) && (sideB+sideC > sideA) && (sideA+sideC > sideB) {
		fmt.Println("Треугольник существует")

		if sideA == sideB && sideB == sideC {
			fmt.Println("Треугольник равносторонний")
		} else if sideA == sideB || sideB == sideC || sideC == sideA {
			fmt.Println("Треугольник равнобедренный")
		} else {
			fmt.Println("Треугольник разносторонний")
		}
	} else {
		fmt.Println("Треугольник не существует")
	}
}
