package main

import "fmt"

// Структура User
type User struct {
	name   string
	rating float64
}

// Метод для структуры User
func (u User) greeting() {
	fmt.Println("Привет, я -", u.name)
	fmt.Println("Мой рейтинг -", u.rating)

	u.name = "Olegafriend"
}

// Обычная функция с User
func goodbye(u User) {
	fmt.Println("Пока")
	fmt.Println(u.name, "Покинул нас")
}

func (u User) goFuckYourself() {
	fmt.Println("Пошел нах")
	fmt.Println("Это тебе говорит", u.name)
}

func (u *User) ratingIncrease(rating float64) {
	if u.rating+rating <= 10 {
		u.rating += rating
		fmt.Println("\nРейтинг повышен")
	} else {
		fmt.Println("Рейтинг максимален")
		fmt.Println("Текущий рейтинг: ", u.rating)
	}

}

// Главная функция
func main() {
	user := User{
		name:   "Oleg",
		rating: 10.0,
	}

	user.greeting()

	user.goFuckYourself()

	goodbye(user)

	user.ratingIncrease(-2)

}
