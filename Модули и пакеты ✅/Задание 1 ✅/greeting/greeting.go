package greeting

import "fmt"

// чтобы скрыть доступ необходимо назвать с маленькой буквы (функци, поля структуры и т.д.)
func SayHello() {
	fmt.Println("Hello")

	num := ReturnInt()
	
	fmt.Println("ReturnInt from greeting.go: ", num)
}