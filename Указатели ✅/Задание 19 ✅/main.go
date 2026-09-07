package main

import "fmt"

func main() {
	num := 10
	numPtr := &num
	var numNilPtr *int = nil
	changeNum(numNilPtr)
	changeNum(numPtr)

	text := "text"
	textPtr := &text
	var textNipPtr *string = nil
	changeText(textNipPtr)
	changeText(textPtr)

	boolean := false
	booleanPtr := &boolean
	var booleanNilPtr *bool = nil
	changeBoolean(booleanNilPtr)
	changeBoolean(booleanPtr)

	float := 1.0
	floatPtr := &float
	var floatNilPtr *float64 = nil
	changeFloat(floatNilPtr)
	changeFloat(floatPtr)
}

func changeNum(ptr *int) {
	if ptr == nil {
		fmt.Println("\nptr num is nil")
		return
	}

	*ptr = 20
	fmt.Println("ptr num is not nil")
	fmt.Println("num: ", *ptr)

}

func changeFloat(ptr *float64) {
	if ptr == nil {
		fmt.Println("\nptr float is nil")
		return
	}

	*ptr = 3.14
	fmt.Println("ptr float is not nil")
	fmt.Println("float: ", *ptr)

}

func changeText(ptr *string) {
	if ptr == nil {
		fmt.Println("\nptr text is nil")
		return
	}

	*ptr = "new text"
	fmt.Println("ptr text is not nil")
	fmt.Println("text: ", *ptr)

}

func changeBoolean(ptr *bool) {
	if ptr == nil {
		fmt.Println("\nptr boolean is nil")
		return
	}

	fmt.Println("ptr boolean is not nil")
	fmt.Println("boolean: ", *ptr)
	*ptr = true
}
