package main

import "fmt"

// ТЗ: реализовать меню банкомата в цикле с начальным балансом 1000:
// пополнение, снятие (с проверкой остатка), просмотр баланса и выход.
func main() {
	choice := 0
	actualBalance := 1000
	desiredAmount := 0

	fmt.Println("== Банкомат ==")
	fmt.Println("Ваш текущий баланс:", actualBalance)

	for {
		fmt.Println("\n1. Пополнить баланс")
		fmt.Println("2. Снять деньги")
		fmt.Println("3. Показать баланс")
		fmt.Println("4. Выйти")

		fmt.Print("\nВыберите пункт: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Введите сумму пополнения: ")
			fmt.Scan(&desiredAmount)
			actualBalance += desiredAmount
			fmt.Println("Ваш текущий баланс:", actualBalance)

		case 2:
			fmt.Print("Введите сумму снятия: ")
			fmt.Scan(&desiredAmount)

			if desiredAmount > actualBalance {
				fmt.Println("Недостаточно средств")
			} else {
				actualBalance -= desiredAmount
				fmt.Println("Вы сняли", desiredAmount, "руб.")
				fmt.Println("Ваш текущий баланс:", actualBalance)
			}

		case 3:
			fmt.Println("Ваш текущий баланс:", actualBalance)

		case 4:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Такого пункта меню нет")
		}
	}
}
