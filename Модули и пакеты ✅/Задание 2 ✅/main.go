package main

import (
	"fmt"
	"task2/greeting"

	"github.com/k0kubun/pp/v3"
)

func main() {
	fmt.Println("hello i'm main() from packet Задание 2")
	greeting.SayHello()

	person1 := newPerson("Oleg", "Solovyov", "Olegovich")
	pp.Println(person1)
}

type person struct {
	name     string
	midname  string
	lastname string
}

func newPerson(name string, midname string, lastname string) person {
	return person{
		name:     name,
		midname:  midname,
		lastname: lastname,
	}

}
