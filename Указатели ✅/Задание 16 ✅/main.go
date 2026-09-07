package main

import "fmt"

func main() {

	var myINT int = 2
	var myFLOAT32 float32 = 4.0
	var myFLOAT64 float64 = 8.0
	var myBOOL bool = false
	var mySTRING string = "string"

	// записываем в новые переменные ссылки/адреса на переменные, в которых хранится значение
	myINT_ptr := &myINT // в переменную ptrInt записали ссылку/адрес, где находится переменная varInt
	myFLOAT32_ptr := &myFLOAT32
	myFLOAT64_ptr := &myFLOAT64
	myBOOL_ptr := &myBOOL
	mySTRING_ptr := &mySTRING

	fmt.Println()
	fmt.Println("Get myINT by myINT_ptr")
	fmt.Println("myINT value:", *myINT_ptr)
	fmt.Println("myINT address:", &myINT_ptr)
	fmt.Println()

	fmt.Println("Get myFLOAT32 by myFLOAT32_ptr")
	fmt.Println("myFLOAT32 value:", *myFLOAT32_ptr)
	fmt.Println("myFLOAT32 address:", &myFLOAT32_ptr)
	fmt.Println()

	fmt.Println("Get myFLOAT64 by myFLOAT64_ptr")
	fmt.Println("myFLOAT64 value:", *myFLOAT64_ptr)
	fmt.Println("myFLOAT64 address:", &myFLOAT64_ptr)
	fmt.Println()

	fmt.Println("Get myBOOL by myBOOL_ptr")
	fmt.Println("myBOOL value:", *myBOOL_ptr)
	fmt.Println("myBOOL address:", &myBOOL_ptr)
	fmt.Println()

	fmt.Println("Get mySTRING by mySTRING_ptr")
	fmt.Println("mySTRING value:", *mySTRING_ptr)
	fmt.Println("mySTRING address:", &mySTRING_ptr)
	fmt.Println()
}
