package condiments

import (
	beverage2 "design-patterns-go/decoratorPattern/beverage"
)

type Milk struct {
	description string
	cost        int
	beverage    beverage2.Beverage
}

func NewMilk(beverage beverage2.Beverage) Milk {
	return Milk{
		description: "Milk",
		cost:        10,
		beverage:    beverage,
	}
}

func (ml Milk)GetDescription() string {
	return ml.beverage.GetDescription() +" "+ ml.description
}

func (ml Milk)Cost() int {
	return ml.beverage.Cost() + ml.cost
}
