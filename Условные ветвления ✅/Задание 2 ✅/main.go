package main

import "fmt"

func main() {
	hour := 0

	fmt.Print("Введите время (0-23): ")
	fmt.Scan(&hour)

	if hour >= 0 && hour <= 5 {
		fmt.Println("Утро")
	} else if hour >= 6 && hour <= 11 {
		fmt.Println("День")
	} else if hour >= 12 && hour <= 17 {
		fmt.Println("Вечер")
	} else if hour >= 18 && hour <= 23 {
		fmt.Println("Ночь")
	} else {
		fmt.Println("Странное какое-то время")
	}
}
