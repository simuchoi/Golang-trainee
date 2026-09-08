package ice

type iceCream struct {
	Flavour string
	Price   float64
}

func NewIceCream(flavour string, price float64) iceCream {
	if flavour == "" {
		return iceCream{}
	}

	if price < 0 {
		return iceCream{}
	}

	return iceCream{
		Flavour: flavour,
		Price:   price,
	}
}

func (i *iceCream) changePrice(price float64) {
	i.Price = 0
	i.Price = price
}
