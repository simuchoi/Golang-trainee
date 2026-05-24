package main

import (
	"errors"
	"fmt"
)

// Задание 3 — Конструкторы (продолжение задания 2)
//
// Цель: научиться писать идиоматичный конструктор в Go — функция возвращает
// значение и error вместо «пустой» структуры и сообщений в консоль.
//
// Ориентир: возьми за основу решение из «Задание 2» (тип User, поля, те же правила валидации).
//
// 1) Объяви тип User с полями:
//    Name (string), Phone (string), Age (int), Rating (float64), isClose (bool, неэкспортируемое).
//
// 2) Реализуй конструктор newUser с сигнатурой:
//    func newUser(name, phone string, age int, rating float64, isClose bool) (User, error)
//
// 3) Валидация — те же правила, что в задании 2, но без fmt.Println внутри конструктора:
//    — имя не пустое;
//    — возраст: age > 0 и age < 150;
//    — телефон не пустой;
//    — рейтинг: от 0.0 до 10.0 включительно.
//    При ошибке возвращай User{} и errors.New("...") с понятным текстом на русском
//    (например: "имя не может быть пустым", "некорректный возраст" и т.д.).
//    При успехе — заполненный User и nil.
//
// 4) Реализуй метод String() string для User (интерфейс fmt.Stringer),
//    чтобы при выводе через fmt.Print / fmt.Println пользователь читался как одна строка,
//    например: {Name: Oleg, Phone: +7915..., Age: 23, Rating: 9.90, isClose: false}
//
// 5) В main:
//    — подключи пакеты fmt и errors;
//    — создай пользователя через newUser с валидными аргументами (как в задании 2);
//    — если err != nil — выведи ошибку (fmt.Println) и заверши main через return;
//    — иначе выведи пользователя через fmt.Println(user).
//
// 6) Дополнительно (по желанию): второй вызов newUser с невалидными данными
//    и проверка, что err != nil и в консоли видно сообщение об ошибке.
//
// Подсказка: конструктор не должен печатать в консоль — только возвращать результат и ошибку.
// Обработка ошибки — ответственность вызывающего кода (main).

// TODO: реализовать по ТЗ выше

func main() {
	user, err := newUser(
		"Oleg",
		23,
		"+79154078632",
		10.0,
		false,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.String())
}

func (u User) String() string {
	return fmt.Sprintf(
		"{Name: %s, Age: %d, Phone: %s, Rating: %.1f, isClose: %v}",
		u.Name, u.Age, u.Phone, u.Rating, u.isClose,
	)
}

type User struct {
	Name    string
	Age     int
	Phone   string
	Rating  float64
	isClose bool
}

func newUser(name string, age int, phone string, rating float64, isClose bool) (User, error) {
	if name == "" {
		return User{}, errors.New("Empty name!")
	}

	if age < 0 || age > 150 {
		return User{}, errors.New("Age out of range!")
	}

	if phone == "" {
		return User{}, errors.New("Empty phone number!")
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}, errors.New("Rating out of range!")
	}

	return User{
		Name:    name,
		Age:     age,
		Phone:   phone,
		Rating:  rating,
		isClose: isClose,
	}, nil
}
