package main

import "fmt"

type myStruct struct {
	Num int
	Str string
}

func main() {
	fmt.Println()
	intSlice := []int{}                    // память не выделена, значений нет
	strSlice := make([]string, 0)          // память не выделена, значений нет
	floatSlice := make([]float64, 10)      // память выделена, 10 нулевых значений
	structSlice := make([]myStruct, 0, 10) // память выделена, значений нет

	intSlice = append(intSlice, 1)
	strSlice = append(strSlice, "a")
	floatSlice = append(floatSlice, 3.14)
	floatSlice = append(floatSlice, 3.14)
	structSlice = append(structSlice, myStruct{Num: 1, Str: "b"})
	structSlice = append(structSlice, myStruct{Num: 1, Str: "b"})

	fmt.Println("intSlice: ", intSlice, len(intSlice), cap(intSlice))            // len=0, cap=0
	fmt.Println("strSlice: ", strSlice, len(strSlice), cap(intSlice))            // len=0, cap=0
	fmt.Println("floatSlice: ", floatSlice, len(floatSlice), cap(floatSlice))    // len=10, cap=10
	fmt.Println("structSlice: ", structSlice, len(structSlice), cap(floatSlice)) // len=0, cap=10
}
