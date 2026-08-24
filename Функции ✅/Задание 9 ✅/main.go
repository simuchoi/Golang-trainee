package main

import "fmt"

var globalVar float64 = 2

func main() {
	fmt.Println("globalVar before multiply", globalVar)

	multiplyGlobalVar()
	fmt.Println("globalVar after multiply", globalVar)

	divideGlobalVar()
	fmt.Println("globalVar after divide", globalVar)
}

func multiplyGlobalVar() float64 {
	globalVar *= 10
	return globalVar
}

func divideGlobalVar() float64 {
	globalVar /= 2
	return globalVar
}
