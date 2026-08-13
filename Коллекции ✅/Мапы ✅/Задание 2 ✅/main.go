package main

import "fmt"

func main() {
	weather := make(map[int]int, 10) // предвыделение памяти для мапы, когда заранее известно сколько пар

	weather[20] = -2 // добавление новой пары ключ: значение
	fmt.Println(weather[20])

	for key, _ := range weather { // цикл для прохода по мапе, работа с оригиналом
		weather[key] += 1
	}

	fmt.Println(weather)
}
