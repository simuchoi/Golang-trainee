package main

import (
	"fmt"
)

// Задание 5 — Структуры (лёгкий+ уровень)
//
// Цель: методы с приёмником по указателю (*T), чтобы менять поля структуры из метода.
// Ошибки (error) не используй; проверять «хватает ли денег» не нужно.
//
// 1) Объяви тип Wallet — структура с полем Rubles (int).
//
// 2) Метод Add(amount int) с приёмником *Wallet: увеличь Rubles на amount.
//
// 3) Метод Spend(amount int) с приёмником *Wallet: уменьши Rubles на amount.
//    Отрицательный баланс допускается — условия в Spend не проверяй.
//
// 4) В main:
//    — подключи пакет fmt;
//    — создай Wallet с начальным Rubles (любое число, например 100);
//    — вызови Add с одним числом, затем Spend с другим;
//    — выведи итоговый Rubles через fmt.Println.
//
// Подсказка: сигнатуры методов начинаются с (w *Wallet).

func main() {
	w := Wallet{
		Rubles: 100,
	}

	fmt.Println("Wallet Rubles Amount: ", w.Rubles)

	w.Add(100)

	fmt.Println("add Wallet Rubles Amount: ", w.Rubles)

	w.Spend(150)

	fmt.Println("spend Wallet Rubles Amount:", w.Rubles)
}

type Wallet struct {
	Rubles int
}

func (w *Wallet) Add(amount int) {
	w.Rubles += amount
}

func (w *Wallet) Spend(amount int) {
	w.Rubles -= amount
}
