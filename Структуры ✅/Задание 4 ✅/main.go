package main

import (
	"fmt"
)

// Задание 4 — Структуры (лёгкий уровень)
//
// Цель: литерал структуры, поля, метод с приёмником по значению. Указатели не нужны.
//
// 1) Объяви тип Rectangle — структура с полями Width и Height (оба int).
//
// 2) Объяви метод Area() int для Rectangle с приёмником по значению:
//    func (r Rectangle) Area() int { ... }
//    Тело: вернуть произведение ширины и высоты.
//
// 3) В main:
//    — подключи пакет fmt;
//    — создай значение Rectangle через литерал полей (любые два положительных int);
//    — выведи Width и Height (fmt.Println — как удобнее: по одному или вместе);
//    — вызови Area() и выведи результат через fmt.Println.

func main() {
	r := Rectangle{
		Height: 10,
		Width:  50,
	}

	fmt.Println("Height: ", r.Height)
	fmt.Println("Width: ", r.Width)
	fmt.Println("Area: ", r.Area())
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Height * r.Width
}
