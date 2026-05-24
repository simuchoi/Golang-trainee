package main

import "fmt"

// ТЗ: принять число от 1 до 9 и вывести для него таблицу умножения от 1 до 10.
func main() {
	num := 0

	fmt.Println("Таблица умножения")
	fmt.Print("Введите число от 1 до 9: ")
	fmt.Scan(&num)

	fmt.Println("Таблица умножения для числа", num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
