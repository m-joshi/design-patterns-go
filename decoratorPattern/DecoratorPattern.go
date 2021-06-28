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
	mocha := condiments.NewMocha()
	finalProduct := condiments.NewWhip()
	finalDescription := finalProduct.GetDescription(mocha.GetDescription(darkRoast.GetDescription))
	finalCost := finalProduct.Cost(mocha.Cost(darkRoast.Cost))

	fmt.Println("Here's your - ", finalDescription())
	fmt.Println("It costs - ", finalCost())

	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("Making an order - House Blend with Double Mocha and Milk")

	houseBlend := beverage.NewHouseBlend()
	mocha1 := condiments.NewMocha()
	mocha2 := condiments.NewMocha()
	milk := condiments.NewMilk()

	houseBlendDescriptionFunc := houseBlend.GetDescription
	mocha1Description := mocha1.GetDescription(houseBlendDescriptionFunc)
	mocha2Description := mocha2.GetDescription(mocha1Description)
	newDescription := milk.GetDescription(mocha2Description)

	houseCostFunc := houseBlend.Cost
	mocha1Cost := mocha1.Cost(houseCostFunc)
	mocha2Cost := mocha2.Cost(mocha1Cost)
	newCost := milk.Cost(mocha2Cost)

	fmt.Println("Here's your - ", newDescription())
	fmt.Println("It costs - ", newCost())
}

