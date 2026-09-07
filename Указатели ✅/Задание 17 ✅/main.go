package main

import "fmt"

func main() {
	number := 32
	pointer := &number

	fmt.Println("\nnumber = ", number, " \n--- значение переменной number")
	fmt.Println("\n&number = ", &number, " \n--- адрес переменной number")
	fmt.Println("\n*pointer = ", *pointer, " \n--- значение переменной number через указатель pointer")
	fmt.Println("\n&pointer = ", &pointer, " \n--- адрес переменной pointer")
}
