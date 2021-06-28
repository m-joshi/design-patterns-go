package condiments

import (
	beverage2 "design-patterns-go/decoratorPattern/beverage"
)

type Whip struct {
	description string
	cost        int
	beverage    beverage2.Beverage
}

func NewWhip(beverage beverage2.Beverage) Whip {
	return Whip{
		description: "Whip",
		cost:        30,
		beverage:    beverage,
	}
}

func (wp Whip)GetDescription() string {
	return wp.beverage.GetDescription() +" "+ wp.description
}

func (wp Whip)Cost() int {
	return wp.beverage.Cost() + wp.cost
}
