package car

import "errors"

type Car struct {
	Model  string
	Engine float64
	Wheels int
	Doors  int
	Color  string
	IsNew  bool
	Price  float64
}

func NewCar(
	model string,
	engine float64,
	wheels int,
	doors int,
	color string,
	isNew bool,
	price float64,
) (Car, error) {
	if model == "" {
		return Car{}, errors.New("Missing model name")
	}

	if engine > 10 {
		return Car{}, errors.New("Too big engine")
	}

	if wheels < 4 {
		return Car{}, errors.New("Missing some wheels")
	}

	if doors < 3 {
		return Car{}, errors.New("Missing some doors")
	}

	if color == "" {
		return Car{}, errors.New("No color")
	}

	if price < 0 {
		return Car{}, errors.New("Is it for free?")
	}

	return Car{
		Model:  model,
		Engine: engine,
		Wheels: wheels,
		Doors:  doors,
		Color:  color,
		IsNew:  isNew,
		Price:  price,
	}, nil
}

func (car *Car) ChangeEngine(volume float64) {
	car.Engine = volume
}
