package main

import "fmt"

type myStruct struct {
	number int
	text   string
}

func main() {
	arrInt := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrInt: ", arrInt)

	arrFloat := [5]float64{1.33, 1.67, 2.00, 2.33, 2.67}
	fmt.Println("arrFloat: ", arrFloat)

	arrBool := [5]bool{true, false, false, true, false}
	fmt.Println("arrBool: ", arrBool)

	arrString := [5]string{"h", "e", "l", "l", "o"}
	fmt.Println("arrString: ", arrString)

	arrStruct := [3]myStruct{
		{
			number: 3,
			text:   "hello",
		},
		{
			number: 6,
			text:   "hello",
		},
		{
			number: 9,
			text:   "hello",
		},
	}
	fmt.Println("arrStruct: ", arrStruct)

	fmt.Println("\nfor i := 0")

	for i := 0; i < len(arrInt); i++ {
		arrInt[i] += 1
	}

	for i := 0; i < len(arrFloat); i++ {
		arrFloat[i] += 0.33
	}

	for i := 0; i < len(arrBool); i++ {
		if arrBool[i] == true {
			arrBool[i] = false
			continue
		}
		arrBool[i] = true
	}

	for i := 0; i < len(arrString); i++ {
		arrString[i] += "+"
	}

	for i := 0; i < len(arrStruct); i++ {
		arrStruct[i] = arrStruct[1]
	}

	fmt.Println("arrInt: ", arrInt)
	fmt.Println("arrInt: ", arrFloat)
	fmt.Println("arrInt: ", arrBool)
	fmt.Println("arrInt: ", arrString)

	fmt.Println("\nfor i, v := range arr")

	fmt.Println("\narrInt: ")
	for i, v := range arrInt {
		arrInt[i] += 10
		fmt.Println(i, v)
	}

	fmt.Println("\narrFLoat: ")
	for i, v := range arrFloat {
		arrFloat[i] += 0.33
		fmt.Println(i, v)
	}

	fmt.Println("\narrString: ")
	for i, v := range arrString {
		arrString[i] += "0w++"
		fmt.Println(i, v)
	}

	fmt.Println("\narrStruct:")
	for i, v := range arrStruct {
		fmt.Println(i, v)
	}

	fmt.Println()
}
