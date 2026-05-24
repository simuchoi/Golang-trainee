package main

import "fmt"

// Задание 4 — Функции (лёгкий уровень)
//
// Три маленькие функции и вывод в main. Без строковых проверок по символам и без ошибок.
//
// 1) double(x int) int
//    Верни x, умноженное на 2.
//
// 2) max2(a, b int) int
//    Верни большее из двух чисел (если равны — можно вернуть любое из них).
//
// 3) hello(name string) string
//    Верни одну строку: слово "Привет, ", потом name, потом "!" (склейка через +).
//
// 4) main()
//    Вызови double и max2 с парой чисел, выведи результаты через fmt.Println.
//    Вызови hello с каким-нибудь именем и выведи возвращённую строку.
//
// Подсказка: в max2 достаточно одного if: если a >= b вернуть a, иначе b.

func main() {
	fmt.Println(double(5))
	fmt.Println(max2(10, 7))
	fmt.Println(hello("Анна"))
}

func double(x int) int {
	return x * 2
}

func max2(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func hello(name string) string {
	return "Привет, " + name + "!"
}
