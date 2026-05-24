package main

import "fmt"

// ТЗ: вывести числа от 1 до 100, заменяя:
// кратные 3 на "Fizz", кратные 5 на "Buzz", кратные 3 и 5 на "FizzBuzz".
func main() {
	fmt.Println("Числа от 1 до 100 (Fizz, Buzz, FizzBuzz)")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
