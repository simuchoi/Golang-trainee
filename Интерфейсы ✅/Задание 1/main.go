package main

import (
	"fmt"
	"reflect"
)

// структура car
type car struct {
	Brand        string  // марка
	Model        string  // модель
	Year         int     // год выпуска
	Color        string  // цвет
	Power        int     // мощность в л.с.
	EngineVolume float64 // объём двигателя в литрах
	FuelType     string  // тип топлива
	Transmission string  // коробка передач
	DriveTrain   string  // привод ("FWD" - передний, "RWD" - задний, "AWD" - полный)
	Doors        int     // количество д
	Mileage      int     // пробег в км
	IsNew        bool    // состояние (новая - true) / (б/у - false)
}

// интерфейс car - объявляет набор методов
type Car interface {
	Gas()
	Brake()
	GetInfo()
}

// метод структуры car - выводит информацию об ускорение машины
func (c car) Gas() {
	fmt.Printf("\nGas, %s. Let's go!\n", c.Brand)
}

// метод структуры car - выводит информацию о тороможении машины
func (c car) Brake() {
	fmt.Printf("\nBrake, %s! Slow down!", c.Brand)
}

// метод структуры car - выводит всю информацию о машине
func (c car) GetInfo() {
	fmt.Print("\n---------------------------------------------------------")
	fmt.Printf("\nHere you are, that's all information about")
	fmt.Printf("\n%s %s %d i've got for today: ", c.Brand, c.Model, c.Year)
	fmt.Print("\n---------------------------------------------------------")
	c.PrintFields()
}

// дополнительный метод структуры car для метода GetInfo() - выводит поля структуры
func (c car) PrintFields() {
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)

	for i := 3; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i)
		fmt.Printf("\n%s: %v", fieldName, fieldValue)
	}
	fmt.Println()
}

// функция, принимающая интерфейс Car, вызывающая его методы
func carControl(vehicle Car) {
	vehicle.Gas()
	vehicle.Brake()
	vehicle.GetInfo()
}

func main() {
	// объявление экземпляра структуры
	honda := car{
		Brand:        "Honda",
		Model:        "Civic Type R",
		Year:         2023,
		Color:        "Yellow",
		Power:        330,
		EngineVolume: 2.0,
		FuelType:     "Benzine",
		Transmission: "Manual",
		DriveTrain:   "FWD",
		Doors:        5,
		Mileage:      12000,
		IsNew:        false,
	}

	carControl(honda)
}
