package main

import "fmt"

func main() {
	num := 10
	fmt.Println("num до изменений: ", num)
	changeNum(&num)
	fmt.Println("num после изменений:", num)

	float := 1.0
	fmt.Println("float до изменений: ", float)
	changeFloat(&float)
	fmt.Println("float после изменений: ", float)

	text := "hello"
	fmt.Println("text до изменений: ", text)
	changeText(&text)
	fmt.Println("text после: ", text)

	boolean := false
	fmt.Println("boolean до изменений: ", boolean)
	changeBoolean(&boolean)
	fmt.Println("boolean после изменений: ", boolean)
}

func changeNum(num *int) {
	*num = 20
}

func changeFloat(float *float64) {
	*float = 3.14
}

func changeText(text *string) {
	*text = "goodbye"
}

func changeBoolean(boolean *bool) {
	*boolean = true
}
