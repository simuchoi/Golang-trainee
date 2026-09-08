package apartment

import "fmt"

type apartment struct {
	number int
	floor  int
	area   float64
}

func NewApartment(
	number int,
	floor int,
	area float64,
) apartment {
	if number < 0 {
		fmt.Println("wront number")
		return apartment{}
	}

	if floor < 0 {
		fmt.Println("wront floor")
		return apartment{}
	}

	if area < 0 {
		fmt.Println("wront area")
		return apartment{}
	}

	fmt.Println("all is correct")
	return apartment{
		number: number,
		floor:  floor,
		area:   area,
	}
}
