package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nВведите команду: ")
		ok := scanner.Scan()

		if !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}

		str := ""

		for i := 1; i < len(fields); i++ {
			str += fields[i]

			if i != len(fields)-1 {
				str += " "
			}
		}

		cmd := fields[0]

		if cmd == "Добавить" {
			fmt.Println("Вы хотите добавить следующие продукты:", str)
		} else if cmd == "Удалить" {
			fmt.Println("Вы хотите удалить следующие элементы:", str)
		} else if cmd == "Помощь" {
			fmt.Println("Доступные команды: \n- Добавить \n- Удалить \n- Помощь")
		} else if cmd == "Выйти" {
			return
		} else {
			fmt.Println("Неизвестная команда")
		}
	}
}
