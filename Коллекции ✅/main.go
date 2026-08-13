package main

import "fmt"

func main() {
	// массив
	myArr := [3]int{1, 2, 3}
	fmt.Println("\nmyArr:", myArr)
	fmt.Println("len:", len(myArr))
	fmt.Println("cap:", cap(myArr))

	// слайс
	mySlice := make([]int, 0, 5)

	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 6)

	fmt.Println("\nmySlice:", mySlice)
	fmt.Println("len:", len(mySlice))
	fmt.Println("cap:", cap(mySlice))

	// мапа
	myMap := make(map[string]int, 10)

	myMap["John"] = 30
	myMap["Tom"] = 35
	myMap["Sam"] = 40

	fmt.Println("\nmyMap:", myMap)
	fmt.Println("len:", len(myMap))
}
