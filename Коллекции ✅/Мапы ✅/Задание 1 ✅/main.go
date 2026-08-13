package main

import "fmt"

func main() {
	weather := map[int]int{ // формат хранения данных парой ключ: значение
		1: +2,
		2: +3,
		3: +4,
	}

	fmt.Println(weather[1])

	c, ok := weather[2] // проверка на значение по умолчанию
	fmt.Println(c, ok)  // false - значение по умолчанию

	weather[20] = -2 // добавление новой пары ключ: значение
	fmt.Println(weather[20])

	for _, value := range weather { // цикл для прохода по мапе, работа с копией
		value += 5
		fmt.Println(value)
	}

	for key, _ := range weather { // цикл для прохода по мапе, работа с оригиналом
		weather[key] += 1
	}

	fmt.Println(weather)
}
