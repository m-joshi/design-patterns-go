package decoratorPattern

import (
	"design-patterns-go/decoratorPattern/beverage"
	"design-patterns-go/decoratorPattern/condiments"
	"fmt"
)

type DecoratorPattern interface {
	Run()
}

type decoratorPattern struct {
	decoratorPattern DecoratorPattern
}

func NewDecorator() decoratorPattern {
	return decoratorPattern{}
}

func (dc decoratorPattern) Run() {
	fmt.Println("~~~~~~~~~~ DECORATOR PATTERN ~~~~~~~~~~~~~")

	fmt.Println("Making an order - Dark Roast with Mocha and Whip")

	darkRoast := beverage.NewDarkRoast()
	mocha := condiments.NewMocha(darkRoast)
	finalProduct := condiments.NewWhip(mocha)

	fmt.Println("Here's your - ", finalProduct.GetDescription())
	fmt.Println("It costs - ", finalProduct.Cost())

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Making an order - House Blend with Double Mocha and Milk")

	newProduct := condiments.NewMilk(condiments.NewMocha(condiments.NewMocha(beverage.NewHouseBlend())))

	fmt.Println("Here's your - ", newProduct.GetDescription())
	fmt.Println("It costs - ", newProduct.Cost())
}

