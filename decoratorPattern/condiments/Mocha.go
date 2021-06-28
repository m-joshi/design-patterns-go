package condiments

import (
	beverage2 "design-patterns-go/decoratorPattern/beverage"
)

type Mocha struct {
	description string
	cost        int
	beverage    beverage2.Beverage
}

func NewMocha(beverage beverage2.Beverage) Mocha {
	return Mocha{
		description: "Mocha",
		cost:        20,
		beverage:    beverage,
	}
}

func (mc Mocha)GetDescription() string {
	return mc.beverage.GetDescription() +" "+ mc.description
}

func (mc Mocha)Cost() int {
	return mc.beverage.Cost() + mc.cost
}

